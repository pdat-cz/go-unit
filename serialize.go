// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode"
)

// SerializationFormat specifies which JSON format to use for serialization
type SerializationFormat int

const (
	// FormatFull includes all unit details: {"value": 25.5, "unit": {"name": "Celsius", "symbol": "°C", "dimension": "temperature"}}
	FormatFull SerializationFormat = iota
	// FormatCompact includes unit key and symbol: {"value": 25.5, "unit": {"key": "temperature_celsius", "symbol": "°C"}}
	FormatCompact
	// FormatMinimal includes only unit key as string: {"value": 25.5, "unit": "temperature_celsius"}
	FormatMinimal
)

// UnitFullJSON is used for full JSON serialization of units (all details nested)
type UnitFullJSON struct {
	Name      string `json:"name" yaml:"name"`
	Symbol    string `json:"symbol" yaml:"symbol"`
	Dimension string `json:"dimension" yaml:"dimension"`
}

// UnitCompactJSON is used for compact JSON serialization of units (key + symbol)
type UnitCompactJSON struct {
	Key    string `json:"key" yaml:"key"`       // "dimension_unitname" format, e.g., "temperature_celsius"
	Symbol string `json:"symbol" yaml:"symbol"` // unit symbol, e.g., "°C"
}

// MeasurementJSON is used for full JSON serialization of measurements
type MeasurementJSON struct {
	Value float64      `json:"value" yaml:"value"`
	Unit  UnitFullJSON `json:"unit" yaml:"unit"`
}

// MeasurementCompactJSON is used for compact JSON serialization of measurements
type MeasurementCompactJSON struct {
	Value float64         `json:"value" yaml:"value"`
	Unit  UnitCompactJSON `json:"unit" yaml:"unit"`
}

// MeasurementMinimalJSON is used for minimal JSON serialization of measurements
type MeasurementMinimalJSON struct {
	Value float64 `json:"value" yaml:"value"`
	Unit  string  `json:"unit" yaml:"unit"` // "dimension_unitname" format, e.g., "temperature_celsius"
}

// Legacy types for backward compatibility during deserialization
type legacyMeasurementJSON struct {
	Value     float64        `json:"value"`
	Unit      legacyUnitJSON `json:"unit"`
	Dimension string         `json:"dimension"`
}

type legacyUnitJSON struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type legacyCompactJSON struct {
	Value  float64 `json:"value"`
	Unit   string  `json:"unit"`
	Symbol string  `json:"symbol,omitempty"`
}

// toSnakeCase converts a string to snake_case
// Handles both PascalCase ("KilometersPerHour") and space-separated ("Kilometers per Hour")
func toSnakeCase(s string) string {
	// First, replace spaces with underscores
	s = strings.ReplaceAll(s, " ", "_")

	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			// Add underscore before uppercase letter if not at start and previous char wasn't underscore
			if i > 0 && result.Len() > 0 {
				lastByte := result.String()[result.Len()-1]
				if lastByte != '_' {
					result.WriteRune('_')
				}
			}
			result.WriteRune(unicode.ToLower(r))
		} else if r == '_' {
			// Avoid double underscores
			if result.Len() > 0 && result.String()[result.Len()-1] != '_' {
				result.WriteRune('_')
			}
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// unitKey returns a snake_case key combining dimension and unit name
// e.g., "temperature", "Celsius" -> "temperature_celsius"
func unitKey(dimension, name string) string {
	return strings.ToLower(dimension) + "_" + toSnakeCase(name)
}

// parseUnitKey splits a compact unit key into dimension and unit name
// e.g., "temperature_celsius" -> "temperature", "celsius"
func parseUnitKey(key string) (dimension, unitName string) {
	idx := strings.Index(key, "_")
	if idx == -1 {
		return key, ""
	}
	return key[:idx], key[idx+1:]
}

// detectFormat determines which JSON format is being used
// Returns: format type, dimension, error
func detectFormat(data []byte) (SerializationFormat, string, error) {
	// Try to parse as a generic map first to inspect structure
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return FormatFull, "", err
	}

	unitRaw, hasUnit := raw["unit"]
	if !hasUnit {
		return FormatFull, "", fmt.Errorf("missing 'unit' field")
	}

	// Check if unit is a string (minimal format) or object
	var unitStr string
	if err := json.Unmarshal(unitRaw, &unitStr); err == nil {
		// It's a string - could be minimal format or legacy compact
		// Check if there's a top-level symbol (legacy compact)
		if _, hasSymbol := raw["symbol"]; hasSymbol {
			// Legacy compact format: {"value": 25.5, "unit": "temperature_celsius", "symbol": "°C"}
			dim, _ := parseUnitKey(unitStr)
			return FormatMinimal, dim, nil // Treat as minimal for parsing purposes
		}
		// Minimal format: {"value": 25.5, "unit": "temperature_celsius"}
		dim, _ := parseUnitKey(unitStr)
		return FormatMinimal, dim, nil
	}

	// Unit is an object - check which type
	var unitObj map[string]json.RawMessage
	if err := json.Unmarshal(unitRaw, &unitObj); err != nil {
		return FormatFull, "", err
	}

	// Check for "key" field (compact format)
	if keyRaw, hasKey := unitObj["key"]; hasKey {
		var keyStr string
		if err := json.Unmarshal(keyRaw, &keyStr); err == nil {
			dim, _ := parseUnitKey(keyStr)
			return FormatCompact, dim, nil
		}
	}

	// Check for "dimension" field (full format)
	if dimRaw, hasDim := unitObj["dimension"]; hasDim {
		var dimStr string
		if err := json.Unmarshal(dimRaw, &dimStr); err == nil {
			return FormatFull, dimStr, nil
		}
	}

	// Check for legacy format with top-level dimension
	if dimRaw, hasDim := raw["dimension"]; hasDim {
		var dimStr string
		if err := json.Unmarshal(dimRaw, &dimStr); err == nil {
			return FormatFull, dimStr, nil // Legacy full format
		}
	}

	return FormatFull, "", fmt.Errorf("could not determine format or dimension")
}

// unmarshalValue extracts the value from any format
func unmarshalValue(data []byte) (float64, error) {
	var obj struct {
		Value float64 `json:"value"`
	}
	if err := json.Unmarshal(data, &obj); err != nil {
		return 0, err
	}
	return obj.Value, nil
}

// unmarshalUnitInfo extracts unit information from any format
// Returns: symbol, name, key (for compact/minimal)
func unmarshalUnitInfo(data []byte) (symbol, name, key string, err error) {
	format, _, err := detectFormat(data)
	if err != nil {
		return "", "", "", err
	}

	switch format {
	case FormatFull:
		// Try new full format first
		var full MeasurementJSON
		if err := json.Unmarshal(data, &full); err == nil && full.Unit.Symbol != "" {
			return full.Unit.Symbol, full.Unit.Name, "", nil
		}
		// Try legacy format
		var legacy legacyMeasurementJSON
		if err := json.Unmarshal(data, &legacy); err == nil {
			return legacy.Unit.Symbol, legacy.Unit.Name, "", nil
		}
	case FormatCompact:
		var compact MeasurementCompactJSON
		if err := json.Unmarshal(data, &compact); err == nil {
			return compact.Unit.Symbol, "", compact.Unit.Key, nil
		}
	case FormatMinimal:
		var minimal MeasurementMinimalJSON
		if err := json.Unmarshal(data, &minimal); err == nil {
			// Check for legacy compact with top-level symbol
			var legacy legacyCompactJSON
			if json.Unmarshal(data, &legacy) == nil && legacy.Symbol != "" {
				return legacy.Symbol, "", minimal.Unit, nil
			}
			return "", "", minimal.Unit, nil
		}
	}

	return "", "", "", fmt.Errorf("could not extract unit info")
}

