//go:build xiao_rp2040 || xiao

package main

import "machine"

var (
	brightnessPin = machine.D10
	huePin        = machine.D9
	neoPin        = machine.D8
)
