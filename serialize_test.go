package unit

import (
	"encoding/json"
	"testing"
)

func TestTemperatureSerialization(t *testing.T) {
	// Create a temperature measurement
	temp := NewTemperature(25.0, Temperature.Celsius)

	// Serialize to JSON
	data, err := MarshalTemperature(temp)
	if err != nil {
		t.Fatalf("Failed to marshal temperature: %v", err)
	}

	// Verify JSON structure
	var jsonM map[string]interface{}
	if err := json.Unmarshal(data, &jsonM); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	// Check fields
	if jsonM["value"] != float64(25.0) {
		t.Errorf("Expected value 25.0, got %v", jsonM["value"])
	}
	unitData := jsonM["unit"].(map[string]interface{})
	if unitData["symbol"] != "°C" {
		t.Errorf("Expected unit symbol °C, got %v", unitData["symbol"])
	}
	if unitData["name"] != "Celsius" {
		t.Errorf("Expected unit name Celsius, got %v", unitData["name"])
	}
	if jsonM["dimension"] != "temperature" {
		t.Errorf("Expected dimension temperature, got %v", jsonM["dimension"])
	}

	// Deserialize back to a measurement
	temp2, err := UnmarshalTemperature(data)
	if err != nil {
		t.Fatalf("Failed to unmarshal temperature: %v", err)
	}

	// Verify the measurement
	if !temp.Equal(temp2) {
		t.Errorf("Round-trip serialization failed: got %v, expected %v", temp2, temp)
	}
}

func TestPressureSerialization(t *testing.T) {
	// Create a pressure measurement
	pressure := NewPressure(101.325, Pressure.Kilopascal)

	// Serialize to JSON
	data, err := MarshalPressure(pressure)
	if err != nil {
		t.Fatalf("Failed to marshal pressure: %v", err)
	}

	// Verify JSON structure
	var jsonM map[string]interface{}
	if err := json.Unmarshal(data, &jsonM); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	// Check fields
	if jsonM["value"] != float64(101.325) {
		t.Errorf("Expected value 101.325, got %v", jsonM["value"])
	}
	unitData := jsonM["unit"].(map[string]interface{})
	if unitData["symbol"] != "kPa" {
		t.Errorf("Expected unit symbol kPa, got %v", unitData["symbol"])
	}
	if unitData["name"] != "Kilopascal" {
		t.Errorf("Expected unit name Kilopascal, got %v", unitData["name"])
	}
	if jsonM["dimension"] != "pressure" {
		t.Errorf("Expected dimension pressure, got %v", jsonM["dimension"])
	}

	// Deserialize back to a measurement
	pressure2, err := UnmarshalPressure(data)
	if err != nil {
		t.Fatalf("Failed to unmarshal pressure: %v", err)
	}

	// Verify the measurement
	if !pressure.Equal(pressure2) {
		t.Errorf("Round-trip serialization failed: got %v, expected %v", pressure2, pressure)
	}
}

func TestFlowRateSerialization(t *testing.T) {
	// Create a flow rate measurement
	flowRate := NewFlowRate(500.0, FlowRate.CubicMetersPerHour)

	// Serialize to JSON
	data, err := MarshalFlowRate(flowRate)
	if err != nil {
		t.Fatalf("Failed to marshal flow rate: %v", err)
	}

	// Deserialize back to a measurement
	flowRate2, err := UnmarshalFlowRate(data)
	if err != nil {
		t.Fatalf("Failed to unmarshal flow rate: %v", err)
	}

	// Verify the measurement
	if !flowRate.Equal(flowRate2) {
		t.Errorf("Round-trip serialization failed: got %v, expected %v", flowRate2, flowRate)
	}
}

func TestPowerSerialization(t *testing.T) {
	// Create a power measurement
	power := NewPower(1000.0, Power.Watt)

	// Serialize to JSON
	data, err := MarshalPower(power)
	if err != nil {
		t.Fatalf("Failed to marshal power: %v", err)
	}

	// Deserialize back to a measurement
	power2, err := UnmarshalPower(data)
	if err != nil {
		t.Fatalf("Failed to unmarshal power: %v", err)
	}

	// Verify the measurement
	if !power.Equal(power2) {
		t.Errorf("Round-trip serialization failed: got %v, expected %v", power2, power)
	}
}

func TestEnergySerialization(t *testing.T) {
	// Create an energy measurement
	energy := NewEnergy(3600000.0, Energy.Joule)

	// Serialize to JSON
	data, err := MarshalEnergy(energy)
	if err != nil {
		t.Fatalf("Failed to marshal energy: %v", err)
	}

	// Deserialize back to a measurement
	energy2, err := UnmarshalEnergy(data)
	if err != nil {
		t.Fatalf("Failed to unmarshal energy: %v", err)
	}

	// Verify the measurement
	if !energy.Equal(energy2) {
		t.Errorf("Round-trip serialization failed: got %v, expected %v", energy2, energy)
	}
}

