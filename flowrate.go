// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

// FlowRateUnit represents a unit of flow rate
type FlowRateUnit struct {
	BaseUnit
}

// FlowRate contains predefined flow rate units
var FlowRate = struct {
	CubicMetersPerHour   FlowRateUnit
	CubicMetersPerSecond FlowRateUnit
	LitersPerSecond      FlowRateUnit
	LitersPerMinute      FlowRateUnit
	LitersPerHour        FlowRateUnit
	GallonsPerMinute     FlowRateUnit
	CFM                  FlowRateUnit
}{
	CubicMetersPerHour: FlowRateUnit{
		BaseUnit: NewBaseUnit(
			"flowrate",
			"m³/h",
			"Cubic Meters per Hour",
			1.0,
			0.0,
			true, // Base unit
		),
	},
	CubicMetersPerSecond: FlowRateUnit{
		BaseUnit: NewBaseUnit(
			"flowrate",
			"m³/s",
			"Cubic Meters per Second",
			3600.0, // 1 m³/s = 3600 m³/h
			0.0,
			false,
		),
	},
	LitersPerSecond: FlowRateUnit{
		BaseUnit: NewBaseUnit(
			"flowrate",
			"L/s",
			"Liters per Second",
			3.6, // 1 L/s = 3.6 m³/h
			0.0,
			false,
		),
	},
	LitersPerMinute: FlowRateUnit{
		BaseUnit: NewBaseUnit(
			"flowrate",
			"L/min",
			"Liters per Minute",
			0.06, // 1 L/min = 0.06 m³/h
			0.0,
			false,
		),
	},
	LitersPerHour: FlowRateUnit{
		BaseUnit: NewBaseUnit(
			"flowrate",
			"L/h",
			"Liters per Hour",
			0.001, // 1 L/h = 0.001 m³/h
			0.0,
			false,
		),
	},
	GallonsPerMinute: FlowRateUnit{
		BaseUnit: NewBaseUnit(
			"flowrate",
			"GPM",
			"US Gallons per Minute",
			0.22712470704, // 1 US gpm = 0.22712470704 m³/h
			0.0,
			false,
		),
	},
	CFM: FlowRateUnit{
		BaseUnit: NewBaseUnit(
			"flowrate",
			"CFM",
			"Cubic Feet per Minute",
			1.699, // 1 CFM = 1.699 m³/h
			0.0,
			false,
		),
	},
}

// NewFlowRate creates a new flow rate quantity
func NewFlowRate(value float64, unit FlowRateUnit) Quantity[FlowRateUnit] {
	return New(value, unit)
}
