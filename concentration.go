// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

// ConcentrationUnit represents a unit of concentration of mass
type ConcentrationUnit struct {
	BaseUnit
}

// Concentration contains predefined concentration units
var Concentration = struct {
	GramsPerLiter      ConcentrationUnit
	MilligramsPerLiter ConcentrationUnit
	PartsPerMillion    ConcentrationUnit
	PartsPerBillion    ConcentrationUnit
}{
	GramsPerLiter: ConcentrationUnit{
		BaseUnit: NewBaseUnit(
			"concentration",
			"g/L",
			"Grams per Liter",
			1.0,
			0.0,
			true, // Base unit
		),
	},
	MilligramsPerLiter: ConcentrationUnit{
		BaseUnit: NewBaseUnit(
			"concentration",
			"mg/L",
			"Milligrams per Liter",
			0.001, // 1 mg/L = 0.001 g/L
			0.0,
			false,
		),
	},
	PartsPerMillion: ConcentrationUnit{
		BaseUnit: NewBaseUnit(
			"concentration",
			"ppm",
			"Parts per Million",
			0.001, // 1 ppm = 0.001 g/L (for water solutions)
			0.0,
			false,
		),
	},
	PartsPerBillion: ConcentrationUnit{
		BaseUnit: NewBaseUnit(
			"concentration",
			"ppb",
			"Parts per Billion",
			0.000001, // 1 ppb = 0.000001 g/L (for water solutions)
			0.0,
			false,
		),
	},
}

// NewConcentration creates a new concentration measurement
func NewConcentration(value float64, unit ConcentrationUnit) Quantity[ConcentrationUnit] {
	return New(value, unit)
}
