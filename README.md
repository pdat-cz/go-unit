# Unit Package

[![Build Status](https://github.com/pdat-cz/go-unit/actions/workflows/go.yml/badge.svg)](https://github.com/pdat-cz/go-unit/actions/workflows/go.yml)
[![Go Version](https://img.shields.io/github/go-mod/go-version/pdat-cz/go-unit)](https://github.com/pdat-cz/go-unit/blob/main/go.mod)
[![License](https://img.shields.io/github/license/pdat-cz/go-unit)](https://github.com/pdat-cz/go-unit/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/pdat-cz/go-unit)](https://goreportcard.com/report/github.com/pdat-cz/go-unit)

The `unit` package provides a system for representing, converting, and operating on physical quantities with
units. It ensures that:

## Installation

```bash
go get github.com/pdat-cz/go-unit
```

## Import

```go
import "github.com/pdat-cz/go-unit"
```

- Units are explicitly tracked with values
- Conversions between compatible units are handled automatically
- Type safety is maintained (preventing temperature from being added to pressure)
- Calculations involving different units work correctly

## Core Concepts

### Quantity

A `Quantity` represents a value with an associated unit. It combines a numeric value with a unit type to create a
complete representation of a physical quantity.

```go
type Quantity[T UnitType] struct {
Value float64
Unit  T
}
```

The `Quantity` type is generic, allowing it to work with different unit types while maintaining type safety.

### Unit Types

The package includes several predefined unit types:

- `TemperatureUnit`: Celsius, Fahrenheit, Kelvin
- `PressureUnit`: Pascal, Kilopascal, Bar, PSI, InchH2O
- `FlowRateUnit`: CubicMetersPerHour, LitersPerSecond, CFM
- `PowerUnit`: Watt, Kilowatt, BTUPerHour
- `EnergyUnit`: Joule, KilowattHour, BTU
- `LengthUnit`: Meter, Kilometer, Centimeter, Millimeter, Micrometer, Nanometer, Inch, Foot, Yard, Mile
- `MassUnit`: Kilogram, Gram, Milligram, Microgram, Pound, Ounce, Stone, MetricTon, Ton
- `DurationUnit`: Second, Minute, Hour, Day, Millisecond, Microsecond, Nanosecond
- `AngleUnit`: Radian, Degree, Arcminute, Arcsecond, Revolution, Gradian
- `AreaUnit`: SquareMeter, SquareKilometer, SquareCentimeter, SquareMillimeter, SquareInch, SquareFoot, SquareYard,
  SquareMile, Acre, Hectare
- `VolumeUnit`: CubicMeter, CubicKilometer, CubicCentimeter, CubicMillimeter, Liter, Milliliter, CubicInch, CubicFoot,
  CubicYard, Gallon, Quart, Pint, Cup, FluidOunce
- `AccelerationUnit`: MetersPerSecondSquared, G, FeetPerSecondSquared
- `ConcentrationUnit`: GramsPerLiter, MilligramsPerLiter, PartsPerMillion, PartsPerBillion
- `DispersionUnit`: PartsPerMillion, PartsPerBillion, PartsPerTrillion, Percent
- `ElectricChargeUnit`: Coulomb, Millicoulomb, Microcoulomb, Ampere_Hour, Milliampere_Hour
- `ElectricCurrentUnit`: Ampere, Milliampere, Microampere, Kiloampere
- `FrequencyUnit`: Hertz, Kilohertz, Megahertz, Gigahertz, Terahertz, RPM
- `FuelEfficiencyUnit`: KilometersPerLiter, MilesPerGallon, LitersPer100Kilometers
- `IlluminanceUnit`: Lux, FootCandle, Phot, Nox
- `InformationUnit`: Bit, Byte, Kilobyte, Megabyte, Gigabyte, Terabyte, Petabyte, Kibibyte, Mebibyte, Gibibyte,
  Tebibyte, Pebibyte
- `GeneralUnit`: A flexible unit type for custom units and project-specific measurements

Each unit type implements the `UnitType` interface, which provides methods for dimension information, unit conversion,
and equality checking.

## Usage Examples

### Creating Quantities

```go
// Create a temperature quantity of 22 degrees Celsius
temp := unit.NewTemperature(22.0, unit.Temperature.Celsius)

// Create a pressure quantity of 101.3 kPa
pressure := unit.NewPressure(101.3, unit.Pressure.Kilopascal)
```

### Converting Between Units

```go
// Convert temperature from Celsius to Fahrenheit
tempF := temp.ConvertTo(unit.Temperature.Fahrenheit)
// tempF.Value is now 71.6

// Convert pressure from kPa to PSI
pressurePSI := pressure.ConvertTo(unit.Pressure.PSI)
// pressurePSI.Value is now 14.69
```

### Arithmetic Operations

```go
// Add two temperatures (must be same dimension)
temp1 := unit.NewTemperature(22.0, unit.Temperature.Celsius)
temp2 := unit.NewTemperature(68.0, unit.Temperature.Fahrenheit)
sum := temp1.Add(temp2) // Automatically converts temp2 to Celsius before adding
// sum.Value is 42.0, sum.Unit is Celsius

// Multiply a quantity by a scalar
doubledTemp := temp1.MultiplyByScalar(2.0)
// doubledTemp.Value is 44.0, doubledTemp.Unit is Celsius
```

### Parsing from Strings

```go
// Parse a temperature from a string
temp, err := unit.ParseTemperature("25°C")
if err != nil {
// Handle error
}

// Parse a pressure from a string
pressure, err := unit.ParsePressure("101.325 kPa")
if err != nil {
// Handle error
}
```

### Serialization and Deserialization

Quantities implement `json.Marshaler` and `json.Unmarshaler` interfaces, so you can use standard Go JSON functions:

```go
temp := unit.NewTemperature(25, unit.Temperature.Celsius)

// Serialize using json.Marshal
data, err := json.Marshal(temp)
// Output: {"value":25,"unit":{"name":"Celsius","symbol":"°C"},"dimension":"temperature"}

// Deserialize using json.Unmarshal
var temp2 unit.Quantity[unit.TemperatureUnit]
err = json.Unmarshal(data, &temp2)
```

#### Explicit Marshal/Unmarshal Functions

You can also use the explicit functions for each dimension:

```go
// Serialize a temperature quantity to JSON
tempJSON, err := unit.MarshalTemperature(temp)
if err != nil {
    // Handle error
}

// Deserialize a temperature quantity from JSON
tempDeserialized, err := unit.UnmarshalTemperature(tempJSON)
if err != nil {
    // Handle error
}
```

#### Generic Deserialization

When you don't know the dimension in advance, you can use the generic `UnmarshalQuantity` function:

```go
// Deserialize a quantity without knowing its dimension in advance
anyQuantity, err := unit.UnmarshalQuantity(jsonData)
if err != nil {
// Handle error
}

// Check the dimension
dimension := anyQuantity.GetDimension()
fmt.Println("Dimension:", dimension)

// Convert to the specific type based on the dimension
switch dimension {
case "temperature":
temp, ok := anyQuantity.AsTemperature()
if ok {
fmt.Println("Temperature:", temp)
}
case "length":
length, ok := anyQuantity.AsLength()
if ok {
fmt.Println("Length:", length)
}
case "volume":
volume, ok := anyQuantity.AsVolume()
if ok {
fmt.Println("Volume:", volume)
}
// ... handle other dimensions
}
```

This approach is useful when processing quantities from external sources where the dimension isn't known until
runtime.

This allows quantities to be easily stored, transmitted, and reconstructed:

```go
// Example of embedding quantities in a larger JSON structure
type SensorReading struct {
ID          string          `json:"id"`
Timestamp   string          `json:"timestamp"`
Temperature json.RawMessage `json:"temperature,omitempty"`
Pressure    json.RawMessage `json:"pressure,omitempty"`
}

// Create a sensor reading with temperature
tempJSON, _ := unit.MarshalTemperature(temp)
reading := SensorReading{
ID:          "sensor1",
Timestamp:   "2023-06-15T12:34:56Z",
Temperature: tempJSON,
}

// Serialize the entire reading
readingJSON, _ := json.Marshal(reading)
// readingJSON contains the sensor ID, timestamp, and serialized temperature
```

### Compact Serialization

For systems that prefer snake_case keys (e.g., NATS subjects), use the `Compact[T]` wrapper type:

```go
temp := unit.NewTemperature(25, unit.Temperature.Celsius)

// Serialize using Compact wrapper with json.Marshal
data, err := json.Marshal(unit.Compact[unit.TemperatureUnit]{temp})
// Output: {"value":25,"unit":"temperature_celsius","symbol":"°C"}

// Deserialize using json.Unmarshal
var compact unit.Compact[unit.TemperatureUnit]
err = json.Unmarshal(data, &compact)
temp2 := compact.Quantity
```

#### Explicit Compact Functions

You can also use the explicit compact functions:

```go
// Compact format: {"value":25,"unit":"temperature_celsius"}
data, err := unit.MarshalCompactTemperature(temp)

// With symbol included: {"value":25,"unit":"temperature_celsius","symbol":"°C"}
data, err := unit.MarshalCompactTemperatureWithSymbol(temp)

// Deserialize compact format
temp, err := unit.UnmarshalCompactTemperature(data)
```

#### Auto-Detection

The generic `UnmarshalMeasurement` function auto-detects the format:

```go
// Works with both standard and compact formats
am, err := unit.UnmarshalMeasurement(jsonData)
```

#### JSON Format Comparison

| Type | JSON Output |
|------|-------------|
| `Quantity[T]` | `{"value":25,"unit":{"name":"Celsius","symbol":"°C"},"dimension":"temperature"}` |
| `Compact[T]` | `{"value":25,"unit":"temperature_celsius","symbol":"°C"}` |

## Custom Units

The package supports defining custom units for project-specific needs using the `GeneralUnit` type:

```go
// Create a custom unit
customUnit := unit.NewGeneralUnit("xyz", "Custom Unit")

// Create a quantity with the custom unit
myQuantity := unit.NewGeneral(42.0, customUnit)

// Create a unit with conversion factors
// For example: 1 abc = 2 xyz + 10
convertibleUnit := unit.NewGeneralUnitWithConversion("abc", "Another Unit", 2.0, 10.0)
```

For more advanced extension options, including creating your own quantity types in your project, see
the [EXTENDING.md](EXTENDING.md) documentation.

## Best Practices

- Always specify units when defining points that represent physical quantities
- Use the appropriate dimension for each physical quantity
- Avoid mixing dimensions in calculations
- When displaying values in UIs, show the unit symbol
- Allow users to choose their preferred units for display

## Future Extensions

- **Compound units**: Support for units like kWh/m² (energy per area)
- **Dimensional analysis**: Automatic tracking of dimensions in calculations
- **Uncertainty**: Track and propagate measurement uncertainty
