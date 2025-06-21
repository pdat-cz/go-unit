// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ParseError represents an error that occurred during parsing a measurement string
type ParseError struct {
	Input string
	Msg   string
}

// Error returns the error message
func (e ParseError) Error() string {
	return fmt.Sprintf("failed to parse measurement '%s': %s", e.Input, e.Msg)
}

// Regular expression to match a measurement string like "22.5°C" or "101.3 kPa"
var measurementRegex = regexp.MustCompile(`^([-+]?\d*\.?\d+)\s*([^\d\s].*)$`)

// ParseTemperature parses a string like "22.5°C" into a Temperature measurement
func ParseTemperature(s string) (Quantity[TemperatureUnit], error) {
	value, unitStr, err := parseValueAndUnit(s)
	if err != nil {
		return Quantity[TemperatureUnit]{}, err
	}

	// Find the matching temperature unit
	var unit TemperatureUnit
	found := false

	switch strings.ToLower(unitStr) {
	case "c", "°c", "celsius":
		unit = Temperature.Celsius
		found = true
	case "f", "°f", "fahrenheit":
		unit = Temperature.Fahrenheit
		found = true
	case "k", "kelvin":
		unit = Temperature.Kelvin
		found = true
	}

	if !found {
		return Quantity[TemperatureUnit]{}, ParseError{
			Input: s,
			Msg:   fmt.Sprintf("unknown temperature unit: %s", unitStr),
		}
	}

	return NewTemperature(value, unit), nil
}

// ParsePressure parses a string like "101.3 kPa" into a Pressure measurement
func ParsePressure(s string) (Quantity[PressureUnit], error) {
	value, unitStr, err := parseValueAndUnit(s)
	if err != nil {
		return Quantity[PressureUnit]{}, err
	}

	// Find the matching pressure unit
	var unit PressureUnit
	found := false

	switch strings.ToLower(unitStr) {
	case "pa", "pascal":
		unit = Pressure.Pascal
		found = true
	case "kpa", "kilopascal":
		unit = Pressure.Kilopascal
		found = true
	case "bar":
		unit = Pressure.Bar
		found = true
	case "psi":
		unit = Pressure.PSI
		found = true
	case "inh2o", "inh₂o", "inch water":
		unit = Pressure.InchH2O
		found = true
	}

	if !found {
		return Quantity[PressureUnit]{}, ParseError{
			Input: s,
			Msg:   fmt.Sprintf("unknown pressure unit: %s", unitStr),
		}
	}

	return NewPressure(value, unit), nil
}

// Helper function to parse a string into a value and unit string
func parseValueAndUnit(s string) (float64, string, error) {
	s = strings.TrimSpace(s)
	matches := measurementRegex.FindStringSubmatch(s)
	if matches == nil {
		return 0, "", ParseError{
			Input: s,
			Msg:   "invalid format, expected '<value><unit>' (e.g., '22.5°C')",
		}
	}

	valueStr := matches[1]
	unitStr := strings.TrimSpace(matches[2])

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return 0, "", ParseError{
			Input: s,
			Msg:   fmt.Sprintf("invalid number: %s", valueStr),
		}
	}

	return value, unitStr, nil
}

// ParseLength parses a string like "10.5 m" into a Length measurement
func ParseLength(s string) (Quantity[LengthUnit], error) {
	value, unitStr, err := parseValueAndUnit(s)
	if err != nil {
		return Quantity[LengthUnit]{}, err
	}

	// Find the matching length unit
	var unit LengthUnit
	found := false

	switch strings.ToLower(unitStr) {
	case "m", "meter", "meters":
		unit = Length.Meter
		found = true
	case "km", "kilometer", "kilometers":
		unit = Length.Kilometer
		found = true
	case "cm", "centimeter", "centimeters":
		unit = Length.Centimeter
		found = true
	case "mm", "millimeter", "millimeters":
		unit = Length.Millimeter
		found = true
	case "µm", "um", "micrometer", "micrometers":
		unit = Length.Micrometer
		found = true
	case "nm", "nanometer", "nanometers":
		unit = Length.Nanometer
		found = true
	case "in", "inch", "inches":
		unit = Length.Inch
		found = true
	case "ft", "foot", "feet":
		unit = Length.Foot
		found = true
	case "yd", "yard", "yards":
		unit = Length.Yard
		found = true
	case "mi", "mile", "miles":
		unit = Length.Mile
		found = true
	}

	if !found {
		return Quantity[LengthUnit]{}, ParseError{
			Input: s,
			Msg:   fmt.Sprintf("unknown length unit: %s", unitStr),
		}
	}

	return NewLength(value, unit), nil
}

