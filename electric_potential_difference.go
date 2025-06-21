// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

// ElectricPotentialDifferenceUnit represents a unit of electric potential difference
type ElectricPotentialDifferenceUnit struct {
	BaseUnit
}

// ElectricPotentialDifference contains predefined electric potential difference units
var ElectricPotentialDifference = struct {
	Volt      ElectricPotentialDifferenceUnit
	Millivolt ElectricPotentialDifferenceUnit
	Microvolt ElectricPotentialDifferenceUnit
	Kilovolt  ElectricPotentialDifferenceUnit
	Megavolt  ElectricPotentialDifferenceUnit
}{
	Volt: ElectricPotentialDifferenceUnit{
		BaseUnit: NewBaseUnit(
			"electric_potential_difference",
			"V",
			"Volt",
			1.0,
			0.0,
			true, // Base unit
		),
	},
	Millivolt: ElectricPotentialDifferenceUnit{
		BaseUnit: NewBaseUnit(
			"electric_potential_difference",
			"mV",
			"Millivolt",
			0.001, // 1 mV = 0.001 V
			0.0,
			false,
		),
	},
	Microvolt: ElectricPotentialDifferenceUnit{
		BaseUnit: NewBaseUnit(
			"electric_potential_difference",
			"µV",
			"Microvolt",
			0.000001, // 1 µV = 0.000001 V
			0.0,
			false,
		),
	},
	Kilovolt: ElectricPotentialDifferenceUnit{
		BaseUnit: NewBaseUnit(
			"electric_potential_difference",
			"kV",
			"Kilovolt",
			1000.0, // 1 kV = 1000 V
			0.0,
			false,
		),
	},
	Megavolt: ElectricPotentialDifferenceUnit{
		BaseUnit: NewBaseUnit(
			"electric_potential_difference",
			"MV",
			"Megavolt",
			1000000.0, // 1 MV = 1,000,000 V
			0.0,
			false,
		),
	},
}

// NewElectricPotentialDifference creates a new electric potential difference measurement
func NewElectricPotentialDifference(value float64, unit ElectricPotentialDifferenceUnit) Quantity[ElectricPotentialDifferenceUnit] {
	return New(value, unit)
}
