// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Unit key registries for compact format deserialization
// Maps "dimension_unitname" -> unit constant

var temperatureUnitsByKey = map[string]TemperatureUnit{
	"temperature_celsius":    Temperature.Celsius,
	"temperature_fahrenheit": Temperature.Fahrenheit,
	"temperature_kelvin":     Temperature.Kelvin,
}

var pressureUnitsByKey = map[string]PressureUnit{
	"pressure_pascal":                    Pressure.Pascal,
	"pressure_kilopascal":                Pressure.Kilopascal,
	"pressure_bar":                       Pressure.Bar,
	"pressure_pounds_per_square_inch":    Pressure.PSI,
	"pressure_inches_of_water_column":    Pressure.InchH2O,
}

var lengthUnitsByKey = map[string]LengthUnit{
	"length_meter":      Length.Meter,
	"length_kilometer":  Length.Kilometer,
	"length_centimeter": Length.Centimeter,
	"length_millimeter": Length.Millimeter,
	"length_micrometer": Length.Micrometer,
	"length_nanometer":  Length.Nanometer,
	"length_inch":       Length.Inch,
	"length_foot":       Length.Foot,
	"length_yard":       Length.Yard,
	"length_mile":       Length.Mile,
}

var massUnitsByKey = map[string]MassUnit{
	"mass_kilogram":   Mass.Kilogram,
	"mass_gram":       Mass.Gram,
	"mass_milligram":  Mass.Milligram,
	"mass_microgram":  Mass.Microgram,
	"mass_pound":      Mass.Pound,
	"mass_ounce":      Mass.Ounce,
	"mass_stone":      Mass.Stone,
	"mass_metric_ton": Mass.MetricTon,
	"mass_ton":        Mass.Ton,
}

var durationUnitsByKey = map[string]DurationUnit{
	"duration_second":      Duration.Second,
	"duration_minute":      Duration.Minute,
	"duration_hour":        Duration.Hour,
	"duration_day":         Duration.Day,
	"duration_millisecond": Duration.Millisecond,
	"duration_microsecond": Duration.Microsecond,
	"duration_nanosecond":  Duration.Nanosecond,
}

var angleUnitsByKey = map[string]AngleUnit{
	"angle_radian":     Angle.Radian,
	"angle_degree":     Angle.Degree,
	"angle_arcminute":  Angle.Arcminute,
	"angle_arcsecond":  Angle.Arcsecond,
	"angle_revolution": Angle.Revolution,
	"angle_gradian":    Angle.Gradian,
}

var areaUnitsByKey = map[string]AreaUnit{
	"area_square_meter":      Area.SquareMeter,
	"area_square_kilometer":  Area.SquareKilometer,
	"area_square_centimeter": Area.SquareCentimeter,
	"area_square_millimeter": Area.SquareMillimeter,
	"area_square_inch":       Area.SquareInch,
	"area_square_foot":       Area.SquareFoot,
	"area_square_yard":       Area.SquareYard,
	"area_square_mile":       Area.SquareMile,
	"area_acre":              Area.Acre,
	"area_hectare":           Area.Hectare,
}

var volumeUnitsByKey = map[string]VolumeUnit{
	"volume_cubic_meter":      Volume.CubicMeter,
	"volume_cubic_kilometer":  Volume.CubicKilometer,
	"volume_cubic_centimeter": Volume.CubicCentimeter,
	"volume_cubic_millimeter": Volume.CubicMillimeter,
	"volume_liter":            Volume.Liter,
	"volume_milliliter":       Volume.Milliliter,
	"volume_cubic_inch":       Volume.CubicInch,
	"volume_cubic_foot":       Volume.CubicFoot,
	"volume_cubic_yard":       Volume.CubicYard,
	"volume_gallon":           Volume.Gallon,
	"volume_quart":            Volume.Quart,
	"volume_pint":             Volume.Pint,
	"volume_cup":              Volume.Cup,
	"volume_fluid_ounce":      Volume.FluidOunce,
}

var speedUnitsByKey = map[string]SpeedUnit{
	"speed_meters_per_second":    Speed.MetersPerSecond,
	"speed_kilometers_per_hour":  Speed.KilometersPerHour,
	"speed_miles_per_hour":       Speed.MilesPerHour,
	"speed_feet_per_second":      Speed.FeetPerSecond,
	"speed_knot":                 Speed.Knot,
}

var accelerationUnitsByKey = map[string]AccelerationUnit{
	"acceleration_meters_per_second_squared": Acceleration.MetersPerSecondSquared,
	"acceleration_g":                         Acceleration.G,
	"acceleration_feet_per_second_squared":   Acceleration.FeetPerSecondSquared,
}

