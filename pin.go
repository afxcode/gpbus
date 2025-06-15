package main

import "periph.io/x/conn/v3/gpio"

type Pin struct {
	Name          string `json:"name"`
	Register      uint16 `json:"register"`
	valueInflight gpio.Level
	gpioPin       gpio.PinIO
}
