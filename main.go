package main

import (
	"flag"
	"fmt"
	"github.com/Lordkissa97/funtemps/conv"
)

var C float64
var F float64
var K float64
var out string

func init() {
	flag.Float64Var(&C, "C", 0, "temperatur i grader Celsius")
	flag.Float64Var(&F, "F", 0, "temperatur i grader Fahrenheit")
	flag.Float64Var(&K, "K", 0, "temperatur i grader Kelvin")
	flag.StringVar(&out, "out", "C", "temperaturskala for resultat")
}

func main() {
	flag.Parse()
	var res float64
	if C != 0 {
		if out == "F" {
			res = conv.CelsiusToFahrenheit(C)
		} else if out == "K" {
			res = C + 273.15
		}
	}
	fmt.Printf("%.2fn\n", res) // output
}
