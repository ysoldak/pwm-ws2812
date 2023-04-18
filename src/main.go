// Connects to an WS2812 RGB LED.
package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

const (
	ledCount         = 10
	adjustPerception = false
)

var ws ws2812.Device

var leds = [ledCount]color.RGBA{}

var (
	bts = int64(0)
	bms = int64(2000)
	hts = int64(0)
	hms = int64(0)
)

// -----------------------------------------------------------------------------

func setup() {
	var err error

	// input (signal, from receiver)
	brightnessPin.Configure(machine.PinConfig{Mode: machine.PinInputPullup}) // full brightness by default
	huePin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})      // green color by default

	// read brightness
	err = brightnessPin.SetInterrupt(machine.PinRising|machine.PinFalling, func(machine.Pin) {
		if brightnessPin.Get() {
			bts = time.Now().UnixMicro()
		} else {
			bms = time.Now().UnixMicro() - bts
		}
	})
	if err != nil {
		println("could not configure brightness pin interrupt:", err.Error())
	}

	// read hue
	err = huePin.SetInterrupt(machine.PinRising|machine.PinFalling, func(machine.Pin) {
		if huePin.Get() {
			hts = time.Now().UnixMicro()
		} else {
			hms = time.Now().UnixMicro() - hts
		}
	})
	if err != nil {
		println("could not configure hue pin interrupt:", err.Error())
	}

	// output (signal to ws2812 led strip)
	neoPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ws = ws2812.New(neoPin)
	for i := range leds {
		leds[i] = color.RGBA{R: 0x00, G: 0x00, B: 0x00}
	}
	ws.WriteColors(leds[:])

}

func main() {

	// setup everything
	setup()

	// test
	// test()

	// main loop
	var hue, brightness uint8
	var red, green, blue uint8
	for {
		hue = scale(hms)
		brightness = scale(bms)
		red, green, blue = colors(hue)
		red, green, blue = dim(red, brightness), dim(green, brightness), dim(blue, brightness)
		for i := range leds {
			leds[i].R = red
			leds[i].G = green
			leds[i].B = blue
		}
		ws.WriteColors(leds[:])
		println(hue, brightness, red, green, blue)
		time.Sleep(10 * time.Millisecond)
	}
}

// -----------------------------------------------------------------------------

// scale (0:255) from RC PWM (1000-2000us, 50Hz)
func scale(ms int64) uint8 {
	if ms < 1000 {
		return 0
	}
	if ms > 2000 {
		return 254
	}
	return uint8(255 * (ms - 1000) / (2000 - 1000))
}

// colors evenly spread on G-R-B slider, every moment sum of all colors equals 0x7F
func colors(hue uint8) (red, green, blue uint8) {
	green = 0x7F - hue
	if hue > 0x7F {
		green = 0
	}
	red = hue
	if red > 0x7F {
		red = 0xFF - hue
	}
	blue = hue - 0x7F
	if hue < 0x7F {
		blue = 0
	}
	return
}

// dim led according to brightness level
func dim(value uint8, brightness uint8) (result uint8) {
	result = uint8((int(value) * int(brightness)) / 0x7F)
	if adjustPerception {
		result = perception(result)
	}
	return
}
