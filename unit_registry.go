// Package unit provides a system for representing, converting, and operating on
// physical quantities with units.
package unit

// UnitRegistry provides lookup functionality for units by symbol
// Each unit type has its own registry map

var temperatureUnitsBySymbol = map[string]TemperatureUnit{
	"°C": Temperature.Celsius,
	"C":  Temperature.Celsius,
	"°F": Temperature.Fahrenheit,
	"F":  Temperature.Fahrenheit,
	"K":  Temperature.Kelvin,
}

var pressureUnitsBySymbol = map[string]PressureUnit{
	"Pa":    Pressure.Pascal,
	"kPa":   Pressure.Kilopascal,
	"bar":   Pressure.Bar,
	"psi":   Pressure.PSI,
	"inH₂O": Pressure.InchH2O,
	"inH2O": Pressure.InchH2O,
}

var flowRateUnitsBySymbol = map[string]FlowRateUnit{
	"m³/h": FlowRate.CubicMetersPerHour,
	"L/s":  FlowRate.LitersPerSecond,
	"CFM":  FlowRate.CFM,
}

var powerUnitsBySymbol = map[string]PowerUnit{
	"W":     Power.Watt,
	"kW":    Power.Kilowatt,
	"BTU/h": Power.BTUPerHour,
}

var energyUnitsBySymbol = map[string]EnergyUnit{
	"J":   Energy.Joule,
	"kWh": Energy.KilowattHour,
	"BTU": Energy.BTU,
}

var lengthUnitsBySymbol = map[string]LengthUnit{
	"m":  Length.Meter,
	"km": Length.Kilometer,
	"cm": Length.Centimeter,
	"mm": Length.Millimeter,
	"µm": Length.Micrometer,
	"nm": Length.Nanometer,
	"in": Length.Inch,
	"ft": Length.Foot,
	"yd": Length.Yard,
	"mi": Length.Mile,
}

var massUnitsBySymbol = map[string]MassUnit{
	"kg":  Mass.Kilogram,
	"g":   Mass.Gram,
	"mg":  Mass.Milligram,
	"µg":  Mass.Microgram,
	"lb":  Mass.Pound,
	"oz":  Mass.Ounce,
	"st":  Mass.Stone,
	"t":   Mass.MetricTon,
	"ton": Mass.Ton,
}

var durationUnitsBySymbol = map[string]DurationUnit{
	"s":   Duration.Second,
	"min": Duration.Minute,
	"h":   Duration.Hour,
	"d":   Duration.Day,
	"ms":  Duration.Millisecond,
	"µs":  Duration.Microsecond,
	"ns":  Duration.Nanosecond,
}

var angleUnitsBySymbol = map[string]AngleUnit{
	"rad":  Angle.Radian,
	"°":    Angle.Degree,
	"′":    Angle.Arcminute,
	"″":    Angle.Arcsecond,
	"rev":  Angle.Revolution,
	"grad": Angle.Gradian,
}

var areaUnitsBySymbol = map[string]AreaUnit{
	"m²":  Area.SquareMeter,
	"km²": Area.SquareKilometer,
	"cm²": Area.SquareCentimeter,
	"mm²": Area.SquareMillimeter,
	"in²": Area.SquareInch,
	"ft²": Area.SquareFoot,
	"yd²": Area.SquareYard,
	"mi²": Area.SquareMile,
	"ac":  Area.Acre,
	"ha":  Area.Hectare,
}

var volumeUnitsBySymbol = map[string]VolumeUnit{
	"m³":    Volume.CubicMeter,
	"km³":   Volume.CubicKilometer,
	"cm³":   Volume.CubicCentimeter,
	"mm³":   Volume.CubicMillimeter,
	"L":     Volume.Liter,
	"mL":    Volume.Milliliter,
	"in³":   Volume.CubicInch,
	"ft³":   Volume.CubicFoot,
	"yd³":   Volume.CubicYard,
	"gal":   Volume.Gallon,
	"qt":    Volume.Quart,
	"pt":    Volume.Pint,
	"cup":   Volume.Cup,
	"fl oz": Volume.FluidOunce,
}

var accelerationUnitsBySymbol = map[string]AccelerationUnit{
	"m/s²":  Acceleration.MetersPerSecondSquared,
	"g":     Acceleration.G,
	"ft/s²": Acceleration.FeetPerSecondSquared,
}