var flowRateUnitsByKey = map[string]FlowRateUnit{
	"flowrate_cubic_meters_per_hour": FlowRate.CubicMetersPerHour,
	"flowrate_liters_per_second":     FlowRate.LitersPerSecond,
	"flowrate_c_f_m":                 FlowRate.CFM,
}

var powerUnitsByKey = map[string]PowerUnit{
	"power_watt":         Power.Watt,
	"power_kilowatt":     Power.Kilowatt,
	"power_b_t_u_per_hour": Power.BTUPerHour,
}

var energyUnitsByKey = map[string]EnergyUnit{
	"energy_joule":         Energy.Joule,
	"energy_kilowatt_hour": Energy.KilowattHour,
	"energy_b_t_u":         Energy.BTU,
}

var concentrationUnitsByKey = map[string]ConcentrationUnit{
	"concentration_grams_per_liter":      Concentration.GramsPerLiter,
	"concentration_milligrams_per_liter": Concentration.MilligramsPerLiter,
	"concentration_parts_per_million":    Concentration.PartsPerMillion,
	"concentration_parts_per_billion":    Concentration.PartsPerBillion,
}

var dispersionUnitsByKey = map[string]DispersionUnit{
	"dispersion_parts_per_million":  Dispersion.PartsPerMillion,
	"dispersion_parts_per_billion":  Dispersion.PartsPerBillion,
	"dispersion_parts_per_trillion": Dispersion.PartsPerTrillion,
	"dispersion_percent":            Dispersion.Percent,
}

var electricChargeUnitsByKey = map[string]ElectricChargeUnit{
	"electric_charge_coulomb":          ElectricCharge.Coulomb,
	"electric_charge_millicoulomb":     ElectricCharge.Millicoulomb,
	"electric_charge_microcoulomb":     ElectricCharge.Microcoulomb,
	"electric_charge_ampere__hour":     ElectricCharge.Ampere_Hour,
	"electric_charge_milliampere__hour": ElectricCharge.Milliampere_Hour,
}

var electricCurrentUnitsByKey = map[string]ElectricCurrentUnit{
	"electric_current_ampere":      ElectricCurrent.Ampere,
	"electric_current_milliampere": ElectricCurrent.Milliampere,
	"electric_current_microampere": ElectricCurrent.Microampere,
	"electric_current_kiloampere":  ElectricCurrent.Kiloampere,
}

var electricPotentialDifferenceUnitsByKey = map[string]ElectricPotentialDifferenceUnit{
	"electric_potential_difference_volt":      ElectricPotentialDifference.Volt,
	"electric_potential_difference_millivolt": ElectricPotentialDifference.Millivolt,
	"electric_potential_difference_microvolt": ElectricPotentialDifference.Microvolt,
	"electric_potential_difference_kilovolt":  ElectricPotentialDifference.Kilovolt,
	"electric_potential_difference_megavolt":  ElectricPotentialDifference.Megavolt,
}

var frequencyUnitsByKey = map[string]FrequencyUnit{
	"frequency_hertz":     Frequency.Hertz,
	"frequency_kilohertz": Frequency.Kilohertz,
	"frequency_megahertz": Frequency.Megahertz,
	"frequency_gigahertz": Frequency.Gigahertz,
	"frequency_terahertz": Frequency.Terahertz,
	"frequency_r_p_m":     Frequency.RPM,
}

var illuminanceUnitsByKey = map[string]IlluminanceUnit{
	"illuminance_lux":         Illuminance.Lux,
	"illuminance_foot_candle": Illuminance.FootCandle,
	"illuminance_phot":        Illuminance.Phot,
	"illuminance_nox":         Illuminance.Nox,
}

var informationUnitsByKey = map[string]InformationUnit{
	"information_bit":      Information.Bit,
	"information_byte":     Information.Byte,
	"information_kilobyte": Information.Kilobyte,
	"information_megabyte": Information.Megabyte,
	"information_gigabyte": Information.Gigabyte,
	"information_terabyte": Information.Terabyte,
	"information_petabyte": Information.Petabyte,
	"information_kibibyte": Information.Kibibyte,
	"information_mebibyte": Information.Mebibyte,
	"information_gibibyte": Information.Gibibyte,
	"information_tebibyte": Information.Tebibyte,
	"information_pebibyte": Information.Pebibyte,
}

var fuelEfficiencyUnitsByKey = map[string]FuelEfficiencyUnit{
	"fuel_efficiency_kilometers_per_liter":      FuelEfficiency.KilometersPerLiter,
	"fuel_efficiency_miles_per_gallon":          FuelEfficiency.MilesPerGallon,
	"fuel_efficiency_liters_per100_kilometers": FuelEfficiency.LitersPer100Kilometers,
}

