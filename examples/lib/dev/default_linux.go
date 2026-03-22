package dev

import (
	"golang.betakappaphi.com/ble"
	"golang.betakappaphi.com/ble/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}
