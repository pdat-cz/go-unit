// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

// PressureUnit represents a unit of pressure
type PressureUnit struct {
	BaseUnit
}

// Pressure contains predefined pressure units
var Pressure = struct {
	Pascal     PressureUnit
	Kilopascal PressureUnit
	Bar        PressureUnit
	PSI        PressureUnit
	InchH2O    PressureUnit
}{
	Pascal: PressureUnit{
		BaseUnit: NewBaseUnit(
			"pressure",
			"Pa",
			"Pascal",
			1.0,
			0.0,
			true, // Base unit
		),
	},
	Kilopascal: PressureUnit{
		BaseUnit: NewBaseUnit(
			"pressure",
			"kPa",
			"Kilopascal",
			1000.0, // 1 kPa = 1000 Pa
			0.0,
			false,
		),
	},
	Bar: PressureUnit{
		BaseUnit: NewBaseUnit(
			"pressure",
			"bar",
			"Bar",
			100000.0, // 1 bar = 100,000 Pa
			0.0,
			false,
		),
	},
	PSI: PressureUnit{
		BaseUnit: NewBaseUnit(
			"pressure",
			"psi",
			"Pounds per Square Inch",
			6894.76, // 1 psi = 6,894.76 Pa
			0.0,
			false,
		),
	},
	InchH2O: PressureUnit{
		BaseUnit: NewBaseUnit(
			"pressure",
			"inH₂O",
			"Inches of Water Column",
			249.089, // 1 inH₂O = 249.089 Pa
			0.0,
			false,
		),
	},
}

// NewPressure creates a new pressure quantity
func NewPressure(value float64, unit PressureUnit) Quantity[PressureUnit] {
	return New(value, unit)
}
