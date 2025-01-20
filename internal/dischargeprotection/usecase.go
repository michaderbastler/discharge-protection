package dischargeprotection

import (
	"fmt"
	"time"

	"discharge-protection/config"
)

func New() DischargeProtection {
	return DischargeProtection{}
}

type DischargeProtection struct {
	dischargeProtectionActive bool
	batteryLow                bool
	timeBatteryLow            time.Time
}

func (d *DischargeProtection) IsDischargeProtectionActive(batteryVoltage float32) bool {
	if d.dischargeProtectionActive {
		fmt.Printf("%.3f -> DischargeProtection\n", batteryVoltage)
		return true
	}

	if batteryVoltage >= config.Parameters.DischargeProtection.BatteryHighVoltage {
		d.batteryLow = false
		fmt.Printf("%.3f -> Battery voltage high\n", batteryVoltage)
		return false
	}

	if batteryVoltage > config.Parameters.DischargeProtection.BatteryLowVoltage {
		fmt.Printf("%.3f -> Battery voltage between high and low\n", batteryVoltage)
	}

	if batteryVoltage <= config.Parameters.DischargeProtection.BatteryLowVoltage {
		fmt.Printf("%.3f -> Battery voltage low\n", batteryVoltage)
		if !d.batteryLow {
			d.batteryLow = true
			d.timeBatteryLow = time.Now()
			fmt.Printf("Staring timer at at %s\n", d.timeBatteryLow.String())
			return false
		}
	}

	if d.batteryLow {
		batteryLowDuration := time.Now().Sub(d.timeBatteryLow)
		fmt.Printf("%.3f -> Timer active since %s\n", batteryVoltage, batteryLowDuration.String())

		if batteryLowDuration >= config.Parameters.DischargeProtection.BatteryLowDuration {
			d.dischargeProtectionActive = true
			fmt.Printf("%.3f -> Activate DischargeProtection\n", batteryVoltage)
			return true
		}
	}

	return false
}

func (d *DischargeProtection) ResetDischargeProtection() {
	d.dischargeProtectionActive = false
	d.batteryLow = false
}
