# Extending the go-unit Package

This document provides detailed information on how to extend the go-unit package with custom quantities and units in your own project.

## Table of Contents

- [Using Custom Units](#using-custom-units)
  - [Basic Custom Units](#basic-custom-units)
  - [Custom Units with Conversion Factors](#custom-units-with-conversion-factors)
  - [Serialization and Deserialization](#serialization-and-deserialization)
- [Creating New Quantity Types](#creating-new-quantity-types)
  - [Creating a New Quantity Type in Your Project](#creating-a-new-quantity-type-in-your-project)
  - [Complete Example: Luminosity Quantity](#complete-example-luminosity-quantity)
- [Best Practices](#best-practices)

## Using Custom Units

The go-unit package provides built-in support for custom units through the `GeneralUnit` type. This is the simplest approach and doesn't require creating new types.

### Basic Custom Units

For simple custom units without specific conversion factors:

```go
package main

import (
    "fmt"
    "github.com/pdat-cz/go-unit"
)

func main() {
    // Create a custom unit
    customUnit := unit.NewGeneralUnit("xyz", "Custom Unit")

    // Create a quantity with the custom unit
    myQuantity := unit.NewGeneral(42.0, customUnit)

    fmt.Println(myQuantity) // Output: 42 xyz
}
```

### Custom Units with Conversion Factors

For units that need to be convertible to other units:

```go
package main

import (
    "fmt"
    "github.com/pdat-cz/go-unit"
)

func main() {
    // Create a base unit
    baseUnit := unit.NewGeneralUnit("base", "Base Unit")

    // Create a unit with conversion factors relative to the base unit
    // For example: 1 derived = 2.5 base + 10
    derivedUnit := unit.NewGeneralUnitWithConversion("derived", "Derived Unit", 2.5, 10.0)

    // Create quantities with these units
    baseQuantity := unit.NewGeneral(100.0, baseUnit)
    derivedQuantity := unit.NewGeneral(50.0, derivedUnit)

    // Convert between units
    convertedToBase := derivedQuantity.ConvertTo(baseUnit)
    fmt.Println(convertedToBase) // Output: 135 base (calculated as: (50*2.5)+10)

    convertedToDerived := baseQuantity.ConvertTo(derivedUnit)
    fmt.Println(convertedToDerived) // Output: 36 derived (calculated as: (100-10)/2.5)
}
```

### Serialization and Deserialization

Custom units work seamlessly with the serialization system:

```go
package main

import (
    "fmt"
    "github.com/pdat-cz/go-unit"
)

func main() {
    // Create a quantity with a custom unit
    customUnit := unit.NewGeneralUnit("xyz", "Custom Unit")
    myQuantity := unit.NewGeneral(42.0, customUnit)

    // Serialize to JSON
    jsonData, err := unit.MarshalGeneral(myQuantity)
    if err != nil {
        // Handle error
    }
    fmt.Println(string(jsonData)) // Output: {"value":42,"unit":"xyz","dimension":"general"}

    // Deserialize from JSON
    deserializedQuantity, err := unit.UnmarshalGeneral(jsonData)
    if err != nil {
        // Handle error
    }
    fmt.Println(deserializedQuantity) // Output: 42 xyz

    // Using the generic deserializer
    anyQuantity, err := unit.UnmarshalQuantity(jsonData)
    if err != nil {
        // Handle error
    }

    // Check if it's a general quantity
    if anyQuantity.GetDimension() == "general" {
        generalQ, ok := anyQuantity.AsGeneral()
        if ok {
            fmt.Println("General quantity:", generalQ)
        }
    }
}
```

## Creating New Quantity Types

For more complex scenarios where you need type safety and dedicated functionality, you can create your own quantity types in your project.

### Creating a New Quantity Type in Your Project

When importing go-unit as a dependency, you can create your own quantity types that leverage the package's infrastructure:

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/pdat-cz/go-unit"
)

// Step 1: Define your unit type
type LuminosityUnit struct {
    unit.BaseUnit
}

// Step 2: Define your predefined units
var Luminosity = struct {
    Candela     LuminosityUnit
    Lumen       LuminosityUnit
    Candlepower LuminosityUnit
}{
    Candela: LuminosityUnit{
        BaseUnit: unit.NewBaseUnit(
            "luminosity",
            "cd",
            "Candela",
            1.0,
            0.0,
            true, // Base unit
        ),
    },
    Lumen: LuminosityUnit{
        BaseUnit: unit.NewBaseUnit(
            "luminosity",
            "lm",
            "Lumen",
            1.0, // Conversion factor
            0.0,
            false,
        ),
    },
    Candlepower: LuminosityUnit{
        BaseUnit: unit.NewBaseUnit(
            "luminosity",
            "cp",
            "Candlepower",
            0.981, // 1 cp ≈ 0.981 cd
            0.0,
            false,
        ),
    },
}

// Step 3: Create a constructor function
func NewLuminosity(value float64, unit LuminosityUnit) unit.Quantity[LuminosityUnit] {
    return unit.New(value, unit)
}

// Step 4: Add serialization support (optional)
func MarshalLuminosity(m unit.Quantity[LuminosityUnit]) ([]byte, error) {
    return json.Marshal(unit.QuantityJSON{
        Value:     m.Value,
        Unit:      m.Unit.Symbol(),
        Dimension: m.Unit.Dimension(),
    })
}

// Step 5: Add deserialization support (optional)
func UnmarshalLuminosity(data []byte) (unit.Quantity[LuminosityUnit], error) {
    var jsonM unit.QuantityJSON
    if err := json.Unmarshal(data, &jsonM); err != nil {
        return unit.Quantity[LuminosityUnit]{}, err
    }

    // Verify dimension
    if jsonM.Dimension != "luminosity" {
        return unit.Quantity[LuminosityUnit]{}, fmt.Errorf("expected dimension 'luminosity', got '%s'", jsonM.Dimension)
    }

    // Find the matching luminosity unit
    var unit LuminosityUnit
    found := false

    switch jsonM.Unit {
    case "cd":
        unit = Luminosity.Candela
        found = true
    case "lm":
        unit = Luminosity.Lumen
        found = true
    case "cp":
        unit = Luminosity.Candlepower
        found = true
    }

    if !found {
        return unit.Quantity[LuminosityUnit]{}, fmt.Errorf("unknown luminosity unit: %s", jsonM.Unit)
    }

    return NewLuminosity(jsonM.Value, unit), nil
}
```

### Complete Example: Luminosity Quantity

Here's a complete example showing how to use your custom quantity type:

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/pdat-cz/go-unit"
)

// LuminosityUnit and related definitions from above...

func main() {
    // Create a luminosity quantity
    candela := NewLuminosity(100.0, Luminosity.Candela)

    // Convert to another unit
    candlepower := candela.ConvertTo(Luminosity.Candlepower)
    fmt.Println(candlepower) // Output: 101.94 cp

    // Perform arithmetic
    doubleCandela := candela.MultiplyByScalar(2.0)
    fmt.Println(doubleCandela) // Output: 200 cd

    // Serialize to JSON
    jsonData, _ := MarshalLuminosity(candela)
    fmt.Println(string(jsonData)) // Output: {"value":100,"unit":"cd","dimension":"luminosity"}

    // Deserialize from JSON
    deserializedQuantity, _ := UnmarshalLuminosity(jsonData)
    fmt.Println(deserializedQuantity) // Output: 100 cd
}
```

## Best Practices

1. **Choose the right approach**:
   - For simple custom units, use `GeneralUnit`
   - For complex domain-specific units with type safety, create a new quantity type

2. **Naming conventions**:
   - Follow the existing pattern: `<Type>Unit` for unit types
   - Use `New<Type>` for constructor functions

3. **Unit definitions**:
   - Make the base unit clear (set `isBase: true`)
   - Document conversion factors and formulas
   - Use SI units as base units when possible

4. **Testing**:
   - Test conversion between units
   - Test serialization and deserialization
   - Test arithmetic operations

5. **Documentation**:
   - Document the physical meaning of each unit
   - Provide examples of typical usage
