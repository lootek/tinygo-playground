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

	//time.Sleep(time.Second * 5)
	//println("start")

	for {
		led.High()
		time.Sleep(time.Millisecond * 500)
		led.Low()

		trigger.High()
		time.Sleep(time.Microsecond * 100)
		trigger.Low()

		eDur := 0
		cnt := 0
		for {
			cnt++
			if cnt > 50*1000 {
				break
			}

			if !echo.Get() {
				if eDur == 0 {
					continue
				}

				break
			}

			eDur++
			time.Sleep(time.Microsecond)
		}

		eDur = eDur / 58
		println(eDur)

		for i := 0; i < eDur/10; i++ {
			led.High()
			time.Sleep(time.Millisecond * 200)
			led.Low()
			time.Sleep(time.Millisecond * 200)
		}

		time.Sleep(time.Second * 5)
	}
}
