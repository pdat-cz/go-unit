// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

// AreaUnit represents a unit of an area
type AreaUnit struct {
	BaseUnit
}

// Area contains predefined area units
var Area = struct {
	SquareMeter      AreaUnit
	SquareKilometer  AreaUnit
	SquareCentimeter AreaUnit
	SquareMillimeter AreaUnit
	SquareInch       AreaUnit
	SquareFoot       AreaUnit
	SquareYard       AreaUnit
	SquareMile       AreaUnit
	Acre             AreaUnit
	Hectare          AreaUnit
}{
	SquareMeter: AreaUnit{
		BaseUnit: NewBaseUnit(
			"area",
			"m²",
			"Square Meter",
			1.0,
			0.0,
			true, // Base unit
		),
	},
	SquareKilometer: AreaUnit{
		BaseUnit: NewBaseUnit(
			"area",
			"km²",
			"Square Kilometer",
			1000000.0, // 1 km² = 1,000,000 m²
			0.0,
			false,
		),
	},
	SquareCentimeter: AreaUnit{
		BaseUnit: NewBaseUnit(
			"area",
			"cm²",
			"Square Centimeter",
			0.0001, // 1 cm² = 0.0001 m²
			0.0,
			false,
		),
	},
	SquareMillimeter: AreaUnit{
		BaseUnit: NewBaseUnit(
			"area",
			"mm²",
			"Square Millimeter",
			0.000001, // 1 mm² = 0.000001 m²
			0.0,
			false,
		),
	},
	SquareInch: AreaUnit{
		BaseUnit: NewBaseUnit(
			"area",
			"in²",
			"Square Inch",
			0.00064516, // 1 in² = 0.00064516 m²
			0.0,
			false,
		),
	},
	SquareFoot: AreaUnit{
		BaseUnit: NewBaseUnit(
			"area",
			"ft²",
			"Square Foot",
			0.09290304, // 1 ft² = 0.09290304 m²
			0.0,
			false,
		),
	},
	SquareYard: AreaUnit{
		BaseUnit: NewBaseUnit(
			"area",
			"yd²",
			"Square Yard",
			0.83612736, // 1 yd² = 0.83612736 m²
			0.0,
			false,
		),
	},
	SquareMile: AreaUnit{
		BaseUnit: NewBaseUnit(
			"area",
			"mi²",
			"Square Mile",
			2589988.11, // 1 mi² = 2,589,988.11 m²
			0.0,
			false,
		),
	},
	Acre: AreaUnit{
		BaseUnit: NewBaseUnit(
			"area",
			"ac",
			"Acre",
			4046.86, // 1 ac = 4,046.86 m²
			0.0,
			false,
		),
	},
	Hectare: AreaUnit{
		BaseUnit: NewBaseUnit(
			"area",
			"ha",
			"Hectare",
			10000.0, // 1 ha = 10,000 m²
			0.0,
			false,
		),
	},
}

// NewArea creates a new area measurement
func NewArea(value float64, unit AreaUnit) Quantity[AreaUnit] {
	return New(value, unit)
}
