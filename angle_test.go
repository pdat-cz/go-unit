package unit

import (
	"math"
	"testing"
)

func TestAngleConversion(t *testing.T) {
	// Create an angle in degrees
	angleDeg := NewAngle(90.0, Angle.Degree)

	// Convert to radians
	angleRad := angleDeg.ConvertTo(Angle.Radian)

	// Expected: 90° = π/2 rad ≈ 1.5708 rad
	expected := math.Pi / 2
	if math.Abs(angleRad.Value-expected) > 0.0001 {
		t.Errorf("Angle conversion failed: got %g rad, expected %g rad", angleRad.Value, expected)
	}

	// Convert to revolutions
	angleRev := angleDeg.ConvertTo(Angle.Revolution)

	// Expected: 90° = 0.25 rev
	expected = 0.25
	if math.Abs(angleRev.Value-expected) > 0.0001 {
		t.Errorf("Angle conversion failed: got %g rev, expected %g rev", angleRev.Value, expected)
	}

	// Convert to gradians
	angleGrad := angleDeg.ConvertTo(Angle.Gradian)

	// Expected: 90° = 100 grad
	expected = 100.0
	if math.Abs(angleGrad.Value-expected) > 0.001 {
		t.Errorf("Angle conversion failed: got %g grad, expected %g grad", angleGrad.Value, expected)
	}

	// Convert back to degrees
	angleDeg2 := angleRad.ConvertTo(Angle.Degree)

	// Should get the original value back
	if math.Abs(angleDeg2.Value-angleDeg.Value) > 0.001 {
		t.Errorf("Round-trip conversion failed: got %g°, expected %g°", angleDeg2.Value, angleDeg.Value)
	}
}

func TestAngleArithmetic(t *testing.T) {
	// Create two angles
	angle1 := NewAngle(90.0, Angle.Degree)
	angle2 := NewAngle(math.Pi/4, Angle.Radian)

	// Add them (should convert angle2 to degrees first)
	sum := angle1.Add(angle2)

	// Expected: 90° + 45° = 135°
	expected := 135.0
	if math.Abs(sum.Value-expected) > 0.001 {
		t.Errorf("Addition failed: got %g°, expected %g°", sum.Value, expected)
	}

	// Subtract
	diff := angle1.Subtract(angle2)

	// Expected: 90° - 45° = 45°
	expected = 45.0
	if math.Abs(diff.Value-expected) > 0.001 {
		t.Errorf("Subtraction failed: got %g°, expected %g°", diff.Value, expected)
	}

	// Multiply by scalar
	doubled := angle1.MultiplyByScalar(2.0)

	// Expected: 90° * 2 = 180°
	expected = 180.0
	if math.Abs(doubled.Value-expected) > 0.001 {
		t.Errorf("Scalar multiplication failed: got %g°, expected %g°", doubled.Value, expected)
	}

	// Divide by scalar
	halved := angle1.DivideByScalar(2.0)

	// Expected: 90° / 2 = 45°
	expected = 45.0
	if math.Abs(halved.Value-expected) > 0.001 {
		t.Errorf("Scalar division failed: got %g°, expected %g°", halved.Value, expected)
	}
}

func TestAngleParsing(t *testing.T) {
	// Test angle parsing
	angleDeg, err := ParseAngle("90°")
	if err != nil {
		t.Errorf("Failed to parse angle: %v", err)
	}
	if math.Abs(angleDeg.Value-90.0) > 0.001 || !angleDeg.Unit.Equals(Angle.Degree) {
		t.Errorf("Parsed angle incorrect: got %v, expected 90°", angleDeg)
	}

	angleRad, err := ParseAngle("3.14159 rad")
	if err != nil {
		t.Errorf("Failed to parse angle: %v", err)
	}
	if math.Abs(angleRad.Value-3.14159) > 0.001 || !angleRad.Unit.Equals(Angle.Radian) {
		t.Errorf("Parsed angle incorrect: got %v, expected 3.14159 rad", angleRad)
	}

	angleRev, err := ParseAngle("0.5 rev")
	if err != nil {
		t.Errorf("Failed to parse angle: %v", err)
	}
	if math.Abs(angleRev.Value-0.5) > 0.001 || !angleRev.Unit.Equals(Angle.Revolution) {
		t.Errorf("Parsed angle incorrect: got %v, expected 0.5 rev", angleRev)
	}
}

func TestAngleSerialization(t *testing.T) {
	// Create an angle
	angle := NewAngle(90.0, Angle.Degree)

	// Serialize to JSON
	data, err := MarshalAngle(angle)
	if err != nil {
		t.Fatalf("Failed to marshal angle: %v", err)
	}

	// Deserialize back to a measurement
	angle2, err := UnmarshalAngle(data)
	if err != nil {
		t.Fatalf("Failed to unmarshal angle: %v", err)
	}

	// Verify the measurement
	if !angle.Equal(angle2) {
		t.Errorf("Round-trip serialization failed: got %v, expected %v", angle2, angle)
	}
}
