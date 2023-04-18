package main

import "time"

// for lab tests to simulate sliders
func test() {
	go func() {
		inc := int64(10)
		for {
			bms += inc
			if bms < 1000 || bms > 2000 {
				inc *= -1
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()
}
