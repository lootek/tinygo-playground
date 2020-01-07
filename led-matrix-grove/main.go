package main

import (
	"machine"
	"time"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			println(r)
		}
	}()

	uart := machine.UART1
	uart.Configure(machine.UARTConfig{
		TX: machine.UART_TX_PIN,
		RX: machine.UART_RX_PIN,
	})

	i2c := machine.I2C0
	i2c.Configure(machine.I2CConfig{
		SCL: machine.SCL_PIN,
		SDA: machine.SDA_PIN,
	})

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	time.Sleep(time.Second * 10)
	println("slept")

	for {
		led.High()

		i2c.Tx()
		if err != nil {
			println(err)
		}

		time.Sleep(time.Millisecond * 150)
		led.Low()
		time.Sleep(time.Millisecond * 150)
	}
}
