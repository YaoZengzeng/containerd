// +build linux darwin freebsd solaris

package main

import (
	"context"
	"os"
	"runtime"

	"github.com/containerd/containerd/log"
	"github.com/containerd/containerd/server"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
)

const defaultConfigPath = "/etc/containerd/config.toml"

var handledSignals = []os.Signal{
	unix.SIGTERM,
	unix.SIGINT,
	unix.SIGUSR1,
	unix.SIGPIPE,
}

func handleSignals(ctx context.Context, signals chan os.Signal, serverC chan *server.Server) chan struct{} {
	done := make(chan struct{}, 1)
	go func() {
		var server *server.Server
		for {
			select {
			case s := <-serverC:
				// 从serverC中获取server结构
				server = s
			case s := <-signals:
				log.G(ctx).WithField("signal", s).Debug("received signal")
				switch s {
				case unix.SIGUSR1:
					dumpStacks()
				case unix.SIGPIPE:
					continue
				default:
					if server == nil {
						// 如果containerd server未启动成功，则只把done关掉
						close(done)
						return
					}
					// 先停止containerd sever
					server.Stop()
					// 关闭done
					close(done)
				}
			}
		}
	}()
	return done
}

func dumpStacks() {
	var (
		buf       []byte
		stackSize int
	)
	bufferLen := 16384
	for stackSize == len(buf) {
		buf = make([]byte, bufferLen)
		stackSize = runtime.Stack(buf, true)
		bufferLen *= 2
	}
	buf = buf[:stackSize]
	logrus.Infof("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf)
}
