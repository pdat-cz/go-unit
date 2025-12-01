package unit

import (
	"encoding/json"
	"testing"
)

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Celsius", "celsius"},
		{"Fahrenheit", "fahrenheit"},
		{"KilometersPerHour", "kilometers_per_hour"},
		{"MetersPerSecondSquared", "meters_per_second_squared"},
		{"BTU", "b_t_u"},
		{"PSI", "p_s_i"},
		{"LitersPer100Kilometers", "liters_per100_kilometers"},
	}

	for _, tt := range tests {
		result := toSnakeCase(tt.input)
		if result != tt.expected {
			t.Errorf("toSnakeCase(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestUnitKey(t *testing.T) {
	tests := []struct {
		dimension string
		name      string
		expected  string
	}{
		{"temperature", "Celsius", "temperature_celsius"},
		{"pressure", "Kilopascal", "pressure_kilopascal"},
		{"speed", "KilometersPerHour", "speed_kilometers_per_hour"},
	}

	for _, tt := range tests {
		result := unitKey(tt.dimension, tt.name)
		if result != tt.expected {
			t.Errorf("unitKey(%q, %q) = %q, want %q", tt.dimension, tt.name, result, tt.expected)
		}
	}
}

func TestParseUnitKey(t *testing.T) {
	tests := []struct {
		key           string
		wantDimension string
		wantUnit      string
	}{
		{"temperature_celsius", "temperature", "celsius"},
		{"pressure_kilopascal", "pressure", "kilopascal"},
		{"electric_current_milliampere", "electric", "current_milliampere"},
	}

	for _, tt := range tests {
		dim, unit := parseUnitKey(tt.key)
		if dim != tt.wantDimension || unit != tt.wantUnit {
			t.Errorf("parseUnitKey(%q) = (%q, %q), want (%q, %q)",
				tt.key, dim, unit, tt.wantDimension, tt.wantUnit)
		}
	}
}

func TestMarshalCompactTemperature(t *testing.T) {
	temp := NewTemperature(25.0, Temperature.Celsius)

	data, err := MarshalCompactTemperature(temp)
	if err != nil {
		t.Fatalf("MarshalCompactTemperature failed: %v", err)
	}

	var cj CompactJSON
	if err := json.Unmarshal(data, &cj); err != nil {
		t.Fatalf("Failed to unmarshal result: %v", err)
	}

	if cj.Value != 25.0 {
		t.Errorf("Value = %v, want 25.0", cj.Value)
	}
	if cj.Unit != "temperature_celsius" {
		t.Errorf("Unit = %q, want %q", cj.Unit, "temperature_celsius")
	}
	if cj.Symbol != "" {
		t.Errorf("Symbol = %q, want empty", cj.Symbol)
	}
}

func TestMarshalCompactTemperatureWithSymbol(t *testing.T) {
	temp := NewTemperature(25.0, Temperature.Celsius)

	data, err := MarshalCompactTemperatureWithSymbol(temp)
	if err != nil {
		t.Fatalf("MarshalCompactTemperatureWithSymbol failed: %v", err)
	}

	var cj CompactJSON
	if err := json.Unmarshal(data, &cj); err != nil {
		t.Fatalf("Failed to unmarshal result: %v", err)
	}

	if cj.Value != 25.0 {
		t.Errorf("Value = %v, want 25.0", cj.Value)
	}
	if cj.Unit != "temperature_celsius" {
		t.Errorf("Unit = %q, want %q", cj.Unit, "temperature_celsius")
	}
	if cj.Symbol != "°C" {
		t.Errorf("Symbol = %q, want %q", cj.Symbol, "°C")
	}
}

func TestUnmarshalCompactTemperature(t *testing.T) {
	data := []byte(`{"value":25,"unit":"temperature_celsius"}`)

	temp, err := UnmarshalCompactTemperature(data)
	if err != nil {
		t.Fatalf("UnmarshalCompactTemperature failed: %v", err)
	}

	if temp.Value != 25.0 {
		t.Errorf("Value = %v, want 25.0", temp.Value)
	}
	if temp.Unit.Symbol() != "°C" {
		t.Errorf("Unit symbol = %q, want %q", temp.Unit.Symbol(), "°C")
	}
}

func TestCompactRoundTrip(t *testing.T) {
	original := NewTemperature(37.5, Temperature.Fahrenheit)

	// Marshal
	data, err := MarshalCompactTemperature(original)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Unmarshal
	restored, err := UnmarshalCompactTemperature(data)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if !original.Equal(restored) {
		t.Errorf("Round trip failed: original=%v, restored=%v", original, restored)
	}
}

func TestUnmarshalMeasurementAutoDetectCompact(t *testing.T) {
	// Compact format
	compactData := []byte(`{"value":100,"unit":"length_meter"}`)

	am, err := UnmarshalMeasurement(compactData)
	if err != nil {
		t.Fatalf("UnmarshalMeasurement (compact) failed: %v", err)
	}

	if am.GetDimension() != "length" {
		t.Errorf("Dimension = %q, want %q", am.GetDimension(), "length")
	}

	length, ok := am.AsLength()
	if !ok {
		t.Fatal("AsLength() returned false")
	}
	if length.Value != 100.0 {
		t.Errorf("Value = %v, want 100.0", length.Value)
	}
}

func TestUnmarshalMeasurementAutoDetectStandard(t *testing.T) {
	// Standard format
	standardData := []byte(`{"value":100,"unit":{"name":"Meter","symbol":"m"},"dimension":"length"}`)

	am, err := UnmarshalMeasurement(standardData)
	if err != nil {
		t.Fatalf("UnmarshalMeasurement (standard) failed: %v", err)
	}

	if am.GetDimension() != "length" {
		t.Errorf("Dimension = %q, want %q", am.GetDimension(), "length")
	}

	length, ok := am.AsLength()
	if !ok {
		t.Fatal("AsLength() returned false")
	}
	if length.Value != 100.0 {
		t.Errorf("Value = %v, want 100.0", length.Value)
	}
}

func TestIsCompactFormat(t *testing.T) {
	tests := []struct {
		data     string
		expected bool
	}{
		{`{"value":25,"unit":"temperature_celsius"}`, true},
		{`{"value":25,"unit":{"name":"Celsius","symbol":"°C"},"dimension":"temperature"}`, false},
		{`{"value":100,"unit":"length_meter"}`, true},
		{`{"value":100,"unit":{"name":"Meter","symbol":"m"},"dimension":"length"}`, false},
	}

	for _, tt := range tests {
		result := isCompactFormat([]byte(tt.data))
		if result != tt.expected {
			t.Errorf("isCompactFormat(%s) = %v, want %v", tt.data, result, tt.expected)
		}
	}
}

func TestCompactRoundTripAllDimensions(t *testing.T) {
	// Test round-trip for each dimension to verify registries are correct
	t.Run("Temperature", func(t *testing.T) {
		original := NewTemperature(25, Temperature.Celsius)
		data, _ := MarshalCompactTemperature(original)
		restored, err := UnmarshalCompactTemperature(data)
		if err != nil {
			t.Fatalf("Unmarshal failed: %v, data: %s", err, string(data))
		}
		if !original.Equal(restored) {
			t.Errorf("Round trip failed")
		}
	})

	t.Run("Pressure", func(t *testing.T) {
		original := NewPressure(101.3, Pressure.Kilopascal)
		data, _ := MarshalCompactPressure(original)
		restored, err := UnmarshalCompactPressure(data)
		if err != nil {
			t.Fatalf("Unmarshal failed: %v, data: %s", err, string(data))
		}
		if !original.Equal(restored) {
			t.Errorf("Round trip failed")
		}
	})

	t.Run("Length", func(t *testing.T) {
		original := NewLength(100, Length.Meter)
		data, _ := MarshalCompactLength(original)
		restored, err := UnmarshalCompactLength(data)
		if err != nil {
			t.Fatalf("Unmarshal failed: %v, data: %s", err, string(data))
		}
		if !original.Equal(restored) {
			t.Errorf("Round trip failed")
		}
	})

	t.Run("Mass", func(t *testing.T) {
		original := NewMass(75, Mass.Kilogram)
		data, _ := MarshalCompactMass(original)
		restored, err := UnmarshalCompactMass(data)
		if err != nil {
			t.Fatalf("Unmarshal failed: %v, data: %s", err, string(data))
		}
		if !original.Equal(restored) {
			t.Errorf("Round trip failed")
		}
	})

	t.Run("Duration", func(t *testing.T) {
		original := NewDuration(60, Duration.Second)
		data, _ := MarshalCompactDuration(original)
		restored, err := UnmarshalCompactDuration(data)
		if err != nil {
			t.Fatalf("Unmarshal failed: %v, data: %s", err, string(data))
		}
		if !original.Equal(restored) {
			t.Errorf("Round trip failed")
		}
	})

	t.Run("Area", func(t *testing.T) {
		original := NewArea(100, Area.SquareMeter)
		data, _ := MarshalCompactArea(original)
		restored, err := UnmarshalCompactArea(data)
		if err != nil {
			t.Fatalf("Unmarshal failed: %v, data: %s", err, string(data))
		}
		if !original.Equal(restored) {
			t.Errorf("Round trip failed")
		}
	})

	t.Run("Volume", func(t *testing.T) {
		original := NewVolume(1, Volume.Liter)
		data, _ := MarshalCompactVolume(original)
		restored, err := UnmarshalCompactVolume(data)
		if err != nil {
			t.Fatalf("Unmarshal failed: %v, data: %s", err, string(data))
		}
		if !original.Equal(restored) {
			t.Errorf("Round trip failed")
		}
	})

	t.Run("Speed", func(t *testing.T) {
		original := NewSpeed(100, Speed.MetersPerSecond)
		data, _ := MarshalCompactSpeed(original)
		restored, err := UnmarshalCompactSpeed(data)
		if err != nil {
			t.Fatalf("Unmarshal failed: %v, data: %s", err, string(data))
		}
		if !original.Equal(restored) {
			t.Errorf("Round trip failed")
		}
	})

	t.Run("FlowRate", func(t *testing.T) {
		original := NewFlowRate(10, FlowRate.CubicMetersPerHour)
		data, _ := MarshalCompactFlowRate(original)
		restored, err := UnmarshalCompactFlowRate(data)
		if err != nil {
			t.Fatalf("Unmarshal failed: %v, data: %s", err, string(data))
		}
		if !original.Equal(restored) {
			t.Errorf("Round trip failed")
		}
	})

	t.Run("Power", func(t *testing.T) {
		original := NewPower(1000, Power.Watt)
		data, _ := MarshalCompactPower(original)
		restored, err := UnmarshalCompactPower(data)
		if err != nil {
			t.Fatalf("Unmarshal failed: %v, data: %s", err, string(data))
		}
		if !original.Equal(restored) {
			t.Errorf("Round trip failed")
		}
	})

	t.Run("Energy", func(t *testing.T) {
		original := NewEnergy(3600, Energy.Joule)
		data, _ := MarshalCompactEnergy(original)
		restored, err := UnmarshalCompactEnergy(data)
		if err != nil {
			t.Fatalf("Unmarshal failed: %v, data: %s", err, string(data))
		}
		if !original.Equal(restored) {
			t.Errorf("Round trip failed")
		}
	})

	t.Run("Frequency", func(t *testing.T) {
		original := NewFrequency(1000, Frequency.Hertz)
		data, _ := MarshalCompactFrequency(original)
		restored, err := UnmarshalCompactFrequency(data)
		if err != nil {
			t.Fatalf("Unmarshal failed: %v, data: %s", err, string(data))
		}
		if !original.Equal(restored) {
			t.Errorf("Round trip failed")
		}
	})

	t.Run("ElectricCurrent", func(t *testing.T) {
		original := NewElectricCurrent(1, ElectricCurrent.Ampere)
		data, _ := MarshalCompactElectricCurrent(original)
		restored, err := UnmarshalCompactElectricCurrent(data)
		if err != nil {
			t.Fatalf("Unmarshal failed: %v, data: %s", err, string(data))
		}
		if !original.Equal(restored) {
			t.Errorf("Round trip failed")
		}
	})

	t.Run("ElectricPotentialDifference", func(t *testing.T) {
		original := NewElectricPotentialDifference(220, ElectricPotentialDifference.Volt)
		data, _ := MarshalCompactElectricPotentialDifference(original)
		restored, err := UnmarshalCompactElectricPotentialDifference(data)
		if err != nil {
			t.Fatalf("Unmarshal failed: %v, data: %s", err, string(data))
		}
		if !original.Equal(restored) {
			t.Errorf("Round trip failed")
		}
	})

	t.Run("Information", func(t *testing.T) {
		original := NewInformation(1024, Information.Byte)
		data, _ := MarshalCompactInformation(original)
		restored, err := UnmarshalCompactInformation(data)
		if err != nil {
			t.Fatalf("Unmarshal failed: %v, data: %s", err, string(data))
		}
		if !original.Equal(restored) {
			t.Errorf("Round trip failed")
		}
	})
}

func TestMarshalCompactAllDimensions(t *testing.T) {
	// Test a sample from each dimension to ensure they all work
	tests := []struct {
		name     string
		marshal  func() ([]byte, error)
		wantUnit string
	}{
		{"Temperature", func() ([]byte, error) { return MarshalCompactTemperature(NewTemperature(25, Temperature.Celsius)) }, "temperature_celsius"},
		{"Pressure", func() ([]byte, error) { return MarshalCompactPressure(NewPressure(101.3, Pressure.Kilopascal)) }, "pressure_kilopascal"},
		{"Length", func() ([]byte, error) { return MarshalCompactLength(NewLength(100, Length.Meter)) }, "length_meter"},
		{"Mass", func() ([]byte, error) { return MarshalCompactMass(NewMass(75, Mass.Kilogram)) }, "mass_kilogram"},
		{"Duration", func() ([]byte, error) { return MarshalCompactDuration(NewDuration(60, Duration.Second)) }, "duration_second"},
		{"Angle", func() ([]byte, error) { return MarshalCompactAngle(NewAngle(90, Angle.Degree)) }, "angle_degree"},
		{"Area", func() ([]byte, error) { return MarshalCompactArea(NewArea(100, Area.SquareMeter)) }, "area_square_meter"},
		{"Volume", func() ([]byte, error) { return MarshalCompactVolume(NewVolume(1, Volume.Liter)) }, "volume_liter"},
		{"Speed", func() ([]byte, error) { return MarshalCompactSpeed(NewSpeed(100, Speed.KilometersPerHour)) }, "speed_kilometers_per_hour"},
		{"Acceleration", func() ([]byte, error) {
			return MarshalCompactAcceleration(NewAcceleration(9.8, Acceleration.MetersPerSecondSquared))
		}, "acceleration_meters_per_second_squared"},
		{"FlowRate", func() ([]byte, error) { return MarshalCompactFlowRate(NewFlowRate(10, FlowRate.CubicMetersPerHour)) }, "flowrate_cubic_meters_per_hour"},
		{"Power", func() ([]byte, error) { return MarshalCompactPower(NewPower(1000, Power.Watt)) }, "power_watt"},
		{"Energy", func() ([]byte, error) { return MarshalCompactEnergy(NewEnergy(3600, Energy.Joule)) }, "energy_joule"},
		{"Concentration", func() ([]byte, error) {
			return MarshalCompactConcentration(NewConcentration(5, Concentration.GramsPerLiter))
		}, "concentration_grams_per_liter"},
		{"Dispersion", func() ([]byte, error) {
			return MarshalCompactDispersion(NewDispersion(100, Dispersion.PartsPerMillion))
		}, "dispersion_parts_per_million"},
		{"ElectricCharge", func() ([]byte, error) {
			return MarshalCompactElectricCharge(NewElectricCharge(1, ElectricCharge.Coulomb))
		}, "electric_charge_coulomb"},
		{"ElectricCurrent", func() ([]byte, error) {
			return MarshalCompactElectricCurrent(NewElectricCurrent(1, ElectricCurrent.Ampere))
		}, "electric_current_ampere"},
		{"ElectricPotentialDifference", func() ([]byte, error) {
			return MarshalCompactElectricPotentialDifference(NewElectricPotentialDifference(220, ElectricPotentialDifference.Volt))
		}, "electric_potential_difference_volt"},
		{"Frequency", func() ([]byte, error) { return MarshalCompactFrequency(NewFrequency(1000, Frequency.Hertz)) }, "frequency_hertz"},
		{"Illuminance", func() ([]byte, error) { return MarshalCompactIlluminance(NewIlluminance(500, Illuminance.Lux)) }, "illuminance_lux"},
		{"Information", func() ([]byte, error) { return MarshalCompactInformation(NewInformation(1024, Information.Byte)) }, "information_byte"},
		{"FuelEfficiency", func() ([]byte, error) {
			return MarshalCompactFuelEfficiency(NewFuelEfficiency(15, FuelEfficiency.KilometersPerLiter))
		}, "fuel_efficiency_kilometers_per_liter"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.marshal()
			if err != nil {
				t.Fatalf("Marshal failed: %v", err)
			}

			var cj CompactJSON
			if err := json.Unmarshal(data, &cj); err != nil {
				t.Fatalf("Failed to unmarshal result: %v", err)
			}

			if cj.Unit != tt.wantUnit {
				t.Errorf("Unit = %q, want %q", cj.Unit, tt.wantUnit)
			}
		})
	}
}
