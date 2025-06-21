// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

// MassUnit represents a unit of mass
type MassUnit struct {
	BaseUnit
}

// Mass contains predefined mass units
var Mass = struct {
	Kilogram  MassUnit
	Gram      MassUnit
	Milligram MassUnit
	Microgram MassUnit
	Pound     MassUnit
	Ounce     MassUnit
	Stone     MassUnit
	MetricTon MassUnit
	Ton       MassUnit
}{
	Kilogram: MassUnit{
		BaseUnit: NewBaseUnit(
			"mass",
			"kg",
			"Kilogram",
			1.0,
			0.0,
			true, // Base unit
		),
	},
	Gram: MassUnit{
		BaseUnit: NewBaseUnit(
			"mass",
			"g",
			"Gram",
			0.001, // 1 g = 0.001 kg
			0.0,
			false,
		),
	},
	Milligram: MassUnit{
		BaseUnit: NewBaseUnit(
			"mass",
			"mg",
			"Milligram",
			0.000001, // 1 mg = 0.000001 kg
			0.0,
			false,
		),
	},
	Microgram: MassUnit{
		BaseUnit: NewBaseUnit(
			"mass",
			"µg",
			"Microgram",
			0.000000001, // 1 µg = 0.000000001 kg
			0.0,
			false,
		),
	},
	Pound: MassUnit{
		BaseUnit: NewBaseUnit(
			"mass",
			"lb",
			"Pound",
			0.45359237, // 1 lb = 0.45359237 kg
			0.0,
			false,
		),
	},
	Ounce: MassUnit{
		BaseUnit: NewBaseUnit(
			"mass",
			"oz",
			"Ounce",
			0.028349523125, // 1 oz = 0.028349523125 kg
			0.0,
			false,
		),
	},
	Stone: MassUnit{
		BaseUnit: NewBaseUnit(
			"mass",
			"st",
			"Stone",
			6.35029318, // 1 st = 6.35029318 kg
			0.0,
			false,
		),
	},
	MetricTon: MassUnit{
		BaseUnit: NewBaseUnit(
			"mass",
			"t",
			"Metric Ton",
			1000.0, // 1 t = 1000 kg
			0.0,
			false,
		),
	},
	Ton: MassUnit{
		BaseUnit: NewBaseUnit(
			"mass",
			"ton",
			"Ton",
			907.18474, // 1 ton = 907.18474 kg (US short ton)
			0.0,
			false,
		),
	},
}

// NewMass creates a new mass quantity
func NewMass(value float64, unit MassUnit) Quantity[MassUnit] {
	return New(value, unit)
}
