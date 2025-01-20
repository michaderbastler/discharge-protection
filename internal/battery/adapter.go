package battery

import (
	"machine"

	"dischargeProtection/config"
)

func NewBatteryAdapter() Battery {
	return Battery{}
}

type Battery struct {
	adc machine.ADC
}

func (b *Battery) Init() {
	machine.InitADC()
	b.adc = machine.ADC{Pin: config.Parameters.Pins.BatVoltage}
	b.adc.Configure(machine.ADCConfig{})
}

func (b *Battery) GetVoltage() float32 {
	adcRawValue := b.adc.Get()
	voltage := float32(adcRawValue) * config.Parameters.Adc.ReferenceVoltage * config.Parameters.Adc.VoltageDividerFactor / 65535
	return voltage
}

func (b *Battery) GetVoltageAndRawValue() (float32, uint16) {
	adcRawValue := b.adc.Get()
	voltage := float32(adcRawValue) * config.Parameters.Adc.ReferenceVoltage * config.Parameters.Adc.VoltageDividerFactor / 65535
	return voltage, adcRawValue
}
