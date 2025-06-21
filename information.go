// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

// InformationUnit represents a unit of information
type InformationUnit struct {
	BaseUnit
}

// Information contains predefined information units
var Information = struct {
	Bit      InformationUnit
	Byte     InformationUnit
	Kilobyte InformationUnit
	Megabyte InformationUnit
	Gigabyte InformationUnit
	Terabyte InformationUnit
	Petabyte InformationUnit
	Kibibyte InformationUnit
	Mebibyte InformationUnit
	Gibibyte InformationUnit
	Tebibyte InformationUnit
	Pebibyte InformationUnit
}{
	Bit: InformationUnit{
		BaseUnit: NewBaseUnit(
			"information",
			"bit",
			"Bit",
			0.125, // 1 bit = 0.125 bytes (base unit)
			0.0,
			false,
		),
	},
	Byte: InformationUnit{
		BaseUnit: NewBaseUnit(
			"information",
			"B",
			"Byte",
			1.0, // 1 B = 1 byte
			0.0,
			true, // Base unit
		),
	},
	Kilobyte: InformationUnit{
		BaseUnit: NewBaseUnit(
			"information",
			"KB",
			"Kilobyte",
			1000.0, // 1 KB = 1000 bytes
			0.0,
			false,
		),
	},
	Megabyte: InformationUnit{
		BaseUnit: NewBaseUnit(
			"information",
			"MB",
			"Megabyte",
			1000000.0, // 1 MB = 1,000,000 bytes
			0.0,
			false,
		),
	},
	Gigabyte: InformationUnit{
		BaseUnit: NewBaseUnit(
			"information",
			"GB",
			"Gigabyte",
			1000000000.0, // 1 GB = 1,000,000,000 bytes
			0.0,
			false,
		),
	},
	Terabyte: InformationUnit{
		BaseUnit: NewBaseUnit(
			"information",
			"TB",
			"Terabyte",
			1000000000000.0, // 1 TB = 1,000,000,000,000 bytes
			0.0,
			false,
		),
	},
	Petabyte: InformationUnit{
		BaseUnit: NewBaseUnit(
			"information",
			"PB",
			"Petabyte",
			1000000000000000.0, // 1 PB = 1,000,000,000,000,000 bytes
			0.0,
			false,
		),
	},
	Kibibyte: InformationUnit{
		BaseUnit: NewBaseUnit(
			"information",
			"KiB",
			"Kibibyte",
			1024.0, // 1 KiB = 1,024 bytes
			0.0,
			false,
		),
	},
	Mebibyte: InformationUnit{
		BaseUnit: NewBaseUnit(
			"information",
			"MiB",
			"Mebibyte",
			1048576.0, // 1 MiB = 1,048,576 bytes
			0.0,
			false,
		),
	},
	Gibibyte: InformationUnit{
		BaseUnit: NewBaseUnit(
			"information",
			"GiB",
			"Gibibyte",
			1073741824.0, // 1 GiB = 1,073,741,824 bytes
			0.0,
			false,
		),
	},
	Tebibyte: InformationUnit{
		BaseUnit: NewBaseUnit(
			"information",
			"TiB",
			"Tebibyte",
			1099511627776.0, // 1 TiB = 1,099,511,627,776 bytes
			0.0,
			false,
		),
	},
	Pebibyte: InformationUnit{
		BaseUnit: NewBaseUnit(
			"information",
			"PiB",
			"Pebibyte",
			1125899906842624.0, // 1 PiB = 1,125,899,906,842,624 bytes
			0.0,
			false,
		),
	},
}

// NewInformation creates a new information measurement
func NewInformation(value float64, unit InformationUnit) Quantity[InformationUnit] {
	return New(value, unit)
}
