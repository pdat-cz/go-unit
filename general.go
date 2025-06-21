// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

// GeneralUnit represents a unit of an unspecified or general dimension
type GeneralUnit struct {
	BaseUnit
}

// General contains predefined general dimension units
var General = struct {
	Unit GeneralUnit // Base unit for general dimensions
}{
	Unit: GeneralUnit{
		BaseUnit: NewBaseUnit(
			"general",
			"unit",
			"General Unit",
			1.0,
			0.0,
			true, // Base unit
		),
	},
}

// NewGeneral creates a new general dimension quantity
func NewGeneral(value float64, unit GeneralUnit) Quantity[GeneralUnit] {
	return New(value, unit)
}

// NewGeneralUnit creates a custom general unit with the specified symbol and name
func NewGeneralUnit(symbol, name string) GeneralUnit {
	return GeneralUnit{
		BaseUnit: NewBaseUnit(
			"general",
			symbol,
			name,
			1.0, // Same conversion as base unit
			0.0,
			false,
		),
	}
}

// NewGeneralUnitWithConversion creates a custom general unit with the specified symbol, name,
// and conversion factors relative to the base general unit
func NewGeneralUnitWithConversion(symbol, name string, coefficient, offset float64) GeneralUnit {
	return GeneralUnit{
		BaseUnit: NewBaseUnit(
			"general",
			symbol,
			name,
			coefficient,
			offset,
			false,
		),
	}
}
