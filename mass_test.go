package unit

import (
	"math"
	"testing"
)

func TestMassConversion(t *testing.T) {
	// Create a mass in kilograms
	massKg := NewMass(75.0, Mass.Kilogram)

	// Convert to grams
	massG := massKg.ConvertTo(Mass.Gram)

	// Expected: 75 kg = 75000 g
	expected := 75000.0
	if math.Abs(massG.Value-expected) > 0.001 {
		t.Errorf("Mass conversion failed: got %g g, expected %g g", massG.Value, expected)
	}

	// Convert to pounds
	massLb := massKg.ConvertTo(Mass.Pound)

	// Expected: 75 kg = 165.347 lb
	expected = 165.347
	if math.Abs(massLb.Value-expected) > 0.01 {
		t.Errorf("Mass conversion failed: got %g lb, expected %g lb", massLb.Value, expected)
	}

	// Convert back to kilograms
	massKg2 := massLb.ConvertTo(Mass.Kilogram)

	// Should get the original value back
	if math.Abs(massKg2.Value-massKg.Value) > 0.001 {
		t.Errorf("Round-trip conversion failed: got %g kg, expected %g kg", massKg2.Value, massKg.Value)
	}
}

func TestMassArithmetic(t *testing.T) {
	// Create two masses
	mass1 := NewMass(75.0, Mass.Kilogram)
	mass2 := NewMass(10.0, Mass.Pound)

	// Add them (should convert mass2 to kilograms first)
	sum := mass1.Add(mass2)

	// Expected: 75 kg + 4.536 kg = 79.536 kg
	expected := 79.536
	if math.Abs(sum.Value-expected) > 0.001 {
		t.Errorf("Addition failed: got %g kg, expected %g kg", sum.Value, expected)
	}

	// Subtract
	diff := mass1.Subtract(mass2)

	// Expected: 75 kg - 4.536 kg = 70.464 kg
	expected = 70.464
	if math.Abs(diff.Value-expected) > 0.001 {
		t.Errorf("Subtraction failed: got %g kg, expected %g kg", diff.Value, expected)
	}

	// Multiply by scalar
	doubled := mass1.MultiplyByScalar(2.0)

	// Expected: 75 kg * 2 = 150 kg
	expected = 150.0
	if math.Abs(doubled.Value-expected) > 0.001 {
		t.Errorf("Scalar multiplication failed: got %g kg, expected %g kg", doubled.Value, expected)
	}

	// Divide by scalar
	halved := mass1.DivideByScalar(2.0)

	// Expected: 75 kg / 2 = 37.5 kg
	expected = 37.5
	if math.Abs(halved.Value-expected) > 0.001 {
		t.Errorf("Scalar division failed: got %g kg, expected %g kg", halved.Value, expected)
	}
}

func TestMassParsing(t *testing.T) {
	// Test mass parsing
	massKg, err := ParseMass("75 kg")
	if err != nil {
		t.Errorf("Failed to parse mass: %v", err)
	}
	if math.Abs(massKg.Value-75.0) > 0.001 || !massKg.Unit.Equals(Mass.Kilogram) {
		t.Errorf("Parsed mass incorrect: got %v, expected 75 kg", massKg)
	}

	massLb, err := ParseMass("150 lb")
	if err != nil {
		t.Errorf("Failed to parse mass: %v", err)
	}
	if math.Abs(massLb.Value-150.0) > 0.001 || !massLb.Unit.Equals(Mass.Pound) {
		t.Errorf("Parsed mass incorrect: got %v, expected 150 lb", massLb)
	}

	massG, err := ParseMass("500 grams")
	if err != nil {
		t.Errorf("Failed to parse mass: %v", err)
	}
	if math.Abs(massG.Value-500.0) > 0.001 || !massG.Unit.Equals(Mass.Gram) {
		t.Errorf("Parsed mass incorrect: got %v, expected 500 g", massG)
	}
}

func TestMassSerialization(t *testing.T) {
	// Create a mass
	mass := NewMass(75.0, Mass.Kilogram)

	// Serialize to JSON
	data, err := MarshalMass(mass)
	if err != nil {
		t.Fatalf("Failed to marshal mass: %v", err)
	}

	// Deserialize back to a measurement
	mass2, err := UnmarshalMass(data)
	if err != nil {
		t.Fatalf("Failed to unmarshal mass: %v", err)
	}

	// Verify the measurement
	if !mass.Equal(mass2) {
		t.Errorf("Round-trip serialization failed: got %v, expected %v", mass2, mass)
	}
}