// ParseMass parses a string like "75 kg" into a Mass measurement
func ParseMass(s string) (Quantity[MassUnit], error) {
	value, unitStr, err := parseValueAndUnit(s)
	if err != nil {
		return Quantity[MassUnit]{}, err
	}

	// Find the matching mass unit
	var unit MassUnit
	found := false

	switch strings.ToLower(unitStr) {
	case "kg", "kilogram", "kilograms":
		unit = Mass.Kilogram
		found = true
	case "g", "gram", "grams":
		unit = Mass.Gram
		found = true
	case "mg", "milligram", "milligrams":
		unit = Mass.Milligram
		found = true
	case "µg", "ug", "microgram", "micrograms":
		unit = Mass.Microgram
		found = true
	case "lb", "pound", "pounds":
		unit = Mass.Pound
		found = true
	case "oz", "ounce", "ounces":
		unit = Mass.Ounce
		found = true
	case "st", "stone", "stones":
		unit = Mass.Stone
		found = true
	case "t", "metric ton", "metric tons":
		unit = Mass.MetricTon
		found = true
	case "ton", "tons", "short ton", "short tons":
		unit = Mass.Ton
		found = true
	}

	if !found {
		return Quantity[MassUnit]{}, ParseError{
			Input: s,
			Msg:   fmt.Sprintf("unknown mass unit: %s", unitStr),
		}
	}

	return NewMass(value, unit), nil
}

// ParseDuration parses a string like "30 min" into a Duration measurement
func ParseDuration(s string) (Quantity[DurationUnit], error) {
	value, unitStr, err := parseValueAndUnit(s)
	if err != nil {
		return Quantity[DurationUnit]{}, err
	}

	// Find the matching duration unit
	var unit DurationUnit
	found := false

	switch strings.ToLower(unitStr) {
	case "s", "sec", "second", "seconds":
		unit = Duration.Second
		found = true
	case "min", "minute", "minutes":
		unit = Duration.Minute
		found = true
	case "h", "hr", "hour", "hours":
		unit = Duration.Hour
		found = true
	case "d", "day", "days":
		unit = Duration.Day
		found = true
	case "ms", "millisecond", "milliseconds":
		unit = Duration.Millisecond
		found = true
	case "µs", "us", "microsecond", "microseconds":
		unit = Duration.Microsecond
		found = true
	case "ns", "nanosecond", "nanoseconds":
		unit = Duration.Nanosecond
		found = true
	}

	if !found {
		return Quantity[DurationUnit]{}, ParseError{
			Input: s,
			Msg:   fmt.Sprintf("unknown duration unit: %s", unitStr),
		}
	}

	return NewDuration(value, unit), nil
}

// ParseAngle parses a string like "90°" into an Angle measurement
func ParseAngle(s string) (Quantity[AngleUnit], error) {
	value, unitStr, err := parseValueAndUnit(s)
	if err != nil {
		return Quantity[AngleUnit]{}, err
	}

	// Find the matching angle unit
	var unit AngleUnit
	found := false

	switch strings.ToLower(unitStr) {
	case "rad", "radian", "radians":
		unit = Angle.Radian
		found = true
	case "°", "deg", "degree", "degrees":
		unit = Angle.Degree
		found = true
	case "′", "arcmin", "arcminute", "arcminutes":
		unit = Angle.Arcminute
		found = true
	case "″", "arcsec", "arcsecond", "arcseconds":
		unit = Angle.Arcsecond
		found = true
	case "rev", "revolution", "revolutions":
		unit = Angle.Revolution
		found = true
	case "grad", "gradian", "gradians":
		unit = Angle.Gradian
		found = true
	}

	if !found {
		return Quantity[AngleUnit]{}, ParseError{
			Input: s,
			Msg:   fmt.Sprintf("unknown angle unit: %s", unitStr),
		}
	}

	return NewAngle(value, unit), nil
}

