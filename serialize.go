// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

import (
	"encoding/json"
	"fmt"
)

// MeasurementJSON is used for JSON serialization of measurements
type MeasurementJSON struct {
	Value     float64 `json:"value" yaml:"value"`
	Unit      string  `json:"unit" yaml:"unit"`
	Dimension string  `json:"dimension" yaml:"dimension"`
}

// QuantityJSON is a non-generic struct used for JSON serialization
type QuantityJSON struct {
	Value     float64 `json:"value" yaml:"value"`
	Unit      string  `json:"unit" yaml:"unit"`
	Dimension string  `json:"dimension" yaml:"dimension"`
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
func UnmarshalMeasurement(data []byte) (*AnyMeasurement, error) {
	var jsonM MeasurementJSON
	if err := json.Unmarshal(data, &jsonM); err != nil {
		return nil, err
	}

	// Based on the dimension, call the appropriate unmarshal function
	switch jsonM.Dimension {
	case "temperature":
		m, err := UnmarshalTemperature(data)
		if err != nil {
			// If there's a problem with the temperature dimension, use general unit
			return fallbackToGeneral(data, jsonM, err)
		}
		return &AnyMeasurement{value: m, dimension: "temperature"}, nil
	case "pressure":
		m, err := UnmarshalPressure(data)
		if err != nil {
			// If there's a problem with the pressure dimension, use general unit
			return fallbackToGeneral(data, jsonM, err)
		}
		return &AnyMeasurement{value: m, dimension: "pressure"}, nil
	case "flowrate":
		m, err := UnmarshalFlowRate(data)
		if err != nil {
			// If there's a problem with the flowrate dimension, use general unit
			return fallbackToGeneral(data, jsonM, err)
		}
		return &AnyMeasurement{value: m, dimension: "flowrate"}, nil
	case "power":
		m, err := UnmarshalPower(data)
		if err != nil {
			// If there's a problem with the power dimension, use general unit
			return fallbackToGeneral(data, jsonM, err)
		}
		return &AnyMeasurement{value: m, dimension: "power"}, nil
	case "energy":
		m, err := UnmarshalEnergy(data)
		if err != nil {
			// If there's a problem with the energy dimension, use general unit
			return fallbackToGeneral(data, jsonM, err)
		}
		return &AnyMeasurement{value: m, dimension: "energy"}, nil
	case "length":
		m, err := UnmarshalLength(data)
		if err != nil {
			// If there's a problem with the length dimension, use general unit
			return fallbackToGeneral(data, jsonM, err)
		}
		return &AnyMeasurement{value: m, dimension: "length"}, nil
	case "mass":
		m, err := UnmarshalMass(data)
		if err != nil {
			// If there's a problem with the mass dimension, use general unit
			return fallbackToGeneral(data, jsonM, err)
		}
		return &AnyMeasurement{value: m, dimension: "mass"}, nil
	case "duration":
		m, err := UnmarshalDuration(data)
		if err != nil {
			// If there's a problem with the duration dimension, use general unit
			return fallbackToGeneral(data, jsonM, err)
		}
		return &AnyMeasurement{value: m, dimension: "duration"}, nil
	case "angle":
		m, err := UnmarshalAngle(data)
		if err != nil {
			// If there's a problem with the angle dimension, use general unit
			return fallbackToGeneral(data, jsonM, err)
		}
		return &AnyMeasurement{value: m, dimension: "angle"}, nil
	case "area":
		m, err := UnmarshalArea(data)
		if err != nil {
			// If there's a problem with the area dimension, use general unit
			return fallbackToGeneral(data, jsonM, err)
		}
		return &AnyMeasurement{value: m, dimension: "area"}, nil
	case "volume":
		m, err := UnmarshalVolume(data)
		if err != nil {
			// If there's a problem with the volume dimension, use general unit
			return fallbackToGeneral(data, jsonM, err)
		}
		return &AnyMeasurement{value: m, dimension: "volume"}, nil
	case "acceleration":
		m, err := UnmarshalAcceleration(data)
		if err != nil {
			// If there's a problem with the acceleration dimension, use general unit
			return fallbackToGeneral(data, jsonM, err)
		}
		return &AnyMeasurement{value: m, dimension: "acceleration"}, nil
	case "concentration":
		m, err := UnmarshalConcentration(data)
		if err != nil {
			// If there's a problem with the concentration dimension, use general unit
			return fallbackToGeneral(data, jsonM, err)
		}
		return &AnyMeasurement{value: m, dimension: "concentration"}, nil
	case "dispersion":
		m, err := UnmarshalDispersion(data)
		if err != nil {
			// If there's a problem with the dispersion dimension, use general unit
			return fallbackToGeneral(data, jsonM, err)
		}
		return &AnyMeasurement{value: m, dimension: "dispersion"}, nil
	case "electric_charge":
		m, err := UnmarshalElectricCharge(data)
		if err != nil {
			// If there's a problem with the electric_charge dimension, use general unit
			return fallbackToGeneral(data, jsonM, err)
		}
		return &AnyMeasurement{value: m, dimension: "electric_charge"}, nil
	case "electric_current":
		m, err := UnmarshalElectricCurrent(data)
		if err != nil {
			// If there's a problem with the electric_current dimension, use general unit
			return fallbackToGeneral(data, jsonM, err)
		}
		return &AnyMeasurement{value: m, dimension: "electric_current"}, nil
	case "electric_potential_difference":
		m, err := UnmarshalElectricPotentialDifference(data)
		if err != nil {
			// If there's a problem with the electric_potential_difference dimension, use general unit
			return fallbackToGeneral(data, jsonM, err)
		}
		return &AnyMeasurement{value: m, dimension: "electric_potential_difference"}, nil
	case "information":
		m, err := UnmarshalInformation(data)
		if err != nil {
			// If there's a problem with the information dimension, use general unit
			return fallbackToGeneral(data, jsonM, err)
		}
		return &AnyMeasurement{value: m, dimension: "information"}, nil
	case "frequency":
		m, err := UnmarshalFrequency(data)
		if err != nil {
			// If there's a problem with the frequency dimension, use general unit
			return fallbackToGeneral(data, jsonM, err)
		}
		return &AnyMeasurement{value: m, dimension: "frequency"}, nil
	case "illuminance":
		m, err := UnmarshalIlluminance(data)
		if err != nil {
			// If there's a problem with the illuminance dimension, use general unit
			return fallbackToGeneral(data, jsonM, err)
		}
		return &AnyMeasurement{value: m, dimension: "illuminance"}, nil
	case "fuel_efficiency":
		m, err := UnmarshalFuelEfficiency(data)
		if err != nil {
			// If there's a problem with the fuel_efficiency dimension, use general unit
			return fallbackToGeneral(data, jsonM, err)
		}
		return &AnyMeasurement{value: m, dimension: "fuel_efficiency"}, nil
	case "general":
		m, err := UnmarshalGeneral(data)
		if err != nil {
			return nil, err
		}
		return &AnyMeasurement{value: m, dimension: "general"}, nil
	default:
		// For unknown dimensions, use general unit
		return fallbackToGeneral(data, jsonM, fmt.Errorf("unknown dimension: %s", jsonM.Dimension))
	}
}

// fallbackToGeneral creates a general measurement from the given JSON data
func fallbackToGeneral(data []byte, jsonM MeasurementJSON, originalErr error) (*AnyMeasurement, error) {
	// Create a new JSON object with the general dimension
	generalJSON := MeasurementJSON{
		Value:     jsonM.Value,
		Unit:      jsonM.Unit,
		Dimension: "general",
	}

	// Marshal the general JSON object
	generalData, err := json.Marshal(generalJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to create general measurement: %w", err)
	}

	// Unmarshal as a general measurement
	m, err := UnmarshalGeneral(generalData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal as general measurement: %w (original error: %v)", err, originalErr)
	}

	return &AnyMeasurement{value: m, dimension: "general"}, nil
}

// marshalGeneric is a helper function to serialize any measurement to JSON
func marshalGeneric[T Category](m Quantity[T]) ([]byte, error) {
	return json.Marshal(MeasurementJSON{
		Value:     m.Value,
		Unit:      m.Unit.Symbol(),
		Dimension: m.Unit.Dimension(),
	})
}

// MarshalTemperature serializes a Temperature measurement to JSON
func MarshalTemperature(m Quantity[TemperatureUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalTemperature deserializes a JSON representation to a Temperature measurement
func UnmarshalTemperature(data []byte) (Quantity[TemperatureUnit], error) {
	var jsonM MeasurementJSON
	if err := json.Unmarshal(data, &jsonM); err != nil {
		return Quantity[TemperatureUnit]{}, err
	}

	// Verify dimension
	if jsonM.Dimension != "temperature" {
		return Quantity[TemperatureUnit]{}, fmt.Errorf("expected dimension 'temperature', got '%s'", jsonM.Dimension)
	}

	// Find the matching temperature unit
	var unit TemperatureUnit
	found := false

	switch jsonM.Unit {
	case "°C", "C":
		unit = Temperature.Celsius
		found = true
	case "°F", "F":
		unit = Temperature.Fahrenheit
		found = true
	case "K":
		unit = Temperature.Kelvin
		found = true
	}

	if !found {
		return Quantity[TemperatureUnit]{}, fmt.Errorf("unknown temperature unit: %s", jsonM.Unit)
	}

	return NewTemperature(jsonM.Value, unit), nil
}

// MarshalPressure serializes a Pressure measurement to JSON
func MarshalPressure(m Quantity[PressureUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalPressure deserializes a JSON representation to a Pressure measurement
func UnmarshalPressure(data []byte) (Quantity[PressureUnit], error) {
	var jsonM MeasurementJSON
	if err := json.Unmarshal(data, &jsonM); err != nil {
		return Quantity[PressureUnit]{}, err
	}

	// Verify dimension
	if jsonM.Dimension != "pressure" {
		return Quantity[PressureUnit]{}, fmt.Errorf("expected dimension 'pressure', got '%s'", jsonM.Dimension)
	}

	// Find the matching pressure unit
	var unit PressureUnit
	found := false

	switch jsonM.Unit {
	case "Pa":
		unit = Pressure.Pascal
		found = true
	case "kPa":
		unit = Pressure.Kilopascal
		found = true
	case "bar":
		unit = Pressure.Bar
		found = true
	case "psi":
		unit = Pressure.PSI
		found = true
	case "inH₂O", "inH2O":
		unit = Pressure.InchH2O
		found = true
	}

	if !found {
		return Quantity[PressureUnit]{}, fmt.Errorf("unknown pressure unit: %s", jsonM.Unit)
	}

	return NewPressure(jsonM.Value, unit), nil
}

// MarshalFlowRate serializes a FlowRate measurement to JSON
func MarshalFlowRate(m Quantity[FlowRateUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalFlowRate deserializes a JSON representation to a FlowRate measurement
func UnmarshalFlowRate(data []byte) (Quantity[FlowRateUnit], error) {
	var jsonM MeasurementJSON
	if err := json.Unmarshal(data, &jsonM); err != nil {
		return Quantity[FlowRateUnit]{}, err
	}

	// Verify dimension
	if jsonM.Dimension != "flowrate" {
		return Quantity[FlowRateUnit]{}, fmt.Errorf("expected dimension 'flowrate', got '%s'", jsonM.Dimension)
	}

	// Find the matching flowrate unit
	var unit FlowRateUnit
	found := false

	switch jsonM.Unit {
	case "m³/h":
		unit = FlowRate.CubicMetersPerHour
		found = true
	case "L/s":
		unit = FlowRate.LitersPerSecond
		found = true
	case "CFM":
		unit = FlowRate.CFM
		found = true
	}

	if !found {
		return Quantity[FlowRateUnit]{}, fmt.Errorf("unknown flowrate unit: %s", jsonM.Unit)
	}

	return NewFlowRate(jsonM.Value, unit), nil
}

// MarshalPower serializes a Power measurement to JSON
func MarshalPower(m Quantity[PowerUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalPower deserializes a JSON representation to a Power measurement
func UnmarshalPower(data []byte) (Quantity[PowerUnit], error) {
	var jsonM MeasurementJSON
	if err := json.Unmarshal(data, &jsonM); err != nil {
		return Quantity[PowerUnit]{}, err
	}

	// Verify dimension
	if jsonM.Dimension != "power" {
		return Quantity[PowerUnit]{}, fmt.Errorf("expected dimension 'power', got '%s'", jsonM.Dimension)
	}

	// Find the matching power unit
	var unit PowerUnit
	found := false

	switch jsonM.Unit {
	case "W":
		unit = Power.Watt
		found = true
	case "kW":
		unit = Power.Kilowatt
		found = true
	case "BTU/h":
		unit = Power.BTUPerHour
		found = true
	}

	if !found {
		return Quantity[PowerUnit]{}, fmt.Errorf("unknown power unit: %s", jsonM.Unit)
	}

	return NewPower(jsonM.Value, unit), nil
}

// MarshalEnergy serializes an Energy measurement to JSON
func MarshalEnergy(m Quantity[EnergyUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalEnergy deserializes a JSON representation to an Energy measurement
func UnmarshalEnergy(data []byte) (Quantity[EnergyUnit], error) {
	var jsonM MeasurementJSON
	if err := json.Unmarshal(data, &jsonM); err != nil {
		return Quantity[EnergyUnit]{}, err
	}

	// Verify dimension
	if jsonM.Dimension != "energy" {
		return Quantity[EnergyUnit]{}, fmt.Errorf("expected dimension 'energy', got '%s'", jsonM.Dimension)
	}

	// Find the matching energy unit
	var unit EnergyUnit
	found := false

	switch jsonM.Unit {
	case "J":
		unit = Energy.Joule
		found = true
	case "kWh":
		unit = Energy.KilowattHour
		found = true
	case "BTU":
		unit = Energy.BTU
		found = true
	}

	if !found {
		return Quantity[EnergyUnit]{}, fmt.Errorf("unknown energy unit: %s", jsonM.Unit)
	}

	return NewEnergy(jsonM.Value, unit), nil
}

// MarshalLength serializes a Length measurement to JSON
func MarshalLength(m Quantity[LengthUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalLength deserializes a JSON representation to a Length measurement
func UnmarshalLength(data []byte) (Quantity[LengthUnit], error) {
	var jsonM MeasurementJSON
	if err := json.Unmarshal(data, &jsonM); err != nil {
		return Quantity[LengthUnit]{}, err
	}

	// Verify dimension
	if jsonM.Dimension != "length" {
		return Quantity[LengthUnit]{}, fmt.Errorf("expected dimension 'length', got '%s'", jsonM.Dimension)
	}

	// Find the matching length unit
	var unit LengthUnit
	found := false

	switch jsonM.Unit {
	case "m":
		unit = Length.Meter
		found = true
	case "km":
		unit = Length.Kilometer
		found = true
	case "cm":
		unit = Length.Centimeter
		found = true
	case "mm":
		unit = Length.Millimeter
		found = true
	case "µm":
		unit = Length.Micrometer
		found = true
	case "nm":
		unit = Length.Nanometer
		found = true
	case "in":
		unit = Length.Inch
		found = true
	case "ft":
		unit = Length.Foot
		found = true
	case "yd":
		unit = Length.Yard
		found = true
	case "mi":
		unit = Length.Mile
		found = true
	}

	if !found {
		return Quantity[LengthUnit]{}, fmt.Errorf("unknown length unit: %s", jsonM.Unit)
	}

	return NewLength(jsonM.Value, unit), nil
}

// MarshalMass serializes a Mass measurement to JSON
func MarshalMass(m Quantity[MassUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalMass deserializes a JSON representation to a Mass measurement
func UnmarshalMass(data []byte) (Quantity[MassUnit], error) {
	var jsonM MeasurementJSON
	if err := json.Unmarshal(data, &jsonM); err != nil {
		return Quantity[MassUnit]{}, err
	}

	// Verify dimension
	if jsonM.Dimension != "mass" {
		return Quantity[MassUnit]{}, fmt.Errorf("expected dimension 'mass', got '%s'", jsonM.Dimension)
	}

	// Find the matching mass unit
	var unit MassUnit
	found := false

	switch jsonM.Unit {
	case "kg":
		unit = Mass.Kilogram
		found = true
	case "g":
		unit = Mass.Gram
		found = true
	case "mg":
		unit = Mass.Milligram
		found = true
	case "µg":
		unit = Mass.Microgram
		found = true
	case "lb":
		unit = Mass.Pound
		found = true
	case "oz":
		unit = Mass.Ounce
		found = true
	case "st":
		unit = Mass.Stone
		found = true
	case "t":
		unit = Mass.MetricTon
		found = true
	case "ton":
		unit = Mass.Ton
		found = true
	}

	if !found {
		return Quantity[MassUnit]{}, fmt.Errorf("unknown mass unit: %s", jsonM.Unit)
	}

	return NewMass(jsonM.Value, unit), nil
}

// MarshalDuration serializes a Duration measurement to JSON
func MarshalDuration(m Quantity[DurationUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalDuration deserializes a JSON representation to a Duration measurement
func UnmarshalDuration(data []byte) (Quantity[DurationUnit], error) {
	var jsonM MeasurementJSON
	if err := json.Unmarshal(data, &jsonM); err != nil {
		return Quantity[DurationUnit]{}, err
	}

	// Verify dimension
	if jsonM.Dimension != "duration" {
		return Quantity[DurationUnit]{}, fmt.Errorf("expected dimension 'duration', got '%s'", jsonM.Dimension)
	}

	// Find the matching duration unit
	var unit DurationUnit
	found := false

	switch jsonM.Unit {
	case "s":
		unit = Duration.Second
		found = true
	case "min":
		unit = Duration.Minute
		found = true
	case "h":
		unit = Duration.Hour
		found = true
	case "d":
		unit = Duration.Day
		found = true
	case "ms":
		unit = Duration.Millisecond
		found = true
	case "µs":
		unit = Duration.Microsecond
		found = true
	case "ns":
		unit = Duration.Nanosecond
		found = true
	}

	if !found {
		return Quantity[DurationUnit]{}, fmt.Errorf("unknown duration unit: %s", jsonM.Unit)
	}

	return NewDuration(jsonM.Value, unit), nil
}

// MarshalAngle serializes an Angle measurement to JSON
func MarshalAngle(m Quantity[AngleUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalAngle deserializes a JSON representation to an Angle measurement
func UnmarshalAngle(data []byte) (Quantity[AngleUnit], error) {
	var jsonM MeasurementJSON
	if err := json.Unmarshal(data, &jsonM); err != nil {
		return Quantity[AngleUnit]{}, err
	}

	// Verify dimension
	if jsonM.Dimension != "angle" {
		return Quantity[AngleUnit]{}, fmt.Errorf("expected dimension 'angle', got '%s'", jsonM.Dimension)
	}

	// Find the matching angle unit
	var unit AngleUnit
	found := false

	switch jsonM.Unit {
	case "rad":
		unit = Angle.Radian
		found = true
	case "°":
		unit = Angle.Degree
		found = true
	case "′":
		unit = Angle.Arcminute
		found = true
	case "″":
		unit = Angle.Arcsecond
		found = true
	case "rev":
		unit = Angle.Revolution
		found = true
	case "grad":
		unit = Angle.Gradian
		found = true
	}

	if !found {
		return Quantity[AngleUnit]{}, fmt.Errorf("unknown angle unit: %s", jsonM.Unit)
	}

	return NewAngle(jsonM.Value, unit), nil
}

// MarshalArea serializes an Area measurement to JSON
func MarshalArea(m Quantity[AreaUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalArea deserializes a JSON representation to an Area measurement
func UnmarshalArea(data []byte) (Quantity[AreaUnit], error) {
	var jsonM MeasurementJSON
	if err := json.Unmarshal(data, &jsonM); err != nil {
		return Quantity[AreaUnit]{}, err
	}

	// Verify dimension
	if jsonM.Dimension != "area" {
		return Quantity[AreaUnit]{}, fmt.Errorf("expected dimension 'area', got '%s'", jsonM.Dimension)
	}

	// Find the matching area unit
	var unit AreaUnit
	found := false

	switch jsonM.Unit {
	case "m²":
		unit = Area.SquareMeter
		found = true
	case "km²":
		unit = Area.SquareKilometer
		found = true
	case "cm²":
		unit = Area.SquareCentimeter
		found = true
	case "mm²":
		unit = Area.SquareMillimeter
		found = true
	case "in²":
		unit = Area.SquareInch
		found = true
	case "ft²":
		unit = Area.SquareFoot
		found = true
	case "yd²":
		unit = Area.SquareYard
		found = true
	case "mi²":
		unit = Area.SquareMile
		found = true
	case "ac":
		unit = Area.Acre
		found = true
	case "ha":
		unit = Area.Hectare
		found = true
	}

	if !found {
		return Quantity[AreaUnit]{}, fmt.Errorf("unknown area unit: %s", jsonM.Unit)
	}

	return NewArea(jsonM.Value, unit), nil
}

// MarshalVolume serializes a Volume measurement to JSON
func MarshalVolume(m Quantity[VolumeUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalVolume deserializes a JSON representation to a Volume measurement
func UnmarshalVolume(data []byte) (Quantity[VolumeUnit], error) {
	var jsonM MeasurementJSON
	if err := json.Unmarshal(data, &jsonM); err != nil {
		return Quantity[VolumeUnit]{}, err
	}

	// Verify dimension
	if jsonM.Dimension != "volume" {
		return Quantity[VolumeUnit]{}, fmt.Errorf("expected dimension 'volume', got '%s'", jsonM.Dimension)
	}

	// Find the matching volume unit
	var unit VolumeUnit
	found := false

	switch jsonM.Unit {
	case "m³":
		unit = Volume.CubicMeter
		found = true
	case "km³":
		unit = Volume.CubicKilometer
		found = true
	case "cm³":
		unit = Volume.CubicCentimeter
		found = true
	case "mm³":
		unit = Volume.CubicMillimeter
		found = true
	case "L":
		unit = Volume.Liter
		found = true
	case "mL":
		unit = Volume.Milliliter
		found = true
	case "in³":
		unit = Volume.CubicInch
		found = true
	case "ft³":
		unit = Volume.CubicFoot
		found = true
	case "yd³":
		unit = Volume.CubicYard
		found = true
	case "gal":
		unit = Volume.Gallon
		found = true
	case "qt":
		unit = Volume.Quart
		found = true
	case "pt":
		unit = Volume.Pint
		found = true
	case "cup":
		unit = Volume.Cup
		found = true
	case "fl oz":
		unit = Volume.FluidOunce
		found = true
	}

	if !found {
		return Quantity[VolumeUnit]{}, fmt.Errorf("unknown volume unit: %s", jsonM.Unit)
	}

	return NewVolume(jsonM.Value, unit), nil
}

// MarshalAcceleration serializes an Acceleration measurement to JSON
func MarshalAcceleration(m Quantity[AccelerationUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalAcceleration deserializes a JSON representation to an Acceleration measurement
func UnmarshalAcceleration(data []byte) (Quantity[AccelerationUnit], error) {
	var jsonM MeasurementJSON
	if err := json.Unmarshal(data, &jsonM); err != nil {
		return Quantity[AccelerationUnit]{}, err
	}

	// Verify dimension
	if jsonM.Dimension != "acceleration" {
		return Quantity[AccelerationUnit]{}, fmt.Errorf("expected dimension 'acceleration', got '%s'", jsonM.Dimension)
	}

	// Find the matching acceleration unit
	var unit AccelerationUnit
	found := false

	switch jsonM.Unit {
	case "m/s²":
		unit = Acceleration.MetersPerSecondSquared
		found = true
	case "g":
		unit = Acceleration.G
		found = true
	case "ft/s²":
		unit = Acceleration.FeetPerSecondSquared
		found = true
	}

	if !found {
		return Quantity[AccelerationUnit]{}, fmt.Errorf("unknown acceleration unit: %s", jsonM.Unit)
	}

	return NewAcceleration(jsonM.Value, unit), nil
}

// MarshalConcentration serializes a Concentration measurement to JSON
func MarshalConcentration(m Quantity[ConcentrationUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalConcentration deserializes a JSON representation to a Concentration measurement
func UnmarshalConcentration(data []byte) (Quantity[ConcentrationUnit], error) {
	var jsonM MeasurementJSON
	if err := json.Unmarshal(data, &jsonM); err != nil {
		return Quantity[ConcentrationUnit]{}, err
	}

	// Verify dimension
	if jsonM.Dimension != "concentration" {
		return Quantity[ConcentrationUnit]{}, fmt.Errorf("expected dimension 'concentration', got '%s'", jsonM.Dimension)
	}

	// Find the matching concentration unit
	var unit ConcentrationUnit
	found := false

	switch jsonM.Unit {
	case "g/L":
		unit = Concentration.GramsPerLiter
		found = true
	case "mg/L":
		unit = Concentration.MilligramsPerLiter
		found = true
	case "ppm":
		unit = Concentration.PartsPerMillion
		found = true
	case "ppb":
		unit = Concentration.PartsPerBillion
		found = true
	}

	if !found {
		return Quantity[ConcentrationUnit]{}, fmt.Errorf("unknown concentration unit: %s", jsonM.Unit)
	}

	return NewConcentration(jsonM.Value, unit), nil
}

// MarshalDispersion serializes a Dispersion measurement to JSON
func MarshalDispersion(m Quantity[DispersionUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalDispersion deserializes a JSON representation to a Dispersion measurement
func UnmarshalDispersion(data []byte) (Quantity[DispersionUnit], error) {
	var jsonM MeasurementJSON
	if err := json.Unmarshal(data, &jsonM); err != nil {
		return Quantity[DispersionUnit]{}, err
	}

	// Verify dimension
	if jsonM.Dimension != "dispersion" {
		return Quantity[DispersionUnit]{}, fmt.Errorf("expected dimension 'dispersion', got '%s'", jsonM.Dimension)
	}

	// Find the matching dispersion unit
	var unit DispersionUnit
	found := false

	switch jsonM.Unit {
	case "ppm":
		unit = Dispersion.PartsPerMillion
		found = true
	case "ppb":
		unit = Dispersion.PartsPerBillion
		found = true
	case "ppt":
		unit = Dispersion.PartsPerTrillion
		found = true
	case "%":
		unit = Dispersion.Percent
		found = true
	}

	if !found {
		return Quantity[DispersionUnit]{}, fmt.Errorf("unknown dispersion unit: %s", jsonM.Unit)
	}

	return NewDispersion(jsonM.Value, unit), nil
}

// MarshalElectricCharge serializes an ElectricCharge measurement to JSON
func MarshalElectricCharge(m Quantity[ElectricChargeUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalElectricCharge deserializes a JSON representation to an ElectricCharge measurement
func UnmarshalElectricCharge(data []byte) (Quantity[ElectricChargeUnit], error) {
	var jsonM MeasurementJSON
	if err := json.Unmarshal(data, &jsonM); err != nil {
		return Quantity[ElectricChargeUnit]{}, err
	}

	// Verify dimension
	if jsonM.Dimension != "electric_charge" {
		return Quantity[ElectricChargeUnit]{}, fmt.Errorf("expected dimension 'electric_charge', got '%s'", jsonM.Dimension)
	}

	// Find the matching electric charge unit
	var unit ElectricChargeUnit
	found := false

	switch jsonM.Unit {
	case "C":
		unit = ElectricCharge.Coulomb
		found = true
	case "mC":
		unit = ElectricCharge.Millicoulomb
		found = true
	case "µC":
		unit = ElectricCharge.Microcoulomb
		found = true
	case "Ah":
		unit = ElectricCharge.Ampere_Hour
		found = true
	case "mAh":
		unit = ElectricCharge.Milliampere_Hour
		found = true
	}

	if !found {
		return Quantity[ElectricChargeUnit]{}, fmt.Errorf("unknown electric charge unit: %s", jsonM.Unit)
	}

	return NewElectricCharge(jsonM.Value, unit), nil
}

// MarshalElectricCurrent serializes an ElectricCurrent measurement to JSON
func MarshalElectricCurrent(m Quantity[ElectricCurrentUnit]) ([]byte, error) {
	return marshalGeneric(m)
}

// UnmarshalElectricCurrent deserializes a JSON representation to an ElectricCurrent measurement
func UnmarshalElectricCurrent(data []byte) (Quantity[ElectricCurrentUnit], error) {
	var jsonM MeasurementJSON
	if err := json.Unmarshal(data, &jsonM); err != nil {
		return Quantity[ElectricCurrentUnit]{}, err
	}

	// Verify dimension
	if jsonM.Dimension != "electric_current" {
		return Quantity[ElectricCurrentUnit]{}, fmt.Errorf("expected dimension 'electric_current', got '%s'", jsonM.Dimension)
	}

	// Find the matching electric current unit
	var unit ElectricCurrentUnit
	found := false

	switch jsonM.Unit {
	case "A":
		unit = ElectricCurrent.Ampere
		found = true
	case "mA":
		unit = ElectricCurrent.Milliampere
		found = true
	case "µA":
		unit = ElectricCurrent.Microampere
		found = true
	case "kA":
		unit = ElectricCurrent.Kiloampere
		found = true
	}

	if !found {
		return Quantity[ElectricCurrentUnit]{}, fmt.Errorf("unknown electric current unit: %s", jsonM.Unit)
	}

	return NewElectricCurrent(jsonM.Value, unit), nil
}

// MarshalSpeed serializes a Speed measurement to JSON
func MarshalSpeed(m Quantity[SpeedUnit]) ([]byte, error) {
	return json.Marshal(MeasurementJSON{
		Value:     m.Value,
		Unit:      m.Unit.Symbol(),
		Dimension: m.Unit.Dimension(),
	})
}

// UnmarshalSpeed deserializes a JSON representation to a Speed measurement
func UnmarshalSpeed(data []byte) (Quantity[SpeedUnit], error) {
	var jsonM MeasurementJSON
	if err := json.Unmarshal(data, &jsonM); err != nil {
		return Quantity[SpeedUnit]{}, err
	}

	// Verify dimension
	if jsonM.Dimension != "speed" {
		return Quantity[SpeedUnit]{}, fmt.Errorf("expected dimension 'speed', got '%s'", jsonM.Dimension)
	}

	// Find the matching speed unit
	var unit SpeedUnit
	found := false

	switch jsonM.Unit {
	case "m/s":
		unit = Speed.MetersPerSecond
		found = true
	case "km/h":
		unit = Speed.KilometersPerHour
		found = true
	case "mph":
		unit = Speed.MilesPerHour
		found = true
	case "ft/s":
		unit = Speed.FeetPerSecond
		found = true
	case "kn":
		unit = Speed.Knot
		found = true
	}

	if !found {
		return Quantity[SpeedUnit]{}, fmt.Errorf("unknown speed unit: %s", jsonM.Unit)
	}

	return NewSpeed(jsonM.Value, unit), nil
}

// MarshalElectricPotentialDifference serializes an ElectricPotentialDifference measurement to JSON
func MarshalElectricPotentialDifference(m Quantity[ElectricPotentialDifferenceUnit]) ([]byte, error) {
	return json.Marshal(MeasurementJSON{
		Value:     m.Value,
		Unit:      m.Unit.Symbol(),
		Dimension: m.Unit.Dimension(),
	})
}

// UnmarshalElectricPotentialDifference deserializes a JSON representation to an ElectricPotentialDifference measurement
func UnmarshalElectricPotentialDifference(data []byte) (Quantity[ElectricPotentialDifferenceUnit], error) {
	var jsonM MeasurementJSON
	if err := json.Unmarshal(data, &jsonM); err != nil {
		return Quantity[ElectricPotentialDifferenceUnit]{}, err
	}

	// Verify dimension
	if jsonM.Dimension != "electric_potential_difference" {
		return Quantity[ElectricPotentialDifferenceUnit]{}, fmt.Errorf("expected dimension 'electric_potential_difference', got '%s'", jsonM.Dimension)
	}

	// Find the matching electric potential difference unit
	var unit ElectricPotentialDifferenceUnit
	found := false

	switch jsonM.Unit {
	case "V":
		unit = ElectricPotentialDifference.Volt
		found = true
	case "mV":
		unit = ElectricPotentialDifference.Millivolt
		found = true
	case "µV":
		unit = ElectricPotentialDifference.Microvolt
		found = true
	case "kV":
		unit = ElectricPotentialDifference.Kilovolt
		found = true
	case "MV":
		unit = ElectricPotentialDifference.Megavolt
		found = true
	}

	if !found {
		return Quantity[ElectricPotentialDifferenceUnit]{}, fmt.Errorf("unknown electric potential difference unit: %s", jsonM.Unit)
	}

	return NewElectricPotentialDifference(jsonM.Value, unit), nil
}

// MarshalInformation serializes an Information measurement to JSON
func MarshalInformation(m Quantity[InformationUnit]) ([]byte, error) {
	return json.Marshal(MeasurementJSON{
		Value:     m.Value,
		Unit:      m.Unit.Symbol(),
		Dimension: m.Unit.Dimension(),
	})
}

// UnmarshalInformation deserializes a JSON representation to an Information measurement
func UnmarshalInformation(data []byte) (Quantity[InformationUnit], error) {
	var jsonM MeasurementJSON
	if err := json.Unmarshal(data, &jsonM); err != nil {
		return Quantity[InformationUnit]{}, err
	}

	// Verify dimension
	if jsonM.Dimension != "information" {
		return Quantity[InformationUnit]{}, fmt.Errorf("expected dimension 'information', got '%s'", jsonM.Dimension)
	}

	// Find the matching information unit
	var unit InformationUnit
	found := false

	switch jsonM.Unit {
	case "bit":
		unit = Information.Bit
		found = true
	case "B":
		unit = Information.Byte
		found = true
	case "KB":
		unit = Information.Kilobyte
		found = true
	case "MB":
		unit = Information.Megabyte
		found = true
	case "GB":
		unit = Information.Gigabyte
		found = true
	case "TB":
		unit = Information.Terabyte
		found = true
	case "PB":
		unit = Information.Petabyte
		found = true
	case "KiB":
		unit = Information.Kibibyte
		found = true
	case "MiB":
		unit = Information.Mebibyte
		found = true
	case "GiB":
		unit = Information.Gibibyte
		found = true
	case "TiB":
		unit = Information.Tebibyte
		found = true
	case "PiB":
		unit = Information.Pebibyte
		found = true
	}

	if !found {
		return Quantity[InformationUnit]{}, fmt.Errorf("unknown information unit: %s", jsonM.Unit)
	}

	return NewInformation(jsonM.Value, unit), nil
}

// MarshalFrequency serializes a Frequency measurement to JSON
func MarshalFrequency(m Quantity[FrequencyUnit]) ([]byte, error) {
	return json.Marshal(MeasurementJSON{
		Value:     m.Value,
		Unit:      m.Unit.Symbol(),
		Dimension: m.Unit.Dimension(),
	})
}

// UnmarshalFrequency deserializes a JSON representation to a Frequency measurement
func UnmarshalFrequency(data []byte) (Quantity[FrequencyUnit], error) {
	var jsonM MeasurementJSON
	if err := json.Unmarshal(data, &jsonM); err != nil {
		return Quantity[FrequencyUnit]{}, err
	}

	// Verify dimension
	if jsonM.Dimension != "frequency" {
		return Quantity[FrequencyUnit]{}, fmt.Errorf("expected dimension 'frequency', got '%s'", jsonM.Dimension)
	}

	// Find the matching frequency unit
	var unit FrequencyUnit
	found := false

	switch jsonM.Unit {
	case "Hz":
		unit = Frequency.Hertz
		found = true
	case "kHz":
		unit = Frequency.Kilohertz
		found = true
	case "MHz":
		unit = Frequency.Megahertz
		found = true
	case "GHz":
		unit = Frequency.Gigahertz
		found = true
	case "THz":
		unit = Frequency.Terahertz
		found = true
	case "rpm":
		unit = Frequency.RPM
		found = true
	}

	if !found {
		return Quantity[FrequencyUnit]{}, fmt.Errorf("unknown frequency unit: %s", jsonM.Unit)
	}

	return NewFrequency(jsonM.Value, unit), nil
}

// MarshalIlluminance serializes an Illuminance measurement to JSON
func MarshalIlluminance(m Quantity[IlluminanceUnit]) ([]byte, error) {
	return json.Marshal(MeasurementJSON{
		Value:     m.Value,
		Unit:      m.Unit.Symbol(),
		Dimension: m.Unit.Dimension(),
	})
}

// UnmarshalIlluminance deserializes a JSON representation to an Illuminance measurement
func UnmarshalIlluminance(data []byte) (Quantity[IlluminanceUnit], error) {
	var jsonM MeasurementJSON
	if err := json.Unmarshal(data, &jsonM); err != nil {
		return Quantity[IlluminanceUnit]{}, err
	}

	// Verify dimension
	if jsonM.Dimension != "illuminance" {
		return Quantity[IlluminanceUnit]{}, fmt.Errorf("expected dimension 'illuminance', got '%s'", jsonM.Dimension)
	}

	// Find the matching illuminance unit
	var unit IlluminanceUnit
	found := false

	switch jsonM.Unit {
	case "lx":
		unit = Illuminance.Lux
		found = true
	case "fc":
		unit = Illuminance.FootCandle
		found = true
	case "ph":
		unit = Illuminance.Phot
		found = true
	case "nx":
		unit = Illuminance.Nox
		found = true
	}

	if !found {
		return Quantity[IlluminanceUnit]{}, fmt.Errorf("unknown illuminance unit: %s", jsonM.Unit)
	}

	return NewIlluminance(jsonM.Value, unit), nil
}

// MarshalGeneral serializes a General measurement to JSON
func MarshalGeneral(m Quantity[GeneralUnit]) ([]byte, error) {
	return json.Marshal(MeasurementJSON{
		Value:     m.Value,
		Unit:      m.Unit.Symbol(),
		Dimension: m.Unit.Dimension(),
	})
}

// MarshalFuelEfficiency serializes a FuelEfficiency measurement to JSON
func MarshalFuelEfficiency(m Quantity[FuelEfficiencyUnit]) ([]byte, error) {
	return json.Marshal(MeasurementJSON{
		Value:     m.Value,
		Unit:      m.Unit.Symbol(),
		Dimension: m.Unit.Dimension(),
	})
}

// UnmarshalGeneral deserializes a JSON representation to a General measurement
func UnmarshalGeneral(data []byte) (Quantity[GeneralUnit], error) {
	var jsonM MeasurementJSON
	if err := json.Unmarshal(data, &jsonM); err != nil {
		return Quantity[GeneralUnit]{}, err
	}

	// Verify dimension
	if jsonM.Dimension != "general" {
		return Quantity[GeneralUnit]{}, fmt.Errorf("expected dimension 'general', got '%s'", jsonM.Dimension)
	}

	// Find the matching general unit
	var unit GeneralUnit
	found := false

	switch jsonM.Unit {
	case "unit":
		unit = General.Unit
		found = true
	default:
		// For custom units, create a new general unit with the given symbol
		unit = NewGeneralUnit(jsonM.Unit, jsonM.Unit)
		found = true
	}

	if !found {
		return Quantity[GeneralUnit]{}, fmt.Errorf("unknown general unit: %s", jsonM.Unit)
	}

	return NewGeneral(jsonM.Value, unit), nil
}

// UnmarshalFuelEfficiency deserializes a JSON representation to a FuelEfficiency measurement
func UnmarshalFuelEfficiency(data []byte) (Quantity[FuelEfficiencyUnit], error) {
	var jsonM MeasurementJSON
	if err := json.Unmarshal(data, &jsonM); err != nil {
		return Quantity[FuelEfficiencyUnit]{}, err
	}

	// Verify dimension
	if jsonM.Dimension != "fuel_efficiency" {
		return Quantity[FuelEfficiencyUnit]{}, fmt.Errorf("expected dimension 'fuel_efficiency', got '%s'", jsonM.Dimension)
	}

	// Find the matching fuel efficiency unit
	var unit FuelEfficiencyUnit
	found := false

	switch jsonM.Unit {
	case "km/L":
		unit = FuelEfficiency.KilometersPerLiter
		found = true
	case "mpg":
		unit = FuelEfficiency.MilesPerGallon
		found = true
	case "L/100km":
		unit = FuelEfficiency.LitersPer100Kilometers
		found = true
	}

	if !found {
		return Quantity[FuelEfficiencyUnit]{}, fmt.Errorf("unknown fuel efficiency unit: %s", jsonM.Unit)
	}

	return NewFuelEfficiency(jsonM.Value, unit), nil
}
