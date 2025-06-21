package unit

import (
	"encoding/json"
	"math"
	"testing"
)

func TestFrequencyConversion(t *testing.T) {
	// Create a frequency in hertz
	freqHz := NewFrequency(1000.0, Frequency.Hertz)

	// Convert to kilohertz
	freqKHz := freqHz.ConvertTo(Frequency.Kilohertz)

	// Expected: 1000 Hz = 1 kHz
	expected := 1.0
	if math.Abs(freqKHz.Value-expected) > 0.0001 {
		t.Errorf("Frequency conversion failed: got %g kHz, expected %g kHz", freqKHz.Value, expected)
	}

	// Convert to megahertz
	freqMHz := freqHz.ConvertTo(Frequency.Megahertz)

	// Expected: 1000 Hz = 0.001 MHz
	expected = 0.001
	if math.Abs(freqMHz.Value-expected) > 0.0001 {
		t.Errorf("Frequency conversion failed: got %g MHz, expected %g MHz", freqMHz.Value, expected)
	}

	// Convert to RPM
	freqRPM := freqHz.ConvertTo(Frequency.RPM)

	// Expected: 1000 Hz = 60000 RPM (1000 * 60)
	expected = 60000.0
	if math.Abs(freqRPM.Value-expected) > 0.001 {
		t.Errorf("Frequency conversion failed: got %g RPM, expected %g RPM", freqRPM.Value, expected)
	}

	// Convert back to hertz
	freqHz2 := freqRPM.ConvertTo(Frequency.Hertz)

	// Should get the original value back
	if math.Abs(freqHz2.Value-freqHz.Value) > 0.001 {
		t.Errorf("Round-trip conversion failed: got %g Hz, expected %g Hz", freqHz2.Value, freqHz.Value)
	}
}

func TestFrequencyArithmetic(t *testing.T) {
	// Create two frequencies
	freq1 := NewFrequency(1000.0, Frequency.Hertz)
	freq2 := NewFrequency(2.0, Frequency.Kilohertz)

	// Add them (should convert freq2 to hertz first)
	sum := freq1.Add(freq2)

	// Expected: 1000 Hz + 2000 Hz = 3000 Hz
	expected := 3000.0
	if math.Abs(sum.Value-expected) > 0.001 {
		t.Errorf("Addition failed: got %g Hz, expected %g Hz", sum.Value, expected)
	}

	// Subtract
	diff := freq2.Subtract(freq1)

	// Expected: 2 kHz - 1000 Hz = 1 kHz
	expected = 1.0
	if math.Abs(diff.Value-expected) > 0.001 {
		t.Errorf("Subtraction failed: got %g kHz, expected %g kHz", diff.Value, expected)
	}

	// Multiply by scalar
	doubled := freq1.MultiplyByScalar(2.0)

	// Expected: 1000 Hz * 2 = 2000 Hz
	expected = 2000.0
	if math.Abs(doubled.Value-expected) > 0.001 {
		t.Errorf("Scalar multiplication failed: got %g Hz, expected %g Hz", doubled.Value, expected)
	}

	// Divide by scalar
	halved := freq1.DivideByScalar(2.0)

	// Expected: 1000 Hz / 2 = 500 Hz
	expected = 500.0
	if math.Abs(halved.Value-expected) > 0.001 {
		t.Errorf("Scalar division failed: got %g Hz, expected %g Hz", halved.Value, expected)
	}
}

func TestFrequencySerialization(t *testing.T) {
	// Create a frequency
	freq := NewFrequency(1000.0, Frequency.Hertz)

	// Serialize to JSON
	data, err := MarshalFrequency(freq)
	if err != nil {
		t.Fatalf("Failed to marshal frequency: %v", err)
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
	if jsonData["unit"].(string) != "Hz" {
		t.Errorf("Expected unit to be Hz, got %s", jsonData["unit"].(string))
	}
	if jsonData["dimension"].(string) != "frequency" {
		t.Errorf("Expected dimension to be frequency, got %s", jsonData["dimension"].(string))
	}

	// Deserialize back to a measurement
	freq2, err := UnmarshalFrequency(data)
	if err != nil {
		t.Fatalf("Failed to unmarshal frequency: %v", err)
	}

	// Verify the measurement
	if !freq.Equal(freq2) {
		t.Errorf("Round-trip serialization failed: got %v, expected %v", freq2, freq)
	}
}
