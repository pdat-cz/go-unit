// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

// SpeedUnit represents a unit of speed
type SpeedUnit struct {
	BaseUnit
}

// Speed contains predefined speed units
var Speed = struct {
	MetersPerSecond   SpeedUnit
	KilometersPerHour SpeedUnit
	MilesPerHour      SpeedUnit
	FeetPerSecond     SpeedUnit
	Knot              SpeedUnit
}{
	MetersPerSecond: SpeedUnit{
		BaseUnit: NewBaseUnit(
			"speed",
			"m/s",
			"Meters per Second",
			1.0,
			0.0,
			true, // Base unit
		),
	},
	KilometersPerHour: SpeedUnit{
		BaseUnit: NewBaseUnit(
			"speed",
			"km/h",
			"Kilometers per Hour",
			1.0/3.6, // 1 km/h = 0.277778 m/s
			0.0,
			false,
		),
	},
	MilesPerHour: SpeedUnit{
		BaseUnit: NewBaseUnit(
			"speed",
			"mph",
			"Miles per Hour",
			0.44704, // 1 mph = 0.44704 m/s
			0.0,
			false,
		),
	},
	FeetPerSecond: SpeedUnit{
		BaseUnit: NewBaseUnit(
			"speed",
			"ft/s",
			"Feet per Second",
			0.3048, // 1 ft/s = 0.3048 m/s
			0.0,
			false,
		),
	},
	Knot: SpeedUnit{
		BaseUnit: NewBaseUnit(
			"speed",
			"kn",
			"Knot",
			0.51444, // 1 knot = 0.51444 m/s
			0.0,
			false,
		),
	},
}

// NewSpeed creates a new speed quantity
func NewSpeed(value float64, unit SpeedUnit) Quantity[SpeedUnit] {
	return New(value, unit)
}
