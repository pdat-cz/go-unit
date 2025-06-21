// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

// FrequencyUnit represents a unit of frequency
type FrequencyUnit struct {
	BaseUnit
}

// Frequency contains predefined frequency units
var Frequency = struct {
	Hertz     FrequencyUnit
	Kilohertz FrequencyUnit
	Megahertz FrequencyUnit
	Gigahertz FrequencyUnit
	Terahertz FrequencyUnit
	RPM       FrequencyUnit // Revolutions per minute
}{
	Hertz: FrequencyUnit{
		BaseUnit: NewBaseUnit(
			"frequency",
			"Hz",
			"Hertz",
			1.0,
			0.0,
			true, // Base unit
		),
	},
	Kilohertz: FrequencyUnit{
		BaseUnit: NewBaseUnit(
			"frequency",
			"kHz",
			"Kilohertz",
			1000.0, // 1 kHz = 1000 Hz
			0.0,
			false,
		),
	},
	Megahertz: FrequencyUnit{
		BaseUnit: NewBaseUnit(
			"frequency",
			"MHz",
			"Megahertz",
			1000000.0, // 1 MHz = 1,000,000 Hz
			0.0,
			false,
		),
	},
	Gigahertz: FrequencyUnit{
		BaseUnit: NewBaseUnit(
			"frequency",
			"GHz",
			"Gigahertz",
			1000000000.0, // 1 GHz = 1,000,000,000 Hz
			0.0,
			false,
		),
	},
	Terahertz: FrequencyUnit{
		BaseUnit: NewBaseUnit(
			"frequency",
			"THz",
			"Terahertz",
			1000000000000.0, // 1 THz = 1,000,000,000,000 Hz
			0.0,
			false,
		),
	},
	RPM: FrequencyUnit{
		BaseUnit: NewBaseUnit(
			"frequency",
			"rpm",
			"Revolutions Per Minute",
			1.0/60.0, // 1 rpm = 1/60 Hz
			0.0,
			false,
		),
	},
}

// NewFrequency creates a new frequency measurement
func NewFrequency(value float64, unit FrequencyUnit) Quantity[FrequencyUnit] {
	return New(value, unit)
}
