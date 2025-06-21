// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

// ElectricChargeUnit represents a unit of electric charge
type ElectricChargeUnit struct {
	BaseUnit
}

// ElectricCharge contains predefined electric charge units
var ElectricCharge = struct {
	Coulomb          ElectricChargeUnit
	Millicoulomb     ElectricChargeUnit
	Microcoulomb     ElectricChargeUnit
	Ampere_Hour      ElectricChargeUnit
	Milliampere_Hour ElectricChargeUnit
}{
	Coulomb: ElectricChargeUnit{
		BaseUnit: NewBaseUnit(
			"electric_charge",
			"C",
			"Coulomb",
			1.0,
			0.0,
			true, // Base unit
		),
	},
	Millicoulomb: ElectricChargeUnit{
		BaseUnit: NewBaseUnit(
			"electric_charge",
			"mC",
			"Millicoulomb",
			0.001, // 1 mC = 0.001 C
			0.0,
			false,
		),
	},
	Microcoulomb: ElectricChargeUnit{
		BaseUnit: NewBaseUnit(
			"electric_charge",
			"µC",
			"Microcoulomb",
			0.000001, // 1 µC = 0.000001 C
			0.0,
			false,
		),
	},
	Ampere_Hour: ElectricChargeUnit{
		BaseUnit: NewBaseUnit(
			"electric_charge",
			"Ah",
			"Ampere-hour",
			3600.0, // 1 Ah = 3600 C
			0.0,
			false,
		),
	},
	Milliampere_Hour: ElectricChargeUnit{
		BaseUnit: NewBaseUnit(
			"electric_charge",
			"mAh",
			"Milliampere-hour",
			3.6, // 1 mAh = 3.6 C
			0.0,
			false,
		),
	},
}

// NewElectricCharge creates a new electric charge measurement
func NewElectricCharge(value float64, unit ElectricChargeUnit) Quantity[ElectricChargeUnit] {
	return New(value, unit)
}
