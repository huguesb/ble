package dev

import (
	"github.com/huguesb/ble"
	"github.com/huguesb/ble/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}
