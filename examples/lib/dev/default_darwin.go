package dev

import (
	"golang.betakappaphi.com/ble"
	"golang.betakappaphi.com/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return darwin.NewDevice(opts...)
}
