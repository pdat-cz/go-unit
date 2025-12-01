// Package unit provides a system for representing, converting, and operating on
// physical quantities with units, inspired by Swift's Measurement framework.
package unit

import (
	"encoding/json"
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

// MarshalJSON implements json.Marshaler interface
func (m Quantity[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"unit"`
		Dimension string `json:"dimension"`
	}{
		Value: m.Value,
		Unit: struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		}{
			Name:   m.Unit.Name(),
			Symbol: m.Unit.Symbol(),
		},
		Dimension: m.Unit.Dimension(),
	})
}

// UnmarshalJSON implements json.Unmarshaler interface
func (m *Quantity[T]) UnmarshalJSON(data []byte) error {
	var raw struct {
		Value float64 `json:"value"`
		Unit  struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"unit"`
		Dimension string `json:"dimension"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	m.Value = raw.Value

	// Look up unit by dimension and symbol
	unit, err := lookupUnit[T](raw.Dimension, raw.Unit.Symbol)
	if err != nil {
		return err
	}

	m.Unit = unit
	return nil
}

// Compact wraps a Quantity for compact JSON serialization
// Use this when you need the compact format: {"value":25,"unit":"temperature_celsius"}
type Compact[T Category] struct {
	Quantity[T]
}

// MarshalJSON implements json.Marshaler for compact format
func (c Compact[T]) MarshalJSON() ([]byte, error) {
	key := toSnakeCase(c.Unit.Dimension()) + "_" + toSnakeCase(c.Unit.Name())
	return json.Marshal(struct {
		Value  float64 `json:"value"`
		Unit   string  `json:"unit"`
		Symbol string  `json:"symbol,omitempty"`
	}{
		Value:  c.Value,
		Unit:   key,
		Symbol: c.Unit.Symbol(),
	})
}

// UnmarshalJSON implements json.Unmarshaler for compact format
func (c *Compact[T]) UnmarshalJSON(data []byte) error {
	var raw struct {
		Value  float64 `json:"value"`
		Unit   string  `json:"unit"`
		Symbol string  `json:"symbol"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	c.Quantity.Value = raw.Value

	// Parse compact unit key (e.g., "temperature_celsius")
	dimension, unitName := parseCompactUnitKey(raw.Unit)

	// Look up unit by dimension and unit name
	unit, err := lookupUnitByName[T](dimension, unitName)
	if err != nil {
		return err
	}

	c.Quantity.Unit = unit
	return nil
}

// parseCompactUnitKey splits "dimension_unit_name" into dimension and unit name
func parseCompactUnitKey(key string) (dimension, unitName string) {
	idx := indexOf(key, "_")
	if idx == -1 {
		return key, ""
	}
	return key[:idx], key[idx+1:]
}

// indexOf returns the index of the first occurrence of sep in s
func indexOf(s, sep string) int {
	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			return i
		}
	}
	return -1
}

// lookupUnitByName finds a unit by dimension and snake_case name
func lookupUnitByName[T Category](dimension, name string) (T, error) {
	var zero T
	var result any

	key := dimension + "_" + name

	switch dimension {
	case "temperature":
		if u, ok := temperatureUnitsByKey[key]; ok {
			result = u
		}
	case "pressure":
		if u, ok := pressureUnitsByKey[key]; ok {
			result = u
		}
	case "flowrate":
		if u, ok := flowRateUnitsByKey[key]; ok {
			result = u
		}
	case "power":
		if u, ok := powerUnitsByKey[key]; ok {
			result = u
		}
	case "energy":
		if u, ok := energyUnitsByKey[key]; ok {
			result = u
		}
	case "length":
		if u, ok := lengthUnitsByKey[key]; ok {
			result = u
		}
	case "mass":
		if u, ok := massUnitsByKey[key]; ok {
			result = u
		}
	case "duration":
		if u, ok := durationUnitsByKey[key]; ok {
			result = u
		}
	case "angle":
		if u, ok := angleUnitsByKey[key]; ok {
			result = u
		}
	case "area":
		if u, ok := areaUnitsByKey[key]; ok {
			result = u
		}
	case "volume":
		if u, ok := volumeUnitsByKey[key]; ok {
			result = u
		}
	case "acceleration":
		if u, ok := accelerationUnitsByKey[key]; ok {
			result = u
		}
	case "concentration":
		if u, ok := concentrationUnitsByKey[key]; ok {
			result = u
		}
	case "dispersion":
		if u, ok := dispersionUnitsByKey[key]; ok {
			result = u
		}
	case "speed":
		if u, ok := speedUnitsByKey[key]; ok {
			result = u
		}
	case "electric_charge":
		if u, ok := electricChargeUnitsByKey[key]; ok {
			result = u
		}
	case "electric_current":
		if u, ok := electricCurrentUnitsByKey[key]; ok {
			result = u
		}
	case "electric_potential_difference":
		if u, ok := electricPotentialDifferenceUnitsByKey[key]; ok {
			result = u
		}
	case "frequency":
		if u, ok := frequencyUnitsByKey[key]; ok {
			result = u
		}
	case "illuminance":
		if u, ok := illuminanceUnitsByKey[key]; ok {
			result = u
		}
	case "information":
		if u, ok := informationUnitsByKey[key]; ok {
			result = u
		}
	case "fuel_efficiency":
		if u, ok := fuelEfficiencyUnitsByKey[key]; ok {
			result = u
		}
	case "general":
		result = NewGeneralUnit(name, name)
	default:
		return zero, fmt.Errorf("unknown dimension: %s", dimension)
	}

	if result == nil {
		return zero, fmt.Errorf("unknown unit key %q", key)
	}

	if typed, ok := result.(T); ok {
		return typed, nil
	}

	return zero, fmt.Errorf("unit type mismatch: expected %T, got %T", zero, result)
}

