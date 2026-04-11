package unit

import (
	"testing"
)

func TestTemperatureUnitConversion(t *testing.T) {
	// Create a temperature in Celsius
	tempC := NewTemperature(100.0, Temperature.Celsius)

	// Define test cases
	testCases := []struct {
		name           string
		input          Quantity[TemperatureUnit]
		targetUnit     TemperatureUnit
		expectedValue  float64
		expectedSymbol string
	}{
		{
			name:           "Celsius to Fahrenheit",
			input:          tempC,
			targetUnit:     Temperature.Fahrenheit,
			expectedValue:  210.6,
			expectedSymbol: "°F",
		},
		{
			name:           "Celsius to Kelvin",
			input:          tempC,
			targetUnit:     Temperature.Kelvin,
			expectedValue:  373.15,
			expectedSymbol: "K",
		},
		{
			name:           "Fahrenheit to Celsius",
			input:          NewTemperature(32.0, Temperature.Fahrenheit),
			targetUnit:     Temperature.Celsius,
			expectedValue:  0.7778,
			expectedSymbol: "°C",
		},
		{
			name:           "Kelvin to Celsius",
			input:          NewTemperature(0.0, Temperature.Kelvin),
			targetUnit:     Temperature.Celsius,
			expectedValue:  -273.15,
			expectedSymbol: "°C",
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

func TestTemperatureUnitArithmetic(t *testing.T) {
	// Create two temperatures
	temp1 := NewTemperature(20.0, Temperature.Celsius)
	temp2 := NewTemperature(68.0, Temperature.Fahrenheit)

	// Define test cases
	testCases := []struct {
		name           string
		operation      string
		result         Quantity[TemperatureUnit]
		expectedValue  float64
		expectedSymbol string
	}{
		{
			name:           "Addition",
			operation:      "20°C + 68°F",
			result:         temp1.Add(temp2),
			expectedValue:  40.7778,
			expectedSymbol: "°C",
		},
		{
			name:           "Subtraction",
			operation:      "20°C - 68°F",
			result:         temp1.Subtract(temp2),
			expectedValue:  -0.7778,
			expectedSymbol: "°C",
		},
		{
			name:           "Multiply by scalar",
			operation:      "20°C * 2",
			result:         temp1.MultiplyByScalar(2.0),
			expectedValue:  40.0,
			expectedSymbol: "°C",
		},
		{
			name:           "Divide by scalar",
			operation:      "20°C / 2",
			result:         temp1.DivideByScalar(2.0),
			expectedValue:  10.0,
			expectedSymbol: "°C",
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

func TestTemperatureUnitParsing(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name           string
		input          string
		expectedValue  float64
		expectedUnit   TemperatureUnit
		expectedSymbol string
	}{
		{
			name:           "Parse Celsius",
			input:          "25 °C",
			expectedValue:  25.0,
			expectedUnit:   Temperature.Celsius,
			expectedSymbol: "°C",
		},
		{
			name:           "Parse Fahrenheit",
			input:          "98.6 °F",
			expectedValue:  98.6,
			expectedUnit:   Temperature.Fahrenheit,
			expectedSymbol: "°F",
		},
		{
			name:           "Parse Kelvin",
			input:          "300 K",
			expectedValue:  300.0,
			expectedUnit:   Temperature.Kelvin,
			expectedSymbol: "K",
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseTemperature(tc.input)
			if err != nil {
				t.Fatalf("Failed to parse temperature '%s': %v", tc.input, err)
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

func TestTemperatureUnitSerialization(t *testing.T) {
	// Define test cases for serialization
	testCases := []struct {
		name  string
		input Quantity[TemperatureUnit]
	}{
		{
			name:  "Serialize and deserialize Celsius",
			input: NewTemperature(25.0, Temperature.Celsius),
		},
		{
			name:  "Serialize and deserialize Fahrenheit",
			input: NewTemperature(98.6, Temperature.Fahrenheit),
		},
		{
			name:  "Serialize and deserialize Kelvin",
			input: NewTemperature(300.0, Temperature.Kelvin),
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Serialize to JSON
			data, err := MarshalTemperature(tc.input)
			if err != nil {
				t.Fatalf("Failed to marshal temperature: %v", err)
			}

			// Deserialize back to a measurement
			result, err := UnmarshalTemperature(data)
			if err != nil {
				t.Fatalf("Failed to unmarshal temperature: %v", err)
			}

			// Verify the measurement
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
