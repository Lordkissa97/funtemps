package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

// funksjon som konverterer Fahrenheit til Celsius
func FahrenheitToCelsius(fahrenheit float64) float64 {
	celsius := (fahrenheit - 32) * 5 / 9
	return celsius
}

// Funksjonen som konverterer Celsius til Fahrenheit
func CelsiusToFahrenheit(celsius float64) float64 {
	fahrenheit := celsius*9/5 + 32
	return fahrenheit
}

// Funksjonen som konverterer Kelvin til Celsius
func KelvinToCelsius(kelvin float64) float64 {
	celsius := kelvin - 273.15
	return celsius
}

// Funksjonen som konverterer Celsius til Kelvin
func CelsiusToKelvin(celsius float64) float64 {
	kelvin := celsius + 273.15
	return kelvin
}

// Funksjonen som konverterer Kelvin til Fahrenheit
func KelvinToFahrenheit(kelvin float64) float64 {
	fahrenheit := (kelvin-273.15)*9/5 + 32
	return fahrenheit
}

// Funksjonen som konverterer Fahrenheit til Kelvin
func FahrenheitToKelvin(fahrenheit float64) float64 {
	kelvin := (fahrenheit-32)*5/9 + 273.15
	return kelvin
}
