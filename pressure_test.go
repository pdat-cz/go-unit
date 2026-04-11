package unit

import (
	"testing"
)

func TestPressureUnitConversion(t *testing.T) {
	// Create a pressure in Pascal
	pressurePa := NewPressure(101325.0, Pressure.Pascal)

	// Define test cases
	testCases := []struct {
		name           string
		input          Quantity[PressureUnit]
		targetUnit     PressureUnit
		expectedValue  float64
		expectedSymbol string
	}{
		{
			name:           "Pascal to Kilopascal",
			input:          pressurePa,
			targetUnit:     Pressure.Kilopascal,
			expectedValue:  101.325,
			expectedSymbol: "kPa",
		},
		{
			name:           "Pascal to Bar",
			input:          pressurePa,
			targetUnit:     Pressure.Bar,
			expectedValue:  1.01325,
			expectedSymbol: "bar",
		},
		{
			name:           "Pascal to PSI",
			input:          pressurePa,
			targetUnit:     Pressure.PSI,
			expectedValue:  14.6959,
			expectedSymbol: "psi",
		},
		{
			name:           "PSI to Pascal",
			input:          NewPressure(14.6959, Pressure.PSI),
			targetUnit:     Pressure.Pascal,
			expectedValue:  101324.7035,
			expectedSymbol: "Pa",
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

func TestPressureUnitArithmetic(t *testing.T) {
	// Create two pressures
	pressure1 := NewPressure(100.0, Pressure.Kilopascal)
	pressure2 := NewPressure(14.5, Pressure.PSI)

	// Define test cases
	testCases := []struct {
		name           string
		operation      string
		result         Quantity[PressureUnit]
		expectedValue  float64
		expectedSymbol string
	}{
		{
			name:           "Addition",
			operation:      "100 kPa + 14.5 PSI",
			result:         pressure1.Add(pressure2),
			expectedValue:  199.9740,
			expectedSymbol: "kPa",
		},
		{
			name:           "Subtraction",
			operation:      "100 kPa - 14.5 PSI",
			result:         pressure1.Subtract(pressure2),
			expectedValue:  0.0260,
			expectedSymbol: "kPa",
		},
		{
			name:           "Multiply by scalar",
			operation:      "100 kPa * 2",
			result:         pressure1.MultiplyByScalar(2.0),
			expectedValue:  200.0,
			expectedSymbol: "kPa",
		},
		{
			name:           "Divide by scalar",
			operation:      "100 kPa / 2",
			result:         pressure1.DivideByScalar(2.0),
			expectedValue:  50.0,
			expectedSymbol: "kPa",
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

func TestPressureUnitParsing(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name           string
		input          string
		expectedValue  float64
		expectedUnit   PressureUnit
		expectedSymbol string
	}{
		{
			name:           "Parse Pascal",
			input:          "101325 Pa",
			expectedValue:  101325.0,
			expectedUnit:   Pressure.Pascal,
			expectedSymbol: "Pa",
		},
		{
			name:           "Parse Kilopascal",
			input:          "101.3 kPa",
			expectedValue:  101.3,
			expectedUnit:   Pressure.Kilopascal,
			expectedSymbol: "kPa",
		},
		{
			name:           "Parse PSI",
			input:          "14.7 psi",
			expectedValue:  14.7,
			expectedUnit:   Pressure.PSI,
			expectedSymbol: "psi",
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParsePressure(tc.input)
			if err != nil {
				t.Fatalf("Failed to parse pressure '%s': %v", tc.input, err)
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

func TestPressureUnitSerialization(t *testing.T) {
	// Define test cases for serialization
	testCases := []struct {
		name  string
		input Quantity[PressureUnit]
	}{
		{
			name:  "Serialize and deserialize Pascal",
			input: NewPressure(101325.0, Pressure.Pascal),
		},
		{
			name:  "Serialize and deserialize Kilopascal",
			input: NewPressure(101.3, Pressure.Kilopascal),
		},
		{
			name:  "Serialize and deserialize PSI",
			input: NewPressure(14.7, Pressure.PSI),
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := MarshalPressure(tc.input)
			if err != nil {
				t.Fatalf("Failed to marshal pressure: %v", err)
			}

			result, err := UnmarshalPressure(data)
			if err != nil {
				t.Fatalf("Failed to unmarshal pressure: %v", err)
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