// parsedMeasurement holds all extracted data from any JSON format
type parsedMeasurement struct {
	Value     float64
	Symbol    string
	Name      string
	Key       string
	Dimension string
	Format    SerializationFormat
}

// parseMeasurement extracts all measurement data from any format
func parseMeasurement(data []byte) (*parsedMeasurement, error) {
	format, dimension, err := detectFormat(data)
	if err != nil {
		return nil, err
	}

	value, err := unmarshalValue(data)
	if err != nil {
		return nil, err
	}

	symbol, name, key, _ := unmarshalUnitInfo(data)

	return &parsedMeasurement{
		Value:     value,
		Symbol:    symbol,
		Name:      name,
		Key:       key,
		Dimension: dimension,
		Format:    format,
	}, nil
}

// matchUnitByKey tries to match a unit name from the key
func (p *parsedMeasurement) matchUnitByKey(unitName string) bool {
	if p.Key == "" {
		return false
	}
	_, keyUnitName := parseUnitKey(p.Key)
	return strings.EqualFold(keyUnitName, unitName)
}

// AnyMeasurement is a wrapper that can hold any type of measurement
// and provides methods to access it based on its dimension
type AnyMeasurement struct {
	value     interface{}
	dimension string
}

// GetDimension returns the dimension of the measurement
func (am *AnyMeasurement) GetDimension() string {
	return am.dimension
}

// AsTemperature attempts to convert the measurement to a Temperature measurement
func (am *AnyMeasurement) AsTemperature() (Quantity[TemperatureUnit], bool) {
	if m, ok := am.value.(Quantity[TemperatureUnit]); ok {
		return m, true
	}
	return Quantity[TemperatureUnit]{}, false
}

// AsPressure attempts to convert the measurement to a Pressure measurement
func (am *AnyMeasurement) AsPressure() (Quantity[PressureUnit], bool) {
	if m, ok := am.value.(Quantity[PressureUnit]); ok {
		return m, true
	}
	return Quantity[PressureUnit]{}, false
}

// AsLength attempts to convert the measurement to a Length measurement
func (am *AnyMeasurement) AsLength() (Quantity[LengthUnit], bool) {
	if m, ok := am.value.(Quantity[LengthUnit]); ok {
		return m, true
	}
	return Quantity[LengthUnit]{}, false
}

// AsVolume attempts to convert the measurement to a Volume measurement
func (am *AnyMeasurement) AsVolume() (Quantity[VolumeUnit], bool) {
	if m, ok := am.value.(Quantity[VolumeUnit]); ok {
		return m, true
	}
	return Quantity[VolumeUnit]{}, false
}

// AsMass attempts to convert the measurement to a Mass measurement
func (am *AnyMeasurement) AsMass() (Quantity[MassUnit], bool) {
	if m, ok := am.value.(Quantity[MassUnit]); ok {
		return m, true
	}
	return Quantity[MassUnit]{}, false
}

// AsSpeed attempts to convert the measurement to a Speed measurement
func (am *AnyMeasurement) AsSpeed() (Quantity[SpeedUnit], bool) {
	if m, ok := am.value.(Quantity[SpeedUnit]); ok {
		return m, true
	}
	return Quantity[SpeedUnit]{}, false
}

// AsAcceleration attempts to convert the measurement to an Acceleration measurement
func (am *AnyMeasurement) AsAcceleration() (Quantity[AccelerationUnit], bool) {
	if m, ok := am.value.(Quantity[AccelerationUnit]); ok {
		return m, true
	}
	return Quantity[AccelerationUnit]{}, false
}

// AsFlowRate attempts to convert the measurement to a FlowRate measurement
func (am *AnyMeasurement) AsFlowRate() (Quantity[FlowRateUnit], bool) {
	if m, ok := am.value.(Quantity[FlowRateUnit]); ok {
		return m, true
	}
	return Quantity[FlowRateUnit]{}, false
}

// AsPower attempts to convert the measurement to a Power measurement
func (am *AnyMeasurement) AsPower() (Quantity[PowerUnit], bool) {
	if m, ok := am.value.(Quantity[PowerUnit]); ok {
		return m, true
	}
	return Quantity[PowerUnit]{}, false
}

// AsEnergy attempts to convert the measurement to an Energy measurement
func (am *AnyMeasurement) AsEnergy() (Quantity[EnergyUnit], bool) {
	if m, ok := am.value.(Quantity[EnergyUnit]); ok {
		return m, true
	}
	return Quantity[EnergyUnit]{}, false
}

// AsDuration attempts to convert the measurement to a Duration measurement
func (am *AnyMeasurement) AsDuration() (Quantity[DurationUnit], bool) {
	if m, ok := am.value.(Quantity[DurationUnit]); ok {
		return m, true
	}
	return Quantity[DurationUnit]{}, false
}

// AsAngle attempts to convert the measurement to an Angle measurement
func (am *AnyMeasurement) AsAngle() (Quantity[AngleUnit], bool) {
	if m, ok := am.value.(Quantity[AngleUnit]); ok {
		return m, true
	}
	return Quantity[AngleUnit]{}, false
}

// AsArea attempts to convert the measurement to an Area measurement
func (am *AnyMeasurement) AsArea() (Quantity[AreaUnit], bool) {
	if m, ok := am.value.(Quantity[AreaUnit]); ok {
		return m, true
	}
	return Quantity[AreaUnit]{}, false
}

// AsConcentration attempts to convert the measurement to a Concentration measurement
func (am *AnyMeasurement) AsConcentration() (Quantity[ConcentrationUnit], bool) {
	if m, ok := am.value.(Quantity[ConcentrationUnit]); ok {
		return m, true
	}
	return Quantity[ConcentrationUnit]{}, false
}

// AsDispersion attempts to convert the measurement to a Dispersion measurement
func (am *AnyMeasurement) AsDispersion() (Quantity[DispersionUnit], bool) {
	if m, ok := am.value.(Quantity[DispersionUnit]); ok {
		return m, true
	}
	return Quantity[DispersionUnit]{}, false
}

// AsElectricCharge attempts to convert the measurement to an ElectricCharge measurement
func (am *AnyMeasurement) AsElectricCharge() (Quantity[ElectricChargeUnit], bool) {
	if m, ok := am.value.(Quantity[ElectricChargeUnit]); ok {
		return m, true
	}
	return Quantity[ElectricChargeUnit]{}, false
}

// AsElectricCurrent attempts to convert the measurement to an ElectricCurrent measurement
func (am *AnyMeasurement) AsElectricCurrent() (Quantity[ElectricCurrentUnit], bool) {
	if m, ok := am.value.(Quantity[ElectricCurrentUnit]); ok {
		return m, true
	}
	return Quantity[ElectricCurrentUnit]{}, false
}

// AsElectricPotentialDifference attempts to convert the measurement to an ElectricPotentialDifference measurement
func (am *AnyMeasurement) AsElectricPotentialDifference() (Quantity[ElectricPotentialDifferenceUnit], bool) {
	if m, ok := am.value.(Quantity[ElectricPotentialDifferenceUnit]); ok {
		return m, true
	}
	return Quantity[ElectricPotentialDifferenceUnit]{}, false
}

