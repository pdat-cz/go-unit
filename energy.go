// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

// EnergyUnit represents a unit of energy
type EnergyUnit struct {
	BaseUnit
}

// Energy contains predefined energy units
var Energy = struct {
	Joule        EnergyUnit
	Kilojoule    EnergyUnit
	Megajoule    EnergyUnit
	Gigajoule    EnergyUnit
	Terajoule    EnergyUnit
	KilowattHour EnergyUnit
	BTU          EnergyUnit
}{
	Joule: EnergyUnit{
		BaseUnit: NewBaseUnit(
			"energy",
			"J",
			"Joule",
			1.0,
			0.0,
			true, // Base unit
		),
	},
	Kilojoule: EnergyUnit{
		BaseUnit: NewBaseUnit(
			"energy",
			"kJ",
			"Kilojoule",
			1e3,
			0.0,
			false,
		),
	},
	Megajoule: EnergyUnit{
		BaseUnit: NewBaseUnit(
			"energy",
			"MJ",
			"Megajoule",
			1e6,
			0.0,
			false,
		),
	},
	Gigajoule: EnergyUnit{
		BaseUnit: NewBaseUnit(
			"energy",
			"GJ",
			"Gigajoule",
			1e9,
			0.0,
			false,
		),
	},
	Terajoule: EnergyUnit{
		BaseUnit: NewBaseUnit(
			"energy",
			"TJ",
			"Terajoule",
			1e12,
			0.0,
			false,
		),
	},
	KilowattHour: EnergyUnit{
		BaseUnit: NewBaseUnit(
			"energy",
			"kWh",
			"Kilowatt-hour",
			3600000.0, // 1 kWh = 3,600,000 J
			0.0,
			false,
		),
	},
	BTU: EnergyUnit{
		BaseUnit: NewBaseUnit(
			"energy",
			"BTU",
			"British Thermal Unit",
			1055.06, // 1 BTU = 1,055.06 J
			0.0,
			false,
		),
	},
}

// NewEnergy creates a new energy quantity
func NewEnergy(value float64, unit EnergyUnit) Quantity[EnergyUnit] {
	return New(value, unit)
}