// marshalCompactGeneric is a helper function to serialize any measurement to compact JSON
func marshalCompactGeneric[T Category](m Quantity[T], includeSymbol bool) ([]byte, error) {
	key := unitKey(m.Unit.Dimension(), m.Unit.Name())
	cj := CompactJSON{
		Value: m.Value,
		Unit:  key,
	}
	if includeSymbol {
		cj.Symbol = m.Unit.Symbol()
	}
	return json.Marshal(cj)
}

// MarshalCompactTemperature serializes a Temperature measurement to compact JSON
func MarshalCompactTemperature(m Quantity[TemperatureUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, false)
}

// MarshalCompactTemperatureWithSymbol serializes a Temperature measurement to compact JSON with symbol
func MarshalCompactTemperatureWithSymbol(m Quantity[TemperatureUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, true)
}

// UnmarshalCompactTemperature deserializes compact JSON to a Temperature measurement
func UnmarshalCompactTemperature(data []byte) (Quantity[TemperatureUnit], error) {
	var cj CompactJSON
	if err := json.Unmarshal(data, &cj); err != nil {
		return Quantity[TemperatureUnit]{}, err
	}
	unit, ok := temperatureUnitsByKey[cj.Unit]
	if !ok {
		return Quantity[TemperatureUnit]{}, fmt.Errorf("unknown temperature unit key: %s", cj.Unit)
	}
	return NewTemperature(cj.Value, unit), nil
}

// MarshalCompactPressure serializes a Pressure measurement to compact JSON
func MarshalCompactPressure(m Quantity[PressureUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, false)
}

// MarshalCompactPressureWithSymbol serializes a Pressure measurement to compact JSON with symbol
func MarshalCompactPressureWithSymbol(m Quantity[PressureUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, true)
}

// UnmarshalCompactPressure deserializes compact JSON to a Pressure measurement
func UnmarshalCompactPressure(data []byte) (Quantity[PressureUnit], error) {
	var cj CompactJSON
	if err := json.Unmarshal(data, &cj); err != nil {
		return Quantity[PressureUnit]{}, err
	}
	unit, ok := pressureUnitsByKey[cj.Unit]
	if !ok {
		return Quantity[PressureUnit]{}, fmt.Errorf("unknown pressure unit key: %s", cj.Unit)
	}
	return NewPressure(cj.Value, unit), nil
}

// MarshalCompactLength serializes a Length measurement to compact JSON
func MarshalCompactLength(m Quantity[LengthUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, false)
}

// MarshalCompactLengthWithSymbol serializes a Length measurement to compact JSON with symbol
func MarshalCompactLengthWithSymbol(m Quantity[LengthUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, true)
}

// UnmarshalCompactLength deserializes compact JSON to a Length measurement
func UnmarshalCompactLength(data []byte) (Quantity[LengthUnit], error) {
	var cj CompactJSON
	if err := json.Unmarshal(data, &cj); err != nil {
		return Quantity[LengthUnit]{}, err
	}
	unit, ok := lengthUnitsByKey[cj.Unit]
	if !ok {
		return Quantity[LengthUnit]{}, fmt.Errorf("unknown length unit key: %s", cj.Unit)
	}
	return NewLength(cj.Value, unit), nil
}

// MarshalCompactMass serializes a Mass measurement to compact JSON
func MarshalCompactMass(m Quantity[MassUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, false)
}

// MarshalCompactMassWithSymbol serializes a Mass measurement to compact JSON with symbol
func MarshalCompactMassWithSymbol(m Quantity[MassUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, true)
}

// UnmarshalCompactMass deserializes compact JSON to a Mass measurement
func UnmarshalCompactMass(data []byte) (Quantity[MassUnit], error) {
	var cj CompactJSON
	if err := json.Unmarshal(data, &cj); err != nil {
		return Quantity[MassUnit]{}, err
	}
	unit, ok := massUnitsByKey[cj.Unit]
	if !ok {
		return Quantity[MassUnit]{}, fmt.Errorf("unknown mass unit key: %s", cj.Unit)
	}
	return NewMass(cj.Value, unit), nil
}

// MarshalCompactDuration serializes a Duration measurement to compact JSON
func MarshalCompactDuration(m Quantity[DurationUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, false)
}

// MarshalCompactDurationWithSymbol serializes a Duration measurement to compact JSON with symbol
func MarshalCompactDurationWithSymbol(m Quantity[DurationUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, true)
}

// UnmarshalCompactDuration deserializes compact JSON to a Duration measurement
func UnmarshalCompactDuration(data []byte) (Quantity[DurationUnit], error) {
	var cj CompactJSON
	if err := json.Unmarshal(data, &cj); err != nil {
		return Quantity[DurationUnit]{}, err
	}
	unit, ok := durationUnitsByKey[cj.Unit]
	if !ok {
		return Quantity[DurationUnit]{}, fmt.Errorf("unknown duration unit key: %s", cj.Unit)
	}
	return NewDuration(cj.Value, unit), nil
}

