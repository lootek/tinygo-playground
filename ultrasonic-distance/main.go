package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {
	machine.InitADC()
	machine.InitPWM()

	uart := machine.UART0
	uart.Configure(machine.UARTConfig{
		//BaudRate: 57600,
		//BaudRate: 9600,
		//TX: machine.UART_TX_PIN,
		//RX: machine.UART_RX_PIN,
		TX: machine.D6,
		RX: machine.D5,
	})

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	//var data []byte

	time.Sleep(time.Second * 5)
	println(".")

	for {
		led.High()

		b, err := uart.ReadByte()
		if err != nil {
			println(err)
		}

		print(fmt.Sprintf("%d", b))

		////if buf >= 4 {
		//n, err := uart.Read(data)
		//if err != nil {
		//	println(err)
		//}
		//
		//println(n)
		//for _, b := range data {
		//	print(fmt.Sprintf("%d", b))
		//}
		//println("")
		////}

		time.Sleep(time.Millisecond * 50)
		led.Low()
		time.Sleep(time.Millisecond * 50)

		//time.Sleep(time.Second)
	}

	//trigger := machine.D4
	//trigger.Configure(machine.PinConfig{Mode: machine.PinInput})
	//
	//echo := machine.D3
	//echo.Configure(machine.PinConfig{Mode: machine.PinOutput})

	//time.Sleep(time.Second * 5)
	//println("start")

	//for {
	//	led.High()
	//	time.Sleep(time.Millisecond * 500)
	//	led.Low()
	//
	//	trigger.High()
	//	time.Sleep(time.Microsecond * 10)
	//	trigger.Low()
	//
	//	eDur := 0
	//	cnt := 0
	//	for {
	//		cnt++
	//		if cnt > 50*1000 {
	//			break
	//		}
	//
	//		if !echo.Get() {
	//			if eDur == 0 {
	//				continue
	//			}
	//
	//			break
	//		}
	//
	//		eDur++
	//		time.Sleep(time.Microsecond)
	//	}
	//
	//	eDur = eDur / 58
	//	println(eDur)
	//
	//	for i := 0; i < eDur/10; i++ {
	//		led.High()
	//		time.Sleep(time.Millisecond * 200)
	//		led.Low()
	//		time.Sleep(time.Millisecond * 200)
	//	}
	//
	//	time.Sleep(time.Second * 5)
	//}
}
