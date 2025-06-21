// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

import "math"

// AngleUnit represents a unit of angle
type AngleUnit struct {
	BaseUnit
}

// Angle contains predefined angle units
var Angle = struct {
	Radian     AngleUnit
	Degree     AngleUnit
	Arcminute  AngleUnit
	Arcsecond  AngleUnit
	Revolution AngleUnit
	Gradian    AngleUnit
}{
	Radian: AngleUnit{
		BaseUnit: NewBaseUnit(
			"angle",
			"rad",
			"Radian",
			1.0,
			0.0,
			true, // Base unit
		),
	},
	Degree: AngleUnit{
		BaseUnit: NewBaseUnit(
			"angle",
			"°",
			"Degree",
			math.Pi/180.0, // 1° = π/180 rad
			0.0,
			false,
		),
	},
	Arcminute: AngleUnit{
		BaseUnit: NewBaseUnit(
			"angle",
			"′",
			"Arcminute",
			math.Pi/(180.0*60.0), // 1′ = π/(180*60) rad
			0.0,
			false,
		),
	},
	Arcsecond: AngleUnit{
		BaseUnit: NewBaseUnit(
			"angle",
			"″",
			"Arcsecond",
			math.Pi/(180.0*60.0*60.0), // 1″ = π/(180*60*60) rad
			0.0,
			false,
		),
	},
	Revolution: AngleUnit{
		BaseUnit: NewBaseUnit(
			"angle",
			"rev",
			"Revolution",
			2.0*math.Pi, // 1 rev = 2π rad
			0.0,
			false,
		),
	},
	Gradian: AngleUnit{
		BaseUnit: NewBaseUnit(
			"angle",
			"grad",
			"Gradian",
			math.Pi/200.0, // 1 grad = π/200 rad
			0.0,
			false,
		),
	},
}

// NewAngle creates a new angle measurement
func NewAngle(value float64, unit AngleUnit) Quantity[AngleUnit] {
	return New(value, unit)
}
