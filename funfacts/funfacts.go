package funfacts

/*
*

	  Implementer funfacts-funksjon:
	    GetFunFacts(about string) []string
	      hvor about kan ha en av tre testverdier, -
	        sun, luna eller terra
	  Sett inn alle Funfucts i en struktur
	  type FunFacts struct {
	      Sun []string
	      Luna []string
	      Terra []string
		  }
*/
package main

import "fmt"

func main() {
	fmt.Println("Temperature fun facts:")
	fmt.Printf("Highest temperature measured on Terra's surface: %.2f °C / %.2f °F / %.2f K\n", fahrenheitToCelsius(134), 134.0, fahrenheitToKelvin(134))
	fmt.Printf("Lowest temperature measured on Terra's surface: %.2f °C / %.2f °F / %.2f K\n", fahrenheitToCelsius(-89.4), -89.4, fahrenheitToKelvin(-89.4))
	fmt.Printf("Temperature in Terra's inner core: %.2f °C / %.2f °F / %.2f K\n", kelvinToCelsius(9392), kelvinToFahrenheit(9392), 9392.0)
	fmt.Printf("Temperature in the Sun's core: %.2f °C / %.2f °F / %.2f K\n", kelvinToCelsius(15000000), kelvinToFahrenheit(15000000), 15000000.0)
	fmt.Printf("Temperature on the outer layer of the Sun: %.2f °C / %.2f °F / %.2f K\n", kelvinToCelsius(5778), kelvinToFahrenheit(5778), 5778.0)
	fmt.Printf("Temperature on the Luna's surface at night: %.2f °C / %.2f °F / %.2f K\n", fahrenheitToCelsius(-183), -183.0, fahrenheitToKelvin(-183))
	fmt.Printf("Temperature on the Luna's surface during the day: %.2f °C / %.2f °F / %.2f K\n", fahrenheitToCelsius(106), 106.0, fahrenheitToKelvin(106))
}

// helper functions for temperature conversions
func celsiusToFahrenheit(c float64) float64 {
	return c*1.8 + 32
}

func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) / 1.8
}

func celsiusToKelvin(c float64) float64 {
	return c + 273.15
}

func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}

func fahrenheitToKelvin(f float64) float64 {
	c := fahrenheitToCelsius(f)
	return celsiusToKelvin(c)
}

func kelvinToFahrenheit(k float64) float64 {
	c := kelvinToCelsius(k)
	return celsiusToFahrenheit(c)
}
