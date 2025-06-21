// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

// ElectricCurrentUnit represents a unit of electric current
type ElectricCurrentUnit struct {
	BaseUnit
}

// ElectricCurrent contains predefined electric current units
var ElectricCurrent = struct {
	Ampere      ElectricCurrentUnit
	Milliampere ElectricCurrentUnit
	Microampere ElectricCurrentUnit
	Kiloampere  ElectricCurrentUnit
}{
	Ampere: ElectricCurrentUnit{
		BaseUnit: NewBaseUnit(
			"electric_current",
			"A",
			"Ampere",
			1.0,
			0.0,
			true, // Base unit
		),
	},
	Milliampere: ElectricCurrentUnit{
		BaseUnit: NewBaseUnit(
			"electric_current",
			"mA",
			"Milliampere",
			0.001, // 1 mA = 0.001 A
			0.0,
			false,
		),
	},
	Microampere: ElectricCurrentUnit{
		BaseUnit: NewBaseUnit(
			"electric_current",
			"µA",
			"Microampere",
			0.000001, // 1 µA = 0.000001 A
			0.0,
			false,
		),
	},
	Kiloampere: ElectricCurrentUnit{
		BaseUnit: NewBaseUnit(
			"electric_current",
			"kA",
			"Kiloampere",
			1000.0, // 1 kA = 1000 A
			0.0,
			false,
		),
	},
}

// NewElectricCurrent creates a new electric current measurement
func NewElectricCurrent(value float64, unit ElectricCurrentUnit) Quantity[ElectricCurrentUnit] {
	return New(value, unit)
}
