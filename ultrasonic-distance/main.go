package main

import (
	"machine"
	"time"
)

func main() {
	machine.UART1.Configure(machine.UARTConfig{TX: machine.PA22, RX: machine.PA23})

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	trigger := machine.D4
	trigger.Configure(machine.PinConfig{Mode: machine.PinInput})

	echo := machine.D3
	echo.Configure(machine.PinConfig{Mode: machine.PinOutput})

	out := machine.D2
	out.Configure(machine.PinConfig{Mode: machine.PinOutput})

	time.Sleep(time.Second * 5)
	println("start")

	for {
		trigger.Configure(machine.PinConfig{Mode: machine.PinInput})
		trigger.High()
		time.Sleep(time.Microsecond * 10)
		trigger.Low()
		trigger.Configure(machine.PinConfig{Mode: machine.PinOutput})

		for {
			time.Sleep(time.Microsecond)

			if out.Get() {
				println("out")
			}
			if echo.Get() {
				println("echo")
			}
			if trigger.Get() {
				println("trigger")
			}
		}

		//eDur := 0
		//for {
		//	if !echo.Get() {
		//		if eDur == 0 {
		//			println("c")
		//			continue
		//		}
		//
		//		println("b")
		//		break
		//	}
		//
		//	println("++")
		//	eDur++
		//	time.Sleep(time.Microsecond)
		//}
		//
		////eDur = 5 * eDur / 58
		//println(eDur)

		//for i := 0; i < eDur; i++ {
		//	led.High()
		//	time.Sleep(time.Millisecond * 100)
		//	led.Low()
		//	time.Sleep(time.Millisecond * 100)
		//}

		time.Sleep(time.Second * 5)
		led.High()
		time.Sleep(time.Millisecond * 500)
		led.Low()
	}
}
