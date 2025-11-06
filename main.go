package main

import (
	"log"
	"github.com/z-sk1/arduino/arduino"
	"github.com/tarm/serial"
	"github.com/ncruses/zenity"
)

var Device *arduino.Device

func main() {
	portName := askForPort()
	openPort(portName)

	go setupTray()
}

func askForPort() string {
	port, err := zenity.Entry("Enter COM Port: (e.g COM3, COM5)")
	if err != nil {
		log.Fatal(err)
	}
	return port
}

func openPort(comPort string) {
	c := &serial.Config{Name: comPort, Baud: 9600}
	port, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	Device = arduino.New(port)
}