// MarshalCompactAngle serializes an Angle measurement to compact JSON
func MarshalCompactAngle(m Quantity[AngleUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, false)
}

// MarshalCompactAngleWithSymbol serializes an Angle measurement to compact JSON with symbol
func MarshalCompactAngleWithSymbol(m Quantity[AngleUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, true)
}

// UnmarshalCompactAngle deserializes compact JSON to an Angle measurement
func UnmarshalCompactAngle(data []byte) (Quantity[AngleUnit], error) {
	var cj CompactJSON
	if err := json.Unmarshal(data, &cj); err != nil {
		return Quantity[AngleUnit]{}, err
	}
	unit, ok := angleUnitsByKey[cj.Unit]
	if !ok {
		return Quantity[AngleUnit]{}, fmt.Errorf("unknown angle unit key: %s", cj.Unit)
	}
	return NewAngle(cj.Value, unit), nil
}

// MarshalCompactArea serializes an Area measurement to compact JSON
func MarshalCompactArea(m Quantity[AreaUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, false)
}

// MarshalCompactAreaWithSymbol serializes an Area measurement to compact JSON with symbol
func MarshalCompactAreaWithSymbol(m Quantity[AreaUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, true)
}

// UnmarshalCompactArea deserializes compact JSON to an Area measurement
func UnmarshalCompactArea(data []byte) (Quantity[AreaUnit], error) {
	var cj CompactJSON
	if err := json.Unmarshal(data, &cj); err != nil {
		return Quantity[AreaUnit]{}, err
	}
	unit, ok := areaUnitsByKey[cj.Unit]
	if !ok {
		return Quantity[AreaUnit]{}, fmt.Errorf("unknown area unit key: %s", cj.Unit)
	}
	return NewArea(cj.Value, unit), nil
}

// MarshalCompactVolume serializes a Volume measurement to compact JSON
func MarshalCompactVolume(m Quantity[VolumeUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, false)
}

// MarshalCompactVolumeWithSymbol serializes a Volume measurement to compact JSON with symbol
func MarshalCompactVolumeWithSymbol(m Quantity[VolumeUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, true)
}

// UnmarshalCompactVolume deserializes compact JSON to a Volume measurement
func UnmarshalCompactVolume(data []byte) (Quantity[VolumeUnit], error) {
	var cj CompactJSON
	if err := json.Unmarshal(data, &cj); err != nil {
		return Quantity[VolumeUnit]{}, err
	}
	unit, ok := volumeUnitsByKey[cj.Unit]
	if !ok {
		return Quantity[VolumeUnit]{}, fmt.Errorf("unknown volume unit key: %s", cj.Unit)
	}
	return NewVolume(cj.Value, unit), nil
}

// MarshalCompactSpeed serializes a Speed measurement to compact JSON
func MarshalCompactSpeed(m Quantity[SpeedUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, false)
}

// MarshalCompactSpeedWithSymbol serializes a Speed measurement to compact JSON with symbol
func MarshalCompactSpeedWithSymbol(m Quantity[SpeedUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, true)
}

// UnmarshalCompactSpeed deserializes compact JSON to a Speed measurement
func UnmarshalCompactSpeed(data []byte) (Quantity[SpeedUnit], error) {
	var cj CompactJSON
	if err := json.Unmarshal(data, &cj); err != nil {
		return Quantity[SpeedUnit]{}, err
	}
	unit, ok := speedUnitsByKey[cj.Unit]
	if !ok {
		return Quantity[SpeedUnit]{}, fmt.Errorf("unknown speed unit key: %s", cj.Unit)
	}
	return NewSpeed(cj.Value, unit), nil
}

// MarshalCompactAcceleration serializes an Acceleration measurement to compact JSON
func MarshalCompactAcceleration(m Quantity[AccelerationUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, false)
}

// MarshalCompactAccelerationWithSymbol serializes an Acceleration measurement to compact JSON with symbol
func MarshalCompactAccelerationWithSymbol(m Quantity[AccelerationUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, true)
}

// UnmarshalCompactAcceleration deserializes compact JSON to an Acceleration measurement
func UnmarshalCompactAcceleration(data []byte) (Quantity[AccelerationUnit], error) {
	var cj CompactJSON
	if err := json.Unmarshal(data, &cj); err != nil {
		return Quantity[AccelerationUnit]{}, err
	}
	unit, ok := accelerationUnitsByKey[cj.Unit]
	if !ok {
		return Quantity[AccelerationUnit]{}, fmt.Errorf("unknown acceleration unit key: %s", cj.Unit)
	}
	return NewAcceleration(cj.Value, unit), nil
}

