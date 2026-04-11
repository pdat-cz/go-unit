package unit

import (
	"testing"
)

func TestSpeedConversion(t *testing.T) {
	// Create a speed in meters per second
	speedMPS := NewSpeed(10.0, Speed.MetersPerSecond)

	// Define test cases
	testCases := []struct {
		name           string
		input          Quantity[SpeedUnit]
		targetUnit     SpeedUnit
		expectedValue  float64
		expectedSymbol string
	}{
		{
			name:           "m/s to km/h",
			input:          speedMPS,
			targetUnit:     Speed.KilometersPerHour,
			expectedValue:  36.0,
			expectedSymbol: "km/h",
		},
		{
			name:           "m/s to mph",
			input:          speedMPS,
			targetUnit:     Speed.MilesPerHour,
			expectedValue:  22.3694,
			expectedSymbol: "mph",
		},
		{
			name:           "m/s to knots",
			input:          speedMPS,
			targetUnit:     Speed.Knot,
			expectedValue:  19.4386,
			expectedSymbol: "kn",
		},
		{
			name:           "km/h to m/s",
			input:          NewSpeed(100.0, Speed.KilometersPerHour),
			targetUnit:     Speed.MetersPerSecond,
			expectedValue:  27.7778,
			expectedSymbol: "m/s",
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

func TestSpeedArithmetic(t *testing.T) {
	// Create two speeds
	speed1 := NewSpeed(10.0, Speed.MetersPerSecond)
	speed2 := NewSpeed(36.0, Speed.KilometersPerHour)

	// Define test cases
	testCases := []struct {
		name           string
		operation      string
		result         Quantity[SpeedUnit]
		expectedValue  float64
		expectedSymbol string
	}{
		{
			name:           "Addition",
			operation:      "10 m/s + 36 km/h",
			result:         speed1.Add(speed2),
			expectedValue:  20.0,
			expectedSymbol: "m/s",
		},
		{
			name:           "Subtraction",
			operation:      "10 m/s - 36 km/h",
			result:         speed1.Subtract(speed2),
			expectedValue:  0.0,
			expectedSymbol: "m/s",
		},
		{
			name:           "Multiply by scalar",
			operation:      "10 m/s * 3",
			result:         speed1.MultiplyByScalar(3.0),
			expectedValue:  30.0,
			expectedSymbol: "m/s",
		},
		{
			name:           "Divide by scalar",
			operation:      "10 m/s / 2",
			result:         speed1.DivideByScalar(2.0),
			expectedValue:  5.0,
			expectedSymbol: "m/s",
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

func TestSpeedParsing(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name           string
		input          string
		expectedValue  float64
		expectedUnit   SpeedUnit
		expectedSymbol string
	}{
		{
			name:           "Parse m/s",
			input:          "10 m/s",
			expectedValue:  10.0,
			expectedUnit:   Speed.MetersPerSecond,
			expectedSymbol: "m/s",
		},
		{
			name:           "Parse km/h",
			input:          "100 km/h",
			expectedValue:  100.0,
			expectedUnit:   Speed.KilometersPerHour,
			expectedSymbol: "km/h",
		},
		{
			name:           "Parse mph",
			input:          "60 mph",
			expectedValue:  60.0,
			expectedUnit:   Speed.MilesPerHour,
			expectedSymbol: "mph",
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseSpeed(tc.input)
			if err != nil {
				t.Fatalf("Failed to parse speed '%s': %v", tc.input, err)
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

func TestSpeedSerialization(t *testing.T) {
	// Define test cases for serialization
	testCases := []struct {
		name  string
		input Quantity[SpeedUnit]
	}{
		{
			name:  "Serialize and deserialize m/s",
			input: NewSpeed(10.0, Speed.MetersPerSecond),
		},
		{
			name:  "Serialize and deserialize km/h",
			input: NewSpeed(100.0, Speed.KilometersPerHour),
		},
		{
			name:  "Serialize and deserialize mph",
			input: NewSpeed(60.0, Speed.MilesPerHour),
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := MarshalSpeed(tc.input)
			if err != nil {
				t.Fatalf("Failed to marshal speed: %v", err)
			}

			result, err := UnmarshalSpeed(data)
			if err != nil {
				t.Fatalf("Failed to unmarshal speed: %v", err)
			}

			if !tc.input.Equal(result) {
				t.Errorf("Round-trip serialization failed: got %v, expected %v", result, tc.input)
			}

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
