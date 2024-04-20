package main
import (
	"math"
	"testing"
)
func ToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

func ToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}
func ToReamur(c float64) float64 {
	return c * 4 / 5
}

func TestToFahrenheit(t *testing.T) {
	celsius := 100.0
	expected := 212.0
	got := ToFahrenheit(celsius)
	if math.Abs(got-expected) > 0.001 {
		t.Errorf("ToFahrenheit(%v) = %v; want %v", celsius, got, expected)
	}
}

// to Celcius


func TestToCelsius(t *testing.T) {
	fahrenheit := 212.0
	expected := 100.0
	got := ToCelsius(fahrenheit)
	if math.Abs(got-expected) > 0.001 {
		t.Errorf("ToCelsius(%v) = %v; want %v", fahrenheit, got, expected)
	}
}


// to Reamur
func TestToReamur(t *testing.T) {
	celsius := 100.0
	expected := 80.0
	got := ToReamur(celsius)
	if math.Abs(got-expected) > 0.001 {
		t.Errorf("ToReamur(%v) = %v; want %v", celsius, got, expected)
	}
}