// MarshalCompactFlowRate serializes a FlowRate measurement to compact JSON
func MarshalCompactFlowRate(m Quantity[FlowRateUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, false)
}

// MarshalCompactFlowRateWithSymbol serializes a FlowRate measurement to compact JSON with symbol
func MarshalCompactFlowRateWithSymbol(m Quantity[FlowRateUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, true)
}

// UnmarshalCompactFlowRate deserializes compact JSON to a FlowRate measurement
func UnmarshalCompactFlowRate(data []byte) (Quantity[FlowRateUnit], error) {
	var cj CompactJSON
	if err := json.Unmarshal(data, &cj); err != nil {
		return Quantity[FlowRateUnit]{}, err
	}
	unit, ok := flowRateUnitsByKey[cj.Unit]
	if !ok {
		return Quantity[FlowRateUnit]{}, fmt.Errorf("unknown flowrate unit key: %s", cj.Unit)
	}
	return NewFlowRate(cj.Value, unit), nil
}

// MarshalCompactPower serializes a Power measurement to compact JSON
func MarshalCompactPower(m Quantity[PowerUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, false)
}

// MarshalCompactPowerWithSymbol serializes a Power measurement to compact JSON with symbol
func MarshalCompactPowerWithSymbol(m Quantity[PowerUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, true)
}

// UnmarshalCompactPower deserializes compact JSON to a Power measurement
func UnmarshalCompactPower(data []byte) (Quantity[PowerUnit], error) {
	var cj CompactJSON
	if err := json.Unmarshal(data, &cj); err != nil {
		return Quantity[PowerUnit]{}, err
	}
	unit, ok := powerUnitsByKey[cj.Unit]
	if !ok {
		return Quantity[PowerUnit]{}, fmt.Errorf("unknown power unit key: %s", cj.Unit)
	}
	return NewPower(cj.Value, unit), nil
}

// MarshalCompactEnergy serializes an Energy measurement to compact JSON
func MarshalCompactEnergy(m Quantity[EnergyUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, false)
}

// MarshalCompactEnergyWithSymbol serializes an Energy measurement to compact JSON with symbol
func MarshalCompactEnergyWithSymbol(m Quantity[EnergyUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, true)
}

// UnmarshalCompactEnergy deserializes compact JSON to an Energy measurement
func UnmarshalCompactEnergy(data []byte) (Quantity[EnergyUnit], error) {
	var cj CompactJSON
	if err := json.Unmarshal(data, &cj); err != nil {
		return Quantity[EnergyUnit]{}, err
	}
	unit, ok := energyUnitsByKey[cj.Unit]
	if !ok {
		return Quantity[EnergyUnit]{}, fmt.Errorf("unknown energy unit key: %s", cj.Unit)
	}
	return NewEnergy(cj.Value, unit), nil
}

// MarshalCompactConcentration serializes a Concentration measurement to compact JSON
func MarshalCompactConcentration(m Quantity[ConcentrationUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, false)
}

// MarshalCompactConcentrationWithSymbol serializes a Concentration measurement to compact JSON with symbol
func MarshalCompactConcentrationWithSymbol(m Quantity[ConcentrationUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, true)
}

// UnmarshalCompactConcentration deserializes compact JSON to a Concentration measurement
func UnmarshalCompactConcentration(data []byte) (Quantity[ConcentrationUnit], error) {
	var cj CompactJSON
	if err := json.Unmarshal(data, &cj); err != nil {
		return Quantity[ConcentrationUnit]{}, err
	}
	unit, ok := concentrationUnitsByKey[cj.Unit]
	if !ok {
		return Quantity[ConcentrationUnit]{}, fmt.Errorf("unknown concentration unit key: %s", cj.Unit)
	}
	return NewConcentration(cj.Value, unit), nil
}

// MarshalCompactDispersion serializes a Dispersion measurement to compact JSON
func MarshalCompactDispersion(m Quantity[DispersionUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, false)
}

// MarshalCompactDispersionWithSymbol serializes a Dispersion measurement to compact JSON with symbol
func MarshalCompactDispersionWithSymbol(m Quantity[DispersionUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, true)
}

// UnmarshalCompactDispersion deserializes compact JSON to a Dispersion measurement
func UnmarshalCompactDispersion(data []byte) (Quantity[DispersionUnit], error) {
	var cj CompactJSON
	if err := json.Unmarshal(data, &cj); err != nil {
		return Quantity[DispersionUnit]{}, err
	}
	unit, ok := dispersionUnitsByKey[cj.Unit]
	if !ok {
		return Quantity[DispersionUnit]{}, fmt.Errorf("unknown dispersion unit key: %s", cj.Unit)
	}
	return NewDispersion(cj.Value, unit), nil
}

