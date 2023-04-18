//go:build pico

package main

import "machine"

var (
	brightnessPin = machine.GPIO28
	huePin        = machine.GPIO27
	neoPin        = machine.GPIO26
)
