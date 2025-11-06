package main

import (
	"github.com/getlantern/systray"
	"log"
	"fmt"
	"github.com/z-sk1/arduino/arduino"
	"os"
	"embed"
)

func setupTray() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("Arduino Control")
	systray.SetTooltip("Control your Arduino from here!")

	// menu items 
	mTurnLEDOn := systray.AddMenuItem("Turn on LED", "Turn on the LED")
	mTurnLEDOff := systray.AddMenuItem("Turn off LED", "Turn off the LED")

	mQuit := systray.AddMenuItem("Quit", "Stop Controlling Arduino and Exit")

	go func() {
		for {
			select {
			case <-mTurnLEDOn:
				if err := Device.Exec("turnLedOn"); err != nil {
					fmt.Println("Error:", err)
				}
			case <-mTurnLEDOff:
				if err := Device.Exec("turnLedOff"); err != nil {
					fmt.Println("Error:", err)
				}
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {

}