// MarshalCompactElectricCharge serializes an ElectricCharge measurement to compact JSON
func MarshalCompactElectricCharge(m Quantity[ElectricChargeUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, false)
}

// MarshalCompactElectricChargeWithSymbol serializes an ElectricCharge measurement to compact JSON with symbol
func MarshalCompactElectricChargeWithSymbol(m Quantity[ElectricChargeUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, true)
}

// UnmarshalCompactElectricCharge deserializes compact JSON to an ElectricCharge measurement
func UnmarshalCompactElectricCharge(data []byte) (Quantity[ElectricChargeUnit], error) {
	var cj CompactJSON
	if err := json.Unmarshal(data, &cj); err != nil {
		return Quantity[ElectricChargeUnit]{}, err
	}
	unit, ok := electricChargeUnitsByKey[cj.Unit]
	if !ok {
		return Quantity[ElectricChargeUnit]{}, fmt.Errorf("unknown electric_charge unit key: %s", cj.Unit)
	}
	return NewElectricCharge(cj.Value, unit), nil
}

// MarshalCompactElectricCurrent serializes an ElectricCurrent measurement to compact JSON
func MarshalCompactElectricCurrent(m Quantity[ElectricCurrentUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, false)
}

// MarshalCompactElectricCurrentWithSymbol serializes an ElectricCurrent measurement to compact JSON with symbol
func MarshalCompactElectricCurrentWithSymbol(m Quantity[ElectricCurrentUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, true)
}

// UnmarshalCompactElectricCurrent deserializes compact JSON to an ElectricCurrent measurement
func UnmarshalCompactElectricCurrent(data []byte) (Quantity[ElectricCurrentUnit], error) {
	var cj CompactJSON
	if err := json.Unmarshal(data, &cj); err != nil {
		return Quantity[ElectricCurrentUnit]{}, err
	}
	unit, ok := electricCurrentUnitsByKey[cj.Unit]
	if !ok {
		return Quantity[ElectricCurrentUnit]{}, fmt.Errorf("unknown electric_current unit key: %s", cj.Unit)
	}
	return NewElectricCurrent(cj.Value, unit), nil
}

// MarshalCompactElectricPotentialDifference serializes an ElectricPotentialDifference measurement to compact JSON
func MarshalCompactElectricPotentialDifference(m Quantity[ElectricPotentialDifferenceUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, false)
}

// MarshalCompactElectricPotentialDifferenceWithSymbol serializes an ElectricPotentialDifference measurement to compact JSON with symbol
func MarshalCompactElectricPotentialDifferenceWithSymbol(m Quantity[ElectricPotentialDifferenceUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, true)
}

// UnmarshalCompactElectricPotentialDifference deserializes compact JSON to an ElectricPotentialDifference measurement
func UnmarshalCompactElectricPotentialDifference(data []byte) (Quantity[ElectricPotentialDifferenceUnit], error) {
	var cj CompactJSON
	if err := json.Unmarshal(data, &cj); err != nil {
		return Quantity[ElectricPotentialDifferenceUnit]{}, err
	}
	unit, ok := electricPotentialDifferenceUnitsByKey[cj.Unit]
	if !ok {
		return Quantity[ElectricPotentialDifferenceUnit]{}, fmt.Errorf("unknown electric_potential_difference unit key: %s", cj.Unit)
	}
	return NewElectricPotentialDifference(cj.Value, unit), nil
}

// MarshalCompactFrequency serializes a Frequency measurement to compact JSON
func MarshalCompactFrequency(m Quantity[FrequencyUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, false)
}

// MarshalCompactFrequencyWithSymbol serializes a Frequency measurement to compact JSON with symbol
func MarshalCompactFrequencyWithSymbol(m Quantity[FrequencyUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, true)
}

// UnmarshalCompactFrequency deserializes compact JSON to a Frequency measurement
func UnmarshalCompactFrequency(data []byte) (Quantity[FrequencyUnit], error) {
	var cj CompactJSON
	if err := json.Unmarshal(data, &cj); err != nil {
		return Quantity[FrequencyUnit]{}, err
	}
	unit, ok := frequencyUnitsByKey[cj.Unit]
	if !ok {
		return Quantity[FrequencyUnit]{}, fmt.Errorf("unknown frequency unit key: %s", cj.Unit)
	}
	return NewFrequency(cj.Value, unit), nil
}

// MarshalCompactIlluminance serializes an Illuminance measurement to compact JSON
func MarshalCompactIlluminance(m Quantity[IlluminanceUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, false)
}

// MarshalCompactIlluminanceWithSymbol serializes an Illuminance measurement to compact JSON with symbol
func MarshalCompactIlluminanceWithSymbol(m Quantity[IlluminanceUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, true)
}

