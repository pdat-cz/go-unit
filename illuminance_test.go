package unit

import (
	"encoding/json"
	"math"
	"testing"
)

func TestIlluminanceConversion(t *testing.T) {
	// Create an illuminance in lux
	illumLux := NewIlluminance(1000.0, Illuminance.Lux)

	// Convert to foot-candles
	illumFc := illumLux.ConvertTo(Illuminance.FootCandle)

	// Expected: 1000 lx ≈ 92.9 fc
	expected := 1000.0 / 10.7639
	if math.Abs(illumFc.Value-expected) > 0.01 {
		t.Errorf("Illuminance conversion failed: got %g fc, expected %g fc", illumFc.Value, expected)
	}

	// Convert to phot
	illumPh := illumLux.ConvertTo(Illuminance.Phot)

	// Expected: 1000 lx = 0.1 ph
	expected = 0.1
	if math.Abs(illumPh.Value-expected) > 0.0001 {
		t.Errorf("Illuminance conversion failed: got %g ph, expected %g ph", illumPh.Value, expected)
	}

	// Convert to nox
	illumNx := illumLux.ConvertTo(Illuminance.Nox)

	// Expected: 1000 lx = 1,000,000 nx
	expected = 1000000.0
	if math.Abs(illumNx.Value-expected) > 0.1 {
		t.Errorf("Illuminance conversion failed: got %g nx, expected %g nx", illumNx.Value, expected)
	}

	// Convert back to lux
	illumLux2 := illumFc.ConvertTo(Illuminance.Lux)

	// Should get the original value back
	if math.Abs(illumLux2.Value-illumLux.Value) > 0.01 {
		t.Errorf("Round-trip conversion failed: got %g lx, expected %g lx", illumLux2.Value, illumLux.Value)
	}
}

func TestIlluminanceArithmetic(t *testing.T) {
	// Create two illuminances
	illum1 := NewIlluminance(1000.0, Illuminance.Lux)
	illum2 := NewIlluminance(10.0, Illuminance.FootCandle)

	// Add them (should convert illum2 to lux first)
	sum := illum1.Add(illum2)

	// Expected: 1000 lx + (10 fc * 10.7639 lx/fc) ≈ 1107.64 lx
	expected := 1000.0 + (10.0 * 10.7639)
	if math.Abs(sum.Value-expected) > 0.01 {
		t.Errorf("Addition failed: got %g lx, expected %g lx", sum.Value, expected)
	}

	// Subtract
	diff := illum1.Subtract(illum2)

	// Expected: 1000 lx - (10 fc * 10.7639 lx/fc) ≈ 892.36 lx
	expected = 1000.0 - (10.0 * 10.7639)
	if math.Abs(diff.Value-expected) > 0.01 {
		t.Errorf("Subtraction failed: got %g lx, expected %g lx", diff.Value, expected)
	}

	// Multiply by scalar
	doubled := illum1.MultiplyByScalar(2.0)

	// Expected: 1000 lx * 2 = 2000 lx
	expected = 2000.0
	if math.Abs(doubled.Value-expected) > 0.001 {
		t.Errorf("Scalar multiplication failed: got %g lx, expected %g lx", doubled.Value, expected)
	}

	// Divide by scalar
	halved := illum1.DivideByScalar(2.0)

	// Expected: 1000 lx / 2 = 500 lx
	expected = 500.0
	if math.Abs(halved.Value-expected) > 0.001 {
		t.Errorf("Scalar division failed: got %g lx, expected %g lx", halved.Value, expected)
	}
}

func TestIlluminanceSerialization(t *testing.T) {
	// Create an illuminance
	illum := NewIlluminance(1000.0, Illuminance.Lux)

	// Serialize to JSON
	data, err := MarshalIlluminance(illum)
	if err != nil {
		t.Fatalf("Failed to marshal illuminance: %v", err)
	}

	// Verify JSON structure
	var jsonData map[string]interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}

	if jsonData["value"].(float64) != 1000.0 {
		t.Errorf("Expected value to be 1000.0, got %f", jsonData["value"].(float64))
	}
	if jsonData["unit"].(string) != "lx" {
		t.Errorf("Expected unit to be lx, got %s", jsonData["unit"].(string))
	}
	if jsonData["dimension"].(string) != "illuminance" {
		t.Errorf("Expected dimension to be illuminance, got %s", jsonData["dimension"].(string))
	}

	// Deserialize back to a measurement
	illum2, err := UnmarshalIlluminance(data)
	if err != nil {
		t.Fatalf("Failed to unmarshal illuminance: %v", err)
	}

	// Verify the measurement
	if !illum.Equal(illum2) {
		t.Errorf("Round-trip serialization failed: got %v, expected %v", illum2, illum)
	}
}
