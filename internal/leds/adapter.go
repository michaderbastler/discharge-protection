package leds

import (
	"machine"

	"discharge-protection/config"
)

func NewLedsAdapter() Leds {
	return Leds{}
}

type Leds struct {
	dischargeProtection machine.Pin
	liveness            machine.Pin
}

func (b *Leds) Init() {
	b.dischargeProtection = config.Parameters.Pins.LedDischargeProtection
	b.dischargeProtection.Configure(machine.PinConfig{Mode: machine.PinOutput})

	b.liveness = config.Parameters.Pins.LedLiveness
	b.liveness.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

func (b *Leds) SwitchLedDischargeProtectionOn() {
	b.dischargeProtection.High()
}

func (b *Leds) SwitchLedDischargeProtectionOff() {
	b.dischargeProtection.Low()
}

func (b *Leds) ToggleLedLiveness() {
	b.liveness.Set(!b.liveness.Get())
}