// UnmarshalCompactIlluminance deserializes compact JSON to an Illuminance measurement
func UnmarshalCompactIlluminance(data []byte) (Quantity[IlluminanceUnit], error) {
	var cj CompactJSON
	if err := json.Unmarshal(data, &cj); err != nil {
		return Quantity[IlluminanceUnit]{}, err
	}
	unit, ok := illuminanceUnitsByKey[cj.Unit]
	if !ok {
		return Quantity[IlluminanceUnit]{}, fmt.Errorf("unknown illuminance unit key: %s", cj.Unit)
	}
	return NewIlluminance(cj.Value, unit), nil
}

// MarshalCompactInformation serializes an Information measurement to compact JSON
func MarshalCompactInformation(m Quantity[InformationUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, false)
}

// MarshalCompactInformationWithSymbol serializes an Information measurement to compact JSON with symbol
func MarshalCompactInformationWithSymbol(m Quantity[InformationUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, true)
}

// UnmarshalCompactInformation deserializes compact JSON to an Information measurement
func UnmarshalCompactInformation(data []byte) (Quantity[InformationUnit], error) {
	var cj CompactJSON
	if err := json.Unmarshal(data, &cj); err != nil {
		return Quantity[InformationUnit]{}, err
	}
	unit, ok := informationUnitsByKey[cj.Unit]
	if !ok {
		return Quantity[InformationUnit]{}, fmt.Errorf("unknown information unit key: %s", cj.Unit)
	}
	return NewInformation(cj.Value, unit), nil
}

// MarshalCompactFuelEfficiency serializes a FuelEfficiency measurement to compact JSON
func MarshalCompactFuelEfficiency(m Quantity[FuelEfficiencyUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, false)
}

// MarshalCompactFuelEfficiencyWithSymbol serializes a FuelEfficiency measurement to compact JSON with symbol
func MarshalCompactFuelEfficiencyWithSymbol(m Quantity[FuelEfficiencyUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, true)
}

// UnmarshalCompactFuelEfficiency deserializes compact JSON to a FuelEfficiency measurement
func UnmarshalCompactFuelEfficiency(data []byte) (Quantity[FuelEfficiencyUnit], error) {
	var cj CompactJSON
	if err := json.Unmarshal(data, &cj); err != nil {
		return Quantity[FuelEfficiencyUnit]{}, err
	}
	unit, ok := fuelEfficiencyUnitsByKey[cj.Unit]
	if !ok {
		return Quantity[FuelEfficiencyUnit]{}, fmt.Errorf("unknown fuel_efficiency unit key: %s", cj.Unit)
	}
	return NewFuelEfficiency(cj.Value, unit), nil
}

// MarshalCompactGeneral serializes a General measurement to compact JSON
func MarshalCompactGeneral(m Quantity[GeneralUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, false)
}

// MarshalCompactGeneralWithSymbol serializes a General measurement to compact JSON with symbol
func MarshalCompactGeneralWithSymbol(m Quantity[GeneralUnit]) ([]byte, error) {
	return marshalCompactGeneric(m, true)
}

// UnmarshalCompactGeneral deserializes compact JSON to a General measurement
func UnmarshalCompactGeneral(data []byte) (Quantity[GeneralUnit], error) {
	var cj CompactJSON
	if err := json.Unmarshal(data, &cj); err != nil {
		return Quantity[GeneralUnit]{}, err
	}
	// Check for predefined general units
	switch cj.Unit {
	case "general_general_unit":
		return NewGeneral(cj.Value, General.Unit), nil
	case "general_percent":
		return NewGeneral(cj.Value, General.Percent), nil
	default:
		// For custom units, create from the key
		_, unitName := parseUnitKey(cj.Unit)
		unit := NewGeneralUnit(unitName, unitName)
		return NewGeneral(cj.Value, unit), nil
	}
}

