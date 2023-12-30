package dev

import (
	"github.com/huguesb/ble"
	"github.com/huguesb/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return darwin.NewDevice(opts...)
}
