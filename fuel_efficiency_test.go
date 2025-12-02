package unit

import (
	"encoding/json"
	"math"
	"testing"
)

func TestFuelEfficiencyConversion(t *testing.T) {
	// Create a fuel efficiency in km/L
	feKmL := NewFuelEfficiency(10.0, FuelEfficiency.KilometersPerLiter)

	// Convert to miles per gallon
	feMpg := feKmL.ConvertTo(FuelEfficiency.MilesPerGallon)

	// Expected: 10 km/L ≈ 23.52 mpg (10 / 0.425144)
	expected := 10.0 / 0.425144
	if math.Abs(feMpg.Value-expected) > 0.01 {
		t.Errorf("Fuel efficiency conversion failed: got %g mpg, expected %g mpg", feMpg.Value, expected)
	}

	// Convert to liters per 100 kilometers
	feL100km := feKmL.ConvertTo(FuelEfficiency.LitersPer100Kilometers)

	// Expected: 10 km/L = 10 L/100km
	expected = 10.0
	if math.Abs(feL100km.Value-expected) > 0.0001 {
		t.Errorf("Fuel efficiency conversion failed: got %g L/100km, expected %g L/100km", feL100km.Value, expected)
	}

	// Convert back to km/L
	feKmL2 := feMpg.ConvertTo(FuelEfficiency.KilometersPerLiter)

	// Should get the original value back
	if math.Abs(feKmL2.Value-feKmL.Value) > 0.01 {
		t.Errorf("Round-trip conversion failed: got %g km/L, expected %g km/L", feKmL2.Value, feKmL.Value)
	}

	// Test inverse relationship: 5 L/100km = 20 km/L
	feL100km = NewFuelEfficiency(5.0, FuelEfficiency.LitersPer100Kilometers)
	feKmL = feL100km.ConvertTo(FuelEfficiency.KilometersPerLiter)
	expected = 20.0
	if math.Abs(feKmL.Value-expected) > 0.0001 {
		t.Errorf("Inverse conversion failed: got %g km/L, expected %g km/L", feKmL.Value, expected)
	}
}

func TestFuelEfficiencyArithmetic(t *testing.T) {
	// Create two fuel efficiencies
	fe1 := NewFuelEfficiency(10.0, FuelEfficiency.KilometersPerLiter)
	fe2 := NewFuelEfficiency(20.0, FuelEfficiency.MilesPerGallon)

	// Add them (should convert fe2 to km/L first)
	sum := fe1.Add(fe2)

	// Expected: 10 km/L + (20 mpg * 0.425144 km/L/mpg) ≈ 18.5 km/L
	expected := 10.0 + (20.0 * 0.425144)
	if math.Abs(sum.Value-expected) > 0.01 {
		t.Errorf("Addition failed: got %g km/L, expected %g km/L", sum.Value, expected)
	}

	// Subtract
	diff := fe1.Subtract(fe2)

	// Expected: 10 km/L - (20 mpg * 0.425144 km/L/mpg) ≈ 1.5 km/L
	expected = 10.0 - (20.0 * 0.425144)
	if math.Abs(diff.Value-expected) > 0.01 {
		t.Errorf("Subtraction failed: got %g km/L, expected %g km/L", diff.Value, expected)
	}

	// Multiply by scalar
	doubled := fe1.MultiplyByScalar(2.0)

	// Expected: 10 km/L * 2 = 20 km/L
	expected = 20.0
	if math.Abs(doubled.Value-expected) > 0.001 {
		t.Errorf("Scalar multiplication failed: got %g km/L, expected %g km/L", doubled.Value, expected)
	}

	// Divide by scalar
	halved := fe1.DivideByScalar(2.0)

	// Expected: 10 km/L / 2 = 5 km/L
	expected = 5.0
	if math.Abs(halved.Value-expected) > 0.001 {
		t.Errorf("Scalar division failed: got %g km/L, expected %g km/L", halved.Value, expected)
	}

	// Test arithmetic with L/100km
	fe3 := NewFuelEfficiency(5.0, FuelEfficiency.LitersPer100Kilometers)
	fe4 := NewFuelEfficiency(10.0, FuelEfficiency.KilometersPerLiter)

	// Add them (should convert fe4 to L/100km first)
	sum = fe3.Add(fe4)

	// Expected: 5 L/100km + (100/10 L/100km) = 15 L/100km
	expected = 15.0
	if math.Abs(sum.Value-expected) > 0.01 {
		t.Errorf("Addition with L/100km failed: got %g L/100km, expected %g L/100km", sum.Value, expected)
	}
}

func TestFuelEfficiencySerialization(t *testing.T) {
	// Create a fuel efficiency
	fe := NewFuelEfficiency(10.0, FuelEfficiency.KilometersPerLiter)

	// Serialize to JSON
	data, err := MarshalFuelEfficiency(fe)
	if err != nil {
		t.Fatalf("Failed to marshal fuel efficiency: %v", err)
	}

	// Verify JSON structure
	var jsonData map[string]interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}

	if jsonData["value"].(float64) != 10.0 {
		t.Errorf("Expected value to be 10.0, got %f", jsonData["value"].(float64))
	}
	unitData := jsonData["unit"].(map[string]interface{})
	if unitData["symbol"].(string) != "km/L" {
		t.Errorf("Expected unit symbol to be km/L, got %s", unitData["symbol"].(string))
	}
	if unitData["name"].(string) != "Kilometers per Liter" {
		t.Errorf("Expected unit name to be Kilometers per Liter, got %s", unitData["name"].(string))
	}
	if unitData["dimension"].(string) != "fuel_efficiency" {
		t.Errorf("Expected dimension to be fuel_efficiency, got %s", unitData["dimension"].(string))
	}

	// Deserialize back to a measurement
	fe2, err := UnmarshalFuelEfficiency(data)
	if err != nil {
		t.Fatalf("Failed to unmarshal fuel efficiency: %v", err)
	}

	// Verify the measurement
	if !fe.Equal(fe2) {
		t.Errorf("Round-trip serialization failed: got %v, expected %v", fe2, fe)
	}

	// Test with L/100km
	fe = NewFuelEfficiency(5.0, FuelEfficiency.LitersPer100Kilometers)
	data, err = MarshalFuelEfficiency(fe)
	if err != nil {
		t.Fatalf("Failed to marshal fuel efficiency: %v", err)
	}

	fe2, err = UnmarshalFuelEfficiency(data)
	if err != nil {
		t.Fatalf("Failed to unmarshal fuel efficiency: %v", err)
	}

	if !fe.Equal(fe2) {
		t.Errorf("Round-trip serialization with L/100km failed: got %v, expected %v", fe2, fe)
	}
}
