// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

// FlowRateUnit represents a unit of flow rate
type FlowRateUnit struct {
	BaseUnit
}

// FlowRate contains predefined flow rate units
var FlowRate = struct {
	CubicMetersPerHour FlowRateUnit
	LitersPerSecond    FlowRateUnit
	CFM                FlowRateUnit
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
