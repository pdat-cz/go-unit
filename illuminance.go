// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

// IlluminanceUnit represents a unit of illuminance
type IlluminanceUnit struct {
	BaseUnit
}

// Illuminance contains predefined illuminance units
var Illuminance = struct {
	Lux        IlluminanceUnit
	FootCandle IlluminanceUnit
	Phot       IlluminanceUnit
	Nox        IlluminanceUnit
}{
	Lux: IlluminanceUnit{
		BaseUnit: NewBaseUnit(
			"illuminance",
			"lx",
			"Lux",
			1.0,
			0.0,
			true, // Base unit
		),
	},
	FootCandle: IlluminanceUnit{
		BaseUnit: NewBaseUnit(
			"illuminance",
			"fc",
			"Foot-candle",
			10.7639, // 1 fc = 10.7639 lx
			0.0,
			false,
		),
	},
	Phot: IlluminanceUnit{
		BaseUnit: NewBaseUnit(
			"illuminance",
			"ph",
			"Phot",
			10000.0, // 1 ph = 10,000 lx
			0.0,
			false,
		),
	},
	Nox: IlluminanceUnit{
		BaseUnit: NewBaseUnit(
			"illuminance",
			"nx",
			"Nox",
			0.001, // 1 nx = 0.001 lx
			0.0,
			false,
		),
	},
}

// NewIlluminance creates a new illuminance measurement
func NewIlluminance(value float64, unit IlluminanceUnit) Quantity[IlluminanceUnit] {
	return New(value, unit)
}