var concentrationUnitsBySymbol = map[string]ConcentrationUnit{
	"g/L":  Concentration.GramsPerLiter,
	"mg/L": Concentration.MilligramsPerLiter,
	"ppm":  Concentration.PartsPerMillion,
	"ppb":  Concentration.PartsPerBillion,
}

var dispersionUnitsBySymbol = map[string]DispersionUnit{
	"ppm": Dispersion.PartsPerMillion,
	"ppb": Dispersion.PartsPerBillion,
	"ppt": Dispersion.PartsPerTrillion,
	"%":   Dispersion.Percent,
}

var speedUnitsBySymbol = map[string]SpeedUnit{
	"m/s":  Speed.MetersPerSecond,
	"km/h": Speed.KilometersPerHour,
	"mph":  Speed.MilesPerHour,
	"ft/s": Speed.FeetPerSecond,
	"kn":   Speed.Knot,
}

var electricChargeUnitsBySymbol = map[string]ElectricChargeUnit{
	"C":   ElectricCharge.Coulomb,
	"mC":  ElectricCharge.Millicoulomb,
	"µC":  ElectricCharge.Microcoulomb,
	"Ah":  ElectricCharge.Ampere_Hour,
	"mAh": ElectricCharge.Milliampere_Hour,
}

var electricCurrentUnitsBySymbol = map[string]ElectricCurrentUnit{
	"A":  ElectricCurrent.Ampere,
	"mA": ElectricCurrent.Milliampere,
	"µA": ElectricCurrent.Microampere,
	"kA": ElectricCurrent.Kiloampere,
}

var electricPotentialDifferenceUnitsBySymbol = map[string]ElectricPotentialDifferenceUnit{
	"V":  ElectricPotentialDifference.Volt,
	"mV": ElectricPotentialDifference.Millivolt,
	"µV": ElectricPotentialDifference.Microvolt,
	"kV": ElectricPotentialDifference.Kilovolt,
	"MV": ElectricPotentialDifference.Megavolt,
}

var frequencyUnitsBySymbol = map[string]FrequencyUnit{
	"Hz":  Frequency.Hertz,
	"kHz": Frequency.Kilohertz,
	"MHz": Frequency.Megahertz,
	"GHz": Frequency.Gigahertz,
	"THz": Frequency.Terahertz,
	"rpm": Frequency.RPM,
}

var illuminanceUnitsBySymbol = map[string]IlluminanceUnit{
	"lx": Illuminance.Lux,
	"fc": Illuminance.FootCandle,
	"ph": Illuminance.Phot,
	"nx": Illuminance.Nox,
}

var informationUnitsBySymbol = map[string]InformationUnit{
	"bit": Information.Bit,
	"B":   Information.Byte,
	"KB":  Information.Kilobyte,
	"MB":  Information.Megabyte,
	"GB":  Information.Gigabyte,
	"TB":  Information.Terabyte,
	"PB":  Information.Petabyte,
	"KiB": Information.Kibibyte,
	"MiB": Information.Mebibyte,
	"GiB": Information.Gibibyte,
	"TiB": Information.Tebibyte,
	"PiB": Information.Pebibyte,
}

var fuelEfficiencyUnitsBySymbol = map[string]FuelEfficiencyUnit{
	"km/L":    FuelEfficiency.KilometersPerLiter,
	"mpg":     FuelEfficiency.MilesPerGallon,
	"L/100km": FuelEfficiency.LitersPer100Kilometers,
}

// LookupTemperatureUnit returns the temperature unit for the given symbol
func LookupTemperatureUnit(symbol string) (TemperatureUnit, bool) {
	u, ok := temperatureUnitsBySymbol[symbol]
	return u, ok
}

// LookupPressureUnit returns the pressure unit for the given symbol
func LookupPressureUnit(symbol string) (PressureUnit, bool) {
	u, ok := pressureUnitsBySymbol[symbol]
	return u, ok
}

// LookupFlowRateUnit returns the flow rate unit for the given symbol
func LookupFlowRateUnit(symbol string) (FlowRateUnit, bool) {
	u, ok := flowRateUnitsBySymbol[symbol]
	return u, ok
}

// LookupPowerUnit returns the power unit for the given symbol
func LookupPowerUnit(symbol string) (PowerUnit, bool) {
	u, ok := powerUnitsBySymbol[symbol]
	return u, ok
}

