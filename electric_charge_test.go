package unit

import (
	"math"
	"testing"
)

func TestElectricChargeConversion(t *testing.T) {
	// Create an electric charge in coulombs
	charge := NewElectricCharge(1000.0, ElectricCharge.Coulomb)

	// Convert to millicoulombs
	chargeMC := charge.ConvertTo(ElectricCharge.Millicoulomb)

	// Expected: 1000 C = 1,000,000 mC
	expected := 1000000.0
	if math.Abs(chargeMC.Value-expected) > 0.001 {
		t.Errorf("Electric charge conversion failed: got %g mC, expected %g mC", chargeMC.Value, expected)
	}

	// Convert to microcoulombs
	chargeUC := charge.ConvertTo(ElectricCharge.Microcoulomb)

	// Expected: 1000 C = 1,000,000,000 µC
	expected = 1000000000.0
	if math.Abs(chargeUC.Value-expected) > 0.001 {
		t.Errorf("Electric charge conversion failed: got %g µC, expected %g µC", chargeUC.Value, expected)
	}

	// Convert to ampere-hours
	chargeAh := charge.ConvertTo(ElectricCharge.Ampere_Hour)

	// Expected: 1000 C = 0.2778 Ah (1000 / 3600)
	expected = 0.2777777777777778
	if math.Abs(chargeAh.Value-expected) > 0.0001 {
		t.Errorf("Electric charge conversion failed: got %g Ah, expected %g Ah", chargeAh.Value, expected)
	}

	// Convert to milliampere-hours
	chargeMah := charge.ConvertTo(ElectricCharge.Milliampere_Hour)

	// Expected: 1000 C = 277.78 mAh (1000 / 3.6)
	expected = 277.77777777777777
	if math.Abs(chargeMah.Value-expected) > 0.0001 {
		t.Errorf("Electric charge conversion failed: got %g mAh, expected %g mAh", chargeMah.Value, expected)
	}

	// Convert back to coulombs
	charge2 := chargeAh.ConvertTo(ElectricCharge.Coulomb)

	// Should get the original value back
	if math.Abs(charge2.Value-charge.Value) > 0.001 {
		t.Errorf("Round-trip conversion failed: got %g C, expected %g C", charge2.Value, charge.Value)
	}
}

func TestElectricChargeArithmetic(t *testing.T) {
	// Create two electric charges
	charge1 := NewElectricCharge(1000.0, ElectricCharge.Coulomb)
	charge2 := NewElectricCharge(1.0, ElectricCharge.Ampere_Hour)

	// Add them (should convert charge2 to coulombs first)
	sum := charge1.Add(charge2)

	// Expected: 1000 C + 3600 C = 4600 C
	expected := 4600.0
	if math.Abs(sum.Value-expected) > 0.001 {
		t.Errorf("Addition failed: got %g C, expected %g C", sum.Value, expected)
	}

	// Subtract
	diff := charge1.Subtract(charge2)

	// Expected: 1000 C - 3600 C = -2600 C
	expected = -2600.0
	if math.Abs(diff.Value-expected) > 0.001 {
		t.Errorf("Subtraction failed: got %g C, expected %g C", diff.Value, expected)
	}

	// Multiply by scalar
	doubled := charge1.MultiplyByScalar(2.0)

	// Expected: 1000 C * 2 = 2000 C
	expected = 2000.0
	if math.Abs(doubled.Value-expected) > 0.001 {
		t.Errorf("Scalar multiplication failed: got %g C, expected %g C", doubled.Value, expected)
	}

	// Divide by scalar
	halved := charge1.DivideByScalar(2.0)

	// Expected: 1000 C / 2 = 500 C
	expected = 500.0
	if math.Abs(halved.Value-expected) > 0.001 {
		t.Errorf("Scalar division failed: got %g C, expected %g C", halved.Value, expected)
	}
}

func TestElectricChargeParsing(t *testing.T) {
	// Test electric charge parsing
	chargeC, err := ParseElectricCharge("1000 C")
	if err != nil {
		t.Errorf("Failed to parse electric charge: %v", err)
	}
	if math.Abs(chargeC.Value-1000.0) > 0.001 || !chargeC.Unit.Equals(ElectricCharge.Coulomb) {
		t.Errorf("Parsed electric charge incorrect: got %v, expected 1000 C", chargeC)
	}

	chargeAh, err := ParseElectricCharge("2.5 Ah")
	if err != nil {
		t.Errorf("Failed to parse electric charge: %v", err)
	}
	if math.Abs(chargeAh.Value-2.5) > 0.001 || !chargeAh.Unit.Equals(ElectricCharge.Ampere_Hour) {
		t.Errorf("Parsed electric charge incorrect: got %v, expected 2.5 Ah", chargeAh)
	}

	chargeMah, err := ParseElectricCharge("2500 mAh")
	if err != nil {
		t.Errorf("Failed to parse electric charge: %v", err)
	}
	if math.Abs(chargeMah.Value-2500.0) > 0.001 || !chargeMah.Unit.Equals(ElectricCharge.Milliampere_Hour) {
		t.Errorf("Parsed electric charge incorrect: got %v, expected 2500 mAh", chargeMah)
	}
}

func TestElectricChargeSerialization(t *testing.T) {
	// Create an electric charge
	charge := NewElectricCharge(1000.0, ElectricCharge.Coulomb)

	// Serialize to JSON
	data, err := MarshalElectricCharge(charge)
	if err != nil {
		t.Fatalf("Failed to marshal electric charge: %v", err)
	}

	// Deserialize back to a measurement
	charge2, err := UnmarshalElectricCharge(data)
	if err != nil {
		t.Fatalf("Failed to unmarshal electric charge: %v", err)
	}

	// Verify the measurement
	if !charge.Equal(charge2) {
		t.Errorf("Round-trip serialization failed: got %v, expected %v", charge2, charge)
	}
}
