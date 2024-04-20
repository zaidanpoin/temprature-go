package main

import (
	"math"
	"testing"
)

func ToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
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

// to Reamur
