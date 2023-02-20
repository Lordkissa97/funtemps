package main

import (
	"Lordkissa97/funtemps/conv/conv"
	"flag"
	"fmt"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var cels float64
var kelv float64
var out string
var funfacts string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&cels, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelv, "K", 0.0, "temperatur i grader kelvin")

	// Du må selv definere flag-variablene for "C" og "K"
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises

}

func main() {

	flag.Parse()

	// Her er noen eksempler du kan bruke i den manuelle testingen
	fmt.Println(fahr, out, funfacts)

	fmt.Println("len(flag.Args())", len(flag.Args()))
	fmt.Println("flag.NFlag()", flag.NFlag())

	fmt.Println(isFlagPassed("out"))

	// Eksempel på enkel logikk
	if out == "C" && isFlagPassed("F") {
		fmt.Println(conv.FahrenheitToCelsius(fahr))
	}

	if out == "F" && isFlagPassed("C") {
		fmt.Println(conv.CelsiusToFahrenheit(cels))
	}

	if out == "K" && isFlagPassed("C") {
		fmt.Println(conv.CelsiusToKelvin(cels))
	}

	if out == "C" && isFlagPassed("K") {
		fmt.Println(conv.KelvinToCelsius(kelv))
	}

	if out == "F" && isFlagPassed("K") {
		fmt.Println(conv.KelvinToFahrenheit(kelv))
	}

	if out == "K" && isFlagPassed("F") {
		fmt.Println(conv.FahrenheitToKelvin(fahr))
	}

}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
