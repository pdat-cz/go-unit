package unit

import (
	"math"
	"testing"
)

func TestElectricPotentialDifferenceConversion(t *testing.T) {
	// Create an electric potential difference in volts
	voltage := NewElectricPotentialDifference(5.0, ElectricPotentialDifference.Volt)

	// Convert to millivolts
	voltageMV := voltage.ConvertTo(ElectricPotentialDifference.Millivolt)

	// Expected: 5.0 V = 5,000 mV
	expected := 5000.0
	if math.Abs(voltageMV.Value-expected) > 0.001 {
		t.Errorf("Electric potential difference conversion failed: got %g mV, expected %g mV", voltageMV.Value, expected)
	}

	// Convert to microvolts
	voltageUV := voltage.ConvertTo(ElectricPotentialDifference.Microvolt)

	// Expected: 5.0 V = 5,000,000 µV
	expected = 5000000.0
	if math.Abs(voltageUV.Value-expected) > 0.001 {
		t.Errorf("Electric potential difference conversion failed: got %g µV, expected %g µV", voltageUV.Value, expected)
	}

	// Convert to kilovolts
	voltageKV := voltage.ConvertTo(ElectricPotentialDifference.Kilovolt)

	// Expected: 5.0 V = 0.005 kV
	expected = 0.005
	if math.Abs(voltageKV.Value-expected) > 0.0001 {
		t.Errorf("Electric potential difference conversion failed: got %g kV, expected %g kV", voltageKV.Value, expected)
	}

	// Convert to megavolts
	voltageMegaV := voltage.ConvertTo(ElectricPotentialDifference.Megavolt)

	// Expected: 5.0 V = 0.000005 MV
	expected = 0.000005
	if math.Abs(voltageMegaV.Value-expected) > 0.0000001 {
		t.Errorf("Electric potential difference conversion failed: got %g MV, expected %g MV", voltageMegaV.Value, expected)
	}

	// Convert back to volts
	voltage2 := voltageMV.ConvertTo(ElectricPotentialDifference.Volt)

	// Should get the original value back
	if math.Abs(voltage2.Value-voltage.Value) > 0.001 {
		t.Errorf("Round-trip conversion failed: got %g V, expected %g V", voltage2.Value, voltage.Value)
	}
}

func TestElectricPotentialDifferenceArithmetic(t *testing.T) {
	// Create two electric potential differences
	voltage1 := NewElectricPotentialDifference(5.0, ElectricPotentialDifference.Volt)
	voltage2 := NewElectricPotentialDifference(500.0, ElectricPotentialDifference.Millivolt)

	// Add them (should convert voltage2 to volts first)
	sum := voltage1.Add(voltage2)

	// Expected: 5.0 V + 0.5 V = 5.5 V
	expected := 5.5
	if math.Abs(sum.Value-expected) > 0.001 {
		t.Errorf("Addition failed: got %g V, expected %g V", sum.Value, expected)
	}

	// Subtract
	diff := voltage1.Subtract(voltage2)

	// Expected: 5.0 V - 0.5 V = 4.5 V
	expected = 4.5
	if math.Abs(diff.Value-expected) > 0.001 {
		t.Errorf("Subtraction failed: got %g V, expected %g V", diff.Value, expected)
	}

	// Multiply by scalar
	doubled := voltage1.MultiplyByScalar(2.0)

	// Expected: 5.0 V * 2 = 10.0 V
	expected = 10.0
	if math.Abs(doubled.Value-expected) > 0.001 {
		t.Errorf("Scalar multiplication failed: got %g V, expected %g V", doubled.Value, expected)
	}

	// Divide by scalar
	halved := voltage1.DivideByScalar(2.0)

	// Expected: 5.0 V / 2 = 2.5 V
	expected = 2.5
	if math.Abs(halved.Value-expected) > 0.001 {
		t.Errorf("Scalar division failed: got %g V, expected %g V", halved.Value, expected)
	}
}

func TestElectricPotentialDifferenceParsing(t *testing.T) {
	// Test electric potential difference parsing
	voltageV, err := ParseElectricPotentialDifference("5.0 V")
	if err != nil {
		t.Errorf("Failed to parse electric potential difference: %v", err)
	}
	if math.Abs(voltageV.Value-5.0) > 0.001 || !voltageV.Unit.Equals(ElectricPotentialDifference.Volt) {
		t.Errorf("Parsed electric potential difference incorrect: got %v, expected 5.0 V", voltageV)
	}

	voltageMV, err := ParseElectricPotentialDifference("500 mV")
	if err != nil {
		t.Errorf("Failed to parse electric potential difference: %v", err)
	}
	if math.Abs(voltageMV.Value-500.0) > 0.001 || !voltageMV.Unit.Equals(ElectricPotentialDifference.Millivolt) {
		t.Errorf("Parsed electric potential difference incorrect: got %v, expected 500 mV", voltageMV)
	}

	voltageKV, err := ParseElectricPotentialDifference("10 kV")
	if err != nil {
		t.Errorf("Failed to parse electric potential difference: %v", err)
	}
	if math.Abs(voltageKV.Value-10.0) > 0.001 || !voltageKV.Unit.Equals(ElectricPotentialDifference.Kilovolt) {
		t.Errorf("Parsed electric potential difference incorrect: got %v, expected 10 kV", voltageKV)
	}
}

func TestElectricPotentialDifferenceSerialization(t *testing.T) {
	// Create an electric potential difference
	voltage := NewElectricPotentialDifference(5.0, ElectricPotentialDifference.Volt)

	// Serialize to JSON
	data, err := MarshalElectricPotentialDifference(voltage)
	if err != nil {
		t.Fatalf("Failed to marshal electric potential difference: %v", err)
	}

	// Deserialize back to a measurement
	voltage2, err := UnmarshalElectricPotentialDifference(data)
	if err != nil {
		t.Fatalf("Failed to unmarshal electric potential difference: %v", err)
	}

	// Verify the measurement
	if !voltage.Equal(voltage2) {
		t.Errorf("Round-trip serialization failed: got %v, expected %v", voltage2, voltage)
	}
}