// ParseArea parses a string like "100 m²" into an Area measurement
func ParseArea(s string) (Quantity[AreaUnit], error) {
	value, unitStr, err := parseValueAndUnit(s)
	if err != nil {
		return Quantity[AreaUnit]{}, err
	}

	// Find the matching area unit
	var unit AreaUnit
	found := false

	switch strings.ToLower(unitStr) {
	case "m²", "m2", "sq m", "square meter", "square meters":
		unit = Area.SquareMeter
		found = true
	case "km²", "km2", "sq km", "square kilometer", "square kilometers":
		unit = Area.SquareKilometer
		found = true
	case "cm²", "cm2", "sq cm", "square centimeter", "square centimeters":
		unit = Area.SquareCentimeter
		found = true
	case "mm²", "mm2", "sq mm", "square millimeter", "square millimeters":
		unit = Area.SquareMillimeter
		found = true
	case "in²", "in2", "sq in", "square inch", "square inches":
		unit = Area.SquareInch
		found = true
	case "ft²", "ft2", "sq ft", "square foot", "square feet":
		unit = Area.SquareFoot
		found = true
	case "yd²", "yd2", "sq yd", "square yard", "square yards":
		unit = Area.SquareYard
		found = true
	case "mi²", "mi2", "sq mi", "square mile", "square miles":
		unit = Area.SquareMile
		found = true
	case "ac", "acre", "acres":
		unit = Area.Acre
		found = true
	case "ha", "hectare", "hectares":
		unit = Area.Hectare
		found = true
	}

	if !found {
		return Quantity[AreaUnit]{}, ParseError{
			Input: s,
			Msg:   fmt.Sprintf("unknown area unit: %s", unitStr),
		}
	}

	return NewArea(value, unit), nil
}

// ParseVolume parses a string like "10 L" into a Volume measurement
func ParseVolume(s string) (Quantity[VolumeUnit], error) {
	value, unitStr, err := parseValueAndUnit(s)
	if err != nil {
		return Quantity[VolumeUnit]{}, err
	}

	// Find the matching volume unit
	var unit VolumeUnit
	found := false

	switch strings.ToLower(unitStr) {
	case "m³", "m3", "cu m", "cubic meter", "cubic meters":
		unit = Volume.CubicMeter
		found = true
	case "km³", "km3", "cu km", "cubic kilometer", "cubic kilometers":
		unit = Volume.CubicKilometer
		found = true
	case "cm³", "cm3", "cc", "cu cm", "cubic centimeter", "cubic centimeters":
		unit = Volume.CubicCentimeter
		found = true
	case "mm³", "mm3", "cu mm", "cubic millimeter", "cubic millimeters":
		unit = Volume.CubicMillimeter
		found = true
	case "l", "liter", "liters", "litre", "litres":
		unit = Volume.Liter
		found = true
	case "ml", "milliliter", "milliliters", "millilitre", "millilitres":
		unit = Volume.Milliliter
		found = true
	case "in³", "in3", "cu in", "cubic inch", "cubic inches":
		unit = Volume.CubicInch
		found = true
	case "ft³", "ft3", "cu ft", "cubic foot", "cubic feet":
		unit = Volume.CubicFoot
		found = true
	case "yd³", "yd3", "cu yd", "cubic yard", "cubic yards":
		unit = Volume.CubicYard
		found = true
	case "gal", "gallon", "gallons":
		unit = Volume.Gallon
		found = true
	case "qt", "quart", "quarts":
		unit = Volume.Quart
		found = true
	case "pt", "pint", "pints":
		unit = Volume.Pint
		found = true
	case "cup", "cups":
		unit = Volume.Cup
		found = true
	case "fl oz", "fluid ounce", "fluid ounces":
		unit = Volume.FluidOunce
		found = true
	}

	if !found {
		return Quantity[VolumeUnit]{}, ParseError{
			Input: s,
			Msg:   fmt.Sprintf("unknown volume unit: %s", unitStr),
		}
	}

	return NewVolume(value, unit), nil
}

