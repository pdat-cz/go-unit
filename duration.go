// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

// DurationUnit represents a unit of time duration
type DurationUnit struct {
	BaseUnit
}

// Duration contains predefined duration units
var Duration = struct {
	Second      DurationUnit
	Minute      DurationUnit
	Hour        DurationUnit
	Day         DurationUnit
	Millisecond DurationUnit
	Microsecond DurationUnit
	Nanosecond  DurationUnit
}{
	Second: DurationUnit{
		BaseUnit: NewBaseUnit(
			"duration",
			"s",
			"Second",
			1.0,
			0.0,
			true, // Base unit
		),
	},
	Minute: DurationUnit{
		BaseUnit: NewBaseUnit(
			"duration",
			"min",
			"Minute",
			60.0, // 1 min = 60 s
			0.0,
			false,
		),
	},
	Hour: DurationUnit{
		BaseUnit: NewBaseUnit(
			"duration",
			"h",
			"Hour",
			3600.0, // 1 h = 3600 s
			0.0,
			false,
		),
	},
	Day: DurationUnit{
		BaseUnit: NewBaseUnit(
			"duration",
			"d",
			"Day",
			86400.0, // 1 d = 86400 s
			0.0,
			false,
		),
	},
	Millisecond: DurationUnit{
		BaseUnit: NewBaseUnit(
			"duration",
			"ms",
			"Millisecond",
			0.001, // 1 ms = 0.001 s
			0.0,
			false,
		),
	},
	Microsecond: DurationUnit{
		BaseUnit: NewBaseUnit(
			"duration",
			"µs",
			"Microsecond",
			0.000001, // 1 µs = 0.000001 s
			0.0,
			false,
		),
	},
	Nanosecond: DurationUnit{
		BaseUnit: NewBaseUnit(
			"duration",
			"ns",
			"Nanosecond",
			0.000000001, // 1 ns = 0.000000001 s
			0.0,
			false,
		),
	},
}

// NewDuration creates a new duration quantity
func NewDuration(value float64, unit DurationUnit) Quantity[DurationUnit] {
	return New(value, unit)
}
