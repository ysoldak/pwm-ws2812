//go:build nano_rp2040

package main

import "machine"

var (
	brightnessPin = machine.D19
	huePin        = machine.D18
	neoPin        = machine.D17
)
