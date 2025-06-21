package unit

import (
	"math"
	"testing"
)

func TestDurationConversion(t *testing.T) {
	// Create a duration in seconds
	durationS := NewDuration(60.0, Duration.Second)

	// Convert to minutes
	durationMin := durationS.ConvertTo(Duration.Minute)

	// Expected: 60 s = 1 min
	expected := 1.0
	if math.Abs(durationMin.Value-expected) > 0.0001 {
		t.Errorf("Duration conversion failed: got %g min, expected %g min", durationMin.Value, expected)
	}

	// Convert to hours
	durationH := durationS.ConvertTo(Duration.Hour)

	// Expected: 60 s = 0.0166667 h
	expected = 0.0166667
	if math.Abs(durationH.Value-expected) > 0.0001 {
		t.Errorf("Duration conversion failed: got %g h, expected %g h", durationH.Value, expected)
	}

	// Convert to milliseconds
	durationMs := durationS.ConvertTo(Duration.Millisecond)

	// Expected: 60 s = 60000 ms
	expected = 60000.0
	if math.Abs(durationMs.Value-expected) > 0.001 {
		t.Errorf("Duration conversion failed: got %g ms, expected %g ms", durationMs.Value, expected)
	}

	// Convert back to seconds
	durationS2 := durationMin.ConvertTo(Duration.Second)

	// Should get the original value back
	if math.Abs(durationS2.Value-durationS.Value) > 0.001 {
		t.Errorf("Round-trip conversion failed: got %g s, expected %g s", durationS2.Value, durationS.Value)
	}
}

func TestDurationArithmetic(t *testing.T) {
	// Create two durations
	duration1 := NewDuration(60.0, Duration.Second)
	duration2 := NewDuration(2.0, Duration.Minute)

	// Add them (should convert duration2 to seconds first)
	sum := duration1.Add(duration2)

	// Expected: 60 s + 120 s = 180 s
	expected := 180.0
	if math.Abs(sum.Value-expected) > 0.001 {
		t.Errorf("Addition failed: got %g s, expected %g s", sum.Value, expected)
	}

	// Subtract
	diff := duration1.Subtract(duration2)

	// Expected: 60 s - 120 s = -60 s
	expected = -60.0
	if math.Abs(diff.Value-expected) > 0.001 {
		t.Errorf("Subtraction failed: got %g s, expected %g s", diff.Value, expected)
	}

	// Multiply by scalar
	doubled := duration1.MultiplyByScalar(2.0)

	// Expected: 60 s * 2 = 120 s
	expected = 120.0
	if math.Abs(doubled.Value-expected) > 0.001 {
		t.Errorf("Scalar multiplication failed: got %g s, expected %g s", doubled.Value, expected)
	}

	// Divide by scalar
	halved := duration1.DivideByScalar(2.0)

	// Expected: 60 s / 2 = 30 s
	expected = 30.0
	if math.Abs(halved.Value-expected) > 0.001 {
		t.Errorf("Scalar division failed: got %g s, expected %g s", halved.Value, expected)
	}
}

func TestDurationParsing(t *testing.T) {
	// Test duration parsing
	durationS, err := ParseDuration("60 s")
	if err != nil {
		t.Errorf("Failed to parse duration: %v", err)
	}
	if math.Abs(durationS.Value-60.0) > 0.001 || !durationS.Unit.Equals(Duration.Second) {
		t.Errorf("Parsed duration incorrect: got %v, expected 60 s", durationS)
	}

	durationMin, err := ParseDuration("2.5 min")
	if err != nil {
		t.Errorf("Failed to parse duration: %v", err)
	}
	if math.Abs(durationMin.Value-2.5) > 0.001 || !durationMin.Unit.Equals(Duration.Minute) {
		t.Errorf("Parsed duration incorrect: got %v, expected 2.5 min", durationMin)
	}

	durationH, err := ParseDuration("1.5 hours")
	if err != nil {
		t.Errorf("Failed to parse duration: %v", err)
	}
	if math.Abs(durationH.Value-1.5) > 0.001 || !durationH.Unit.Equals(Duration.Hour) {
		t.Errorf("Parsed duration incorrect: got %v, expected 1.5 h", durationH)
	}
}

func TestDurationSerialization(t *testing.T) {
	// Create a duration
	duration := NewDuration(60.0, Duration.Second)

	// Serialize to JSON
	data, err := MarshalDuration(duration)
	if err != nil {
		t.Fatalf("Failed to marshal duration: %v", err)
	}

	// Deserialize back to a measurement
	duration2, err := UnmarshalDuration(data)
	if err != nil {
		t.Fatalf("Failed to unmarshal duration: %v", err)
	}

	// Verify the measurement
	if !duration.Equal(duration2) {
		t.Errorf("Round-trip serialization failed: got %v, expected %v", duration2, duration)
	}
}
