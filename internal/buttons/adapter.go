package buttons

import (
	"fmt"
	"machine"
	"time"

	"discharge-protection/config"
)

func NewButtonsAdapter(callbackButtonInverterOn func(), callbackButtonInverterOff func()) Buttons {
	return Buttons{
		callbackButtonInverterOn:  callbackButtonInverterOn,
		callbackButtonInverterOff: callbackButtonInverterOff,
	}
}

const debounceTime = 500 * time.Millisecond

type Buttons struct {
	buttonInverterOn             machine.Pin
	buttonInverterOff            machine.Pin
	timeButtonInverterOnPressed  time.Time // For debouncing
	timeButtonInverterOffPressed time.Time // For debouncing
	callbackButtonInverterOn     func()
	callbackButtonInverterOff    func()
}

func (b *Buttons) Init() {
	b.buttonInverterOn = config.Parameters.Pins.ButtonInverterOn
	b.buttonInverterOn.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	err := b.buttonInverterOn.SetInterrupt(machine.PinFalling, func(p machine.Pin) {
		now := time.Now()
		if now.Sub(b.timeButtonInverterOnPressed) > debounceTime {
			b.timeButtonInverterOnPressed = now
			fmt.Println("buttonInverterOn pressed")
			b.callbackButtonInverterOn()
		}
	})
	if err != nil {
		println(err.Error())
		return
	}

	b.buttonInverterOff = config.Parameters.Pins.ButtonInverterOff
	b.buttonInverterOff.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	err = b.buttonInverterOff.SetInterrupt(machine.PinFalling, func(p machine.Pin) {
		now := time.Now()
		if now.Sub(b.timeButtonInverterOffPressed) > debounceTime {
			b.timeButtonInverterOffPressed = now
			fmt.Println("buttonInverterOff pressed")
			b.callbackButtonInverterOff()
		}
	})
	if err != nil {
		println(err.Error())
		return
	}

}
