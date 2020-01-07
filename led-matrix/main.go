package main

import (
	"machine"
	"math/rand"
	"time"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			println(r)
		}
	}()

	uart := machine.UART2
	uart.Configure(machine.UARTConfig{
		TX: machine.UART_TX_PIN,
		RX: machine.UART_RX_PIN,
	})

	spi := machine.SPI0
	spi.Configure(machine.SPIConfig{
		SCK:  machine.SCL_PIN,
		MOSI: machine.D11,
	})

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	time.Sleep(time.Second * 5)
	println("slept")

	for {
		led.High()

		for i := 0; i < 64; i++ {
			val := rand.Intn(256)
			_, err := spi.Transfer(byte(val))
			if err != nil {
				println(err)
			}
		}
		_, err := spi.Transfer(byte(12))
		if err != nil {
			println(err)
		}

		time.Sleep(time.Millisecond * 150)
		led.Low()
		time.Sleep(time.Millisecond * 150)
	}
}
