package unit

import (
	"math"
	"testing"
)

func TestElectricCurrentConversion(t *testing.T) {
	// Create an electric current in amperes
	current := NewElectricCurrent(2.5, ElectricCurrent.Ampere)

	// Convert to milliamperes
	currentMA := current.ConvertTo(ElectricCurrent.Milliampere)

	// Expected: 2.5 A = 2,500 mA
	expected := 2500.0
	if math.Abs(currentMA.Value-expected) > 0.001 {
		t.Errorf("Electric current conversion failed: got %g mA, expected %g mA", currentMA.Value, expected)
	}

	// Convert to microamperes
	currentUA := current.ConvertTo(ElectricCurrent.Microampere)

	// Expected: 2.5 A = 2,500,000 µA
	expected = 2500000.0
	if math.Abs(currentUA.Value-expected) > 0.001 {
		t.Errorf("Electric current conversion failed: got %g µA, expected %g µA", currentUA.Value, expected)
	}

	// Convert to kiloamperes
	currentKA := current.ConvertTo(ElectricCurrent.Kiloampere)

	// Expected: 2.5 A = 0.0025 kA
	expected = 0.0025
	if math.Abs(currentKA.Value-expected) > 0.0001 {
		t.Errorf("Electric current conversion failed: got %g kA, expected %g kA", currentKA.Value, expected)
	}

	// Convert back to amperes
	current2 := currentMA.ConvertTo(ElectricCurrent.Ampere)

	// Should get the original value back
	if math.Abs(current2.Value-current.Value) > 0.001 {
		t.Errorf("Round-trip conversion failed: got %g A, expected %g A", current2.Value, current.Value)
	}
}

func TestElectricCurrentArithmetic(t *testing.T) {
	// Create two electric currents
	current1 := NewElectricCurrent(2.5, ElectricCurrent.Ampere)
	current2 := NewElectricCurrent(500.0, ElectricCurrent.Milliampere)

	// Add them (should convert current2 to amperes first)
	sum := current1.Add(current2)

	// Expected: 2.5 A + 0.5 A = 3.0 A
	expected := 3.0
	if math.Abs(sum.Value-expected) > 0.001 {
		t.Errorf("Addition failed: got %g A, expected %g A", sum.Value, expected)
	}

	// Subtract
	diff := current1.Subtract(current2)

	// Expected: 2.5 A - 0.5 A = 2.0 A
	expected = 2.0
	if math.Abs(diff.Value-expected) > 0.001 {
		t.Errorf("Subtraction failed: got %g A, expected %g A", diff.Value, expected)
	}

	// Multiply by scalar
	doubled := current1.MultiplyByScalar(2.0)

	// Expected: 2.5 A * 2 = 5.0 A
	expected = 5.0
	if math.Abs(doubled.Value-expected) > 0.001 {
		t.Errorf("Scalar multiplication failed: got %g A, expected %g A", doubled.Value, expected)
	}

	// Divide by scalar
	halved := current1.DivideByScalar(2.0)

	// Expected: 2.5 A / 2 = 1.25 A
	expected = 1.25
	if math.Abs(halved.Value-expected) > 0.001 {
		t.Errorf("Scalar division failed: got %g A, expected %g A", halved.Value, expected)
	}
}

func TestElectricCurrentParsing(t *testing.T) {
	// Test electric current parsing
	currentA, err := ParseElectricCurrent("2.5 A")
	if err != nil {
		t.Errorf("Failed to parse electric current: %v", err)
	}
	if math.Abs(currentA.Value-2.5) > 0.001 || !currentA.Unit.Equals(ElectricCurrent.Ampere) {
		t.Errorf("Parsed electric current incorrect: got %v, expected 2.5 A", currentA)
	}

	currentMA, err := ParseElectricCurrent("500 mA")
	if err != nil {
		t.Errorf("Failed to parse electric current: %v", err)
	}
	if math.Abs(currentMA.Value-500.0) > 0.001 || !currentMA.Unit.Equals(ElectricCurrent.Milliampere) {
		t.Errorf("Parsed electric current incorrect: got %v, expected 500 mA", currentMA)
	}

	currentKA, err := ParseElectricCurrent("0.01 kA")
	if err != nil {
		t.Errorf("Failed to parse electric current: %v", err)
	}
	if math.Abs(currentKA.Value-0.01) > 0.001 || !currentKA.Unit.Equals(ElectricCurrent.Kiloampere) {
		t.Errorf("Parsed electric current incorrect: got %v, expected 0.01 kA", currentKA)
	}
}

func TestElectricCurrentSerialization(t *testing.T) {
	// Create an electric current
	current := NewElectricCurrent(2.5, ElectricCurrent.Ampere)

	// Serialize to JSON
	data, err := MarshalElectricCurrent(current)
	if err != nil {
		t.Fatalf("Failed to marshal electric current: %v", err)
	}

	// Deserialize back to a measurement
	current2, err := UnmarshalElectricCurrent(data)
	if err != nil {
		t.Fatalf("Failed to unmarshal electric current: %v", err)
	}

	// Verify the measurement
	if !current.Equal(current2) {
		t.Errorf("Round-trip serialization failed: got %v, expected %v", current2, current)
	}
}
