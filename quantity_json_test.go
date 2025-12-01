package unit

import (
	"encoding/json"
	"testing"
)

func TestQuantityMarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		quantity any
		wantJSON string
	}{
		{
			name:     "Temperature Celsius",
			quantity: NewTemperature(25, Temperature.Celsius),
			wantJSON: `{"value":25,"unit":{"name":"Celsius","symbol":"°C"},"dimension":"temperature"}`,
		},
		{
			name:     "Pressure Kilopascal",
			quantity: NewPressure(101.325, Pressure.Kilopascal),
			wantJSON: `{"value":101.325,"unit":{"name":"Kilopascal","symbol":"kPa"},"dimension":"pressure"}`,
		},
		{
			name:     "Length Meter",
			quantity: NewLength(100, Length.Meter),
			wantJSON: `{"value":100,"unit":{"name":"Meter","symbol":"m"},"dimension":"length"}`,
		},
		{
			name:     "Mass Kilogram",
			quantity: NewMass(75.5, Mass.Kilogram),
			wantJSON: `{"value":75.5,"unit":{"name":"Kilogram","symbol":"kg"},"dimension":"mass"}`,
		},
		{
			name:     "Duration Hour",
			quantity: NewDuration(2.5, Duration.Hour),
			wantJSON: `{"value":2.5,"unit":{"name":"Hour","symbol":"h"},"dimension":"duration"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.quantity)
			if err != nil {
				t.Fatalf("json.Marshal failed: %v", err)
			}

			if string(data) != tt.wantJSON {
				t.Errorf("json.Marshal() = %s, want %s", string(data), tt.wantJSON)
			}
		})
	}
}

func TestQuantityUnmarshalJSON(t *testing.T) {
	t.Run("Temperature", func(t *testing.T) {
		jsonData := `{"value":25,"unit":{"name":"Celsius","symbol":"°C"},"dimension":"temperature"}`

		var temp Quantity[TemperatureUnit]
		err := json.Unmarshal([]byte(jsonData), &temp)
		if err != nil {
			t.Fatalf("json.Unmarshal failed: %v", err)
		}

		if temp.Value != 25 {
			t.Errorf("Value = %v, want 25", temp.Value)
		}
		if temp.Unit.Symbol() != "°C" {
			t.Errorf("Unit.Symbol() = %v, want °C", temp.Unit.Symbol())
		}
	})

	t.Run("Pressure", func(t *testing.T) {
		jsonData := `{"value":101.325,"unit":{"name":"Kilopascal","symbol":"kPa"},"dimension":"pressure"}`

		var pressure Quantity[PressureUnit]
		err := json.Unmarshal([]byte(jsonData), &pressure)
		if err != nil {
			t.Fatalf("json.Unmarshal failed: %v", err)
		}

		if pressure.Value != 101.325 {
			t.Errorf("Value = %v, want 101.325", pressure.Value)
		}
		if pressure.Unit.Symbol() != "kPa" {
			t.Errorf("Unit.Symbol() = %v, want kPa", pressure.Unit.Symbol())
		}
	})

	t.Run("Length", func(t *testing.T) {
		jsonData := `{"value":100,"unit":{"name":"Meter","symbol":"m"},"dimension":"length"}`

		var length Quantity[LengthUnit]
		err := json.Unmarshal([]byte(jsonData), &length)
		if err != nil {
			t.Fatalf("json.Unmarshal failed: %v", err)
		}

		if length.Value != 100 {
			t.Errorf("Value = %v, want 100", length.Value)
		}
		if length.Unit.Symbol() != "m" {
			t.Errorf("Unit.Symbol() = %v, want m", length.Unit.Symbol())
		}
	})
}

func TestQuantityJSONRoundTrip(t *testing.T) {
	tests := []struct {
		name     string
		original any
		target   any
		equal    func() bool
	}{
		{
			name:     "Temperature",
			original: NewTemperature(25, Temperature.Celsius),
			target:   new(Quantity[TemperatureUnit]),
			equal: func() bool {
				return NewTemperature(25, Temperature.Celsius).Equal(*new(Quantity[TemperatureUnit]))
			},
		},
	}

	// Temperature round-trip
	t.Run("Temperature", func(t *testing.T) {
		original := NewTemperature(25, Temperature.Celsius)

		data, err := json.Marshal(original)
		if err != nil {
			t.Fatalf("json.Marshal failed: %v", err)
		}

		var restored Quantity[TemperatureUnit]
		err = json.Unmarshal(data, &restored)
		if err != nil {
			t.Fatalf("json.Unmarshal failed: %v", err)
		}

		if !original.Equal(restored) {
			t.Errorf("Round-trip failed: original=%v, restored=%v", original, restored)
		}
	})

	// Pressure round-trip
	t.Run("Pressure", func(t *testing.T) {
		original := NewPressure(101.325, Pressure.Kilopascal)

		data, err := json.Marshal(original)
		if err != nil {
			t.Fatalf("json.Marshal failed: %v", err)
		}

		var restored Quantity[PressureUnit]
		err = json.Unmarshal(data, &restored)
		if err != nil {
			t.Fatalf("json.Unmarshal failed: %v", err)
		}

		if !original.Equal(restored) {
			t.Errorf("Round-trip failed: original=%v, restored=%v", original, restored)
		}
	})

	// Length round-trip
	t.Run("Length", func(t *testing.T) {
		original := NewLength(100, Length.Meter)

		data, err := json.Marshal(original)
		if err != nil {
			t.Fatalf("json.Marshal failed: %v", err)
		}

		var restored Quantity[LengthUnit]
		err = json.Unmarshal(data, &restored)
		if err != nil {
			t.Fatalf("json.Unmarshal failed: %v", err)
		}

		if !original.Equal(restored) {
			t.Errorf("Round-trip failed: original=%v, restored=%v", original, restored)
		}
	})

	// Ignore the tests slice, we're using subtests instead
	_ = tests
}

func TestQuantityUnmarshalJSONErrors(t *testing.T) {
	t.Run("Invalid JSON", func(t *testing.T) {
		var temp Quantity[TemperatureUnit]
		err := json.Unmarshal([]byte(`{invalid`), &temp)
		if err == nil {
			t.Error("Expected error for invalid JSON")
		}
	})

	t.Run("Wrong dimension", func(t *testing.T) {
		// Try to unmarshal pressure JSON into temperature
		jsonData := `{"value":101.325,"unit":{"name":"Kilopascal","symbol":"kPa"},"dimension":"pressure"}`

		var temp Quantity[TemperatureUnit]
		err := json.Unmarshal([]byte(jsonData), &temp)
		if err == nil {
			t.Error("Expected error for dimension mismatch")
		}
	})

	t.Run("Unknown unit symbol", func(t *testing.T) {
		jsonData := `{"value":25,"unit":{"name":"Unknown","symbol":"XX"},"dimension":"temperature"}`

		var temp Quantity[TemperatureUnit]
		err := json.Unmarshal([]byte(jsonData), &temp)
		if err == nil {
			t.Error("Expected error for unknown unit symbol")
		}
	})
}

func TestCompactMarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		compact  any
		contains string
	}{
		{
			name:     "Temperature Celsius",
			compact:  Compact[TemperatureUnit]{NewTemperature(25, Temperature.Celsius)},
			contains: `"unit":"temperature_celsius"`,
		},
		{
			name:     "Pressure Kilopascal",
			compact:  Compact[PressureUnit]{NewPressure(101.325, Pressure.Kilopascal)},
			contains: `"unit":"pressure_kilopascal"`,
		},
		{
			name:     "Length Meter",
			compact:  Compact[LengthUnit]{NewLength(100, Length.Meter)},
			contains: `"unit":"length_meter"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.compact)
			if err != nil {
				t.Fatalf("json.Marshal failed: %v", err)
			}

			jsonStr := string(data)
			if !contains(jsonStr, tt.contains) {
				t.Errorf("json.Marshal() = %s, should contain %s", jsonStr, tt.contains)
			}
		})
	}
}

