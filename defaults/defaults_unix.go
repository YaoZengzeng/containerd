// +build !windows

package defaults

const (
	// DefaultRootDir is the default location used by containerd to store
	// persistent data
	// DefaultRootDir用来存储持久数据
	DefaultRootDir = "/var/lib/containerd"
	// DefaultStateDir is the default location used by containerd to store
	// transient data
	// DefaultStateDir用来存储临时数据
	DefaultStateDir = "/run/containerd"
	// DefaultAddress is the default unix socket address
	DefaultAddress = "/run/containerd/containerd.sock"
	// DefaultDebugAddress is the default unix socket address for pprof data
	DefaultDebugAddress = "/run/containerd/debug.sock"
)
