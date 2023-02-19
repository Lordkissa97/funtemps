package conv

import (
	"reflect"
	"testing"
)

// funksjon som konverterer Fahrenheit til Celsius
func FahrenheitToCelsius(fahrenheit float64) float64 {
	celsius := (fahrenheit - 32) * 5 / 9
	return celsius
}

func TestFahrenheitToCelsius(t *testing.T) {
	type test struct {
		Fahrenheit float64
		Celsius    float64
	}

	var tests = []test{
		{Fahrenheit: 134, Celsius: 56.67},
		{Fahrenheit: 32, Celsius: 0},
	}
	for _, v := range tests {
		got := FahrenheitToCelsius(v.Fahrenheit)
		if !reflect.DeepEqual(got, v.Celsius) {
			t.Errorf("FahrenheitToCelsius(%v) = %v, want %v", v.Fahrenheit, got, v.Celsius)
		}
	}
}

// Funksjonen som konverterer Celsius til Fahrenheit
func CelsiusToFahrenheit(celsius float64) float64 {
	fahrenheit := celsius*9/5 + 32
	return fahrenheit
}

func TestCelsiusToFahrenheit(t *testing.T) {
	type test struct {
		Celsius    float64
		Fahrenheit float64
	}

	var tests = []test{
		{Celsius: 56.67, Fahrenheit: 134},
		{Celsius: 0, Fahrenheit: 32},
	}
	for _, v := range tests {
		got := CelsiusToFahrenheit(v.Celsius)
		if !reflect.DeepEqual(got, v.Fahrenheit) {
			t.Errorf("CelsiusToFahrenheit(%v) = %v, want %v", v.Celsius, got, v.Fahrenheit)
		}
	}
}

// Funksjonen som konverterer Kelvin til Celsius
func KelvinToCelsius(kelvin float64) float64 {
	celsius := kelvin - 273.15
	return celsius
}

func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		Kelvin  float64
		Celsius float64
	}

	var tests = []test{
		{Kelvin: 329.82, Celsius: 56.67},
		{Kelvin: 273.15, Celsius: 0},
	}
	for _, v := range tests {
		got := KelvinToCelsius(v.Kelvin)
		if !reflect.DeepEqual(got, v.Celsius) {
			t.Errorf("KelvinToCelsius(%v) = %v, want %v", v.Kelvin, got, v.Celsius)
		}
	}
}

// Funksjonen som konverterer Celsius til Kelvin
func CelsiusToKelvin(celsius float64) float64 {
	kelvin := celsius + 273.15
	return kelvin
}

func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		Celsius float64
		Kelvin  float64
	}

	var tests = []test{
		{Celsius: 56.67, Kelvin: 329.82},
		{Celsius: 0, Kelvin: 273.15},
	}
	for _, v := range tests {
		got := CelsiusToKelvin(v.Celsius)
		if !reflect.DeepEqual(got, v.Kelvin) {
			t.Errorf("CelsiusToKelvin(%v) = %v, want %v", v.Celsius, got, v.Kelvin)
		}
	}
}

// Funksjonen som konverterer Kelvin til Fahrenheit
func KelvinToFahrenheit(kelvin float64) float64 {
	fahrenheit := (kelvin-273.15)*9/5 + 32
	return fahrenheit
}

func TestKelvinToFahrenheit(t *testing.T) {
	type test struct {
		Kelvin     float64
		Fahrenheit float64
	}

	var tests = []test{
		{Kelvin: 329.82, Fahrenheit: 134},
		{Kelvin: 273.15, Fahrenheit: 32},
	}
	for _, v := range tests {
		got := KelvinToFahrenheit(v.Kelvin)
		if !reflect.DeepEqual(got, v.Fahrenheit) {
			t.Errorf("KelvinToFahrenheit(%v) = %v, want %v", v.Kelvin, got, v.Fahrenheit)
		}
	}
}

// Funksjonen som konverterer Fahrenheit til Kelvin
func FahrenheitToKelvin(fahrenheit float64) float64 {
	kelvin := (fahrenheit-32)*5/9 + 273.15
	return kelvin
}

func TestFahrenheitToKelvin(t *testing.T) {
	type test struct {
		Fahrenheit float64
		Kelvin     float64
	}

	var tests = []test{
		{Fahrenheit: 134, Kelvin: 329.82},
		{Fahrenheit: 32, Kelvin: 273.15},
	}
	for _, v := range tests {
		got := FahrenheitToKelvin(v.Fahrenheit)
		if !reflect.DeepEqual(got, v.Kelvin) {
			t.Errorf("FahrenheitToKelvin(%v) = %v, want %v", v.Fahrenheit, got, v.Kelvin)
		}
	}
}
