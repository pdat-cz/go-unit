// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

// TemperatureUnit represents a unit of temperature
type TemperatureUnit struct {
	BaseUnit
}

// Temperature contains predefined temperature units
var Temperature = struct {
	Celsius    TemperatureUnit
	Fahrenheit TemperatureUnit
	Kelvin     TemperatureUnit
}{
	Celsius: TemperatureUnit{
		BaseUnit: NewBaseUnit(
			"temperature",
			"°C",
			"Celsius",
			1.0,
			0.0,
			true, // Base unit
		),
	},
	Fahrenheit: TemperatureUnit{
		BaseUnit: NewBaseUnit(
			"temperature",
			"°F",
			"Fahrenheit",
			5.0/9.0, // Conversion factor: (F - 32) * 5/9 = C
			-32*5/9, // Offset adjustment
			false,
		),
	},
	Kelvin: TemperatureUnit{
		BaseUnit: NewBaseUnit(
			"temperature",
			"K",
			"Kelvin",
			1.0,
			-273.15, // Offset: K - 273.15 = C
			false,
		),
	},
}

// NewTemperature creates a new temperature quantity
func NewTemperature(value float64, unit TemperatureUnit) Quantity[TemperatureUnit] {
	return New(value, unit)
}
