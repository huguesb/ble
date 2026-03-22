package dev

import (
	"golang.betakappaphi.com/ble"
)

// NewDevice ...
func NewDevice(impl string, opts ...ble.Option) (d ble.Device, err error) {
	return DefaultDevice(opts...)
}
