package plugins

import (
	"github.com/docker/docker/plugins"
	"github.com/docker/docker/volume"
)

func init() {
	plugins.Handle("VolumeDriver", func(name string, client *plugins.Client) {
		proxy := &volumeDriverProxy{name, client}
		adapter := &volumeDriverAdapter{proxy}
		volume.Drivers.Register(adapter, name)
	})
}

type VolumeDriver interface {
	Create(name string) (err error)
	Remove(name string) (err error)
	Mount(name string) (mountpoint string, err error)
	Unmount(name string) (err error)
}
