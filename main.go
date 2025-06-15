package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tbrandon/mbserver"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

var version = "v0.1.0"

var inputs = make([]Pin, 0)
var outputs = make([]Pin, 0)

func main() {
	slog.Info("Starting GPBUS", "version", version)

	argConfFile := flag.String("config", "", "config.json file")
	showVersion := flag.Bool("version", false, "Show version and exit")
	flag.Parse()

	if *showVersion {
		slog.Info("GPBUS version", "version", version)
		os.Exit(0)
	}

	if argConfFile != nil && *argConfFile != "" {
		slog.Info("Using config file:", *argConfFile)
		configFile = *argConfFile
	}

	conf, err := ReadConfig()
	if err != nil {
		slog.Error("Failed reading config: ", err)
		return
	}

	inputs = conf.Inputs
	outputs = conf.Outputs

	signalQuit := make(chan os.Signal, 2)
	signal.Notify(signalQuit, os.Interrupt, os.Kill, syscall.SIGTERM)

	modbus := mbserver.NewServer()
	if err = modbus.ListenTCP(fmt.Sprintf(":%d", conf.Port)); err != nil {
		slog.Error("Failed to start server", "err", err)
		return
	}
	defer modbus.Close()

	if _, err = host.Init(); err != nil {
		slog.Error("Failed to initialize gpio", "err", err)
		return
	}

	for i, pin := range inputs {
		gpioPin := gpioreg.ByName(pin.Name)
		if gpioPin == nil {
			slog.Error("Failed to find pin", "pin", pin.Name)
			return
		}

		inputs[i].gpioPin = gpioPin

		if err = gpioPin.In(gpio.PullNoChange, gpio.NoEdge); err != nil {
			slog.Error("Failed to set gpio input mode", "pin", pin.Name, "err", err)
			return
		}
	}

	for i, pin := range outputs {
		gpioPin := gpioreg.ByName(pin.Name)
		if gpioPin == nil {
			slog.Error("Failed to find pin", "pin", pin.Name)
			return
		}

		outputs[i].gpioPin = gpioPin

		vInit := gpioPin.Read()
		outputs[i].valueInflight = vInit

		vInitByte := byte(0)
		if vInit {
			vInitByte = byte(1)
		}
		modbus.Coils[pin.Register] = vInitByte

		if err = gpioPin.Out(vInit); err != nil {
			slog.Error("Failed to set gpio output mode", "pin", pin.Name, "err", err)
			return
		}
	}

	go func() {
		for {
			for i, pin := range inputs {
				val := pin.gpioPin.Read()
				if val == pin.valueInflight {
					continue
				}

				inputs[i].valueInflight = val

				vRead := byte(0)
				if val {
					vRead = byte(1)
				}

				modbus.DiscreteInputs[pin.Register] = vRead
			}

			for i, pin := range outputs {
				val := gpio.Level(modbus.Coils[pin.Register] > 0)
				if val != pin.valueInflight {
					outputs[i].valueInflight = val

					if err = pin.gpioPin.Out(val); err != nil {
						slog.Error("Failed to write to pin", "pin", pin.Name, "err", err)
						continue
					}
				}
			}

			time.Sleep(5 * time.Millisecond)
		}
	}()

	select {
	case <-signalQuit:
	}
	slog.Info("GPBUS Stopped Successfully")
}