// AsInformation attempts to convert the measurement to an Information measurement
func (am *AnyMeasurement) AsInformation() (Quantity[InformationUnit], bool) {
	if m, ok := am.value.(Quantity[InformationUnit]); ok {
		return m, true
	}
	return Quantity[InformationUnit]{}, false
}

// AsFrequency attempts to convert the measurement to a Frequency measurement
func (am *AnyMeasurement) AsFrequency() (Quantity[FrequencyUnit], bool) {
	if m, ok := am.value.(Quantity[FrequencyUnit]); ok {
		return m, true
	}
	return Quantity[FrequencyUnit]{}, false
}

// AsIlluminance attempts to convert the measurement to an Illuminance measurement
func (am *AnyMeasurement) AsIlluminance() (Quantity[IlluminanceUnit], bool) {
	if m, ok := am.value.(Quantity[IlluminanceUnit]); ok {
		return m, true
	}
	return Quantity[IlluminanceUnit]{}, false
}

// AsFuelEfficiency attempts to convert the measurement to a FuelEfficiency measurement
func (am *AnyMeasurement) AsFuelEfficiency() (Quantity[FuelEfficiencyUnit], bool) {
	if m, ok := am.value.(Quantity[FuelEfficiencyUnit]); ok {
		return m, true
	}
	return Quantity[FuelEfficiencyUnit]{}, false
}

// AsGeneral attempts to convert the measurement to a General measurement
func (am *AnyMeasurement) AsGeneral() (Quantity[GeneralUnit], bool) {
	if m, ok := am.value.(Quantity[GeneralUnit]); ok {
		return m, true
	}
	return Quantity[GeneralUnit]{}, false
}

// UnmarshalMeasurement deserializes a JSON representation to an AnyMeasurement
// without requiring knowledge of the dimension in advance
// Supports all three formats: full, compact, and minimal
func UnmarshalMeasurement(data []byte) (*AnyMeasurement, error) {
	// Detect format and extract dimension
	_, dimension, err := detectFormat(data)
	if err != nil {
		return nil, err
	}

	// Helper to create fallback
	createFallback := func(origErr error) (*AnyMeasurement, error) {
		value, _ := unmarshalValue(data)
		symbol, name, key, _ := unmarshalUnitInfo(data)
		if key != "" && symbol == "" {
			_, unitName := parseUnitKey(key)
			symbol = unitName
		}
		if name == "" {
			name = symbol
		}
		return fallbackToGeneral(value, symbol, name, origErr)
	}

	// Based on the dimension, call the appropriate unmarshal function
	switch dimension {
	case "temperature":
		m, err := UnmarshalTemperature(data)
		if err != nil {
			return createFallback(err)
		}
		return &AnyMeasurement{value: m, dimension: "temperature"}, nil
	case "pressure":
		m, err := UnmarshalPressure(data)
		if err != nil {
			return createFallback(err)
		}
		return &AnyMeasurement{value: m, dimension: "pressure"}, nil
	case "flowrate":
		m, err := UnmarshalFlowRate(data)
		if err != nil {
			return createFallback(err)
		}
		return &AnyMeasurement{value: m, dimension: "flowrate"}, nil
	case "power":
		m, err := UnmarshalPower(data)
		if err != nil {
			return createFallback(err)
		}
		return &AnyMeasurement{value: m, dimension: "power"}, nil
	case "energy":
		m, err := UnmarshalEnergy(data)
		if err != nil {
			return createFallback(err)
		}
		return &AnyMeasurement{value: m, dimension: "energy"}, nil
	case "length":
		m, err := UnmarshalLength(data)
		if err != nil {
			return createFallback(err)
		}
		return &AnyMeasurement{value: m, dimension: "length"}, nil
	case "mass":
		m, err := UnmarshalMass(data)
		if err != nil {
			return createFallback(err)
		}
		return &AnyMeasurement{value: m, dimension: "mass"}, nil
	case "duration":
		m, err := UnmarshalDuration(data)
		if err != nil {
			return createFallback(err)
		}
		return &AnyMeasurement{value: m, dimension: "duration"}, nil
	case "angle":
		m, err := UnmarshalAngle(data)
		if err != nil {
			return createFallback(err)
		}
		return &AnyMeasurement{value: m, dimension: "angle"}, nil
	case "area":
		m, err := UnmarshalArea(data)
		if err != nil {
			return createFallback(err)
		}
		return &AnyMeasurement{value: m, dimension: "area"}, nil
	case "volume":
		m, err := UnmarshalVolume(data)
		if err != nil {
			return createFallback(err)
		}
		return &AnyMeasurement{value: m, dimension: "volume"}, nil
	case "acceleration":
		m, err := UnmarshalAcceleration(data)
		if err != nil {
			return createFallback(err)
		}
		return &AnyMeasurement{value: m, dimension: "acceleration"}, nil
	case "concentration":
		m, err := UnmarshalConcentration(data)
		if err != nil {
			return createFallback(err)
		}
		return &AnyMeasurement{value: m, dimension: "concentration"}, nil
	case "dispersion":
		m, err := UnmarshalDispersion(data)
		if err != nil {
			return createFallback(err)
		}
		return &AnyMeasurement{value: m, dimension: "dispersion"}, nil
	case "electric_charge":
		m, err := UnmarshalElectricCharge(data)
		if err != nil {
			return createFallback(err)
		}
		return &AnyMeasurement{value: m, dimension: "electric_charge"}, nil
	case "electric_current":
		m, err := UnmarshalElectricCurrent(data)
		if err != nil {
			return createFallback(err)
		}
		return &AnyMeasurement{value: m, dimension: "electric_current"}, nil
	case "electric_potential_difference":
		m, err := UnmarshalElectricPotentialDifference(data)
		if err != nil {
			return createFallback(err)
		}
		return &AnyMeasurement{value: m, dimension: "electric_potential_difference"}, nil
	case "information":
		m, err := UnmarshalInformation(data)
		if err != nil {
			return createFallback(err)
		}
		return &AnyMeasurement{value: m, dimension: "information"}, nil
	case "frequency":
		m, err := UnmarshalFrequency(data)
		if err != nil {
			return createFallback(err)
		}
		return &AnyMeasurement{value: m, dimension: "frequency"}, nil
	case "illuminance":
		m, err := UnmarshalIlluminance(data)
		if err != nil {
			return createFallback(err)
		}
		return &AnyMeasurement{value: m, dimension: "illuminance"}, nil
	case "fuel_efficiency":
		m, err := UnmarshalFuelEfficiency(data)
		if err != nil {
			return createFallback(err)
		}
		return &AnyMeasurement{value: m, dimension: "fuel_efficiency"}, nil
	case "speed":
		m, err := UnmarshalSpeed(data)
		if err != nil {
			return createFallback(err)
		}
		return &AnyMeasurement{value: m, dimension: "speed"}, nil
	case "general":
		m, err := UnmarshalGeneral(data)
		if err != nil {
			return nil, err
		}
		return &AnyMeasurement{value: m, dimension: "general"}, nil
	default:
		// For unknown dimensions, use general unit
		return createFallback(fmt.Errorf("unknown dimension: %s", dimension))
	}
}

