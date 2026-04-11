package unit

import (
	"testing"
)

func TestEnergyConversion(t *testing.T) {
	// Create an energy in Joules
	energyJ := NewEnergy(3600000.0, Energy.Joule)

	// Define test cases
	testCases := []struct {
		name           string
		input          Quantity[EnergyUnit]
		targetUnit     EnergyUnit
		expectedValue  float64
		expectedSymbol string
	}{
		{
			name:           "Joules to Kilowatt-hours",
			input:          energyJ,
			targetUnit:     Energy.KilowattHour,
			expectedValue:  1.0,
			expectedSymbol: "kWh",
		},
		{
			name:           "Joules to BTU",
			input:          NewEnergy(1055.06, Energy.Joule),
			targetUnit:     Energy.BTU,
			expectedValue:  1.0,
			expectedSymbol: "BTU",
		},
		{
			name:           "Kilowatt-hours to Joules",
			input:          NewEnergy(1.0, Energy.KilowattHour),
			targetUnit:     Energy.Joule,
			expectedValue:  3600000.0,
			expectedSymbol: "J",
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

func TestEnergyArithmetic(t *testing.T) {
	// Create two energies
	energy1 := NewEnergy(1000.0, Energy.Joule)
	energy2 := NewEnergy(1.0, Energy.KilowattHour)

	// Define test cases
	testCases := []struct {
		name           string
		operation      string
		result         Quantity[EnergyUnit]
		expectedValue  float64
		expectedSymbol string
	}{
		{
			name:           "Addition",
			operation:      "1000 J + 1 kWh",
			result:         energy1.Add(energy2),
			expectedValue:  3601000.0,
			expectedSymbol: "J",
		},
		{
			name:           "Subtraction",
			operation:      "1 kWh - 1000 J",
			result:         energy2.Subtract(energy1),
			expectedValue:  0.9997,
			expectedSymbol: "kWh",
		},
		{
			name:           "Multiply by scalar",
			operation:      "1000 J * 3",
			result:         energy1.MultiplyByScalar(3.0),
			expectedValue:  3000.0,
			expectedSymbol: "J",
		},
		{
			name:           "Divide by scalar",
			operation:      "1000 J / 2",
			result:         energy1.DivideByScalar(2.0),
			expectedValue:  500.0,
			expectedSymbol: "J",
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

func TestEnergyParsing(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name           string
		input          string
		expectedValue  float64
		expectedUnit   EnergyUnit
		expectedSymbol string
	}{
		{
			name:           "Parse Joules",
			input:          "1000 J",
			expectedValue:  1000.0,
			expectedUnit:   Energy.Joule,
			expectedSymbol: "J",
		},
		{
			name:           "Parse Kilowatt-hours",
			input:          "1.5 kWh",
			expectedValue:  1.5,
			expectedUnit:   Energy.KilowattHour,
			expectedSymbol: "kWh",
		},
		{
			name:           "Parse BTU",
			input:          "100 BTU",
			expectedValue:  100.0,
			expectedUnit:   Energy.BTU,
			expectedSymbol: "BTU",
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseEnergy(tc.input)
			if err != nil {
				t.Fatalf("Failed to parse energy '%s': %v", tc.input, err)
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

func TestEnergyUnitSerialization(t *testing.T) {
	// Define test cases for serialization
	testCases := []struct {
		name  string
		input Quantity[EnergyUnit]
	}{
		{
			name:  "Serialize and deserialize Joules",
			input: NewEnergy(1000.0, Energy.Joule),
		},
		{
			name:  "Serialize and deserialize Kilowatt-hours",
			input: NewEnergy(1.5, Energy.KilowattHour),
		},
		{
			name:  "Serialize and deserialize BTU",
			input: NewEnergy(100.0, Energy.BTU),
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := MarshalEnergy(tc.input)
			if err != nil {
				t.Fatalf("Failed to marshal energy: %v", err)
			}

			result, err := UnmarshalEnergy(data)
			if err != nil {
				t.Fatalf("Failed to unmarshal energy: %v", err)
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
