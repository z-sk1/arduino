package main

import (
	_ "embed"

	"log"

	"strconv"

	"fmt"

	"github.com/getlantern/systray"
	"github.com/ncruces/zenity"
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
	sLights := systray.AddMenuItem("Lights", "Lights sections")
	mToggleBlueLED := sLights.AddSubMenuItem("Turn on Blue LED", "Turn on the Blue LED")
	mToggleGreenLED := sLights.AddSubMenuItem("Turn on Green LED", "Turn on the Green LED")
	mToggleRedLED := sLights.AddSubMenuItem("Turn on Red LED", "Turn on the Red LED")
	mToggleYellowLED := sLights.AddSubMenuItem("Turn on Yellow LED", "Turn on the Yellow LED")
	mLightShow := sLights.AddSubMenuItem("Turn on RGB Light Show", "Turn on an RGB Light Show using the LEDS")
	mLightControlJoystick := sLights.AddSubMenuItem("Control the light with Joystick", "Control your light using a joystick")

	sBuzz := systray.AddMenuItem("Buzzers", "Buzzer section")

	mToggleBuzz := sBuzz.AddSubMenuItem("Turn on Buzz", "Turn on the Buzz")
	mToggleMelody := sBuzz.AddSubMenuItem("Play a Melody!", "Play a simple note sequence")
	mToggleSiren := sBuzz.AddSubMenuItem("Play a Siren", "Play a Siren sound!")
	mToggleMegalovania := sBuzz.AddSubMenuItem("Play MEGALOVANIA", "PLAY MEGALOVANIA :3")
	mBuzzControlJoystick := sBuzz.AddSubMenuItem("Control Buzzer using Joystick", "Control the Buzzer using a Joystick")

	sPortal := sBuzz.AddSubMenuItem("Portal Themes", "Portal Themes Section for Buzz")
	mTogglePortalThemeMain := sPortal.AddSubMenuItem("Play Main Theme", "Play the Main Theme of Portal 2")
	mToggleApertureTheme := sPortal.AddSubMenuItem("Play Aperture Science Theme", "Play the Theme of Aperture Science found in Portal 2")

	sMario := sBuzz.AddSubMenuItem("Mario Themes", "Mario Themes Sections for Buzz")
	mToggleOverworldTheme := sMario.AddSubMenuItem("Play Super Mario Bros. Overworld Theme", "Play the iconic 1-1 overworld theme in the og Mario")
	mToggleUndergroundTheme := sMario.AddSubMenuItem("Play Super Mario Bros. Underground Theme", "Play the iconic 1-2 underground them in the og Mario")

	sServo := systray.AddMenuItem("Servos", "Servos section")
	mRotate90 := sServo.AddSubMenuItem("Rotate 90 Degrees", "Rotate the servo by 90 Degrees")
	mRotateNeg90 := sServo.AddSubMenuItem("Rotate Negative 90 Degrees", "Rotate the servo by Negative 90 Degrees")
	mToggleSpin := sServo.AddSubMenuItem("Spin the Servo", "Keep on Spinning the Servo")
	mRotatePrecise := sServo.AddSubMenuItem("Enter a Precise Angle to Rotate", "Enter an exact degree to spin the servo")
	mControlJoystick := sServo.AddSubMenuItem("Control the Servo with Joystick", "Control your Servo using a Joystick on the X-Axis")

	systray.AddSeparator()

	mQuit := systray.AddMenuItem("Quit", "Stop Controlling Arduino and Exit")

	var (
		blueLedOn          = false
		greenLedOn         = false
		redLedOn           = false
		yellowLedOn        = false
		buzzOn             = false
		rgbShowOn          = false
		melodyOn           = false
		sirenOn            = false
		megalovaniaOn      = false
		portalThemeMainOn  = false
		apertureThemeOn    = false
		overworldThemeOn   = false
		undergroundThemeOn = false
		servoSpinOn        = false
		servoJoystickOn    = false
		lightsJoystickOn   = false
		buzzerJoystickOn   = false
	)

	go func() {
		for {
			select {
			case <-mToggleBlueLED.ClickedCh:
				if blueLedOn {
					Device.Exec("turnBlueLedOff")
					mToggleBlueLED.SetTitle("Turn on Blue LED")
					blueLedOn = false
				} else {
					Device.Exec("turnBlueLedOn")
					mToggleBlueLED.SetTitle("Turn off Blue LED")
					blueLedOn = true
				}

			case <-mToggleGreenLED.ClickedCh:
				if greenLedOn {
					Device.Exec("turnGreenLedOff")
					mToggleGreenLED.SetTitle("Turn on Green LED")
					greenLedOn = false
				} else {
					Device.Exec("turnGreenLedOn")
					mToggleGreenLED.SetTitle("Turn off Green LED")
					greenLedOn = true
				}

			case <-mToggleRedLED.ClickedCh:
				if redLedOn {
					Device.Exec("turnRedLedOff")
					mToggleRedLED.SetTitle("Turn on Red LED")
					redLedOn = false
				} else {
					Device.Exec("turnRedLedOn")
					mToggleRedLED.SetTitle("Turn off Red LED")
					redLedOn = true
				}

			case <-mToggleYellowLED.ClickedCh:
				if yellowLedOn {
					Device.Exec("turnYellowLedOff")
					mToggleYellowLED.SetTitle("Turn on Yellow LED")
					yellowLedOn = false
				} else {
					Device.Exec("turnYellowLedOn")
					mToggleYellowLED.SetTitle("Turn off Yellow LED")
					yellowLedOn = true
				}

			case <-mToggleBuzz.ClickedCh:
				if buzzOn {
					Device.Exec("turnBuzzOff")
					mToggleBuzz.SetTitle("Turn on Buzz")
					buzzOn = false
				} else {
					Device.Exec("turnBuzzOn")
					mToggleBuzz.SetTitle("Turn off Buzz")
					buzzOn = true
				}

			case <-mLightShow.ClickedCh:
				if rgbShowOn {
					Device.Exec("rgbShowOff")
					mLightShow.SetTitle("Turn on RGB Light Show")
					rgbShowOn = false
				} else {
					Device.Exec("rgbShowOn")
					mLightShow.SetTitle("Turn off RGB Light Show")
					rgbShowOn = true
				}

			case <-mToggleMelody.ClickedCh:
				if melodyOn {
					Device.Exec("melodyOff")
					mToggleMelody.SetTitle("Play a Melody!")
					melodyOn = false
				} else {
					Device.Exec("melodyOn")
					mToggleMelody.SetTitle("Stop playing Melody")
					melodyOn = true
				}

			case <-mToggleSiren.ClickedCh:
				if sirenOn {
					Device.Exec("sirenOff")
					mToggleSiren.SetTitle("Play a Siren")
					sirenOn = false
				} else {
					Device.Exec("sirenOn")
					mToggleSiren.SetTitle("Stop playing Siren")
					sirenOn = true
				}

			case <-mToggleMegalovania.ClickedCh:
				if megalovaniaOn {
					Device.Exec("megalovaniaOff")
					mToggleMegalovania.SetTitle("Play MEGALOVANIA")
					megalovaniaOn = false
				} else {
					Device.Exec("megalovaniaOn")
					mToggleMegalovania.SetTitle("Stop Playing MEGALOVANIA")
					megalovaniaOn = true
				}

			case <-mTogglePortalThemeMain.ClickedCh:
				if portalThemeMainOn {
					Device.Exec("portalMainThemeOff")
					mTogglePortalThemeMain.SetTitle("Play Main Theme")
					portalThemeMainOn = false
				} else {
					Device.Exec("portalMainThemeOn")
					mTogglePortalThemeMain.SetTitle("Stop Playing Main Theme")
					portalThemeMainOn = true
				}

			case <-mToggleApertureTheme.ClickedCh:
				if apertureThemeOn {
					Device.Exec("apertureThemeOff")
					mToggleApertureTheme.SetTitle("Play Aperture Theme")
					apertureThemeOn = false
				} else {
					Device.Exec("apertureThemeOn")
					mToggleApertureTheme.SetTitle("Stop Playing Aperture Theme")
					apertureThemeOn = true
				}

			case <-mToggleOverworldTheme.ClickedCh:
				if overworldThemeOn {
					Device.Exec("overworldThemeOff")
					mToggleOverworldTheme.SetTitle("Play Super Mario Bros. Overworld Theme")
					overworldThemeOn = false
				} else {
					Device.Exec("overworldThemeOn")
					mToggleOverworldTheme.SetTitle("Stop playing Super Mario Bros. Overworld Theme")
					overworldThemeOn = true
				}

			case <-mToggleUndergroundTheme.ClickedCh:
				if undergroundThemeOn {
					Device.Exec("undergroundThemeOff")
					mToggleUndergroundTheme.SetTitle("Play Super Mario Bros. Underground Theme")
					undergroundThemeOn = false
				} else {
					Device.Exec("undergroundThemeOn")
					mToggleUndergroundTheme.SetTitle("Stop playing Super Mario Bros. Underground Theme")
					undergroundThemeOn = true
				}

			case <-mRotate90.ClickedCh:
				Device.Exec("rotateServo90")

			case <-mRotateNeg90.ClickedCh:
				Device.Exec("rotateServo-90")

			case <-mToggleSpin.ClickedCh:
				if servoSpinOn {
					Device.Exec("servoSpinOff")
					mToggleSpin.SetTitle("Spin the Servo")
					servoSpinOn = false
				} else {
					Device.Exec("servoSpinOn")
					mToggleSpin.SetTitle("Stop spinning the Servo")
					servoSpinOn = true
				}

			case <-mRotatePrecise.ClickedCh:
				deg, err := zenity.Entry("Enter an angle for the servo to rotate: (0-180)")
				if err != nil {
					log.Fatal(err)
				}

				intDeg, err := strconv.Atoi(deg)
				if err != nil {
					log.Println(err)
					return
				}

				if intDeg >= 180 {
					deg = "180"
				}

				if intDeg <= 0 {
					deg = "0"
				}

				deg = deg + "\n"

				Device.Exec(fmt.Sprintf("rotatePrecise %s", deg))

			case <-mControlJoystick.ClickedCh:
				if servoJoystickOn {
					Device.Exec("servoJoyControlOff")
					mControlJoystick.SetTitle("Control Servo with Joystick")
					servoJoystickOn = false
				} else {
					Device.Exec("servoJoyControlOn")
					mControlJoystick.SetTitle("Stop Controlling Servo with Joystick")
					servoJoystickOn = true
				}

			case <-mLightControlJoystick.ClickedCh:
				if lightsJoystickOn {
					Device.Exec("lightJoyControlOff")
					mLightControlJoystick.SetTitle("Control Lights with Joystick")
					lightsJoystickOn = false
				} else {
					Device.Exec("lightJoyControlOn")
					mLightControlJoystick.SetTitle("Stop Controlling Lights with Joystick")
					lightsJoystickOn = true
				}

			case <-mBuzzControlJoystick.ClickedCh:
				if buzzerJoystickOn {
					Device.Exec("buzzerJoyControlOff")
					mBuzzControlJoystick.SetTitle("Control Buzzer with Joystick")
					buzzerJoystickOn = false
				} else {
					Device.Exec("buzzerJoyControlOn")
					mBuzzControlJoystick.SetTitle("Stop Controlling Buzzer with Joystick")
					buzzerJoystickOn = true
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