// lookupUnit finds a unit by dimension and symbol
func lookupUnit[T Category](dimension, symbol string) (T, error) {
	var zero T
	var result any

	switch dimension {
	case "temperature":
		if u, ok := LookupTemperatureUnit(symbol); ok {
			result = u
		}
	case "pressure":
		if u, ok := LookupPressureUnit(symbol); ok {
			result = u
		}
	case "flowrate":
		if u, ok := LookupFlowRateUnit(symbol); ok {
			result = u
		}
	case "power":
		if u, ok := LookupPowerUnit(symbol); ok {
			result = u
		}
	case "energy":
		if u, ok := LookupEnergyUnit(symbol); ok {
			result = u
		}
	case "length":
		if u, ok := LookupLengthUnit(symbol); ok {
			result = u
		}
	case "mass":
		if u, ok := LookupMassUnit(symbol); ok {
			result = u
		}
	case "duration":
		if u, ok := LookupDurationUnit(symbol); ok {
			result = u
		}
	case "angle":
		if u, ok := LookupAngleUnit(symbol); ok {
			result = u
		}
	case "area":
		if u, ok := LookupAreaUnit(symbol); ok {
			result = u
		}
	case "volume":
		if u, ok := LookupVolumeUnit(symbol); ok {
			result = u
		}
	case "acceleration":
		if u, ok := LookupAccelerationUnit(symbol); ok {
			result = u
		}
	case "concentration":
		if u, ok := LookupConcentrationUnit(symbol); ok {
			result = u
		}
	case "dispersion":
		if u, ok := LookupDispersionUnit(symbol); ok {
			result = u
		}
	case "speed":
		if u, ok := LookupSpeedUnit(symbol); ok {
			result = u
		}
	case "electric_charge":
		if u, ok := LookupElectricChargeUnit(symbol); ok {
			result = u
		}
	case "electric_current":
		if u, ok := LookupElectricCurrentUnit(symbol); ok {
			result = u
		}
	case "electric_potential_difference":
		if u, ok := LookupElectricPotentialDifferenceUnit(symbol); ok {
			result = u
		}
	case "frequency":
		if u, ok := LookupFrequencyUnit(symbol); ok {
			result = u
		}
	case "illuminance":
		if u, ok := LookupIlluminanceUnit(symbol); ok {
			result = u
		}
	case "information":
		if u, ok := LookupInformationUnit(symbol); ok {
			result = u
		}
	case "fuel_efficiency":
		if u, ok := LookupFuelEfficiencyUnit(symbol); ok {
			result = u
		}
	case "general":
		result = NewGeneralUnit(symbol, symbol)
	default:
		return zero, fmt.Errorf("unknown dimension: %s", dimension)
	}

	if result == nil {
		return zero, fmt.Errorf("unknown unit symbol %q for dimension %q", symbol, dimension)
	}

	// Type assert to T
	if typed, ok := result.(T); ok {
		return typed, nil
	}

	return zero, fmt.Errorf("unit type mismatch: expected %T, got %T", zero, result)
}
