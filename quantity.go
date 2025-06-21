// Package unit provides a system for representing, converting, and operating on
// physical quantities with units, inspired by Swift's Measurement framework.
package unit

import (
	"fmt"
	"math"
)

// Category is an interface that all unit types must implement
type Category interface {
	// Dimension returns the physical dimension of the unit (e.g., "temperature", "pressure")
	Dimension() string

	// Symbol returns the unit symbol (e.g., "°C", "Pa")
	Symbol() string

	// Name returns the human-readable name of the unit (e.g., "Celsius", "Pascal")
	Name() string

	// IsBaseUnit returns true if this is the base unit for its dimension
	IsBaseUnit() bool

	// ConvertToBaseUnit converts a value in this unit to the base unit
	ConvertToBaseUnit(value float64) float64

	// ConvertFromBaseUnit converts a value from the base unit to this unit
	ConvertFromBaseUnit(value float64) float64

	// Equals checks if this unit is equal to another unit
	Equals(other Category) bool
}

// Quantity represents a value with an associated unit.
// It is a generic type that can work with any unit type that implements Category.
type Quantity[T Category] struct {
	Value float64
	Unit  T
}

// Equal checks if two quantities are equal
func (m Quantity[T]) Equal(other Quantity[T]) bool {
	// Check if the dimensions are compatible
	if m.Unit.Dimension() != other.Unit.Dimension() {
		return false
	}

	// Convert both to base unit and compare values
	mBaseValue := m.Unit.ConvertToBaseUnit(m.Value)
	otherBaseValue := other.Unit.ConvertToBaseUnit(other.Value)

	// Use a small epsilon for floating point comparison
	const epsilon = 1e-9
	return math.Abs(mBaseValue-otherBaseValue) < epsilon
}

// New creates a new quantity with the given value and unit
func New[T Category](value float64, unit T) Quantity[T] {
	return Quantity[T]{
		Value: value,
		Unit:  unit,
	}
}

// ConvertTo converts this quantity to the specified unit
func (m Quantity[T]) ConvertTo(unit T) Quantity[T] {
	// If the units are the same, return a copy of the quantity
	if m.Unit.Equals(unit) {
		return Quantity[T]{
			Value: m.Value,
			Unit:  unit,
		}
	}

	// Check if the dimensions are compatible
	if m.Unit.Dimension() != unit.Dimension() {
		panic(fmt.Sprintf("Cannot convert from %s to %s: incompatible dimensions",
			m.Unit.Dimension(), unit.Dimension()))
	}

	// Convert to base unit first, then to the target unit
	valueInBaseUnit := m.Unit.ConvertToBaseUnit(m.Value)
	valueInTargetUnit := unit.ConvertFromBaseUnit(valueInBaseUnit)

	return Quantity[T]{
		Value: valueInTargetUnit,
		Unit:  unit,
	}
}

// Add adds another quantity to this one, converting if necessary
func (m Quantity[T]) Add(other Quantity[T]) Quantity[T] {
	// Check if the dimensions are compatible
	if m.Unit.Dimension() != other.Unit.Dimension() {
		panic(fmt.Sprintf("Cannot add %s and %s: incompatible dimensions",
			m.Unit.Dimension(), other.Unit.Dimension()))
	}

	// Convert the other quantity to this unit
	otherConverted := other.ConvertTo(m.Unit)

	// Add the values
	return Quantity[T]{
		Value: m.Value + otherConverted.Value,
		Unit:  m.Unit,
	}
}

// Subtract subtracts another quantity from this one, converting if necessary
func (m Quantity[T]) Subtract(other Quantity[T]) Quantity[T] {
	// Check if the dimensions are compatible
	if m.Unit.Dimension() != other.Unit.Dimension() {
		panic(fmt.Sprintf("Cannot subtract %s from %s: incompatible dimensions",
			other.Unit.Dimension(), m.Unit.Dimension()))
	}

	// Convert the other quantity to this unit
	otherConverted := other.ConvertTo(m.Unit)

	// Subtract the values
	return Quantity[T]{
		Value: m.Value - otherConverted.Value,
		Unit:  m.Unit,
	}
}

// MultiplyByScalar multiplies this quantity by a scalar value
func (m Quantity[T]) MultiplyByScalar(scalar float64) Quantity[T] {
	return Quantity[T]{
		Value: m.Value * scalar,
		Unit:  m.Unit,
	}
}

// DivideByScalar divides this quantity by a scalar value
func (m Quantity[T]) DivideByScalar(scalar float64) Quantity[T] {
	if scalar == 0 {
		panic("Cannot divide by zero")
	}

	return Quantity[T]{
		Value: m.Value / scalar,
		Unit:  m.Unit,
	}
}

// String returns a string representation of the quantity
func (m Quantity[T]) String() string {
	return fmt.Sprintf("%g %s", m.Value, m.Unit.Symbol())
}
