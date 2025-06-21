// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

// LengthUnit represents a unit of length
type LengthUnit struct {
	BaseUnit
}

// Length contains predefined length units
var Length = struct {
	Meter      LengthUnit
	Kilometer  LengthUnit
	Centimeter LengthUnit
	Millimeter LengthUnit
	Micrometer LengthUnit
	Nanometer  LengthUnit
	Inch       LengthUnit
	Foot       LengthUnit
	Yard       LengthUnit
	Mile       LengthUnit
}{
	Meter: LengthUnit{
		BaseUnit: NewBaseUnit(
			"length",
			"m",
			"Meter",
			1.0,
			0.0,
			true, // Base unit
		),
	},
	Kilometer: LengthUnit{
		BaseUnit: NewBaseUnit(
			"length",
			"km",
			"Kilometer",
			1000.0, // 1 km = 1000 m
			0.0,
			false,
		),
	},
	Centimeter: LengthUnit{
		BaseUnit: NewBaseUnit(
			"length",
			"cm",
			"Centimeter",
			0.01, // 1 cm = 0.01 m
			0.0,
			false,
		),
	},
	Millimeter: LengthUnit{
		BaseUnit: NewBaseUnit(
			"length",
			"mm",
			"Millimeter",
			0.001, // 1 mm = 0.001 m
			0.0,
			false,
		),
	},
	Micrometer: LengthUnit{
		BaseUnit: NewBaseUnit(
			"length",
			"µm",
			"Micrometer",
			0.000001, // 1 µm = 0.000001 m
			0.0,
			false,
		),
	},
	Nanometer: LengthUnit{
		BaseUnit: NewBaseUnit(
			"length",
			"nm",
			"Nanometer",
			0.000000001, // 1 nm = 0.000000001 m
			0.0,
			false,
		),
	},
	Inch: LengthUnit{
		BaseUnit: NewBaseUnit(
			"length",
			"in",
			"Inch",
			0.0254, // 1 in = 0.0254 m
			0.0,
			false,
		),
	},
	Foot: LengthUnit{
		BaseUnit: NewBaseUnit(
			"length",
			"ft",
			"Foot",
			0.3048, // 1 ft = 0.3048 m
			0.0,
			false,
		),
	},
	Yard: LengthUnit{
		BaseUnit: NewBaseUnit(
			"length",
			"yd",
			"Yard",
			0.9144, // 1 yd = 0.9144 m
			0.0,
			false,
		),
	},
	Mile: LengthUnit{
		BaseUnit: NewBaseUnit(
			"length",
			"mi",
			"Mile",
			1609.34, // 1 mi = 1609.34 m
			0.0,
			false,
		),
	},
}

// NewLength creates a new length quantity
func NewLength(value float64, unit LengthUnit) Quantity[LengthUnit] {
	return New(value, unit)
}
