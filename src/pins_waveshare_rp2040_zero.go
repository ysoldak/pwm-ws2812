//go:build waveshare_rp2040_zero

package main

import "machine"

var (
	brightnessPin = machine.GPIO29
	huePin        = machine.GPIO28
	neoPin        = machine.GPIO27
)
