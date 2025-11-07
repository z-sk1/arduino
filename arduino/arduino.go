package arduino

import (
	"fmt"

	"github.com/tarm/serial"
)

type Device struct {
	Port *serial.Port
}

func New(port *serial.Port) *Device {
	return &Device{Port: port}
}

func (d *Device) Close() error {
	return d.Port.Close()
}

func (d *Device) Exec(cmd string) error {
	cmd = cmd + "\n"
	_, err := d.Port.Write([]byte(cmd))
	if err != nil {
		return fmt.Errorf("failed to send command: %w", err)
	}
	return nil
}
