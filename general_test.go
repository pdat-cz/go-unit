package unit

import (
	"testing"
)

func TestGeneralUnit(t *testing.T) {
	// Test the base unit
	baseUnit := General.Unit
	if baseUnit.Dimension() != "general" {
		t.Errorf("Expected dimension 'general', got '%s'", baseUnit.Dimension())
	}
	if baseUnit.Symbol() != "unit" {
		t.Errorf("Expected symbol 'unit', got '%s'", baseUnit.Symbol())
	}
	if baseUnit.Name() != "General Unit" {
		t.Errorf("Expected name 'General Unit', got '%s'", baseUnit.Name())
	}
	if !baseUnit.IsBaseUnit() {
		t.Error("Expected base unit to be true")
	}

	// Test creating a custom unit
	customUnit := NewGeneralUnit("cu", "Custom Unit")
	if customUnit.Dimension() != "general" {
		t.Errorf("Expected dimension 'general', got '%s'", customUnit.Dimension())
	}
	if customUnit.Symbol() != "cu" {
		t.Errorf("Expected symbol 'cu', got '%s'", customUnit.Symbol())
	}
	if customUnit.Name() != "Custom Unit" {
		t.Errorf("Expected name 'Custom Unit', got '%s'", customUnit.Name())
	}
	if customUnit.IsBaseUnit() {
		t.Error("Expected custom unit to not be a base unit")
	}

	// Test creating a unit with conversion
	convUnit := NewGeneralUnitWithConversion("xu", "X Unit", 2.0, 10.0)
	if convUnit.Dimension() != "general" {
		t.Errorf("Expected dimension 'general', got '%s'", convUnit.Dimension())
	}
	if convUnit.Symbol() != "xu" {
		t.Errorf("Expected symbol 'xu', got '%s'", convUnit.Symbol())
	}
	if convUnit.Name() != "X Unit" {
		t.Errorf("Expected name 'X Unit', got '%s'", convUnit.Name())
	}

	// Test conversion
	baseValue := 5.0
	convValue := convUnit.ConvertFromBaseUnit(baseValue)
	expectedConvValue := (baseValue - 10.0) / 2.0
	if convValue != expectedConvValue {
		t.Errorf("Expected conversion value %f, got %f", expectedConvValue, convValue)
	}

	// Convert back to base
	reconvValue := convUnit.ConvertToBaseUnit(convValue)
	if reconvValue != baseValue {
		t.Errorf("Expected reconversion value %f, got %f", baseValue, reconvValue)
	}
}

func TestGeneralMeasurement(t *testing.T) {
	// Create measurements
	baseMeasurement := NewGeneral(10.0, General.Unit)
	customUnit := NewGeneralUnit("cu", "Custom Unit")
	customMeasurement := NewGeneral(10.0, customUnit)

	// Test equality of same measurements
	if !baseMeasurement.Equal(baseMeasurement) {
		t.Error("Expected base measurement to equal itself")
	}

	// Test equality of measurements with different units but same value
	if !baseMeasurement.Equal(customMeasurement) {
		t.Error("Expected base measurement to equal custom measurement with same value")
	}

	// Test conversion
	convUnit := NewGeneralUnitWithConversion("xu", "X Unit", 2.0, 10.0)
	convMeasurement := NewGeneral(5.0, convUnit)

	// Convert to base unit
	baseConverted := convMeasurement.ConvertTo(General.Unit)
	expectedBaseValue := (5.0 * 2.0) + 10.0 // Using the conversion formula
	if baseConverted.Value != expectedBaseValue {
		t.Errorf("Expected converted value %f, got %f", expectedBaseValue, baseConverted.Value)
	}

	// Convert back to original unit
	reconverted := baseConverted.ConvertTo(convUnit)
	if reconverted.Value != convMeasurement.Value {
		t.Errorf("Expected reconverted value %f, got %f", convMeasurement.Value, reconverted.Value)
	}

	// Test addition
	sum := baseMeasurement.Add(customMeasurement)
	if sum.Value != 20.0 {
		t.Errorf("Expected sum value 20.0, got %f", sum.Value)
	}

	// Test subtraction
	diff := baseMeasurement.Subtract(customMeasurement)
	if diff.Value != 0.0 {
		t.Errorf("Expected difference value 0.0, got %f", diff.Value)
	}

	// Test multiplication by scalar
	mult := baseMeasurement.MultiplyByScalar(2.5)
	if mult.Value != 25.0 {
		t.Errorf("Expected product value 25.0, got %f", mult.Value)
	}

	// Test division by scalar
	div := baseMeasurement.DivideByScalar(2.0)
	if div.Value != 5.0 {
		t.Errorf("Expected quotient value 5.0, got %f", div.Value)
	}
}
