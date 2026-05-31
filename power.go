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
	Milliwatt  PowerUnit
	Kilowatt   PowerUnit
	Megawatt   PowerUnit
	Horsepower PowerUnit
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
	Milliwatt: PowerUnit{
		BaseUnit: NewBaseUnit(
			"power",
			"mW",
			"Milliwatt",
			1e-3,
			0.0,
			false,
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
	Megawatt: PowerUnit{
		BaseUnit: NewBaseUnit(
			"power",
			"MW",
			"Megawatt",
			1e6,
			0.0,
			false,
		),
	},
	Horsepower: PowerUnit{
		BaseUnit: NewBaseUnit(
			"power",
			"hp",
			"Horsepower",
			745.6998715822702, // mechanical horsepower
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
