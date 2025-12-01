package unit

import (
	"encoding/json"
	"testing"
)

func TestInformationConversion(t *testing.T) {
	// Test conversion from bytes to kilobytes
	bytes := NewInformation(1000, Information.Byte)
	kilobytes := bytes.ConvertTo(Information.Kilobyte)
	if kilobytes.Value != 1.0 {
		t.Errorf("Expected 1000 bytes to be 1.0 kilobytes, got %f", kilobytes.Value)
	}

	// Test conversion from kilobytes to bytes
	kb := NewInformation(1.0, Information.Kilobyte)
	b := kb.ConvertTo(Information.Byte)
	if b.Value != 1000.0 {
		t.Errorf("Expected 1.0 kilobytes to be 1000.0 bytes, got %f", b.Value)
	}

	// Test conversion from megabytes to kilobytes
	mb := NewInformation(1.0, Information.Megabyte)
	kb = mb.ConvertTo(Information.Kilobyte)
	if kb.Value != 1000.0 {
		t.Errorf("Expected 1.0 megabytes to be 1000.0 kilobytes, got %f", kb.Value)
	}

	// Test conversion from kibibytes to bytes
	kib := NewInformation(1.0, Information.Kibibyte)
	b = kib.ConvertTo(Information.Byte)
	if b.Value != 1024.0 {
		t.Errorf("Expected 1.0 kibibytes to be 1024.0 bytes, got %f", b.Value)
	}

	// Test conversion from mebibytes to kibibytes
	mib := NewInformation(1.0, Information.Mebibyte)
	kib = mib.ConvertTo(Information.Kibibyte)
	if kib.Value != 1024.0 {
		t.Errorf("Expected 1.0 mebibytes to be 1024.0 kibibytes, got %f", kib.Value)
	}

	// Test conversion from bits to bytes
	bits := NewInformation(8.0, Information.Bit)
	b = bits.ConvertTo(Information.Byte)
	if b.Value != 1.0 {
		t.Errorf("Expected 8.0 bits to be 1.0 bytes, got %f", b.Value)
	}

	// Test conversion from bytes to bits
	b = NewInformation(1.0, Information.Byte)
	bits = b.ConvertTo(Information.Bit)
	if bits.Value != 8.0 {
		t.Errorf("Expected 1.0 bytes to be 8.0 bits, got %f", bits.Value)
	}
}

func TestInformationArithmetic(t *testing.T) {
	// Test addition
	a := NewInformation(1024.0, Information.Byte)
	b := NewInformation(1.0, Information.Kibibyte)
	sum := a.Add(b)
	if sum.Value != 2048.0 || sum.Unit != Information.Byte {
		t.Errorf("Expected 1024.0 bytes + 1.0 kibibytes to be 2048.0 bytes, got %f %s", sum.Value, sum.Unit.Symbol())
	}

	// Test subtraction
	diff := b.Subtract(a)
	if diff.Value != 0.0 || diff.Unit != Information.Kibibyte {
		t.Errorf("Expected 1.0 kibibytes - 1024.0 bytes to be 0.0 kibibytes, got %f %s", diff.Value, diff.Unit.Symbol())
	}

	// Test multiplication by scalar
	doubled := a.MultiplyByScalar(2.0)
	if doubled.Value != 2048.0 || doubled.Unit != Information.Byte {
		t.Errorf("Expected 1024.0 bytes * 2.0 to be 2048.0 bytes, got %f %s", doubled.Value, doubled.Unit.Symbol())
	}

	// Test division by scalar
	halved := a.DivideByScalar(2.0)
	if halved.Value != 512.0 || halved.Unit != Information.Byte {
		t.Errorf("Expected 1024.0 bytes / 2.0 to be 512.0 bytes, got %f %s", halved.Value, halved.Unit.Symbol())
	}
}

func TestInformationSerialization(t *testing.T) {
	// Test serialization
	info := NewInformation(1024.0, Information.Kibibyte)
	data, err := MarshalInformation(info)
	if err != nil {
		t.Errorf("Failed to marshal information: %v", err)
	}

	// Verify JSON structure
	var jsonData map[string]interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}

	if jsonData["value"].(float64) != 1024.0 {
		t.Errorf("Expected value to be 1024.0, got %f", jsonData["value"].(float64))
	}
	unitData := jsonData["unit"].(map[string]interface{})
	if unitData["symbol"].(string) != "KiB" {
		t.Errorf("Expected unit symbol to be KiB, got %s", unitData["symbol"].(string))
	}
	if unitData["name"].(string) != "Kibibyte" {
		t.Errorf("Expected unit name to be Kibibyte, got %s", unitData["name"].(string))
	}
	if jsonData["dimension"].(string) != "information" {
		t.Errorf("Expected dimension to be information, got %s", jsonData["dimension"].(string))
	}

	// Test deserialization
	infoDeserialized, err := UnmarshalInformation(data)
	if err != nil {
		t.Errorf("Failed to unmarshal information: %v", err)
	}

	if infoDeserialized.Value != info.Value {
		t.Errorf("Expected deserialized value to be %f, got %f", info.Value, infoDeserialized.Value)
	}
	if !infoDeserialized.Unit.Equals(info.Unit) {
		t.Errorf("Expected deserialized unit to be %s, got %s", info.Unit.Symbol(), infoDeserialized.Unit.Symbol())
	}
}
