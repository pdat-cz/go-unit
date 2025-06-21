// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

// AccelerationUnit represents a unit of acceleration
type AccelerationUnit struct {
	BaseUnit
}

// Acceleration contains predefined acceleration units
var Acceleration = struct {
	MetersPerSecondSquared AccelerationUnit
	G                      AccelerationUnit
	FeetPerSecondSquared   AccelerationUnit
}{
	MetersPerSecondSquared: AccelerationUnit{
		BaseUnit: NewBaseUnit(
			"acceleration",
			"m/s²",
			"Meters per Second Squared",
			1.0,
			0.0,
			true, // Base unit
		),
	},
	G: AccelerationUnit{
		BaseUnit: NewBaseUnit(
			"acceleration",
			"g",
			"G-force",
			9.80665, // 1 g = 9.80665 m/s²
			0.0,
			false,
		),
	},
	FeetPerSecondSquared: AccelerationUnit{
		BaseUnit: NewBaseUnit(
			"acceleration",
			"ft/s²",
			"Feet per Second Squared",
			0.3048, // 1 ft/s² = 0.3048 m/s²
			0.0,
			false,
		),
	},
}

// NewAcceleration creates a new acceleration measurement
func NewAcceleration(value float64, unit AccelerationUnit) Quantity[AccelerationUnit] {
	return New(value, unit)
}
