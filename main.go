package main

import (
	"fmt"
	"log"

	"os"

	"github.com/ncruces/zenity"
	"github.com/tarm/serial"
	"github.com/z-sk1/arduino/arduino"
)

var Device *arduino.Device

func main() {
	portName := askForPort()
	openPort(portName)

	setupTray()

	go func() {
		buf := make([]byte, 128)
		for {
			n, err := Device.Port.Read(buf)
			if err != nil {
				fmt.Println("Read error:", err)
				return
			}
			if n > 0 {
				fmt.Print(string(buf[:n]))
				// Force flush immediately
				os.Stdout.Sync()
			}
		}
	}()
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
