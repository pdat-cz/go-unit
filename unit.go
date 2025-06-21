// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

// BaseUnit provides a common implementation of the Category interface
type BaseUnit struct {
	dimension   string
	symbol      string
	name        string
	coefficient float64
	offset      float64
	isBase      bool
}

// NewBaseUnit creates a new BaseUnit
func NewBaseUnit(dimension, symbol, name string, coefficient, offset float64, isBase bool) BaseUnit {
	return BaseUnit{
		dimension:   dimension,
		symbol:      symbol,
		name:        name,
		coefficient: coefficient,
		offset:      offset,
		isBase:      isBase,
	}
}

// Dimension returns the physical dimension of the unit
func (u BaseUnit) Dimension() string {
	return u.dimension
}

// Symbol returns the unit symbol
func (u BaseUnit) Symbol() string {
	return u.symbol
}

// Name returns the human-readable name of the unit
func (u BaseUnit) Name() string {
	return u.name
}

// IsBaseUnit returns true if this is the base unit for its dimension
func (u BaseUnit) IsBaseUnit() bool {
	return u.isBase
}

// ConvertToBaseUnit converts a value in this unit to the base unit
func (u BaseUnit) ConvertToBaseUnit(value float64) float64 {
	if u.isBase {
		return value
	}

	// For temperature and similar units with an offset
	if u.offset != 0 {
		return (value * u.coefficient) + u.offset
	}

	// For most units, just multiply by the coefficient
	return value * u.coefficient
}

// ConvertFromBaseUnit converts a value from the base unit to this unit
func (u BaseUnit) ConvertFromBaseUnit(value float64) float64 {
	if u.isBase {
		return value
	}

	// For temperature and similar units with an offset
	if u.offset != 0 {
		return (value - u.offset) / u.coefficient
	}

	// For most units, just divide by the coefficient
	return value / u.coefficient
}

// Equals checks if this unit is equal to another unit
func (u BaseUnit) Equals(other Category) bool {
	// Check if dimensions match
	if u.dimension != other.Dimension() {
		return false
	}

	// Check if symbols match
	if u.symbol != other.Symbol() {
		return false
	}

	return true
}
