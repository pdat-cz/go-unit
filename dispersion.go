// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

// DispersionUnit represents a unit of dispersion
type DispersionUnit struct {
	BaseUnit
}

// Dispersion contains predefined dispersion units
var Dispersion = struct {
	PartsPerMillion  DispersionUnit
	PartsPerBillion  DispersionUnit
	PartsPerTrillion DispersionUnit
	Percent          DispersionUnit
}{
	PartsPerMillion: DispersionUnit{
		BaseUnit: NewBaseUnit(
			"dispersion",
			"ppm",
			"Parts per Million",
			1.0,
			0.0,
			true, // Base unit
		),
	},
	PartsPerBillion: DispersionUnit{
		BaseUnit: NewBaseUnit(
			"dispersion",
			"ppb",
			"Parts per Billion",
			0.001, // 1 ppb = 0.001 ppm
			0.0,
			false,
		),
	},
	PartsPerTrillion: DispersionUnit{
		BaseUnit: NewBaseUnit(
			"dispersion",
			"ppt",
			"Parts per Trillion",
			0.000001, // 1 ppt = 0.000001 ppm
			0.0,
			false,
		),
	},
	Percent: DispersionUnit{
		BaseUnit: NewBaseUnit(
			"dispersion",
			"%",
			"Percent",
			10000.0, // 1% = 10,000 ppm
			0.0,
			false,
		),
	},
}

// NewDispersion creates a new dispersion quantity
func NewDispersion(value float64, unit DispersionUnit) Quantity[DispersionUnit] {
	return New(value, unit)
}
