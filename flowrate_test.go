package unit

import (
	"testing"
)

func TestFlowRateConversion(t *testing.T) {
	// Create a flow rate in cubic meters per hour
	flowM3H := NewFlowRate(10.0, FlowRate.CubicMetersPerHour)

	// Define test cases
	testCases := []struct {
		name           string
		input          Quantity[FlowRateUnit]
		targetUnit     FlowRateUnit
		expectedValue  float64
		expectedSymbol string
	}{
		{
			name:           "m³/h to L/s",
			input:          flowM3H,
			targetUnit:     FlowRate.LitersPerSecond,
			expectedValue:  2.7778,
			expectedSymbol: "L/s",
		},
		{
			name:           "m³/h to CFM",
			input:          flowM3H,
			targetUnit:     FlowRate.CFM,
			expectedValue:  5.8858,
			expectedSymbol: "CFM",
		},
		{
			name:           "CFM to m³/h",
			input:          NewFlowRate(100.0, FlowRate.CFM),
			targetUnit:     FlowRate.CubicMetersPerHour,
			expectedValue:  169.9,
			expectedSymbol: "m³/h",
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

func TestFlowRateArithmetic(t *testing.T) {
	// Create two flow rates
	flow1 := NewFlowRate(10.0, FlowRate.CubicMetersPerHour)
	flow2 := NewFlowRate(1.0, FlowRate.LitersPerSecond)

	// Define test cases
	testCases := []struct {
		name           string
		operation      string
		result         Quantity[FlowRateUnit]
		expectedValue  float64
		expectedSymbol string
	}{
		{
			name:           "Addition",
			operation:      "10 m³/h + 1 L/s",
			result:         flow1.Add(flow2),
			expectedValue:  13.6,
			expectedSymbol: "m³/h",
		},
		{
			name:           "Subtraction",
			operation:      "10 m³/h - 1 L/s",
			result:         flow1.Subtract(flow2),
			expectedValue:  6.4,
			expectedSymbol: "m³/h",
		},
		{
			name:           "Multiply by scalar",
			operation:      "10 m³/h * 2",
			result:         flow1.MultiplyByScalar(2.0),
			expectedValue:  20.0,
			expectedSymbol: "m³/h",
		},
		{
			name:           "Divide by scalar",
			operation:      "10 m³/h / 5",
			result:         flow1.DivideByScalar(5.0),
			expectedValue:  2.0,
			expectedSymbol: "m³/h",
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

func TestFlowRateParsing(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name           string
		input          string
		expectedValue  float64
		expectedUnit   FlowRateUnit
		expectedSymbol string
	}{
		{
			name:           "Parse m³/h",
			input:          "10 m³/h",
			expectedValue:  10.0,
			expectedUnit:   FlowRate.CubicMetersPerHour,
			expectedSymbol: "m³/h",
		},
		{
			name:           "Parse L/s",
			input:          "2.5 L/s",
			expectedValue:  2.5,
			expectedUnit:   FlowRate.LitersPerSecond,
			expectedSymbol: "L/s",
		},
		{
			name:           "Parse CFM",
			input:          "100 CFM",
			expectedValue:  100.0,
			expectedUnit:   FlowRate.CFM,
			expectedSymbol: "CFM",
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseFlowRate(tc.input)
			if err != nil {
				t.Fatalf("Failed to parse flow rate '%s': %v", tc.input, err)
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

func TestFlowRateUnitSerialization(t *testing.T) {
	// Define test cases for serialization
	testCases := []struct {
		name  string
		input Quantity[FlowRateUnit]
	}{
		{
			name:  "Serialize and deserialize m³/h",
			input: NewFlowRate(10.0, FlowRate.CubicMetersPerHour),
		},
		{
			name:  "Serialize and deserialize L/s",
			input: NewFlowRate(2.5, FlowRate.LitersPerSecond),
		},
		{
			name:  "Serialize and deserialize CFM",
			input: NewFlowRate(100.0, FlowRate.CFM),
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := MarshalFlowRate(tc.input)
			if err != nil {
				t.Fatalf("Failed to marshal flow rate: %v", err)
			}

			result, err := UnmarshalFlowRate(data)
			if err != nil {
				t.Fatalf("Failed to unmarshal flow rate: %v", err)
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