// ParseAcceleration parses a string like "9.8 m/s²" into an Acceleration measurement
func ParseAcceleration(s string) (Quantity[AccelerationUnit], error) {
	value, unitStr, err := parseValueAndUnit(s)
	if err != nil {
		return Quantity[AccelerationUnit]{}, err
	}

	// Find the matching acceleration unit
	var unit AccelerationUnit
	found := false

	switch strings.ToLower(unitStr) {
	case "m/s²", "m/s2", "meters per second squared":
		unit = Acceleration.MetersPerSecondSquared
		found = true
	case "g", "g-force":
		unit = Acceleration.G
		found = true
	case "ft/s²", "ft/s2", "feet per second squared":
		unit = Acceleration.FeetPerSecondSquared
		found = true
	}

	if !found {
		return Quantity[AccelerationUnit]{}, ParseError{
			Input: s,
			Msg:   fmt.Sprintf("unknown acceleration unit: %s", unitStr),
		}
	}

	return NewAcceleration(value, unit), nil
}

// ParseConcentration parses a string like "5 g/L" into a Concentration measurement
func ParseConcentration(s string) (Quantity[ConcentrationUnit], error) {
	value, unitStr, err := parseValueAndUnit(s)
	if err != nil {
		return Quantity[ConcentrationUnit]{}, err
	}

	// Find the matching concentration unit
	var unit ConcentrationUnit
	found := false

	switch strings.ToLower(unitStr) {
	case "g/l", "grams per liter":
		unit = Concentration.GramsPerLiter
		found = true
	case "mg/l", "milligrams per liter":
		unit = Concentration.MilligramsPerLiter
		found = true
	case "ppm", "parts per million":
		unit = Concentration.PartsPerMillion
		found = true
	case "ppb", "parts per billion":
		unit = Concentration.PartsPerBillion
		found = true
	}

	if !found {
		return Quantity[ConcentrationUnit]{}, ParseError{
			Input: s,
			Msg:   fmt.Sprintf("unknown concentration unit: %s", unitStr),
		}
	}

	return NewConcentration(value, unit), nil
}

// ParseDispersion parses a string like "5 ppm" into a Dispersion measurement
func ParseDispersion(s string) (Quantity[DispersionUnit], error) {
	value, unitStr, err := parseValueAndUnit(s)
	if err != nil {
		return Quantity[DispersionUnit]{}, err
	}

	// Find the matching dispersion unit
	var unit DispersionUnit
	found := false

	switch strings.ToLower(unitStr) {
	case "ppm", "parts per million":
		unit = Dispersion.PartsPerMillion
		found = true
	case "ppb", "parts per billion":
		unit = Dispersion.PartsPerBillion
		found = true
	case "ppt", "parts per trillion":
		unit = Dispersion.PartsPerTrillion
		found = true
	case "%", "percent":
		unit = Dispersion.Percent
		found = true
	}

	if !found {
		return Quantity[DispersionUnit]{}, ParseError{
			Input: s,
			Msg:   fmt.Sprintf("unknown dispersion unit: %s", unitStr),
		}
	}

	return NewDispersion(value, unit), nil
}

// ParseElectricCharge parses a string like "5 C" into an ElectricCharge measurement
func ParseElectricCharge(s string) (Quantity[ElectricChargeUnit], error) {
	value, unitStr, err := parseValueAndUnit(s)
	if err != nil {
		return Quantity[ElectricChargeUnit]{}, err
	}

	// Find the matching electric charge unit
	var unit ElectricChargeUnit
	found := false

	switch strings.ToLower(unitStr) {
	case "c", "coulomb", "coulombs":
		unit = ElectricCharge.Coulomb
		found = true
	case "mc", "millicoulomb", "millicoulombs":
		unit = ElectricCharge.Millicoulomb
		found = true
	case "µc", "uc", "microcoulomb", "microcoulombs":
		unit = ElectricCharge.Microcoulomb
		found = true
	case "ah", "ampere-hour", "ampere-hours":
		unit = ElectricCharge.Ampere_Hour
		found = true
	case "mah", "milliampere-hour", "milliampere-hours":
		unit = ElectricCharge.Milliampere_Hour
		found = true
	}

	if !found {
		return Quantity[ElectricChargeUnit]{}, ParseError{
			Input: s,
			Msg:   fmt.Sprintf("unknown electric charge unit: %s", unitStr),
		}
	}

	return NewElectricCharge(value, unit), nil
}

