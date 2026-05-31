// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

// PressureUnit represents a unit of pressure
type PressureUnit struct {
	BaseUnit
}

// Pressure contains predefined pressure units
var Pressure = struct {
	Pascal      PressureUnit
	Hectopascal PressureUnit
	Kilopascal  PressureUnit
	Megapascal  PressureUnit
	Bar         PressureUnit
	Millibar    PressureUnit
	PSI         PressureUnit
	Atmosphere  PressureUnit
	MmHg        PressureUnit
	InchH2O     PressureUnit
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
	Hectopascal: PressureUnit{
		BaseUnit: NewBaseUnit(
			"pressure",
			"hPa",
			"Hectopascal",
			100.0, // 1 hPa = 100 Pa
			0.0,
			false,
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
	Megapascal: PressureUnit{
		BaseUnit: NewBaseUnit(
			"pressure",
			"MPa",
			"Megapascal",
			1e6, // 1 MPa = 1,000,000 Pa
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
	Millibar: PressureUnit{
		BaseUnit: NewBaseUnit(
			"pressure",
			"mbar",
			"Millibar",
			100.0, // 1 mbar = 100 Pa
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
	Atmosphere: PressureUnit{
		BaseUnit: NewBaseUnit(
			"pressure",
			"atm",
			"Atmosphere",
			101325.0, // 1 atm = 101,325 Pa
			0.0,
			false,
		),
	},
	MmHg: PressureUnit{
		BaseUnit: NewBaseUnit(
			"pressure",
			"mmHg",
			"Millimeters of Mercury",
			133.322387415, // 1 mmHg = 133.322387415 Pa (Torr)
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
