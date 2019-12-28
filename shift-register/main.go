package main

import (
	"machine"
	"math/rand"
	"time"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	sleepTime := time.Millisecond * 500

	data := machine.D2
	data.Configure(machine.PinConfig{Mode: machine.PinOutput})

	latch := machine.D3
	latch.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		time.Sleep(sleepTime)
		led.High()

		b := rand.Intn(10) >= 5
		data.Set(b)

		latch.High()
		time.Sleep(sleepTime)
		latch.Low()

		led.Low()
	}
}
