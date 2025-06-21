package unit

import (
	"testing"
)

func TestAccelerationConversion(t *testing.T) {
	// Create an acceleration in m/s²
	accel := NewAcceleration(9.8, Acceleration.MetersPerSecondSquared)
	accelG := accel.ConvertTo(Acceleration.G)

	// Define test cases
	testCases := []struct {
		name           string
		input          Quantity[AccelerationUnit]
		targetUnit     AccelerationUnit
		expectedValue  float64
		expectedSymbol string
	}{
		{
			name:           "m/s² to g",
			input:          accel,
			targetUnit:     Acceleration.G,
			expectedValue:  0.9993218887183698,
			expectedSymbol: "g",
		},
		{
			name:           "m/s² to ft/s²",
			input:          accel,
			targetUnit:     Acceleration.FeetPerSecondSquared,
			expectedValue:  32.15223097112861,
			expectedSymbol: "ft/s²",
		},
		{
			name:           "g to m/s²",
			input:          accelG,
			targetUnit:     Acceleration.MetersPerSecondSquared,
			expectedValue:  9.8,
			expectedSymbol: "m/s²",
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.input.ConvertTo(tc.targetUnit)

			if !approxEqual(result.Value, tc.expectedValue) {
				t.Errorf("Conversion failed: got %g %s, expected %g %s",
					result.Value, result.Unit.Symbol(), tc.expectedValue, tc.expectedSymbol)
			}

			if result.Unit.Symbol() != tc.expectedSymbol {
				t.Errorf("Unit symbol mismatch: got %s, expected %s",
					result.Unit.Symbol(), tc.expectedSymbol)
			}
		})
	}
}

func TestAccelerationArithmetic(t *testing.T) {
	// Create two accelerations
	accel1 := NewAcceleration(9.8, Acceleration.MetersPerSecondSquared)
	accel2 := NewAcceleration(1.0, Acceleration.G)

	// Define test cases
	testCases := []struct {
		name           string
		operation      string
		result         Quantity[AccelerationUnit]
		expectedValue  float64
		expectedSymbol string
	}{
		{
			name:           "Addition",
			operation:      "9.8 m/s² + 1.0 g",
			result:         accel1.Add(accel2),
			expectedValue:  19.60665,
			expectedSymbol: "m/s²",
		},
		{
			name:           "Subtraction",
			operation:      "9.8 m/s² - 1.0 g",
			result:         accel1.Subtract(accel2),
			expectedValue:  -0.00665,
			expectedSymbol: "m/s²",
		},
		{
			name:           "Multiply by scalar",
			operation:      "9.8 m/s² * 2",
			result:         accel1.MultiplyByScalar(2.0),
			expectedValue:  19.6,
			expectedSymbol: "m/s²",
		},
		{
			name:           "Divide by scalar",
			operation:      "9.8 m/s² / 2",
			result:         accel1.DivideByScalar(2.0),
			expectedValue:  4.9,
			expectedSymbol: "m/s²",
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if !approxEqual(tc.result.Value, tc.expectedValue) {
				t.Errorf("%s failed: got %g %s, expected %g %s",
					tc.operation, tc.result.Value, tc.result.Unit.Symbol(),
					tc.expectedValue, tc.expectedSymbol)
			}

			if tc.result.Unit.Symbol() != tc.expectedSymbol {
				t.Errorf("Unit symbol mismatch: got %s, expected %s",
					tc.result.Unit.Symbol(), tc.expectedSymbol)
			}
		})
	}
}

func TestAccelerationParsing(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name           string
		input          string
		expectedValue  float64
		expectedUnit   AccelerationUnit
		expectedSymbol string
	}{
		{
			name:           "Parse m/s²",
			input:          "9.8 m/s²",
			expectedValue:  9.8,
			expectedUnit:   Acceleration.MetersPerSecondSquared,
			expectedSymbol: "m/s²",
		},
		{
			name:           "Parse g",
			input:          "1 g",
			expectedValue:  1.0,
			expectedUnit:   Acceleration.G,
			expectedSymbol: "g",
		},
		{
			name:           "Parse ft/s²",
			input:          "32.2 ft/s²",
			expectedValue:  32.2,
			expectedUnit:   Acceleration.FeetPerSecondSquared,
			expectedSymbol: "ft/s²",
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseAcceleration(tc.input)
			if err != nil {
				t.Fatalf("Failed to parse acceleration '%s': %v", tc.input, err)
			}

			if !approxEqual(result.Value, tc.expectedValue) {
				t.Errorf("Parsed value incorrect: got %g, expected %g",
					result.Value, tc.expectedValue)
			}

			if !result.Unit.Equals(tc.expectedUnit) {
				t.Errorf("Parsed unit incorrect: got %s, expected %s",
					result.Unit.Symbol(), tc.expectedSymbol)
			}
		})
	}
}

func TestAccelerationSerialization(t *testing.T) {
	// Define test cases for serialization
	testCases := []struct {
		name  string
		input Quantity[AccelerationUnit]
	}{
		{
			name:  "Serialize and deserialize m/s²",
			input: NewAcceleration(9.8, Acceleration.MetersPerSecondSquared),
		},
		{
			name:  "Serialize and deserialize g",
			input: NewAcceleration(1.0, Acceleration.G),
		},
		{
			name:  "Serialize and deserialize ft/s²",
			input: NewAcceleration(32.2, Acceleration.FeetPerSecondSquared),
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Serialize to JSON
			data, err := MarshalAcceleration(tc.input)
			if err != nil {
				t.Fatalf("Failed to marshal acceleration: %v", err)
			}

			// Deserialize back to a measurement
			result, err := UnmarshalAcceleration(data)
			if err != nil {
				t.Fatalf("Failed to unmarshal acceleration: %v", err)
			}

			// Verify the measurement
			if !tc.input.Equal(result) {
				t.Errorf("Round-trip serialization failed: got %v, expected %v", result, tc.input)
			}

			// Additional check for value and unit separately for better error messages
			if !approxEqual(result.Value, tc.input.Value) {
				t.Errorf("Value mismatch after serialization: got %g, expected %g",
					result.Value, tc.input.Value)
			}

			if !result.Unit.Equals(tc.input.Unit) {
				t.Errorf("Unit mismatch after serialization: got %s, expected %s",
					result.Unit.Symbol(), tc.input.Unit.Symbol())
			}
		})
	}
}
