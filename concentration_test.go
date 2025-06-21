package unit

import (
	"math"
	"testing"
)

func TestConcentrationConversion(t *testing.T) {
	// Create a concentration in g/L
	conc := NewConcentration(5.0, Concentration.GramsPerLiter)

	// Convert to mg/L
	concMg := conc.ConvertTo(Concentration.MilligramsPerLiter)

	// Expected: 5 g/L = 5000 mg/L
	expected := 5000.0
	if math.Abs(concMg.Value-expected) > 0.001 {
		t.Errorf("Concentration conversion failed: got %g mg/L, expected %g mg/L", concMg.Value, expected)
	}

	// Convert to ppm
	concPpm := conc.ConvertTo(Concentration.PartsPerMillion)

	// Expected: 5 g/L = 5000 ppm
	expected = 5000.0
	if math.Abs(concPpm.Value-expected) > 0.001 {
		t.Errorf("Concentration conversion failed: got %g ppm, expected %g ppm", concPpm.Value, expected)
	}

	// Convert to ppb
	concPpb := conc.ConvertTo(Concentration.PartsPerBillion)

	// Expected: 5 g/L = 5,000,000 ppb
	expected = 5000000.0
	if math.Abs(concPpb.Value-expected) > 0.001 {
		t.Errorf("Concentration conversion failed: got %g ppb, expected %g ppb", concPpb.Value, expected)
	}

	// Convert back to g/L
	conc2 := concMg.ConvertTo(Concentration.GramsPerLiter)

	// Should get the original value back
	if math.Abs(conc2.Value-conc.Value) > 0.001 {
		t.Errorf("Round-trip conversion failed: got %g g/L, expected %g g/L", conc2.Value, conc.Value)
	}
}

func TestConcentrationArithmetic(t *testing.T) {
	// Create two concentrations
	conc1 := NewConcentration(5.0, Concentration.GramsPerLiter)
	conc2 := NewConcentration(2000.0, Concentration.MilligramsPerLiter)

	// Add them (should convert conc2 to g/L first)
	sum := conc1.Add(conc2)

	// Expected: 5 g/L + 2 g/L = 7 g/L
	expected := 7.0
	if math.Abs(sum.Value-expected) > 0.001 {
		t.Errorf("Addition failed: got %g g/L, expected %g g/L", sum.Value, expected)
	}

	// Subtract
	diff := conc1.Subtract(conc2)

	// Expected: 5 g/L - 2 g/L = 3 g/L
	expected = 3.0
	if math.Abs(diff.Value-expected) > 0.001 {
		t.Errorf("Subtraction failed: got %g g/L, expected %g g/L", diff.Value, expected)
	}

	// Multiply by scalar
	doubled := conc1.MultiplyByScalar(2.0)

	// Expected: 5 g/L * 2 = 10 g/L
	expected = 10.0
	if math.Abs(doubled.Value-expected) > 0.001 {
		t.Errorf("Scalar multiplication failed: got %g g/L, expected %g g/L", doubled.Value, expected)
	}

	// Divide by scalar
	halved := conc1.DivideByScalar(2.0)

	// Expected: 5 g/L / 2 = 2.5 g/L
	expected = 2.5
	if math.Abs(halved.Value-expected) > 0.001 {
		t.Errorf("Scalar division failed: got %g g/L, expected %g g/L", halved.Value, expected)
	}
}

func TestConcentrationParsing(t *testing.T) {
	// Test concentration parsing
	concG, err := ParseConcentration("5 g/L")
	if err != nil {
		t.Errorf("Failed to parse concentration: %v", err)
	}
	if math.Abs(concG.Value-5.0) > 0.001 || !concG.Unit.Equals(Concentration.GramsPerLiter) {
		t.Errorf("Parsed concentration incorrect: got %v, expected 5 g/L", concG)
	}

	concMg, err := ParseConcentration("500 mg/L")
	if err != nil {
		t.Errorf("Failed to parse concentration: %v", err)
	}
	if math.Abs(concMg.Value-500.0) > 0.001 || !concMg.Unit.Equals(Concentration.MilligramsPerLiter) {
		t.Errorf("Parsed concentration incorrect: got %v, expected 500 mg/L", concMg)
	}

	concPpm, err := ParseConcentration("100 ppm")
	if err != nil {
		t.Errorf("Failed to parse concentration: %v", err)
	}
	if math.Abs(concPpm.Value-100.0) > 0.001 || !concPpm.Unit.Equals(Concentration.PartsPerMillion) {
		t.Errorf("Parsed concentration incorrect: got %v, expected 100 ppm", concPpm)
	}
}

func TestConcentrationSerialization(t *testing.T) {
	// Create a concentration
	conc := NewConcentration(5.0, Concentration.GramsPerLiter)

	// Serialize to JSON
	data, err := MarshalConcentration(conc)
	if err != nil {
		t.Fatalf("Failed to marshal concentration: %v", err)
	}

	// Deserialize back to a measurement
	conc2, err := UnmarshalConcentration(data)
	if err != nil {
		t.Fatalf("Failed to unmarshal concentration: %v", err)
	}

	// Verify the measurement
	if !conc.Equal(conc2) {
		t.Errorf("Round-trip serialization failed: got %v, expected %v", conc2, conc)
	}
}
