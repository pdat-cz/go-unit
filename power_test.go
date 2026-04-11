package unit

import (
	"testing"
)

func TestPowerConversion(t *testing.T) {
	// Create a power in Watts
	powerW := NewPower(1000.0, Power.Watt)

	// Define test cases
	testCases := []struct {
		name           string
		input          Quantity[PowerUnit]
		targetUnit     PowerUnit
		expectedValue  float64
		expectedSymbol string
	}{
		{
			name:           "Watts to Kilowatts",
			input:          powerW,
			targetUnit:     Power.Kilowatt,
			expectedValue:  1.0,
			expectedSymbol: "kW",
		},
		{
			name:           "Watts to BTU/h",
			input:          powerW,
			targetUnit:     Power.BTUPerHour,
			expectedValue:  3412.1416,
			expectedSymbol: "BTU/h",
		},
		{
			name:           "Kilowatts to Watts",
			input:          NewPower(2.5, Power.Kilowatt),
			targetUnit:     Power.Watt,
			expectedValue:  2500.0,
			expectedSymbol: "W",
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

func TestPowerArithmetic(t *testing.T) {
	// Create two powers
	power1 := NewPower(500.0, Power.Watt)
	power2 := NewPower(1.0, Power.Kilowatt)

	// Define test cases
	testCases := []struct {
		name           string
		operation      string
		result         Quantity[PowerUnit]
		expectedValue  float64
		expectedSymbol string
	}{
		{
			name:           "Addition",
			operation:      "500 W + 1 kW",
			result:         power1.Add(power2),
			expectedValue:  1500.0,
			expectedSymbol: "W",
		},
		{
			name:           "Subtraction",
			operation:      "1 kW - 500 W",
			result:         power2.Subtract(power1),
			expectedValue:  0.5,
			expectedSymbol: "kW",
		},
		{
			name:           "Multiply by scalar",
			operation:      "500 W * 3",
			result:         power1.MultiplyByScalar(3.0),
			expectedValue:  1500.0,
			expectedSymbol: "W",
		},
		{
			name:           "Divide by scalar",
			operation:      "500 W / 2",
			result:         power1.DivideByScalar(2.0),
			expectedValue:  250.0,
			expectedSymbol: "W",
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

func TestPowerParsing(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name           string
		input          string
		expectedValue  float64
		expectedUnit   PowerUnit
		expectedSymbol string
	}{
		{
			name:           "Parse Watts",
			input:          "1000 W",
			expectedValue:  1000.0,
			expectedUnit:   Power.Watt,
			expectedSymbol: "W",
		},
		{
			name:           "Parse Kilowatts",
			input:          "2.5 kW",
			expectedValue:  2.5,
			expectedUnit:   Power.Kilowatt,
			expectedSymbol: "kW",
		},
		{
			name:           "Parse BTU/h",
			input:          "5000 BTU/h",
			expectedValue:  5000.0,
			expectedUnit:   Power.BTUPerHour,
			expectedSymbol: "BTU/h",
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParsePower(tc.input)
			if err != nil {
				t.Fatalf("Failed to parse power '%s': %v", tc.input, err)
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

func TestPowerUnitSerialization(t *testing.T) {
	// Define test cases for serialization
	testCases := []struct {
		name  string
		input Quantity[PowerUnit]
	}{
		{
			name:  "Serialize and deserialize Watts",
			input: NewPower(1000.0, Power.Watt),
		},
		{
			name:  "Serialize and deserialize Kilowatts",
			input: NewPower(2.5, Power.Kilowatt),
		},
		{
			name:  "Serialize and deserialize BTU/h",
			input: NewPower(5000.0, Power.BTUPerHour),
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := MarshalPower(tc.input)
			if err != nil {
				t.Fatalf("Failed to marshal power: %v", err)
			}

			result, err := UnmarshalPower(data)
			if err != nil {
				t.Fatalf("Failed to unmarshal power: %v", err)
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