// ParseElectricCurrent parses a string like "5 A" into an ElectricCurrent measurement
func ParseElectricCurrent(s string) (Quantity[ElectricCurrentUnit], error) {
	value, unitStr, err := parseValueAndUnit(s)
	if err != nil {
		return Quantity[ElectricCurrentUnit]{}, err
	}

	// Find the matching electric current unit
	var unit ElectricCurrentUnit
	found := false

	switch strings.ToLower(unitStr) {
	case "a", "ampere", "amperes":
		unit = ElectricCurrent.Ampere
		found = true
	case "ma", "milliampere", "milliamperes":
		unit = ElectricCurrent.Milliampere
		found = true
	case "µa", "ua", "microampere", "microamperes":
		unit = ElectricCurrent.Microampere
		found = true
	case "ka", "kiloampere", "kiloamperes":
		unit = ElectricCurrent.Kiloampere
		found = true
	}

	if !found {
		return Quantity[ElectricCurrentUnit]{}, ParseError{
			Input: s,
			Msg:   fmt.Sprintf("unknown electric current unit: %s", unitStr),
		}
	}

	return NewElectricCurrent(value, unit), nil
}

// ParseSpeed parses a string like "10 m/s" into a Speed measurement
func ParseSpeed(s string) (Quantity[SpeedUnit], error) {
	value, unitStr, err := parseValueAndUnit(s)
	if err != nil {
		return Quantity[SpeedUnit]{}, err
	}

	// Find the matching speed unit
	var unit SpeedUnit
	found := false

	switch strings.ToLower(unitStr) {
	case "m/s", "meters per second", "metres per second":
		unit = Speed.MetersPerSecond
		found = true
	case "km/h", "kilometers per hour", "kilometres per hour", "kph":
		unit = Speed.KilometersPerHour
		found = true
	case "mph", "miles per hour":
		unit = Speed.MilesPerHour
		found = true
	case "ft/s", "feet per second", "foot per second":
		unit = Speed.FeetPerSecond
		found = true
	case "kn", "knot", "knots":
		unit = Speed.Knot
		found = true
	}

	if !found {
		return Quantity[SpeedUnit]{}, ParseError{
			Input: s,
			Msg:   fmt.Sprintf("unknown speed unit: %s", unitStr),
		}
	}

	return NewSpeed(value, unit), nil
}

// ParseElectricPotentialDifference parses a string like "5 V" into an ElectricPotentialDifference measurement
func ParseElectricPotentialDifference(s string) (Quantity[ElectricPotentialDifferenceUnit], error) {
	value, unitStr, err := parseValueAndUnit(s)
	if err != nil {
		return Quantity[ElectricPotentialDifferenceUnit]{}, err
	}

	// Find the matching electric potential difference unit
	var unit ElectricPotentialDifferenceUnit
	found := false

	switch strings.ToLower(unitStr) {
	case "v", "volt", "volts":
		unit = ElectricPotentialDifference.Volt
		found = true
	case "mv", "millivolt", "millivolts":
		unit = ElectricPotentialDifference.Millivolt
		found = true
	case "µv", "uv", "microvolt", "microvolts":
		unit = ElectricPotentialDifference.Microvolt
		found = true
	case "kv", "kilovolt", "kilovolts":
		unit = ElectricPotentialDifference.Kilovolt
		found = true
	case "megav", "megavolt", "megavolts":
		unit = ElectricPotentialDifference.Megavolt
		found = true
	}

	if !found {
		return Quantity[ElectricPotentialDifferenceUnit]{}, ParseError{
			Input: s,
			Msg:   fmt.Sprintf("unknown electric potential difference unit: %s", unitStr),
		}
	}

	return NewElectricPotentialDifference(value, unit), nil
}

// FormatWithUnit formats a value with its unit symbol
func FormatWithUnit(value float64, unitSymbol string) string {
	return fmt.Sprintf("%g %s", value, unitSymbol)
}
