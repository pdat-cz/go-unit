// Example program demonstrating the use of the measurement package
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/pdat-cz/go-unit"
)

func main() {
	// Create temperature measurements
	tempC := unit.NewTemperature(25.0, unit.Temperature.Celsius)
	tempF := unit.NewTemperature(68.0, unit.Temperature.Fahrenheit)

	fmt.Printf("Temperature 1: %s\n", tempC)
	fmt.Printf("Temperature 2: %s\n", tempF)

	// Convert between units
	tempCtoF := tempC.ConvertTo(unit.Temperature.Fahrenheit)
	tempFtoC := tempF.ConvertTo(unit.Temperature.Celsius)

	fmt.Printf("Temperature 1 in Fahrenheit: %s\n", tempCtoF)
	fmt.Printf("Temperature 2 in Celsius: %s\n", tempFtoC)

	// Arithmetic operations
	sum := tempC.Add(tempF)
	fmt.Printf("Sum of temperatures: %s\n", sum)

	doubled := tempC.MultiplyByScalar(2.0)
	fmt.Printf("Doubled temperature: %s\n", doubled)

	// Create pressure measurements
	pressurePa := unit.NewPressure(101325.0, unit.Pressure.Pascal)
	pressureBar := pressurePa.ConvertTo(unit.Pressure.Bar)
	pressurePSI := pressurePa.ConvertTo(unit.Pressure.PSI)

	fmt.Printf("Pressure in Pascal: %s\n", pressurePa)
	fmt.Printf("Pressure in Bar: %s\n", pressureBar)
	fmt.Printf("Pressure in PSI: %s\n", pressurePSI)

	// Parse measurements from strings
	tempFromStr, err := unit.ParseTemperature("22.5°C")
	if err != nil {
		log.Fatalf("Failed to parse temperature: %v", err)
	}
	fmt.Printf("Parsed temperature: %s\n", tempFromStr)

	pressureFromStr, err := unit.ParsePressure("14.7 psi")
	if err != nil {
		log.Fatalf("Failed to parse pressure: %v", err)
	}
	fmt.Printf("Parsed pressure: %s\n", pressureFromStr)

	// Create flow rate measurements
	flowRate := unit.NewFlowRate(500.0, unit.FlowRate.CubicMetersPerHour)
	flowRateLPS := flowRate.ConvertTo(unit.FlowRate.LitersPerSecond)
	flowRateCFM := flowRate.ConvertTo(unit.FlowRate.CFM)

	fmt.Printf("Flow rate in m³/h: %s\n", flowRate)
	fmt.Printf("Flow rate in L/s: %s\n", flowRateLPS)
	fmt.Printf("Flow rate in CFM: %s\n", flowRateCFM)

	// Create power measurements
	power := unit.NewPower(1000.0, unit.Power.Watt)
	powerKW := power.ConvertTo(unit.Power.Kilowatt)
	powerBTU := power.ConvertTo(unit.Power.BTUPerHour)

	fmt.Printf("Power in Watt: %s\n", power)
	fmt.Printf("Power in Kilowatt: %s\n", powerKW)
	fmt.Printf("Power in BTU/h: %s\n", powerBTU)

	// Create energy measurements
	energy := unit.NewEnergy(3600000.0, unit.Energy.Joule)
	energyKWh := energy.ConvertTo(unit.Energy.KilowattHour)
	energyBTU := energy.ConvertTo(unit.Energy.BTU)

	fmt.Printf("Energy in Joule: %s\n", energy)
	fmt.Printf("Energy in kWh: %s\n", energyKWh)
	fmt.Printf("Energy in BTU: %s\n", energyBTU)

	// Create length measurements
	length := unit.NewLength(10.0, unit.Length.Meter)
	lengthKm := length.ConvertTo(unit.Length.Kilometer)
	lengthFt := length.ConvertTo(unit.Length.Foot)
	lengthMi := length.ConvertTo(unit.Length.Mile)

	fmt.Printf("\nLength in Meter: %s\n", length)
	fmt.Printf("Length in Kilometer: %s\n", lengthKm)
	fmt.Printf("Length in Foot: %s\n", lengthFt)
	fmt.Printf("Length in Mile: %s\n", lengthMi)

	// Create mass measurements
	mass := unit.NewMass(75.0, unit.Mass.Kilogram)
	massG := mass.ConvertTo(unit.Mass.Gram)
	massLb := mass.ConvertTo(unit.Mass.Pound)
	massOz := mass.ConvertTo(unit.Mass.Ounce)

	fmt.Printf("\nMass in Kilogram: %s\n", mass)
	fmt.Printf("Mass in Gram: %s\n", massG)
	fmt.Printf("Mass in Pound: %s\n", massLb)
	fmt.Printf("Mass in Ounce: %s\n", massOz)

	// Create duration measurements
	duration := unit.NewDuration(60.0, unit.Duration.Second)
	durationMin := duration.ConvertTo(unit.Duration.Minute)
	durationHr := duration.ConvertTo(unit.Duration.Hour)
	durationMs := duration.ConvertTo(unit.Duration.Millisecond)

	fmt.Printf("\nDuration in Second: %s\n", duration)
	fmt.Printf("Duration in Minute: %s\n", durationMin)
	fmt.Printf("Duration in Hour: %s\n", durationHr)
	fmt.Printf("Duration in Millisecond: %s\n", durationMs)

	// Create angle measurements
	angle := unit.NewAngle(90.0, unit.Angle.Degree)
	angleRad := angle.ConvertTo(unit.Angle.Radian)
	angleRev := angle.ConvertTo(unit.Angle.Revolution)
	angleGrad := angle.ConvertTo(unit.Angle.Gradian)

	fmt.Printf("\nAngle in Degree: %s\n", angle)
	fmt.Printf("Angle in Radian: %s\n", angleRad)
	fmt.Printf("Angle in Revolution: %s\n", angleRev)
	fmt.Printf("Angle in Gradian: %s\n", angleGrad)

	// Create area measurements
	area := unit.NewArea(100.0, unit.Area.SquareMeter)
	areaHa := area.ConvertTo(unit.Area.Hectare)
	areaAc := area.ConvertTo(unit.Area.Acre)
	areaSqFt := area.ConvertTo(unit.Area.SquareFoot)

	fmt.Printf("\nArea in Square Meter: %s\n", area)
	fmt.Printf("Area in Hectare: %s\n", areaHa)
	fmt.Printf("Area in Acre: %s\n", areaAc)
	fmt.Printf("Area in Square Foot: %s\n", areaSqFt)

	// Create volume measurements
	volume := unit.NewVolume(1.0, unit.Volume.CubicMeter)
	volumeL := volume.ConvertTo(unit.Volume.Liter)
	volumeGal := volume.ConvertTo(unit.Volume.Gallon)
	volumeCuFt := volume.ConvertTo(unit.Volume.CubicFoot)

	fmt.Printf("\nVolume in Cubic Meter: %s\n", volume)
	fmt.Printf("Volume in Liter: %s\n", volumeL)
	fmt.Printf("Volume in Gallon: %s\n", volumeGal)
	fmt.Printf("Volume in Cubic Foot: %s\n", volumeCuFt)

	// Parse new unit types from strings
	lengthFromStr, err := unit.ParseLength("5.5 km")
	if err != nil {
		log.Fatalf("Failed to parse length: %v", err)
	}
	fmt.Printf("\nParsed length: %s\n", lengthFromStr)

	massFromStr, err := unit.ParseMass("150 lb")
	if err != nil {
		log.Fatalf("Failed to parse mass: %v", err)
	}
	fmt.Printf("Parsed mass: %s\n", massFromStr)

	durationFromStr, err := unit.ParseDuration("2.5 h")
	if err != nil {
		log.Fatalf("Failed to parse duration: %v", err)
	}
	fmt.Printf("Parsed duration: %s\n", durationFromStr)

	angleFromStr, err := unit.ParseAngle("45°")
	if err != nil {
		log.Fatalf("Failed to parse angle: %v", err)
	}
	fmt.Printf("Parsed angle: %s\n", angleFromStr)

	areaFromStr, err := unit.ParseArea("2 ha")
	if err != nil {
		log.Fatalf("Failed to parse area: %v", err)
	}
	fmt.Printf("Parsed area: %s\n", areaFromStr)

	volumeFromStr, err := unit.ParseVolume("500 mL")
	if err != nil {
		log.Fatalf("Failed to parse volume: %v", err)
	}
	fmt.Printf("Parsed volume: %s\n", volumeFromStr)

	// Serialization and deserialization examples
	fmt.Println("\n--- Serialization and Deserialization ---")

	// Temperature serialization
	tempJSON, err := unit.MarshalTemperature(tempC)
	if err != nil {
		log.Fatalf("Failed to marshal temperature: %v", err)
	}
	fmt.Printf("Temperature JSON: %s\n", tempJSON)

	// Temperature deserialization
	tempDeserialized, err := unit.UnmarshalTemperature(tempJSON)
	if err != nil {
		log.Fatalf("Failed to unmarshal temperature: %v", err)
	}
	fmt.Printf("Deserialized temperature: %s\n", tempDeserialized)

	// Length serialization/deserialization
	lengthJSON, _ := unit.MarshalLength(length)
	lengthDeserialized, _ := unit.UnmarshalLength(lengthJSON)
	fmt.Printf("Length serialization/deserialization: %s -> %s\n", length, lengthDeserialized)

	// Mass serialization/deserialization
	massJSON, _ := unit.MarshalMass(mass)
	massDeserialized, _ := unit.UnmarshalMass(massJSON)
	fmt.Printf("Mass serialization/deserialization: %s -> %s\n", mass, massDeserialized)

	// Duration serialization/deserialization
	durationJSON, _ := unit.MarshalDuration(duration)
	durationDeserialized, _ := unit.UnmarshalDuration(durationJSON)
	fmt.Printf("Duration serialization/deserialization: %s -> %s\n", duration, durationDeserialized)

	// Angle serialization/deserialization
	angleJSON, _ := unit.MarshalAngle(angle)
	angleDeserialized, _ := unit.UnmarshalAngle(angleJSON)
	fmt.Printf("Angle serialization/deserialization: %s -> %s\n", angle, angleDeserialized)

	// Area serialization/deserialization
	areaJSON, _ := unit.MarshalArea(area)
	areaDeserialized, _ := unit.UnmarshalArea(areaJSON)
	fmt.Printf("Area serialization/deserialization: %s -> %s\n", area, areaDeserialized)

	// Volume serialization/deserialization
	volumeJSON, _ := unit.MarshalVolume(volume)
	volumeDeserialized, _ := unit.UnmarshalVolume(volumeJSON)
	fmt.Printf("Volume serialization/deserialization: %s -> %s\n", volume, volumeDeserialized)

	// Example of using serialized data in a struct
	type SensorReading struct {
		ID          string          `json:"id"`
		Timestamp   string          `json:"timestamp"`
		Temperature json.RawMessage `json:"temperature,omitempty"`
		Pressure    json.RawMessage `json:"pressure,omitempty"`
		Length      json.RawMessage `json:"length,omitempty"`
		Volume      json.RawMessage `json:"volume,omitempty"`
	}

	// Create a sensor reading with multiple measurements
	combinedReading := SensorReading{
		ID:          "sensor3",
		Timestamp:   "2023-06-15T12:36:00Z",
		Temperature: tempJSON,
		Length:      lengthJSON,
		Volume:      volumeJSON,
	}

	// Serialize the sensor reading
	readingJSON, err := json.MarshalIndent(combinedReading, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal sensor reading: %v", err)
	}
	fmt.Printf("\nSensor reading with multiple measurements:\n%s\n", readingJSON)

	// Create acceleration measurements
	accel := unit.NewAcceleration(9.8, unit.Acceleration.MetersPerSecondSquared)
	accelG := accel.ConvertTo(unit.Acceleration.G)
	accelFt := accel.ConvertTo(unit.Acceleration.FeetPerSecondSquared)

	fmt.Printf("\n--- Acceleration Examples ---\n")
	fmt.Printf("Acceleration in m/s²: %s\n", accel)
	fmt.Printf("Acceleration in g: %s\n", accelG)
	fmt.Printf("Acceleration in ft/s²: %s\n", accelFt)

	// Parse acceleration from string
	accelFromStr, err := unit.ParseAcceleration("32.2 ft/s²")
	if err != nil {
		log.Fatalf("Failed to parse acceleration: %v", err)
	}
	fmt.Printf("Parsed acceleration: %s\n", accelFromStr)

	// Serialize/deserialize acceleration
	accelJSON, _ := unit.MarshalAcceleration(accel)
	accelDeserialized, _ := unit.UnmarshalAcceleration(accelJSON)
	fmt.Printf("Acceleration serialization/deserialization: %s -> %s\n", accel, accelDeserialized)

	// Create concentration measurements
	conc := unit.NewConcentration(5.0, unit.Concentration.GramsPerLiter)
	concMg := conc.ConvertTo(unit.Concentration.MilligramsPerLiter)
	concPpm := conc.ConvertTo(unit.Concentration.PartsPerMillion)

	fmt.Printf("\n--- Concentration Examples ---\n")
	fmt.Printf("Concentration in g/L: %s\n", conc)
	fmt.Printf("Concentration in mg/L: %s\n", concMg)
	fmt.Printf("Concentration in ppm: %s\n", concPpm)

	// Parse concentration from string
	concFromStr, err := unit.ParseConcentration("500 mg/L")
	if err != nil {
		log.Fatalf("Failed to parse concentration: %v", err)
	}
	fmt.Printf("Parsed concentration: %s\n", concFromStr)

	// Serialize/deserialize concentration
	concJSON, _ := unit.MarshalConcentration(conc)
	concDeserialized, _ := unit.UnmarshalConcentration(concJSON)
	fmt.Printf("Concentration serialization/deserialization: %s -> %s\n", conc, concDeserialized)

	// Create dispersion measurements
	disp := unit.NewDispersion(100.0, unit.Dispersion.PartsPerMillion)
	dispPpb := disp.ConvertTo(unit.Dispersion.PartsPerBillion)
	dispPercent := disp.ConvertTo(unit.Dispersion.Percent)

	fmt.Printf("\n--- Dispersion Examples ---\n")
	fmt.Printf("Dispersion in ppm: %s\n", disp)
	fmt.Printf("Dispersion in ppb: %s\n", dispPpb)
	fmt.Printf("Dispersion in percent: %s\n", dispPercent)

	// Parse dispersion from string
	dispFromStr, err := unit.ParseDispersion("0.01 %")
	if err != nil {
		log.Fatalf("Failed to parse dispersion: %v", err)
	}
	fmt.Printf("Parsed dispersion: %s\n", dispFromStr)

	// Serialize/deserialize dispersion
	dispJSON, _ := unit.MarshalDispersion(disp)
	dispDeserialized, _ := unit.UnmarshalDispersion(dispJSON)
	fmt.Printf("Dispersion serialization/deserialization: %s -> %s\n", disp, dispDeserialized)

	// Create electric charge measurements
	charge := unit.NewElectricCharge(1000.0, unit.ElectricCharge.Coulomb)
	chargeAh := charge.ConvertTo(unit.ElectricCharge.Ampere_Hour)
	chargeMah := charge.ConvertTo(unit.ElectricCharge.Milliampere_Hour)

	fmt.Printf("\n--- Electric Charge Examples ---\n")
	fmt.Printf("Electric charge in C: %s\n", charge)
	fmt.Printf("Electric charge in Ah: %s\n", chargeAh)
	fmt.Printf("Electric charge in mAh: %s\n", chargeMah)

	// Parse electric charge from string
	chargeFromStr, err := unit.ParseElectricCharge("2500 mAh")
	if err != nil {
		log.Fatalf("Failed to parse electric charge: %v", err)
	}
	fmt.Printf("Parsed electric charge: %s\n", chargeFromStr)

	// Serialize/deserialize electric charge
	chargeJSON, _ := unit.MarshalElectricCharge(charge)
	chargeDeserialized, _ := unit.UnmarshalElectricCharge(chargeJSON)
	fmt.Printf("Electric charge serialization/deserialization: %s -> %s\n", charge, chargeDeserialized)

	// Create electric current measurements
	current := unit.NewElectricCurrent(2.5, unit.ElectricCurrent.Ampere)
	currentMa := current.ConvertTo(unit.ElectricCurrent.Milliampere)
	currentKa := current.ConvertTo(unit.ElectricCurrent.Kiloampere)

	fmt.Printf("\n--- Electric Current Examples ---\n")
	fmt.Printf("Electric current in A: %s\n", current)
	fmt.Printf("Electric current in mA: %s\n", currentMa)
	fmt.Printf("Electric current in kA: %s\n", currentKa)

	// Parse electric current from string
	currentFromStr, err := unit.ParseElectricCurrent("500 mA")
	if err != nil {
		log.Fatalf("Failed to parse electric current: %v", err)
	}
	fmt.Printf("Parsed electric current: %s\n", currentFromStr)

	// Serialize/deserialize electric current
	currentJSON, _ := unit.MarshalElectricCurrent(current)
	currentDeserialized, _ := unit.UnmarshalElectricCurrent(currentJSON)
	fmt.Printf("Electric current serialization/deserialization: %s -> %s\n", current, currentDeserialized)
}
