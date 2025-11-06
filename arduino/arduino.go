package arduino

import (
	"fmt"
	"github.com/tarm/serial"
)

type Device struct {
	port *serial.Port
}

func New(port *serial.Port) *Device {
	return &Device{port: port}
}

func (d *Device) Exec(cmd string) error {
	cmd = cmd + "\n"
	_, err := d.Port.Write([]byte(cmd))
	if err != nil {
		return fmt.Errorf("failed to send command: %w", err)
	}
	return nil
}