func TestInvalidSerialization(t *testing.T) {
	// Test invalid JSON
	_, err := UnmarshalTemperature([]byte(`{"value": "not a number", "unit": {"name": "Celsius", "symbol": "°C"}, "dimension": "temperature"}`))
	if err == nil {
		t.Errorf("Expected error for invalid JSON, but got nil")
	}

	// Test invalid dimension
	_, err = UnmarshalTemperature([]byte(`{"value": 25.0, "unit": {"name": "Celsius", "symbol": "°C"}, "dimension": "pressure"}`))
	if err == nil {
		t.Errorf("Expected error for invalid dimension, but got nil")
	}

	// Test invalid unit
	_, err = UnmarshalTemperature([]byte(`{"value": 25.0, "unit": {"name": "invalid", "symbol": "invalid"}, "dimension": "temperature"}`))
	if err == nil {
		t.Errorf("Expected error for invalid unit, but got nil")
	}
}

func TestGenericUnmarshalMeasurement(t *testing.T) {
	// Test temperature measurement
	tempJSON := []byte(`{"value": 25.0, "unit": {"name": "Celsius", "symbol": "°C"}, "dimension": "temperature"}`)
	anyTemp, err := UnmarshalMeasurement(tempJSON)
	if err != nil {
		t.Fatalf("Failed to unmarshal temperature: %v", err)
	}

	// Check dimension
	if anyTemp.GetDimension() != "temperature" {
		t.Errorf("Expected dimension 'temperature', got '%s'", anyTemp.GetDimension())
	}

	// Convert to specific type
	temp, ok := anyTemp.AsTemperature()
	if !ok {
		t.Fatalf("Failed to convert to Temperature measurement")
	}

	// Verify the measurement
	expectedTemp := NewTemperature(25.0, Temperature.Celsius)
	if !temp.Equal(expectedTemp) {
		t.Errorf("Expected %v, got %v", expectedTemp, temp)
	}

	// Test length measurement
	lengthJSON := []byte(`{"value": 10.0, "unit": {"name": "Meter", "symbol": "m"}, "dimension": "length"}`)
	anyLength, err := UnmarshalMeasurement(lengthJSON)
	if err != nil {
		t.Fatalf("Failed to unmarshal length: %v", err)
	}

	// Check dimension
	if anyLength.GetDimension() != "length" {
		t.Errorf("Expected dimension 'length', got '%s'", anyLength.GetDimension())
	}

	// Convert to specific type
	length, ok := anyLength.AsLength()
	if !ok {
		t.Fatalf("Failed to convert to Length measurement")
	}

	// Verify the measurement
	expectedLength := NewLength(10.0, Length.Meter)
	if !length.Equal(expectedLength) {
		t.Errorf("Expected %v, got %v", expectedLength, length)
	}

	// Test volume measurement
	volumeJSON := []byte(`{"value": 5.0, "unit": {"name": "Liter", "symbol": "L"}, "dimension": "volume"}`)
	anyVolume, err := UnmarshalMeasurement(volumeJSON)
	if err != nil {
		t.Fatalf("Failed to unmarshal volume: %v", err)
	}

	// Check dimension
	if anyVolume.GetDimension() != "volume" {
		t.Errorf("Expected dimension 'volume', got '%s'", anyVolume.GetDimension())
	}

	// Convert to specific type
	volume, ok := anyVolume.AsVolume()
	if !ok {
		t.Fatalf("Failed to convert to Volume measurement")
	}

	// Verify the measurement
	expectedVolume := NewVolume(5.0, Volume.Liter)
	if !volume.Equal(expectedVolume) {
		t.Errorf("Expected %v, got %v", expectedVolume, volume)
	}

	// Test invalid dimension (should now fall back to general unit)
	invalidJSON := []byte(`{"value": 10.0, "unit": {"name": "m", "symbol": "m"}, "dimension": "invalid_dimension"}`)
	anyInvalid, err := UnmarshalMeasurement(invalidJSON)
	if err != nil {
		t.Errorf("Expected successful unmarshal with general fallback, got error: %v", err)
	}

	// Check that it was converted to a general measurement
	if anyInvalid.GetDimension() != "general" {
		t.Errorf("Expected dimension to be 'general', got '%s'", anyInvalid.GetDimension())
	}

	// Check if we can get it as a general measurement
	generalM, ok := anyInvalid.AsGeneral()
	if !ok {
		t.Error("Expected to be able to get measurement as general type")
	}

	// Verify the measurement values
	if generalM.Value != 10.0 {
		t.Errorf("Expected value to be 10.0, got %f", generalM.Value)
	}
	if generalM.Unit.Symbol() != "m" {
		t.Errorf("Expected unit symbol to be 'm', got '%s'", generalM.Unit.Symbol())
	}

	// Test type conversion failure
	// Try to convert a temperature measurement to a length measurement
	_, ok = anyTemp.AsLength()
	if ok {
		t.Errorf("Expected conversion from temperature to length to fail, but it succeeded")
	}
}
