package main

import (
	"github.com/containerd/containerd/defaults"
	"github.com/containerd/containerd/server"
)

func defaultConfig() *server.Config {
	return &server.Config{
		Root:  defaults.DefaultRootDir,
		State: defaults.DefaultStateDir,
		GRPC: server.GRPCConfig{
			// 默认的grpc地址为"/run/containerd/containerd.sock"
			Address: defaults.DefaultAddress,
		},
		Debug: server.Debug{
			Level:   "info",
			Address: defaults.DefaultDebugAddress,
		},
	}
}
