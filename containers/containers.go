package containers

import (
	"context"
	"time"

	"github.com/gogo/protobuf/types"
)

// Container represents the set of data pinned by a container. Unless otherwise
// noted, the resources here are considered in use by the container.
// Container代表一系列和container绑定的数据，除非特殊说明，这里的资源都被认为容器正在使用
//
// The resources specified in this object are used to create tasks from the container.
// 本对象中声明的资源用于从container中创建task
type Container struct {
	// ID uniquely identifies the container in a nameapace.
	//
	// This property is required and cannot be changed after creation.
	ID string

	// Labels provide metadata extension for a contaienr.
	//
	// These are optional and fully mutable.
	Labels map[string]string

	// Image specifies the image reference used for a container.
	//
	// This property is optional but immutable.
	Image string

	// Runtime specifies which runtime should be used when launching container
	// tasks.
	//
	// This property is required and immutable.
	// Runtime指定了启动容器task时使用哪种容器运行时
	Runtime RuntimeInfo

	// Spec should carry the the runtime specification used to implement the
	// container.
	//
	// This field is required but mutable.
	Spec *types.Any

	// SnapshotKey specifies the snapshot key to use for the container's root
	// filesystem. When starting a task from this container, a caller should
	// look up the mounts from the snapshot service and include those on the
	// task create request.
	// SnapshotKey指定了容器的rootfs使用的snapshot key
	// 当从该容器启动一个task时，caller应该从snapshot service中寻找mounts并且将它们包含在
	// task create request中
	//
	// This field is not required but immutable.
	SnapshotKey string

	// Snapshotter specifies the snapshotter name used for rootfs
	//
	// This field is not required but immutable.
	// Snapshotter指定了rootfs使用的snapshotter name
	Snapshotter string

	// CreatedAt is the time at which the container was created.
	CreatedAt time.Time

	// UpdatedAt is the time at which the container was updated.
	UpdatedAt time.Time

	// Extensions stores client-specified metadata
	Extensions map[string]types.Any
}

// RuntimeInfo holds runtime specific information
type RuntimeInfo struct {
	Name    string
	Options *types.Any
}

// Store interacts with the underlying container storage
type Store interface {
	Get(ctx context.Context, id string) (Container, error)

	// List returns containers that match one or more of the provided filters.
	List(ctx context.Context, filters ...string) ([]Container, error)

	// Create a container in the store from the provided container.
	Create(ctx context.Context, container Container) (Container, error)

	// Update the container with the provided container object. ID must be set.
	//
	// If one or more fieldpaths are provided, only the field corresponding to
	// the fieldpaths will be mutated.
	Update(ctx context.Context, container Container, fieldpaths ...string) (Container, error)

	// Delete a container using the id.
	//
	// nil will be returned on success. If the container is not known to the
	// store, ErrNotFound will be returned.
	Delete(ctx context.Context, id string) error
}
