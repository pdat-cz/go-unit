package unit

import (
	"testing"
)

func TestVolumeConversion(t *testing.T) {
	// Create a volume in liters
	volumeL := NewVolume(1.0, Volume.Liter)

	// Define test cases
	testCases := []struct {
		name           string
		input          Quantity[VolumeUnit]
		targetUnit     VolumeUnit
		expectedValue  float64
		expectedSymbol string
	}{
		{
			name:           "Liters to Cubic Meters",
			input:          volumeL,
			targetUnit:     Volume.CubicMeter,
			expectedValue:  0.001,
			expectedSymbol: "m³",
		},
		{
			name:           "Liters to Gallons",
			input:          volumeL,
			targetUnit:     Volume.Gallon,
			expectedValue:  0.2642,
			expectedSymbol: "gal",
		},
		{
			name:           "Liters to Milliliters",
			input:          volumeL,
			targetUnit:     Volume.Milliliter,
			expectedValue:  1000.0,
			expectedSymbol: "mL",
		},
		{
			name:           "Gallons to Liters",
			input:          NewVolume(1.0, Volume.Gallon),
			targetUnit:     Volume.Liter,
			expectedValue:  3.7854,
			expectedSymbol: "L",
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

func TestVolumeArithmetic(t *testing.T) {
	// Create two volumes
	volume1 := NewVolume(2.0, Volume.Liter)
	volume2 := NewVolume(500.0, Volume.Milliliter)

	// Define test cases
	testCases := []struct {
		name           string
		operation      string
		result         Quantity[VolumeUnit]
		expectedValue  float64
		expectedSymbol string
	}{
		{
			name:           "Addition",
			operation:      "2 L + 500 mL",
			result:         volume1.Add(volume2),
			expectedValue:  2.5,
			expectedSymbol: "L",
		},
		{
			name:           "Subtraction",
			operation:      "2 L - 500 mL",
			result:         volume1.Subtract(volume2),
			expectedValue:  1.5,
			expectedSymbol: "L",
		},
		{
			name:           "Multiply by scalar",
			operation:      "2 L * 3",
			result:         volume1.MultiplyByScalar(3.0),
			expectedValue:  6.0,
			expectedSymbol: "L",
		},
		{
			name:           "Divide by scalar",
			operation:      "2 L / 4",
			result:         volume1.DivideByScalar(4.0),
			expectedValue:  0.5,
			expectedSymbol: "L",
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

func TestVolumeParsing(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name           string
		input          string
		expectedValue  float64
		expectedUnit   VolumeUnit
		expectedSymbol string
	}{
		{
			name:           "Parse liters",
			input:          "5 L",
			expectedValue:  5.0,
			expectedUnit:   Volume.Liter,
			expectedSymbol: "L",
		},
		{
			name:           "Parse gallons",
			input:          "2.5 gal",
			expectedValue:  2.5,
			expectedUnit:   Volume.Gallon,
			expectedSymbol: "gal",
		},
		{
			name:           "Parse cubic meters",
			input:          "1 m³",
			expectedValue:  1.0,
			expectedUnit:   Volume.CubicMeter,
			expectedSymbol: "m³",
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseVolume(tc.input)
			if err != nil {
				t.Fatalf("Failed to parse volume '%s': %v", tc.input, err)
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

func TestVolumeSerialization(t *testing.T) {
	// Define test cases for serialization
	testCases := []struct {
		name  string
		input Quantity[VolumeUnit]
	}{
		{
			name:  "Serialize and deserialize liters",
			input: NewVolume(5.0, Volume.Liter),
		},
		{
			name:  "Serialize and deserialize gallons",
			input: NewVolume(2.5, Volume.Gallon),
		},
		{
			name:  "Serialize and deserialize cubic meters",
			input: NewVolume(1.0, Volume.CubicMeter),
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := MarshalVolume(tc.input)
			if err != nil {
				t.Fatalf("Failed to marshal volume: %v", err)
			}

			result, err := UnmarshalVolume(data)
			if err != nil {
				t.Fatalf("Failed to unmarshal volume: %v", err)
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
