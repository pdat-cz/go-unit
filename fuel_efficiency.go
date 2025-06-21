// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

// FuelEfficiencyUnit represents a unit of fuel efficiency
type FuelEfficiencyUnit struct {
	BaseUnit
}

// FuelEfficiency contains predefined fuel efficiency units
var FuelEfficiency = struct {
	KilometersPerLiter     FuelEfficiencyUnit
	MilesPerGallon         FuelEfficiencyUnit
	LitersPer100Kilometers FuelEfficiencyUnit
}{
	KilometersPerLiter: FuelEfficiencyUnit{
		BaseUnit: NewBaseUnit(
			"fuel_efficiency",
			"km/L",
			"Kilometers per Liter",
			1.0,
			0.0,
			true, // Base unit
		),
	},
	MilesPerGallon: FuelEfficiencyUnit{
		BaseUnit: NewBaseUnit(
			"fuel_efficiency",
			"mpg",
			"Miles per Gallon",
			0.425144, // 1 mpg ≈ 0.425144 km/L
			0.0,
			false,
		),
	},
	LitersPer100Kilometers: FuelEfficiencyUnit{
		BaseUnit: NewBaseUnit(
			"fuel_efficiency",
			"L/100km",
			"Liters per 100 Kilometers",
			-100.0, // Special case: inverse relationship
			0.0,
			false,
		),
	},
}

// NewFuelEfficiency creates a new fuel efficiency quantity
func NewFuelEfficiency(value float64, unit FuelEfficiencyUnit) Quantity[FuelEfficiencyUnit] {
	// Special handling for L/100km since it's an inverse measure
	if unit.Equals(FuelEfficiency.LitersPer100Kilometers) && value == 0 {
		panic("Cannot create fuel efficiency with 0 L/100km (infinite efficiency)")
	}
	return New(value, unit)
}

// ConvertToBaseUnit Override ConvertToBaseUnit for FuelEfficiencyUnit to handle the inverse relationship of L/100km
func (u FuelEfficiencyUnit) ConvertToBaseUnit(value float64) float64 {
	if u.Symbol() == "L/100km" {
		// For L/100km, we need to convert to km/L which is the inverse
		if value == 0 {
			panic("Cannot convert 0 L/100km to km/L (infinite efficiency)")
		}
		return 100.0 / value
	}

	// For other units, use the standard conversion
	return u.BaseUnit.ConvertToBaseUnit(value)
}

// Override ConvertFromBaseUnit for FuelEfficiencyUnit to handle the inverse relationship of L/100km
func (u FuelEfficiencyUnit) ConvertFromBaseUnit(value float64) float64 {
	if u.Symbol() == "L/100km" {
		// For L/100km, we need to convert from km/L which is the inverse
		if value == 0 {
			panic("Cannot convert 0 km/L to L/100km (infinite consumption)")
		}
		return 100.0 / value
	}

	// For other units, use the standard conversion
	return u.BaseUnit.ConvertFromBaseUnit(value)
}