// LookupEnergyUnit returns the energy unit for the given symbol
func LookupEnergyUnit(symbol string) (EnergyUnit, bool) {
	u, ok := energyUnitsBySymbol[symbol]
	return u, ok
}

// LookupLengthUnit returns the length unit for the given symbol
func LookupLengthUnit(symbol string) (LengthUnit, bool) {
	u, ok := lengthUnitsBySymbol[symbol]
	return u, ok
}

// LookupMassUnit returns the mass unit for the given symbol
func LookupMassUnit(symbol string) (MassUnit, bool) {
	u, ok := massUnitsBySymbol[symbol]
	return u, ok
}

// LookupDurationUnit returns the duration unit for the given symbol
func LookupDurationUnit(symbol string) (DurationUnit, bool) {
	u, ok := durationUnitsBySymbol[symbol]
	return u, ok
}

// LookupAngleUnit returns the angle unit for the given symbol
func LookupAngleUnit(symbol string) (AngleUnit, bool) {
	u, ok := angleUnitsBySymbol[symbol]
	return u, ok
}

// LookupAreaUnit returns the area unit for the given symbol
func LookupAreaUnit(symbol string) (AreaUnit, bool) {
	u, ok := areaUnitsBySymbol[symbol]
	return u, ok
}

// LookupVolumeUnit returns the volume unit for the given symbol
func LookupVolumeUnit(symbol string) (VolumeUnit, bool) {
	u, ok := volumeUnitsBySymbol[symbol]
	return u, ok
}

// LookupAccelerationUnit returns the acceleration unit for the given symbol
func LookupAccelerationUnit(symbol string) (AccelerationUnit, bool) {
	u, ok := accelerationUnitsBySymbol[symbol]
	return u, ok
}

// LookupConcentrationUnit returns the concentration unit for the given symbol
func LookupConcentrationUnit(symbol string) (ConcentrationUnit, bool) {
	u, ok := concentrationUnitsBySymbol[symbol]
	return u, ok
}

// LookupDispersionUnit returns the dispersion unit for the given symbol
func LookupDispersionUnit(symbol string) (DispersionUnit, bool) {
	u, ok := dispersionUnitsBySymbol[symbol]
	return u, ok
}

// LookupSpeedUnit returns the speed unit for the given symbol
func LookupSpeedUnit(symbol string) (SpeedUnit, bool) {
	u, ok := speedUnitsBySymbol[symbol]
	return u, ok
}

// LookupElectricChargeUnit returns the electric charge unit for the given symbol
func LookupElectricChargeUnit(symbol string) (ElectricChargeUnit, bool) {
	u, ok := electricChargeUnitsBySymbol[symbol]
	return u, ok
}

// LookupElectricCurrentUnit returns the electric current unit for the given symbol
func LookupElectricCurrentUnit(symbol string) (ElectricCurrentUnit, bool) {
	u, ok := electricCurrentUnitsBySymbol[symbol]
	return u, ok
}

// LookupElectricPotentialDifferenceUnit returns the electric potential difference unit for the given symbol
func LookupElectricPotentialDifferenceUnit(symbol string) (ElectricPotentialDifferenceUnit, bool) {
	u, ok := electricPotentialDifferenceUnitsBySymbol[symbol]
	return u, ok
}

// LookupFrequencyUnit returns the frequency unit for the given symbol
func LookupFrequencyUnit(symbol string) (FrequencyUnit, bool) {
	u, ok := frequencyUnitsBySymbol[symbol]
	return u, ok
}

// LookupIlluminanceUnit returns the illuminance unit for the given symbol
func LookupIlluminanceUnit(symbol string) (IlluminanceUnit, bool) {
	u, ok := illuminanceUnitsBySymbol[symbol]
	return u, ok
}

// LookupInformationUnit returns the information unit for the given symbol
func LookupInformationUnit(symbol string) (InformationUnit, bool) {
	u, ok := informationUnitsBySymbol[symbol]
	return u, ok
}

// LookupFuelEfficiencyUnit returns the fuel efficiency unit for the given symbol
func LookupFuelEfficiencyUnit(symbol string) (FuelEfficiencyUnit, bool) {
	u, ok := fuelEfficiencyUnitsBySymbol[symbol]
	return u, ok
}
