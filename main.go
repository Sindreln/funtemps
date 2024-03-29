package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

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
		fmt.Printf("%.3f F er %.2f C\n", fahr, celsius)
		converted = true
	}

	if isFlagPassed("C") && out == "F" {
		fahr = conv.CelsiusToFahrenheit(celsius)
		fmt.Printf("%.2f C er %.2f F\n", celsius, fahr)
		converted = true
	}

	if isFlagPassed("C") && out == "K" {
		kelvin = conv.CelsiusToKelvin(celsius)
		fmt.Printf("%.2f C er %.2f K\n", celsius, kelvin)
		converted = true
	}

	if isFlagPassed("K") && out == "C" {
		celsius = conv.KelvinToCelsius(kelvin)
		fmt.Printf("%.2f K er %.2f C\n", kelvin, celsius)
		converted = true
	}

	if isFlagPassed("K") && out == "F" {
		fahr = conv.KelvinToFahrenheit(kelvin)
		fmt.Printf("%.2f K er %.2f F\n", kelvin, fahr)
		converted = true
	}

	if isFlagPassed("F") && out == "K" {
		kelvin = conv.FahrenheitToKelvin(fahr)
		fmt.Printf("%.3f F er %.2f K\n", fahr, kelvin)
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

func formatFloat(val float64) string {
	// Sjekk om verdien er et heltall
	if val == float64(int(val)) {
		return strconv.Itoa(int(val))
	}

	// Ellers formater med 2 desimaler
	formatted := strconv.FormatFloat(val, 'f', 2, 64)

	// Sjekk om verdien er et stort tall
	if val >= 1000 {
		// Finn posisjonen til desimaltegnet
		decimalIndex := strings.Index(formatted, ".")
		if decimalIndex == -1 {
			decimalIndex = len(formatted)
		}

		// Legg til mellomrom mellom hver tredje siffer, start fra desimaltegnet og gå bakover
		for i := decimalIndex - 3; i > 0; i -= 3 {
			formatted = formatted[:i] + " " + formatted[i:]
		}
	}

	return formatted
}
