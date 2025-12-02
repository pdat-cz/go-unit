package unit

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalMeasurementWithGeneralFallback(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name               string
		jsonInput          string
		expectedValue      float64
		expectedUnitSymbol string
		expectedDimension  string
	}{
		{
			name:               "Unknown dimension should fall back to general",
			jsonInput:          `{"value": 42, "unit": {"name": "custom", "symbol": "custom"}, "dimension": "unknown_dimension"}`,
			expectedValue:      42,
			expectedUnitSymbol: "custom",
			expectedDimension:  "general",
		},
		{
			name:               "Known dimension but unknown unit should fall back to general",
			jsonInput:          `{"value": 25, "unit": {"name": "unknown_unit", "symbol": "unknown_unit"}, "dimension": "temperature"}`,
			expectedValue:      25,
			expectedUnitSymbol: "unknown_unit",
			expectedDimension:  "general",
		},
		{
			name:               "Direct general dimension should work normally",
			jsonInput:          `{"value": 100, "unit": {"name": "Unit", "symbol": "unit"}, "dimension": "general"}`,
			expectedValue:      100,
			expectedUnitSymbol: "unit",
			expectedDimension:  "general",
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Unmarshal the JSON
			anyM, err := UnmarshalMeasurement([]byte(tc.jsonInput))
			if err != nil {
				t.Fatalf("Expected successful unmarshal, got error: %v", err)
			}

			// Check dimension
			if anyM.GetDimension() != tc.expectedDimension {
				t.Errorf("Expected dimension to be '%s', got '%s'",
					tc.expectedDimension, anyM.GetDimension())
			}

			// Check if we can get it as a general measurement
			generalM, ok := anyM.AsGeneral()
			if !ok {
				t.Fatal("Expected to be able to get measurement as general type")
			}

			// Check value
			if generalM.Value != tc.expectedValue {
				t.Errorf("Expected value to be %f, got %f",
					tc.expectedValue, generalM.Value)
			}

			// Check unit symbol
			if generalM.Unit.Symbol() != tc.expectedUnitSymbol {
				t.Errorf("Expected unit symbol to be '%s', got '%s'",
					tc.expectedUnitSymbol, generalM.Unit.Symbol())
			}
		})
	}
}

func TestMarshalAndUnmarshalGeneral(t *testing.T) {
	// Define custom unit for testing
	customUnit := NewGeneralUnit("xyz", "Custom XYZ")

	// Define test cases
	testCases := []struct {
		name              string
		input             Quantity[GeneralUnit]
		expectedValue     float64
		expectedSymbol    string
		expectedDimension string
	}{
		{
			name:              "Standard general unit",
			input:             NewGeneral(42, General.Unit),
			expectedValue:     42,
			expectedSymbol:    "unit",
			expectedDimension: "general",
		},
		{
			name:              "Custom general unit",
			input:             NewGeneral(123, customUnit),
			expectedValue:     123,
			expectedSymbol:    "xyz",
			expectedDimension: "general",
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Marshal to JSON
			data, err := MarshalGeneral(tc.input)
			if err != nil {
				t.Fatalf("Failed to marshal general measurement: %v", err)
			}

			// Verify JSON structure
			var jsonM MeasurementJSON
			if err := json.Unmarshal(data, &jsonM); err != nil {
				t.Fatalf("Failed to unmarshal JSON: %v", err)
			}

			// Check JSON structure
			if jsonM.Value != tc.expectedValue {
				t.Errorf("Expected value to be %f, got %f", tc.expectedValue, jsonM.Value)
			}
			if jsonM.Unit.Symbol != tc.expectedSymbol {
				t.Errorf("Expected unit symbol to be '%s', got '%s'", tc.expectedSymbol, jsonM.Unit.Symbol)
			}
			if jsonM.Unit.Dimension != tc.expectedDimension {
				t.Errorf("Expected dimension to be '%s', got '%s'", tc.expectedDimension, jsonM.Unit.Dimension)
			}

			// Unmarshal back to a measurement
			unmarshaledM, err := UnmarshalGeneral(data)
			if err != nil {
				t.Fatalf("Failed to unmarshal general measurement: %v", err)
			}

			// Verify the unmarshaled measurement
			if unmarshaledM.Value != tc.input.Value {
				t.Errorf("Expected value to be %f, got %f", tc.input.Value, unmarshaledM.Value)
			}
			if !unmarshaledM.Unit.Equals(tc.input.Unit) {
				t.Errorf("Expected unit to be %s, got %s",
					tc.input.Unit.Symbol(), unmarshaledM.Unit.Symbol())
			}
		})
	}
}
