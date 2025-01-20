# Discharge protection
This software runs on a microcontroller as a discharge protection for 12V batteries.
It can control an inverter based on buttons (manual interaction) or depending on the 12V battery voltage 
(automatic discharge protection).

## Functionality
- An inverter can be switched on or off with two (turn on / turn off) buttons.
- The status of the inverter is shown by a LED.  
- If the 12V Battery is to too low for a too long period of time, the inverter get switched off automatically 
(-> discharge protection).
- The status of the discharge protection is shown by a LED. 

## Installation
1. Install the [tinygo-plugin](https://plugins.jetbrains.com/plugin/16915-tinygo) for the Goland IDE.
2. If still necessary, install tinygo
3. Run `tinygo flash -target=arduino-nano33 -monitor` to flash the e.g. an ARDUINO NANO 33 IOT