// unmarshalCompactMeasurement deserializes compact JSON to an AnyMeasurement
func unmarshalCompactMeasurement(data []byte) (*AnyMeasurement, error) {
	var cj CompactJSON
	if err := json.Unmarshal(data, &cj); err != nil {
		return nil, err
	}

	// Extract dimension from the unit key
	dimension, _ := parseUnitKey(cj.Unit)

	switch dimension {
	case "temperature":
		m, err := UnmarshalCompactTemperature(data)
		if err != nil {
			return nil, err
		}
		return &AnyMeasurement{value: m, dimension: "temperature"}, nil
	case "pressure":
		m, err := UnmarshalCompactPressure(data)
		if err != nil {
			return nil, err
		}
		return &AnyMeasurement{value: m, dimension: "pressure"}, nil
	case "length":
		m, err := UnmarshalCompactLength(data)
		if err != nil {
			return nil, err
		}
		return &AnyMeasurement{value: m, dimension: "length"}, nil
	case "mass":
		m, err := UnmarshalCompactMass(data)
		if err != nil {
			return nil, err
		}
		return &AnyMeasurement{value: m, dimension: "mass"}, nil
	case "duration":
		m, err := UnmarshalCompactDuration(data)
		if err != nil {
			return nil, err
		}
		return &AnyMeasurement{value: m, dimension: "duration"}, nil
	case "angle":
		m, err := UnmarshalCompactAngle(data)
		if err != nil {
			return nil, err
		}
		return &AnyMeasurement{value: m, dimension: "angle"}, nil
	case "area":
		m, err := UnmarshalCompactArea(data)
		if err != nil {
			return nil, err
		}
		return &AnyMeasurement{value: m, dimension: "area"}, nil
	case "volume":
		m, err := UnmarshalCompactVolume(data)
		if err != nil {
			return nil, err
		}
		return &AnyMeasurement{value: m, dimension: "volume"}, nil
	case "speed":
		m, err := UnmarshalCompactSpeed(data)
		if err != nil {
			return nil, err
		}
		return &AnyMeasurement{value: m, dimension: "speed"}, nil
	case "acceleration":
		m, err := UnmarshalCompactAcceleration(data)
		if err != nil {
			return nil, err
		}
		return &AnyMeasurement{value: m, dimension: "acceleration"}, nil
	case "flowrate":
		m, err := UnmarshalCompactFlowRate(data)
		if err != nil {
			return nil, err
		}
		return &AnyMeasurement{value: m, dimension: "flowrate"}, nil
	case "power":
		m, err := UnmarshalCompactPower(data)
		if err != nil {
			return nil, err
		}
		return &AnyMeasurement{value: m, dimension: "power"}, nil
	case "energy":
		m, err := UnmarshalCompactEnergy(data)
		if err != nil {
			return nil, err
		}
		return &AnyMeasurement{value: m, dimension: "energy"}, nil
	case "concentration":
		m, err := UnmarshalCompactConcentration(data)
		if err != nil {
			return nil, err
		}
		return &AnyMeasurement{value: m, dimension: "concentration"}, nil
	case "dispersion":
		m, err := UnmarshalCompactDispersion(data)
		if err != nil {
			return nil, err
		}
		return &AnyMeasurement{value: m, dimension: "dispersion"}, nil
	case "electric_charge":
		m, err := UnmarshalCompactElectricCharge(data)
		if err != nil {
			return nil, err
		}
		return &AnyMeasurement{value: m, dimension: "electric_charge"}, nil
	case "electric_current":
		m, err := UnmarshalCompactElectricCurrent(data)
		if err != nil {
			return nil, err
		}
		return &AnyMeasurement{value: m, dimension: "electric_current"}, nil
	case "electric_potential_difference":
		m, err := UnmarshalCompactElectricPotentialDifference(data)
		if err != nil {
			return nil, err
		}
		return &AnyMeasurement{value: m, dimension: "electric_potential_difference"}, nil
	case "frequency":
		m, err := UnmarshalCompactFrequency(data)
		if err != nil {
			return nil, err
		}
		return &AnyMeasurement{value: m, dimension: "frequency"}, nil
	case "illuminance":
		m, err := UnmarshalCompactIlluminance(data)
		if err != nil {
			return nil, err
		}
		return &AnyMeasurement{value: m, dimension: "illuminance"}, nil
	case "information":
		m, err := UnmarshalCompactInformation(data)
		if err != nil {
			return nil, err
		}
		return &AnyMeasurement{value: m, dimension: "information"}, nil
	case "fuel_efficiency":
		m, err := UnmarshalCompactFuelEfficiency(data)
		if err != nil {
			return nil, err
		}
		return &AnyMeasurement{value: m, dimension: "fuel_efficiency"}, nil
	case "general":
		m, err := UnmarshalCompactGeneral(data)
		if err != nil {
			return nil, err
		}
		return &AnyMeasurement{value: m, dimension: "general"}, nil
	default:
		// Fallback to general for unknown dimensions
		m, err := UnmarshalCompactGeneral(data)
		if err != nil {
			return nil, err
		}
		return &AnyMeasurement{value: m, dimension: "general"}, nil
	}
}

// isCompactFormat checks if the JSON data is in compact format
func isCompactFormat(data []byte) bool {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return false
	}
	// Compact format has "unit" with underscore but no "dimension" field
	if _, hasDimension := raw["dimension"]; hasDimension {
		return false
	}
	if unitVal, hasUnit := raw["unit"]; hasUnit {
		if unitStr, ok := unitVal.(string); ok {
			return strings.Contains(unitStr, "_")
		}
	}
	return false
}
