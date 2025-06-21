// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

// VolumeUnit represents a unit of volume
type VolumeUnit struct {
	BaseUnit
}

// Volume contains predefined volume units
var Volume = struct {
	CubicMeter      VolumeUnit
	CubicKilometer  VolumeUnit
	CubicCentimeter VolumeUnit
	CubicMillimeter VolumeUnit
	Liter           VolumeUnit
	Milliliter      VolumeUnit
	CubicInch       VolumeUnit
	CubicFoot       VolumeUnit
	CubicYard       VolumeUnit
	Gallon          VolumeUnit
	Quart           VolumeUnit
	Pint            VolumeUnit
	Cup             VolumeUnit
	FluidOunce      VolumeUnit
}{
	CubicMeter: VolumeUnit{
		BaseUnit: NewBaseUnit(
			"volume",
			"m³",
			"Cubic Meter",
			1.0,
			0.0,
			true, // Base unit
		),
	},
	CubicKilometer: VolumeUnit{
		BaseUnit: NewBaseUnit(
			"volume",
			"km³",
			"Cubic Kilometer",
			1000000000.0, // 1 km³ = 1,000,000,000 m³
			0.0,
			false,
		),
	},
	CubicCentimeter: VolumeUnit{
		BaseUnit: NewBaseUnit(
			"volume",
			"cm³",
			"Cubic Centimeter",
			0.000001, // 1 cm³ = 0.000001 m³
			0.0,
			false,
		),
	},
	CubicMillimeter: VolumeUnit{
		BaseUnit: NewBaseUnit(
			"volume",
			"mm³",
			"Cubic Millimeter",
			0.000000001, // 1 mm³ = 0.000000001 m³
			0.0,
			false,
		),
	},
	Liter: VolumeUnit{
		BaseUnit: NewBaseUnit(
			"volume",
			"L",
			"Liter",
			0.001, // 1 L = 0.001 m³
			0.0,
			false,
		),
	},
	Milliliter: VolumeUnit{
		BaseUnit: NewBaseUnit(
			"volume",
			"mL",
			"Milliliter",
			0.000001, // 1 mL = 0.000001 m³
			0.0,
			false,
		),
	},
	CubicInch: VolumeUnit{
		BaseUnit: NewBaseUnit(
			"volume",
			"in³",
			"Cubic Inch",
			0.000016387064, // 1 in³ = 0.000016387064 m³
			0.0,
			false,
		),
	},
	CubicFoot: VolumeUnit{
		BaseUnit: NewBaseUnit(
			"volume",
			"ft³",
			"Cubic Foot",
			0.028316846592, // 1 ft³ = 0.028316846592 m³
			0.0,
			false,
		),
	},
	CubicYard: VolumeUnit{
		BaseUnit: NewBaseUnit(
			"volume",
			"yd³",
			"Cubic Yard",
			0.764554857984, // 1 yd³ = 0.764554857984 m³
			0.0,
			false,
		),
	},
	Gallon: VolumeUnit{
		BaseUnit: NewBaseUnit(
			"volume",
			"gal",
			"Gallon",
			0.003785411784, // 1 gal = 0.003785411784 m³ (US gallon)
			0.0,
			false,
		),
	},
	Quart: VolumeUnit{
		BaseUnit: NewBaseUnit(
			"volume",
			"qt",
			"Quart",
			0.000946352946, // 1 qt = 0.000946352946 m³ (US quart)
			0.0,
			false,
		),
	},
	Pint: VolumeUnit{
		BaseUnit: NewBaseUnit(
			"volume",
			"pt",
			"Pint",
			0.000473176473, // 1 pt = 0.000473176473 m³ (US pint)
			0.0,
			false,
		),
	},
	Cup: VolumeUnit{
		BaseUnit: NewBaseUnit(
			"volume",
			"cup",
			"Cup",
			0.000236588236, // 1 cup = 0.000236588236 m³ (US cup)
			0.0,
			false,
		),
	},
	FluidOunce: VolumeUnit{
		BaseUnit: NewBaseUnit(
			"volume",
			"fl oz",
			"Fluid Ounce",
			0.0000295735295625, // 1 fl oz = 0.0000295735295625 m³ (US fluid ounce)
			0.0,
			false,
		),
	},
}

// NewVolume creates a new volume quantity
func NewVolume(value float64, unit VolumeUnit) Quantity[VolumeUnit] {
	return New(value, unit)
}
