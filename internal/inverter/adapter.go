package inverter

import (
	"machine"

	"dischargeProtection/config"
)

func NewInverterAdapter() Inverter {
	return Inverter{}
}

type Inverter struct {
	inverter machine.Pin
}

func (b *Inverter) Init() {
	b.inverter = config.Parameters.Pins.Inverter
	b.inverter.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

func (b *Inverter) SwitchOn() {
	b.inverter.High()
}

func (b *Inverter) SwitchOff() {
	b.inverter.Low()
}
