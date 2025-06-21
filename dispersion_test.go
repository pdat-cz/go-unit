package unit

import (
	"math"
	"testing"
)

func TestDispersionConversion(t *testing.T) {
	// Create a dispersion in ppm
	disp := NewDispersion(100.0, Dispersion.PartsPerMillion)

	// Convert to ppb
	dispPpb := disp.ConvertTo(Dispersion.PartsPerBillion)

	// Expected: 100 ppm = 100,000 ppb
	expected := 100000.0
	if math.Abs(dispPpb.Value-expected) > 0.001 {
		t.Errorf("Dispersion conversion failed: got %g ppb, expected %g ppb", dispPpb.Value, expected)
	}

	// Convert to ppt
	dispPpt := disp.ConvertTo(Dispersion.PartsPerTrillion)

	// Expected: 100 ppm = 100,000,000 ppt
	expected = 100000000.0
	if math.Abs(dispPpt.Value-expected) > 0.001 {
		t.Errorf("Dispersion conversion failed: got %g ppt, expected %g ppt", dispPpt.Value, expected)
	}

	// Convert to percent
	dispPercent := disp.ConvertTo(Dispersion.Percent)

	// Expected: 100 ppm = 0.01 %
	expected = 0.01
	if math.Abs(dispPercent.Value-expected) > 0.0001 {
		t.Errorf("Dispersion conversion failed: got %g %%, expected %g %%", dispPercent.Value, expected)
	}

	// Convert back to ppm
	disp2 := dispPpb.ConvertTo(Dispersion.PartsPerMillion)

	// Should get the original value back
	if math.Abs(disp2.Value-disp.Value) > 0.001 {
		t.Errorf("Round-trip conversion failed: got %g ppm, expected %g ppm", disp2.Value, disp.Value)
	}
}

func TestDispersionArithmetic(t *testing.T) {
	// Create two dispersions
	disp1 := NewDispersion(100.0, Dispersion.PartsPerMillion)
	disp2 := NewDispersion(0.005, Dispersion.Percent)

	// Add them (should convert disp2 to ppm first)
	sum := disp1.Add(disp2)

	// Expected: 100 ppm + 50 ppm = 150 ppm
	expected := 150.0
	if math.Abs(sum.Value-expected) > 0.001 {
		t.Errorf("Addition failed: got %g ppm, expected %g ppm", sum.Value, expected)
	}

	// Subtract
	diff := disp1.Subtract(disp2)

	// Expected: 100 ppm - 50 ppm = 50 ppm
	expected = 50.0
	if math.Abs(diff.Value-expected) > 0.001 {
		t.Errorf("Subtraction failed: got %g ppm, expected %g ppm", diff.Value, expected)
	}

	// Multiply by scalar
	doubled := disp1.MultiplyByScalar(2.0)

	// Expected: 100 ppm * 2 = 200 ppm
	expected = 200.0
	if math.Abs(doubled.Value-expected) > 0.001 {
		t.Errorf("Scalar multiplication failed: got %g ppm, expected %g ppm", doubled.Value, expected)
	}

	// Divide by scalar
	halved := disp1.DivideByScalar(2.0)

	// Expected: 100 ppm / 2 = 50 ppm
	expected = 50.0
	if math.Abs(halved.Value-expected) > 0.001 {
		t.Errorf("Scalar division failed: got %g ppm, expected %g ppm", halved.Value, expected)
	}
}

func TestDispersionParsing(t *testing.T) {
	// Test dispersion parsing
	dispPpm, err := ParseDispersion("100 ppm")
	if err != nil {
		t.Errorf("Failed to parse dispersion: %v", err)
	}
	if math.Abs(dispPpm.Value-100.0) > 0.001 || !dispPpm.Unit.Equals(Dispersion.PartsPerMillion) {
		t.Errorf("Parsed dispersion incorrect: got %v, expected 100 ppm", dispPpm)
	}

	dispPpb, err := ParseDispersion("500 ppb")
	if err != nil {
		t.Errorf("Failed to parse dispersion: %v", err)
	}
	if math.Abs(dispPpb.Value-500.0) > 0.001 || !dispPpb.Unit.Equals(Dispersion.PartsPerBillion) {
		t.Errorf("Parsed dispersion incorrect: got %v, expected 500 ppb", dispPpb)
	}

	dispPercent, err := ParseDispersion("0.01 %")
	if err != nil {
		t.Errorf("Failed to parse dispersion: %v", err)
	}
	if math.Abs(dispPercent.Value-0.01) > 0.001 || !dispPercent.Unit.Equals(Dispersion.Percent) {
		t.Errorf("Parsed dispersion incorrect: got %v, expected 0.01 %%", dispPercent)
	}
}

func TestDispersionSerialization(t *testing.T) {
	// Create a dispersion
	disp := NewDispersion(100.0, Dispersion.PartsPerMillion)

	// Serialize to JSON
	data, err := MarshalDispersion(disp)
	if err != nil {
		t.Fatalf("Failed to marshal dispersion: %v", err)
	}

	// Deserialize back to a measurement
	disp2, err := UnmarshalDispersion(data)
	if err != nil {
		t.Fatalf("Failed to unmarshal dispersion: %v", err)
	}

	// Verify the measurement
	if !disp.Equal(disp2) {
		t.Errorf("Round-trip serialization failed: got %v, expected %v", disp2, disp)
	}
}