func TestCompactUnmarshalJSON(t *testing.T) {
	t.Run("Temperature", func(t *testing.T) {
		jsonData := `{"value":25,"unit":"temperature_celsius","symbol":"°C"}`

		var compact Compact[TemperatureUnit]
		err := json.Unmarshal([]byte(jsonData), &compact)
		if err != nil {
			t.Fatalf("json.Unmarshal failed: %v", err)
		}

		if compact.Value != 25 {
			t.Errorf("Value = %v, want 25", compact.Value)
		}
		if compact.Unit.Symbol() != "°C" {
			t.Errorf("Unit.Symbol() = %v, want °C", compact.Unit.Symbol())
		}
	})

	t.Run("Pressure", func(t *testing.T) {
		jsonData := `{"value":101.325,"unit":"pressure_kilopascal","symbol":"kPa"}`

		var compact Compact[PressureUnit]
		err := json.Unmarshal([]byte(jsonData), &compact)
		if err != nil {
			t.Fatalf("json.Unmarshal failed: %v", err)
		}

		if compact.Value != 101.325 {
			t.Errorf("Value = %v, want 101.325", compact.Value)
		}
		if compact.Unit.Symbol() != "kPa" {
			t.Errorf("Unit.Symbol() = %v, want kPa", compact.Unit.Symbol())
		}
	})

	t.Run("Length", func(t *testing.T) {
		jsonData := `{"value":100,"unit":"length_meter","symbol":"m"}`

		var compact Compact[LengthUnit]
		err := json.Unmarshal([]byte(jsonData), &compact)
		if err != nil {
			t.Fatalf("json.Unmarshal failed: %v", err)
		}

		if compact.Value != 100 {
			t.Errorf("Value = %v, want 100", compact.Value)
		}
		if compact.Unit.Symbol() != "m" {
			t.Errorf("Unit.Symbol() = %v, want m", compact.Unit.Symbol())
		}
	})
}

func TestCompactJSONRoundTrip(t *testing.T) {
	t.Run("Temperature", func(t *testing.T) {
		original := Compact[TemperatureUnit]{NewTemperature(25, Temperature.Celsius)}

		data, err := json.Marshal(original)
		if err != nil {
			t.Fatalf("json.Marshal failed: %v", err)
		}

		var restored Compact[TemperatureUnit]
		err = json.Unmarshal(data, &restored)
		if err != nil {
			t.Fatalf("json.Unmarshal failed: %v", err)
		}

		if !original.Quantity.Equal(restored.Quantity) {
			t.Errorf("Round-trip failed: original=%v, restored=%v", original, restored)
		}
	})

	t.Run("Pressure", func(t *testing.T) {
		original := Compact[PressureUnit]{NewPressure(101.325, Pressure.Kilopascal)}

		data, err := json.Marshal(original)
		if err != nil {
			t.Fatalf("json.Marshal failed: %v", err)
		}

		var restored Compact[PressureUnit]
		err = json.Unmarshal(data, &restored)
		if err != nil {
			t.Fatalf("json.Unmarshal failed: %v", err)
		}

		if !original.Quantity.Equal(restored.Quantity) {
			t.Errorf("Round-trip failed: original=%v, restored=%v", original, restored)
		}
	})
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
