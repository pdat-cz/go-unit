// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

// PowerUnit represents a unit of power
type PowerUnit struct {
	BaseUnit
}

// Power contains predefined power units
var Power = struct {
	Watt       PowerUnit
	Kilowatt   PowerUnit
	BTUPerHour PowerUnit
}{
	Watt: PowerUnit{
		BaseUnit: NewBaseUnit(
			"power",
			"W",
			"Watt",
			1.0,
			0.0,
			true, // Base unit
		),
	},
	Kilowatt: PowerUnit{
		BaseUnit: NewBaseUnit(
			"power",
			"kW",
			"Kilowatt",
			1000.0, // 1 kW = 1000 W
			0.0,
			false,
		),
	},
	BTUPerHour: PowerUnit{
		BaseUnit: NewBaseUnit(
			"power",
			"BTU/h",
			"British Thermal Unit per Hour",
			0.29307107, // 1 BTU/h = 0.29307107 W
			0.0,
			false,
		),
	},
}

// NewPower creates a new power quantity
func NewPower(value float64, unit PowerUnit) Quantity[PowerUnit] {
	return New(value, unit)
}
