package unit

import (
	"math"
	"testing"
)

// Define a standard tolerance for floating-point comparisons
const tolerance = 0.0001

// Helper function to check if two float values are approximately equal
func approxEqual(a, b float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestLengthConversion(t *testing.T) {
	// Create a length in meters
	lengthM := NewLength(10.0, Length.Meter)
	lengthFt := lengthM.ConvertTo(Length.Foot)

	// Define test cases
	testCases := []struct {
		name           string
		input          Quantity[LengthUnit]
		targetUnit     LengthUnit
		expectedValue  float64
		expectedSymbol string
	}{
		{
			name:           "Meters to Kilometers",
			input:          lengthM,
			targetUnit:     Length.Kilometer,
			expectedValue:  0.01,
			expectedSymbol: "km",
		},
		{
			name:           "Meters to Feet",
			input:          lengthM,
			targetUnit:     Length.Foot,
			expectedValue:  32.8084,
			expectedSymbol: "ft",
		},
		{
			name:           "Feet to Meters",
			input:          lengthFt,
			targetUnit:     Length.Meter,
			expectedValue:  10.0,
			expectedSymbol: "m",
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

func TestLengthArithmetic(t *testing.T) {
	// Create two lengths
	length1 := NewLength(10.0, Length.Meter)
	length2 := NewLength(5.0, Length.Foot)

	// Define test cases
	testCases := []struct {
		name           string
		operation      string
		result         Quantity[LengthUnit]
		expectedValue  float64
		expectedSymbol string
	}{
		{
			name:           "Addition",
			operation:      "10 m + 5 ft",
			result:         length1.Add(length2),
			expectedValue:  11.524,
			expectedSymbol: "m",
		},
		{
			name:           "Subtraction",
			operation:      "10 m - 5 ft",
			result:         length1.Subtract(length2),
			expectedValue:  8.476,
			expectedSymbol: "m",
		},
		{
			name:           "Multiply by scalar",
			operation:      "10 m * 2",
			result:         length1.MultiplyByScalar(2.0),
			expectedValue:  20.0,
			expectedSymbol: "m",
		},
		{
			name:           "Divide by scalar",
			operation:      "10 m / 2",
			result:         length1.DivideByScalar(2.0),
			expectedValue:  5.0,
			expectedSymbol: "m",
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

func TestLengthParsing(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name           string
		input          string
		expectedValue  float64
		expectedUnit   LengthUnit
		expectedSymbol string
	}{
		{
			name:           "Parse meters",
			input:          "10 m",
			expectedValue:  10.0,
			expectedUnit:   Length.Meter,
			expectedSymbol: "m",
		},
		{
			name:           "Parse kilometers",
			input:          "5.5 km",
			expectedValue:  5.5,
			expectedUnit:   Length.Kilometer,
			expectedSymbol: "km",
		},
		{
			name:           "Parse feet",
			input:          "6 feet",
			expectedValue:  6.0,
			expectedUnit:   Length.Foot,
			expectedSymbol: "ft",
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseLength(tc.input)
			if err != nil {
				t.Fatalf("Failed to parse length '%s': %v", tc.input, err)
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

func TestLengthSerialization(t *testing.T) {
	// Define test cases for serialization
	testCases := []struct {
		name  string
		input Quantity[LengthUnit]
	}{
		{
			name:  "Serialize and deserialize meters",
			input: NewLength(10.0, Length.Meter),
		},
		{
			name:  "Serialize and deserialize kilometers",
			input: NewLength(5.5, Length.Kilometer),
		},
		{
			name:  "Serialize and deserialize feet",
			input: NewLength(6.0, Length.Foot),
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Serialize to JSON
			data, err := MarshalLength(tc.input)
			if err != nil {
				t.Fatalf("Failed to marshal length: %v", err)
			}

			// Deserialize back to a measurement
			result, err := UnmarshalLength(data)
			if err != nil {
				t.Fatalf("Failed to unmarshal length: %v", err)
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
