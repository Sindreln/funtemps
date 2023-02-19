package main

import (
	"flag"
	"fmt"

	"github.com/Sindreln/funtemps/conv"
	"github.com/Sindreln/funtemps/funfacts"
)

var fahr float64
var out string
var funfact string
var t string
var celsius float64
var kelvin float64

func init() {
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfact, "funfacts", "", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	flag.StringVar(&t, "t", "", "temperaturskala når funfacts skal vises")
	flag.Float64Var(&kelvin, "K", 0.0, "temperatur i grader Kelvin")
	flag.Float64Var(&celsius, "C", 0.0, "temperatur i grader Celsius")
}

func main() {
	flag.Parse()

	converted := false

	if isFlagPassed("F") && out == "C" {
		celsius = conv.FahrenheitToCelsius(fahr)
		fmt.Println(fahr, "F er", celsius, "C")
		converted = true
	}

	if isFlagPassed("C") && out == "F" {
		fahr = conv.CelsiusToFahrenheit(celsius)
		fmt.Println(celsius, "C er", fahr, "F")
		converted = true
	}

	if isFlagPassed("C") && out == "K" {
		kelvin = conv.CelsiusToKelvin(celsius)
		fmt.Println(celsius, "C er", kelvin, "K")
		converted = true
	}

	if isFlagPassed("K") && out == "C" {
		celsius = conv.KelvinToCelsius(kelvin)
		fmt.Println(kelvin, "K er", celsius, "C")
		converted = true
	}

	if isFlagPassed("K") && out == "F" {
		fahr = conv.KelvinToFahrenheit(kelvin)
		fmt.Println(kelvin, "K er", fahr, "F")
		converted = true
	}

	if isFlagPassed("F") && out == "K" {
		kelvin = conv.FahrenheitToKelvin(fahr)
		fmt.Println(fahr, "F er", kelvin, "K")
		converted = true
	}

	if !converted && isFlagPassed("funfacts") && isFlagPassed("t") && t == "C" {
		facts := funfacts.GetFunFacts(funfact)
		for i, fact := range facts {
			fmt.Printf("%d: %s\n", i+1, fact)
		}
	}
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
