package main

import (
	"time"

	"discharge-protection/internal/battery"
	"discharge-protection/internal/buttons"
	"discharge-protection/internal/dischargeprotection"
	"discharge-protection/internal/inverter"
	"discharge-protection/internal/leds"
)

func main() {
	// Battery
	batteryAdapter := battery.NewBatteryAdapter()
	batteryAdapter.Init()

	// Discharge protection
	dischargeProtection := dischargeprotection.New()

	// Inverter
	inverterAdapter := inverter.NewInverterAdapter()
	inverterAdapter.Init()

	// Buttons
	callbackButtonInverterOn := func() {
		dischargeProtection.ResetDischargeProtection()
		inverterAdapter.SwitchOn()
	}
	callbackButtonInverterOff := func() {
		inverterAdapter.SwitchOff()
	}
	buttonsAdapter := buttons.NewButtonsAdapter(callbackButtonInverterOn, callbackButtonInverterOff)
	buttonsAdapter.Init()

	// LEDs
	ledsAdapter := leds.NewLedsAdapter()
	ledsAdapter.Init()

	for {
		batVolt := batteryAdapter.GetVoltage()

		dischargeProtectionActive := dischargeProtection.IsDischargeProtectionActive(batVolt)

		if dischargeProtectionActive {
			inverterAdapter.SwitchOff()
			ledsAdapter.SwitchLedDischargeProtectionOn()
		} else {
			ledsAdapter.SwitchLedDischargeProtectionOff()
		}

		ledsAdapter.ToggleLedLiveness()
		time.Sleep(1 * time.Second)
	}
}