// fallbackToGeneral creates a general measurement from the given JSON data
func fallbackToGeneral(value float64, symbol, name string, originalErr error) (*AnyMeasurement, error) {
	// Create a general unit with the given symbol and name
	unit := NewGeneralUnit(symbol, name)
	m := NewGeneral(value, unit)
	return &AnyMeasurement{value: m, dimension: "general"}, nil
}

// marshalGeneric is a helper function to serialize any measurement to JSON (full format)
func marshalGeneric[T Category](m Quantity[T]) ([]byte, error) {
	return json.Marshal(MeasurementJSON{
		Value: m.Value,
		Unit: UnitFullJSON{
			Name:      m.Unit.Name(),
			Symbol:    m.Unit.Symbol(),
			Dimension: m.Unit.Dimension(),
		},
	})
}

// marshalGenericCompact is a helper function to serialize any measurement to compact JSON
func marshalGenericCompact[T Category](m Quantity[T]) ([]byte, error) {
	return json.Marshal(MeasurementCompactJSON{
		Value: m.Value,
		Unit: UnitCompactJSON{
			Key:    unitKey(m.Unit.Dimension(), m.Unit.Name()),
			Symbol: m.Unit.Symbol(),
		},
	})
}

// marshalGenericMinimal is a helper function to serialize any measurement to minimal JSON
func marshalGenericMinimal[T Category](m Quantity[T]) ([]byte, error) {
	return json.Marshal(MeasurementMinimalJSON{
		Value: m.Value,
		Unit:  unitKey(m.Unit.Dimension(), m.Unit.Name()),
	})
}

// MarshalWithFormat serializes any measurement to JSON with the specified format
func MarshalWithFormat[T Category](m Quantity[T], format SerializationFormat) ([]byte, error) {
	switch format {
	case FormatFull:
		return marshalGeneric(m)
	case FormatCompact:
		return marshalGenericCompact(m)
	case FormatMinimal:
		return marshalGenericMinimal(m)
	default:
		return marshalGeneric(m)
	}
}

