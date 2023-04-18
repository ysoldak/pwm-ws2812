package main

import "math"

var perceptionAdjustmentFactor = (255 * math.Log(2)) / math.Log(255)

// perception adjustment of brightness for human eye, see https://diarmuid.ie/blog/pwm-exponential-led-fading-on-arduino-or-other-platforms
func perception(actualBrightness uint8) uint8 {
	return uint8(math.Pow(2, float64(actualBrightness)/perceptionAdjustmentFactor) - 1)
}
