package unit

import (
	"testing"
)

func TestAreaConversion(t *testing.T) {
	// Create an area in square meters
	areaSqM := NewArea(1000.0, Area.SquareMeter)

	// Define test cases
	testCases := []struct {
		name           string
		input          Quantity[AreaUnit]
		targetUnit     AreaUnit
		expectedValue  float64
		expectedSymbol string
	}{
		{
			name:           "Square Meters to Square Feet",
			input:          areaSqM,
			targetUnit:     Area.SquareFoot,
			expectedValue:  10763.9104,
			expectedSymbol: "ft²",
		},
		{
			name:           "Square Meters to Acres",
			input:          areaSqM,
			targetUnit:     Area.Acre,
			expectedValue:  0.2471,
			expectedSymbol: "ac",
		},
		{
			name:           "Square Meters to Hectares",
			input:          areaSqM,
			targetUnit:     Area.Hectare,
			expectedValue:  0.1,
			expectedSymbol: "ha",
		},
		{
			name:           "Acres to Square Meters",
			input:          NewArea(1.0, Area.Acre),
			targetUnit:     Area.SquareMeter,
			expectedValue:  4046.86,
			expectedSymbol: "m²",
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

func TestAreaArithmetic(t *testing.T) {
	// Create two areas
	area1 := NewArea(100.0, Area.SquareMeter)
	area2 := NewArea(500.0, Area.SquareFoot)

	// Define test cases
	testCases := []struct {
		name           string
		operation      string
		result         Quantity[AreaUnit]
		expectedValue  float64
		expectedSymbol string
	}{
		{
			name:           "Addition",
			operation:      "100 m² + 500 ft²",
			result:         area1.Add(area2),
			expectedValue:  146.4515,
			expectedSymbol: "m²",
		},
		{
			name:           "Subtraction",
			operation:      "100 m² - 500 ft²",
			result:         area1.Subtract(area2),
			expectedValue:  53.5485,
			expectedSymbol: "m²",
		},
		{
			name:           "Multiply by scalar",
			operation:      "100 m² * 2",
			result:         area1.MultiplyByScalar(2.0),
			expectedValue:  200.0,
			expectedSymbol: "m²",
		},
		{
			name:           "Divide by scalar",
			operation:      "100 m² / 4",
			result:         area1.DivideByScalar(4.0),
			expectedValue:  25.0,
			expectedSymbol: "m²",
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

func TestAreaParsing(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name           string
		input          string
		expectedValue  float64
		expectedUnit   AreaUnit
		expectedSymbol string
	}{
		{
			name:           "Parse square meters",
			input:          "100 m²",
			expectedValue:  100.0,
			expectedUnit:   Area.SquareMeter,
			expectedSymbol: "m²",
		},
		{
			name:           "Parse acres",
			input:          "5 ac",
			expectedValue:  5.0,
			expectedUnit:   Area.Acre,
			expectedSymbol: "ac",
		},
		{
			name:           "Parse hectares",
			input:          "2.5 ha",
			expectedValue:  2.5,
			expectedUnit:   Area.Hectare,
			expectedSymbol: "ha",
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseArea(tc.input)
			if err != nil {
				t.Fatalf("Failed to parse area '%s': %v", tc.input, err)
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

func TestAreaSerialization(t *testing.T) {
	// Define test cases for serialization
	testCases := []struct {
		name  string
		input Quantity[AreaUnit]
	}{
		{
			name:  "Serialize and deserialize square meters",
			input: NewArea(100.0, Area.SquareMeter),
		},
		{
			name:  "Serialize and deserialize acres",
			input: NewArea(5.0, Area.Acre),
		},
		{
			name:  "Serialize and deserialize hectares",
			input: NewArea(2.5, Area.Hectare),
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := MarshalArea(tc.input)
			if err != nil {
				t.Fatalf("Failed to marshal area: %v", err)
			}

			result, err := UnmarshalArea(data)
			if err != nil {
				t.Fatalf("Failed to unmarshal area: %v", err)
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