// MarshalTemperature serializes a Temperature measurement to JSON
func MarshalTemperature(m Quantity[TemperatureUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalTemperature deserializes a JSON representation to a Temperature measurement
// Supports all three formats: full, compact, and minimal
func UnmarshalTemperature(data []byte) (Quantity[TemperatureUnit], error) {
	p, err := parseMeasurement(data)
	if err != nil {
		return Quantity[TemperatureUnit]{}, err
	}

	if p.Dimension != "temperature" {
		return Quantity[TemperatureUnit]{}, fmt.Errorf("expected dimension 'temperature', got '%s'", p.Dimension)
	}

	var unit TemperatureUnit
	switch {
	case p.Symbol == "°C" || p.Symbol == "C" || p.matchUnitByKey("celsius"):
		unit = Temperature.Celsius
	case p.Symbol == "°F" || p.Symbol == "F" || p.matchUnitByKey("fahrenheit"):
		unit = Temperature.Fahrenheit
	case p.Symbol == "K" || p.matchUnitByKey("kelvin"):
		unit = Temperature.Kelvin
	default:
		return Quantity[TemperatureUnit]{}, fmt.Errorf("unknown temperature unit: symbol=%s, key=%s", p.Symbol, p.Key)
	}

	return NewTemperature(p.Value, unit), nil
}

// MarshalPressure serializes a Pressure measurement to JSON
func MarshalPressure(m Quantity[PressureUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalPressure deserializes a JSON representation to a Pressure measurement
func UnmarshalPressure(data []byte) (Quantity[PressureUnit], error) {
	p, err := parseMeasurement(data)
	if err != nil {
		return Quantity[PressureUnit]{}, err
	}

	if p.Dimension != "pressure" {
		return Quantity[PressureUnit]{}, fmt.Errorf("expected dimension 'pressure', got '%s'", p.Dimension)
	}

	var unit PressureUnit
	switch {
	case p.Symbol == "Pa" || p.matchUnitByKey("pascal"):
		unit = Pressure.Pascal
	case p.Symbol == "kPa" || p.matchUnitByKey("kilopascal"):
		unit = Pressure.Kilopascal
	case p.Symbol == "bar" || p.matchUnitByKey("bar"):
		unit = Pressure.Bar
	case p.Symbol == "psi" || p.matchUnitByKey("psi"):
		unit = Pressure.PSI
	case p.Symbol == "inH₂O" || p.Symbol == "inH2O" || p.matchUnitByKey("inch_h2o"):
		unit = Pressure.InchH2O
	default:
		return Quantity[PressureUnit]{}, fmt.Errorf("unknown pressure unit: symbol=%s, key=%s", p.Symbol, p.Key)
	}

	return NewPressure(p.Value, unit), nil
}

// MarshalFlowRate serializes a FlowRate measurement to JSON
func MarshalFlowRate(m Quantity[FlowRateUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalFlowRate deserializes a JSON representation to a FlowRate measurement
func UnmarshalFlowRate(data []byte) (Quantity[FlowRateUnit], error) {
	p, err := parseMeasurement(data)
	if err != nil {
		return Quantity[FlowRateUnit]{}, err
	}

	if p.Dimension != "flowrate" {
		return Quantity[FlowRateUnit]{}, fmt.Errorf("expected dimension 'flowrate', got '%s'", p.Dimension)
	}

	var unit FlowRateUnit
	switch {
	case p.Symbol == "m³/h" || p.matchUnitByKey("cubic_meters_per_hour"):
		unit = FlowRate.CubicMetersPerHour
	case p.Symbol == "L/s" || p.matchUnitByKey("liters_per_second"):
		unit = FlowRate.LitersPerSecond
	case p.Symbol == "CFM" || p.matchUnitByKey("cfm"):
		unit = FlowRate.CFM
	default:
		return Quantity[FlowRateUnit]{}, fmt.Errorf("unknown flowrate unit: symbol=%s, key=%s", p.Symbol, p.Key)
	}

	return NewFlowRate(p.Value, unit), nil
}

// MarshalPower serializes a Power measurement to JSON
func MarshalPower(m Quantity[PowerUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalPower deserializes a JSON representation to a Power measurement
func UnmarshalPower(data []byte) (Quantity[PowerUnit], error) {
	p, err := parseMeasurement(data)
	if err != nil {
		return Quantity[PowerUnit]{}, err
	}

	if p.Dimension != "power" {
		return Quantity[PowerUnit]{}, fmt.Errorf("expected dimension 'power', got '%s'", p.Dimension)
	}

	var unit PowerUnit
	switch {
	case p.Symbol == "W" || p.matchUnitByKey("watt"):
		unit = Power.Watt
	case p.Symbol == "kW" || p.matchUnitByKey("kilowatt"):
		unit = Power.Kilowatt
	case p.Symbol == "BTU/h" || p.matchUnitByKey("btu_per_hour"):
		unit = Power.BTUPerHour
	default:
		return Quantity[PowerUnit]{}, fmt.Errorf("unknown power unit: symbol=%s, key=%s", p.Symbol, p.Key)
	}

	return NewPower(p.Value, unit), nil
}

// MarshalEnergy serializes an Energy measurement to JSON
func MarshalEnergy(m Quantity[EnergyUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalEnergy deserializes a JSON representation to an Energy measurement
func UnmarshalEnergy(data []byte) (Quantity[EnergyUnit], error) {
	p, err := parseMeasurement(data)
	if err != nil {
		return Quantity[EnergyUnit]{}, err
	}

	if p.Dimension != "energy" {
		return Quantity[EnergyUnit]{}, fmt.Errorf("expected dimension 'energy', got '%s'", p.Dimension)
	}

	var unit EnergyUnit
	switch {
	case p.Symbol == "J" || p.matchUnitByKey("joule"):
		unit = Energy.Joule
	case p.Symbol == "kWh" || p.matchUnitByKey("kilowatt_hour"):
		unit = Energy.KilowattHour
	case p.Symbol == "BTU" || p.matchUnitByKey("btu"):
		unit = Energy.BTU
	default:
		return Quantity[EnergyUnit]{}, fmt.Errorf("unknown energy unit: symbol=%s, key=%s", p.Symbol, p.Key)
	}

	return NewEnergy(p.Value, unit), nil
}

// MarshalLength serializes a Length measurement to JSON
func MarshalLength(m Quantity[LengthUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalLength deserializes a JSON representation to a Length measurement
func UnmarshalLength(data []byte) (Quantity[LengthUnit], error) {
	p, err := parseMeasurement(data)
	if err != nil {
		return Quantity[LengthUnit]{}, err
	}

	if p.Dimension != "length" {
		return Quantity[LengthUnit]{}, fmt.Errorf("expected dimension 'length', got '%s'", p.Dimension)
	}

	var unit LengthUnit
	switch {
	case p.Symbol == "m" || p.matchUnitByKey("meter"):
		unit = Length.Meter
	case p.Symbol == "km" || p.matchUnitByKey("kilometer"):
		unit = Length.Kilometer
	case p.Symbol == "cm" || p.matchUnitByKey("centimeter"):
		unit = Length.Centimeter
	case p.Symbol == "mm" || p.matchUnitByKey("millimeter"):
		unit = Length.Millimeter
	case p.Symbol == "µm" || p.matchUnitByKey("micrometer"):
		unit = Length.Micrometer
	case p.Symbol == "nm" || p.matchUnitByKey("nanometer"):
		unit = Length.Nanometer
	case p.Symbol == "in" || p.matchUnitByKey("inch"):
		unit = Length.Inch
	case p.Symbol == "ft" || p.matchUnitByKey("foot"):
		unit = Length.Foot
	case p.Symbol == "yd" || p.matchUnitByKey("yard"):
		unit = Length.Yard
	case p.Symbol == "mi" || p.matchUnitByKey("mile"):
		unit = Length.Mile
	default:
		return Quantity[LengthUnit]{}, fmt.Errorf("unknown length unit: symbol=%s, key=%s", p.Symbol, p.Key)
	}

	return NewLength(p.Value, unit), nil
}

// MarshalMass serializes a Mass measurement to JSON
func MarshalMass(m Quantity[MassUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalMass deserializes a JSON representation to a Mass measurement
func UnmarshalMass(data []byte) (Quantity[MassUnit], error) {
	p, err := parseMeasurement(data)
	if err != nil {
		return Quantity[MassUnit]{}, err
	}

	if p.Dimension != "mass" {
		return Quantity[MassUnit]{}, fmt.Errorf("expected dimension 'mass', got '%s'", p.Dimension)
	}

	var unit MassUnit
	switch {
	case p.Symbol == "kg" || p.matchUnitByKey("kilogram"):
		unit = Mass.Kilogram
	case p.Symbol == "g" || p.matchUnitByKey("gram"):
		unit = Mass.Gram
	case p.Symbol == "mg" || p.matchUnitByKey("milligram"):
		unit = Mass.Milligram
	case p.Symbol == "µg" || p.matchUnitByKey("microgram"):
		unit = Mass.Microgram
	case p.Symbol == "lb" || p.matchUnitByKey("pound"):
		unit = Mass.Pound
	case p.Symbol == "oz" || p.matchUnitByKey("ounce"):
		unit = Mass.Ounce
	case p.Symbol == "st" || p.matchUnitByKey("stone"):
		unit = Mass.Stone
	case p.Symbol == "t" || p.matchUnitByKey("metric_ton"):
		unit = Mass.MetricTon
	case p.Symbol == "ton" || p.matchUnitByKey("ton"):
		unit = Mass.Ton
	default:
		return Quantity[MassUnit]{}, fmt.Errorf("unknown mass unit: symbol=%s, key=%s", p.Symbol, p.Key)
	}

	return NewMass(p.Value, unit), nil
}

// MarshalDuration serializes a Duration measurement to JSON
func MarshalDuration(m Quantity[DurationUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalDuration deserializes a JSON representation to a Duration measurement
func UnmarshalDuration(data []byte) (Quantity[DurationUnit], error) {
	p, err := parseMeasurement(data)
	if err != nil {
		return Quantity[DurationUnit]{}, err
	}

	if p.Dimension != "duration" {
		return Quantity[DurationUnit]{}, fmt.Errorf("expected dimension 'duration', got '%s'", p.Dimension)
	}

	var unit DurationUnit
	switch {
	case p.Symbol == "s" || p.matchUnitByKey("second"):
		unit = Duration.Second
	case p.Symbol == "min" || p.matchUnitByKey("minute"):
		unit = Duration.Minute
	case p.Symbol == "h" || p.matchUnitByKey("hour"):
		unit = Duration.Hour
	case p.Symbol == "d" || p.matchUnitByKey("day"):
		unit = Duration.Day
	case p.Symbol == "ms" || p.matchUnitByKey("millisecond"):
		unit = Duration.Millisecond
	case p.Symbol == "µs" || p.matchUnitByKey("microsecond"):
		unit = Duration.Microsecond
	case p.Symbol == "ns" || p.matchUnitByKey("nanosecond"):
		unit = Duration.Nanosecond
	default:
		return Quantity[DurationUnit]{}, fmt.Errorf("unknown duration unit: symbol=%s, key=%s", p.Symbol, p.Key)
	}

	return NewDuration(p.Value, unit), nil
}

// MarshalAngle serializes an Angle measurement to JSON
func MarshalAngle(m Quantity[AngleUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalAngle deserializes a JSON representation to an Angle measurement
func UnmarshalAngle(data []byte) (Quantity[AngleUnit], error) {
	p, err := parseMeasurement(data)
	if err != nil {
		return Quantity[AngleUnit]{}, err
	}

	if p.Dimension != "angle" {
		return Quantity[AngleUnit]{}, fmt.Errorf("expected dimension 'angle', got '%s'", p.Dimension)
	}

	var unit AngleUnit
	switch {
	case p.Symbol == "rad" || p.matchUnitByKey("radian"):
		unit = Angle.Radian
	case p.Symbol == "°" || p.matchUnitByKey("degree"):
		unit = Angle.Degree
	case p.Symbol == "′" || p.matchUnitByKey("arcminute"):
		unit = Angle.Arcminute
	case p.Symbol == "″" || p.matchUnitByKey("arcsecond"):
		unit = Angle.Arcsecond
	case p.Symbol == "rev" || p.matchUnitByKey("revolution"):
		unit = Angle.Revolution
	case p.Symbol == "grad" || p.matchUnitByKey("gradian"):
		unit = Angle.Gradian
	default:
		return Quantity[AngleUnit]{}, fmt.Errorf("unknown angle unit: symbol=%s, key=%s", p.Symbol, p.Key)
	}

	return NewAngle(p.Value, unit), nil
}

// MarshalArea serializes an Area measurement to JSON
func MarshalArea(m Quantity[AreaUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalArea deserializes a JSON representation to an Area measurement
func UnmarshalArea(data []byte) (Quantity[AreaUnit], error) {
	p, err := parseMeasurement(data)
	if err != nil {
		return Quantity[AreaUnit]{}, err
	}

	if p.Dimension != "area" {
		return Quantity[AreaUnit]{}, fmt.Errorf("expected dimension 'area', got '%s'", p.Dimension)
	}

	var unit AreaUnit
	switch {
	case p.Symbol == "m²" || p.matchUnitByKey("square_meter"):
		unit = Area.SquareMeter
	case p.Symbol == "km²" || p.matchUnitByKey("square_kilometer"):
		unit = Area.SquareKilometer
	case p.Symbol == "cm²" || p.matchUnitByKey("square_centimeter"):
		unit = Area.SquareCentimeter
	case p.Symbol == "mm²" || p.matchUnitByKey("square_millimeter"):
		unit = Area.SquareMillimeter
	case p.Symbol == "in²" || p.matchUnitByKey("square_inch"):
		unit = Area.SquareInch
	case p.Symbol == "ft²" || p.matchUnitByKey("square_foot"):
		unit = Area.SquareFoot
	case p.Symbol == "yd²" || p.matchUnitByKey("square_yard"):
		unit = Area.SquareYard
	case p.Symbol == "mi²" || p.matchUnitByKey("square_mile"):
		unit = Area.SquareMile
	case p.Symbol == "ac" || p.matchUnitByKey("acre"):
		unit = Area.Acre
	case p.Symbol == "ha" || p.matchUnitByKey("hectare"):
		unit = Area.Hectare
	default:
		return Quantity[AreaUnit]{}, fmt.Errorf("unknown area unit: symbol=%s, key=%s", p.Symbol, p.Key)
	}

	return NewArea(p.Value, unit), nil
}

// MarshalVolume serializes a Volume measurement to JSON
func MarshalVolume(m Quantity[VolumeUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalVolume deserializes a JSON representation to a Volume measurement
func UnmarshalVolume(data []byte) (Quantity[VolumeUnit], error) {
	p, err := parseMeasurement(data)
	if err != nil {
		return Quantity[VolumeUnit]{}, err
	}

	if p.Dimension != "volume" {
		return Quantity[VolumeUnit]{}, fmt.Errorf("expected dimension 'volume', got '%s'", p.Dimension)
	}

	var unit VolumeUnit
	switch {
	case p.Symbol == "m³" || p.matchUnitByKey("cubic_meter"):
		unit = Volume.CubicMeter
	case p.Symbol == "km³" || p.matchUnitByKey("cubic_kilometer"):
		unit = Volume.CubicKilometer
	case p.Symbol == "cm³" || p.matchUnitByKey("cubic_centimeter"):
		unit = Volume.CubicCentimeter
	case p.Symbol == "mm³" || p.matchUnitByKey("cubic_millimeter"):
		unit = Volume.CubicMillimeter
	case p.Symbol == "L" || p.matchUnitByKey("liter"):
		unit = Volume.Liter
	case p.Symbol == "mL" || p.matchUnitByKey("milliliter"):
		unit = Volume.Milliliter
	case p.Symbol == "in³" || p.matchUnitByKey("cubic_inch"):
		unit = Volume.CubicInch
	case p.Symbol == "ft³" || p.matchUnitByKey("cubic_foot"):
		unit = Volume.CubicFoot
	case p.Symbol == "yd³" || p.matchUnitByKey("cubic_yard"):
		unit = Volume.CubicYard
	case p.Symbol == "gal" || p.matchUnitByKey("gallon"):
		unit = Volume.Gallon
	case p.Symbol == "qt" || p.matchUnitByKey("quart"):
		unit = Volume.Quart
	case p.Symbol == "pt" || p.matchUnitByKey("pint"):
		unit = Volume.Pint
	case p.Symbol == "cup" || p.matchUnitByKey("cup"):
		unit = Volume.Cup
	case p.Symbol == "fl oz" || p.matchUnitByKey("fluid_ounce"):
		unit = Volume.FluidOunce
	default:
		return Quantity[VolumeUnit]{}, fmt.Errorf("unknown volume unit: symbol=%s, key=%s", p.Symbol, p.Key)
	}

	return NewVolume(p.Value, unit), nil
}

// MarshalAcceleration serializes an Acceleration measurement to JSON
func MarshalAcceleration(m Quantity[AccelerationUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalAcceleration deserializes a JSON representation to an Acceleration measurement
func UnmarshalAcceleration(data []byte) (Quantity[AccelerationUnit], error) {
	p, err := parseMeasurement(data)
	if err != nil {
		return Quantity[AccelerationUnit]{}, err
	}

	if p.Dimension != "acceleration" {
		return Quantity[AccelerationUnit]{}, fmt.Errorf("expected dimension 'acceleration', got '%s'", p.Dimension)
	}

	var unit AccelerationUnit
	switch {
	case p.Symbol == "m/s²" || p.matchUnitByKey("meters_per_second_squared"):
		unit = Acceleration.MetersPerSecondSquared
	case p.Symbol == "g" || p.matchUnitByKey("g"):
		unit = Acceleration.G
	case p.Symbol == "ft/s²" || p.matchUnitByKey("feet_per_second_squared"):
		unit = Acceleration.FeetPerSecondSquared
	default:
		return Quantity[AccelerationUnit]{}, fmt.Errorf("unknown acceleration unit: symbol=%s, key=%s", p.Symbol, p.Key)
	}

	return NewAcceleration(p.Value, unit), nil
}

// MarshalConcentration serializes a Concentration measurement to JSON
func MarshalConcentration(m Quantity[ConcentrationUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalConcentration deserializes a JSON representation to a Concentration measurement
func UnmarshalConcentration(data []byte) (Quantity[ConcentrationUnit], error) {
	p, err := parseMeasurement(data)
	if err != nil {
		return Quantity[ConcentrationUnit]{}, err
	}

	if p.Dimension != "concentration" {
		return Quantity[ConcentrationUnit]{}, fmt.Errorf("expected dimension 'concentration', got '%s'", p.Dimension)
	}

	var unit ConcentrationUnit
	switch {
	case p.Symbol == "g/L" || p.matchUnitByKey("grams_per_liter"):
		unit = Concentration.GramsPerLiter
	case p.Symbol == "mg/L" || p.matchUnitByKey("milligrams_per_liter"):
		unit = Concentration.MilligramsPerLiter
	case p.Symbol == "ppm" || p.matchUnitByKey("parts_per_million"):
		unit = Concentration.PartsPerMillion
	case p.Symbol == "ppb" || p.matchUnitByKey("parts_per_billion"):
		unit = Concentration.PartsPerBillion
	default:
		return Quantity[ConcentrationUnit]{}, fmt.Errorf("unknown concentration unit: symbol=%s, key=%s", p.Symbol, p.Key)
	}

	return NewConcentration(p.Value, unit), nil
}

// MarshalDispersion serializes a Dispersion measurement to JSON
func MarshalDispersion(m Quantity[DispersionUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalDispersion deserializes a JSON representation to a Dispersion measurement
func UnmarshalDispersion(data []byte) (Quantity[DispersionUnit], error) {
	p, err := parseMeasurement(data)
	if err != nil {
		return Quantity[DispersionUnit]{}, err
	}

	if p.Dimension != "dispersion" {
		return Quantity[DispersionUnit]{}, fmt.Errorf("expected dimension 'dispersion', got '%s'", p.Dimension)
	}

	var unit DispersionUnit
	switch {
	case p.Symbol == "ppm" || p.matchUnitByKey("parts_per_million"):
		unit = Dispersion.PartsPerMillion
	case p.Symbol == "ppb" || p.matchUnitByKey("parts_per_billion"):
		unit = Dispersion.PartsPerBillion
	case p.Symbol == "ppt" || p.matchUnitByKey("parts_per_trillion"):
		unit = Dispersion.PartsPerTrillion
	case p.Symbol == "%" || p.matchUnitByKey("percent"):
		unit = Dispersion.Percent
	default:
		return Quantity[DispersionUnit]{}, fmt.Errorf("unknown dispersion unit: symbol=%s, key=%s", p.Symbol, p.Key)
	}

	return NewDispersion(p.Value, unit), nil
}

// MarshalElectricCharge serializes an ElectricCharge measurement to JSON
func MarshalElectricCharge(m Quantity[ElectricChargeUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalElectricCharge deserializes a JSON representation to an ElectricCharge measurement
func UnmarshalElectricCharge(data []byte) (Quantity[ElectricChargeUnit], error) {
	p, err := parseMeasurement(data)
	if err != nil {
		return Quantity[ElectricChargeUnit]{}, err
	}

	if p.Dimension != "electric_charge" {
		return Quantity[ElectricChargeUnit]{}, fmt.Errorf("expected dimension 'electric_charge', got '%s'", p.Dimension)
	}

	var unit ElectricChargeUnit
	switch {
	case p.Symbol == "C" || p.matchUnitByKey("coulomb"):
		unit = ElectricCharge.Coulomb
	case p.Symbol == "mC" || p.matchUnitByKey("millicoulomb"):
		unit = ElectricCharge.Millicoulomb
	case p.Symbol == "µC" || p.matchUnitByKey("microcoulomb"):
		unit = ElectricCharge.Microcoulomb
	case p.Symbol == "Ah" || p.matchUnitByKey("ampere_hour"):
		unit = ElectricCharge.Ampere_Hour
	case p.Symbol == "mAh" || p.matchUnitByKey("milliampere_hour"):
		unit = ElectricCharge.Milliampere_Hour
	default:
		return Quantity[ElectricChargeUnit]{}, fmt.Errorf("unknown electric charge unit: symbol=%s, key=%s", p.Symbol, p.Key)
	}

	return NewElectricCharge(p.Value, unit), nil
}

// MarshalElectricCurrent serializes an ElectricCurrent measurement to JSON
func MarshalElectricCurrent(m Quantity[ElectricCurrentUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalElectricCurrent deserializes a JSON representation to an ElectricCurrent measurement
func UnmarshalElectricCurrent(data []byte) (Quantity[ElectricCurrentUnit], error) {
	p, err := parseMeasurement(data)
	if err != nil {
		return Quantity[ElectricCurrentUnit]{}, err
	}

	if p.Dimension != "electric_current" {
		return Quantity[ElectricCurrentUnit]{}, fmt.Errorf("expected dimension 'electric_current', got '%s'", p.Dimension)
	}

	var unit ElectricCurrentUnit
	switch {
	case p.Symbol == "A" || p.matchUnitByKey("ampere"):
		unit = ElectricCurrent.Ampere
	case p.Symbol == "mA" || p.matchUnitByKey("milliampere"):
		unit = ElectricCurrent.Milliampere
	case p.Symbol == "µA" || p.matchUnitByKey("microampere"):
		unit = ElectricCurrent.Microampere
	case p.Symbol == "kA" || p.matchUnitByKey("kiloampere"):
		unit = ElectricCurrent.Kiloampere
	default:
		return Quantity[ElectricCurrentUnit]{}, fmt.Errorf("unknown electric current unit: symbol=%s, key=%s", p.Symbol, p.Key)
	}

	return NewElectricCurrent(p.Value, unit), nil
}

// MarshalSpeed serializes a Speed measurement to JSON
func MarshalSpeed(m Quantity[SpeedUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalSpeed deserializes a JSON representation to a Speed measurement
func UnmarshalSpeed(data []byte) (Quantity[SpeedUnit], error) {
	p, err := parseMeasurement(data)
	if err != nil {
		return Quantity[SpeedUnit]{}, err
	}

	if p.Dimension != "speed" {
		return Quantity[SpeedUnit]{}, fmt.Errorf("expected dimension 'speed', got '%s'", p.Dimension)
	}

	var unit SpeedUnit
	switch {
	case p.Symbol == "m/s" || p.matchUnitByKey("meters_per_second"):
		unit = Speed.MetersPerSecond
	case p.Symbol == "km/h" || p.matchUnitByKey("kilometers_per_hour"):
		unit = Speed.KilometersPerHour
	case p.Symbol == "mph" || p.matchUnitByKey("miles_per_hour"):
		unit = Speed.MilesPerHour
	case p.Symbol == "ft/s" || p.matchUnitByKey("feet_per_second"):
		unit = Speed.FeetPerSecond
	case p.Symbol == "kn" || p.matchUnitByKey("knot"):
		unit = Speed.Knot
	default:
		return Quantity[SpeedUnit]{}, fmt.Errorf("unknown speed unit: symbol=%s, key=%s", p.Symbol, p.Key)
	}

	return NewSpeed(p.Value, unit), nil
}

// MarshalElectricPotentialDifference serializes an ElectricPotentialDifference measurement to JSON
func MarshalElectricPotentialDifference(m Quantity[ElectricPotentialDifferenceUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalElectricPotentialDifference deserializes a JSON representation to an ElectricPotentialDifference measurement
func UnmarshalElectricPotentialDifference(data []byte) (Quantity[ElectricPotentialDifferenceUnit], error) {
	p, err := parseMeasurement(data)
	if err != nil {
		return Quantity[ElectricPotentialDifferenceUnit]{}, err
	}

	if p.Dimension != "electric_potential_difference" {
		return Quantity[ElectricPotentialDifferenceUnit]{}, fmt.Errorf("expected dimension 'electric_potential_difference', got '%s'", p.Dimension)
	}

	var unit ElectricPotentialDifferenceUnit
	switch {
	case p.Symbol == "V" || p.matchUnitByKey("volt"):
		unit = ElectricPotentialDifference.Volt
	case p.Symbol == "mV" || p.matchUnitByKey("millivolt"):
		unit = ElectricPotentialDifference.Millivolt
	case p.Symbol == "µV" || p.matchUnitByKey("microvolt"):
		unit = ElectricPotentialDifference.Microvolt
	case p.Symbol == "kV" || p.matchUnitByKey("kilovolt"):
		unit = ElectricPotentialDifference.Kilovolt
	case p.Symbol == "MV" || p.matchUnitByKey("megavolt"):
		unit = ElectricPotentialDifference.Megavolt
	default:
		return Quantity[ElectricPotentialDifferenceUnit]{}, fmt.Errorf("unknown electric potential difference unit: symbol=%s, key=%s", p.Symbol, p.Key)
	}

	return NewElectricPotentialDifference(p.Value, unit), nil
}

// MarshalInformation serializes an Information measurement to JSON
func MarshalInformation(m Quantity[InformationUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalInformation deserializes a JSON representation to an Information measurement
func UnmarshalInformation(data []byte) (Quantity[InformationUnit], error) {
	p, err := parseMeasurement(data)
	if err != nil {
		return Quantity[InformationUnit]{}, err
	}

	if p.Dimension != "information" {
		return Quantity[InformationUnit]{}, fmt.Errorf("expected dimension 'information', got '%s'", p.Dimension)
	}

	var unit InformationUnit
	switch {
	case p.Symbol == "bit" || p.matchUnitByKey("bit"):
		unit = Information.Bit
	case p.Symbol == "B" || p.matchUnitByKey("byte"):
		unit = Information.Byte
	case p.Symbol == "KB" || p.matchUnitByKey("kilobyte"):
		unit = Information.Kilobyte
	case p.Symbol == "MB" || p.matchUnitByKey("megabyte"):
		unit = Information.Megabyte
	case p.Symbol == "GB" || p.matchUnitByKey("gigabyte"):
		unit = Information.Gigabyte
	case p.Symbol == "TB" || p.matchUnitByKey("terabyte"):
		unit = Information.Terabyte
	case p.Symbol == "PB" || p.matchUnitByKey("petabyte"):
		unit = Information.Petabyte
	case p.Symbol == "KiB" || p.matchUnitByKey("kibibyte"):
		unit = Information.Kibibyte
	case p.Symbol == "MiB" || p.matchUnitByKey("mebibyte"):
		unit = Information.Mebibyte
	case p.Symbol == "GiB" || p.matchUnitByKey("gibibyte"):
		unit = Information.Gibibyte
	case p.Symbol == "TiB" || p.matchUnitByKey("tebibyte"):
		unit = Information.Tebibyte
	case p.Symbol == "PiB" || p.matchUnitByKey("pebibyte"):
		unit = Information.Pebibyte
	default:
		return Quantity[InformationUnit]{}, fmt.Errorf("unknown information unit: symbol=%s, key=%s", p.Symbol, p.Key)
	}

	return NewInformation(p.Value, unit), nil
}

// MarshalFrequency serializes a Frequency measurement to JSON
func MarshalFrequency(m Quantity[FrequencyUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalFrequency deserializes a JSON representation to a Frequency measurement
func UnmarshalFrequency(data []byte) (Quantity[FrequencyUnit], error) {
	p, err := parseMeasurement(data)
	if err != nil {
		return Quantity[FrequencyUnit]{}, err
	}

	if p.Dimension != "frequency" {
		return Quantity[FrequencyUnit]{}, fmt.Errorf("expected dimension 'frequency', got '%s'", p.Dimension)
	}

	var unit FrequencyUnit
	switch {
	case p.Symbol == "Hz" || p.matchUnitByKey("hertz"):
		unit = Frequency.Hertz
	case p.Symbol == "kHz" || p.matchUnitByKey("kilohertz"):
		unit = Frequency.Kilohertz
	case p.Symbol == "MHz" || p.matchUnitByKey("megahertz"):
		unit = Frequency.Megahertz
	case p.Symbol == "GHz" || p.matchUnitByKey("gigahertz"):
		unit = Frequency.Gigahertz
	case p.Symbol == "THz" || p.matchUnitByKey("terahertz"):
		unit = Frequency.Terahertz
	case p.Symbol == "rpm" || p.matchUnitByKey("rpm"):
		unit = Frequency.RPM
	default:
		return Quantity[FrequencyUnit]{}, fmt.Errorf("unknown frequency unit: symbol=%s, key=%s", p.Symbol, p.Key)
	}

	return NewFrequency(p.Value, unit), nil
}

// MarshalIlluminance serializes an Illuminance measurement to JSON
func MarshalIlluminance(m Quantity[IlluminanceUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalIlluminance deserializes a JSON representation to an Illuminance measurement
func UnmarshalIlluminance(data []byte) (Quantity[IlluminanceUnit], error) {
	p, err := parseMeasurement(data)
	if err != nil {
		return Quantity[IlluminanceUnit]{}, err
	}

	if p.Dimension != "illuminance" {
		return Quantity[IlluminanceUnit]{}, fmt.Errorf("expected dimension 'illuminance', got '%s'", p.Dimension)
	}

	var unit IlluminanceUnit
	switch {
	case p.Symbol == "lx" || p.matchUnitByKey("lux"):
		unit = Illuminance.Lux
	case p.Symbol == "fc" || p.matchUnitByKey("foot_candle"):
		unit = Illuminance.FootCandle
	case p.Symbol == "ph" || p.matchUnitByKey("phot"):
		unit = Illuminance.Phot
	case p.Symbol == "nx" || p.matchUnitByKey("nox"):
		unit = Illuminance.Nox
	default:
		return Quantity[IlluminanceUnit]{}, fmt.Errorf("unknown illuminance unit: symbol=%s, key=%s", p.Symbol, p.Key)
	}

	return NewIlluminance(p.Value, unit), nil
}

// MarshalGeneral serializes a General measurement to JSON
func MarshalGeneral(m Quantity[GeneralUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// MarshalFuelEfficiency serializes a FuelEfficiency measurement to JSON
func MarshalFuelEfficiency(m Quantity[FuelEfficiencyUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalGeneral deserializes a JSON representation to a General measurement
func UnmarshalGeneral(data []byte) (Quantity[GeneralUnit], error) {
	p, err := parseMeasurement(data)
	if err != nil {
		return Quantity[GeneralUnit]{}, err
	}

	if p.Dimension != "general" {
		return Quantity[GeneralUnit]{}, fmt.Errorf("expected dimension 'general', got '%s'", p.Dimension)
	}

	var unit GeneralUnit
	switch {
	case p.Symbol == "unit" || p.matchUnitByKey("unit"):
		unit = General.Unit
	default:
		// For custom units, create a new general unit with the given symbol and name
		symbolOrKey := p.Symbol
		if symbolOrKey == "" && p.Key != "" {
			_, symbolOrKey = parseUnitKey(p.Key)
		}
		name := p.Name
		if name == "" {
			name = symbolOrKey
		}
		unit = NewGeneralUnit(symbolOrKey, name)
	}

	return NewGeneral(p.Value, unit), nil
}

// UnmarshalFuelEfficiency deserializes a JSON representation to a FuelEfficiency measurement
func UnmarshalFuelEfficiency(data []byte) (Quantity[FuelEfficiencyUnit], error) {
	p, err := parseMeasurement(data)
	if err != nil {
		return Quantity[FuelEfficiencyUnit]{}, err
	}

	if p.Dimension != "fuel_efficiency" {
		return Quantity[FuelEfficiencyUnit]{}, fmt.Errorf("expected dimension 'fuel_efficiency', got '%s'", p.Dimension)
	}

	var unit FuelEfficiencyUnit
	switch {
	case p.Symbol == "km/L" || p.matchUnitByKey("kilometers_per_liter"):
		unit = FuelEfficiency.KilometersPerLiter
	case p.Symbol == "mpg" || p.matchUnitByKey("miles_per_gallon"):
		unit = FuelEfficiency.MilesPerGallon
	case p.Symbol == "L/100km" || p.matchUnitByKey("liters_per_100_kilometers"):
		unit = FuelEfficiency.LitersPer100Kilometers
	default:
		return Quantity[FuelEfficiencyUnit]{}, fmt.Errorf("unknown fuel efficiency unit: symbol=%s, key=%s", p.Symbol, p.Key)
	}

	return NewFuelEfficiency(p.Value, unit), nil
}
