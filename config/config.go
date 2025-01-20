package config

import (
	"machine"
	"time"
)

type configuration struct {
	Pins                pinConfig
	DischargeProtection dischargeProtectionConfig
	Adc                 adcConfig
}

type pinConfig struct {
	BatVoltage             machine.Pin
	ButtonInverterOn       machine.Pin
	ButtonInverterOff      machine.Pin
	Inverter               machine.Pin
	LedDischargeProtection machine.Pin
	LedLiveness            machine.Pin
}

type dischargeProtectionConfig struct {
	BatteryLowVoltage  float32
	BatteryHighVoltage float32
	BatteryLowDuration time.Duration
}

type adcConfig struct {
	ReferenceVoltage     float32
	VoltageDividerFactor float32
}

var Parameters = configuration{
	Pins: pinConfig{
		BatVoltage:             machine.A1,
		ButtonInverterOn:       machine.D2,
		ButtonInverterOff:      machine.D3,
		Inverter:               machine.D4,
		LedDischargeProtection: machine.D5,
		LedLiveness:            machine.LED,
	},
	// TODO optimize these Parameters
	DischargeProtection: dischargeProtectionConfig{
		BatteryLowVoltage:  11.0,
		BatteryHighVoltage: 12.0,
		BatteryLowDuration: 10 * time.Second,
	},
	Adc: adcConfig{
		ReferenceVoltage:     3.3,
		VoltageDividerFactor: 5.0,
	},
}
