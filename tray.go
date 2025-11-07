package main

import (
	_ "embed"

	"github.com/getlantern/systray"
)

//go:embed icon.ico
var iconData []byte

func setupTray() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("Arduino Control")
	systray.SetTooltip("Control your Arduino from here!")

	systray.SetIcon(iconData)
	// menu items
	mToggleLED := systray.AddMenuItem("Turn on LED", "Turn on the LED")
	mToggleFan := systray.AddMenuItem("Turn on Fan", "Turn on the Fan")

	mQuit := systray.AddMenuItem("Quit", "Stop Controlling Arduino and Exit")

	var (
		ledOn = false
		fanOn = false
	)

	go func() {
		for {
			select {
			case <-mToggleLED.ClickedCh:
				if ledOn {
					Device.Exec("turnLedOff")
					mToggleLED.SetTitle("Turn on LED")
					ledOn = false
				} else {
					Device.Exec("turnLedOn")
					mToggleLED.SetTitle("Turn off LED")
					ledOn = true
				}

			case <-mToggleFan.ClickedCh:
				if fanOn {
					Device.Exec("turnFanOff")
					mToggleFan.SetTitle("Turn on Fan")
					fanOn = false
				} else {
					Device.Exec("turnFanOn")
					mToggleFan.SetTitle("Turn off Fan")
					fanOn = true
				}

			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	if Device != nil {
		Device.Close()
	}
}
