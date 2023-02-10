package main

import (
	"flag"
	"fmt"
)

var fahr float64
var out string
var funfacts string
var t string

func init() {
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	flag.StringVar(&t, "t", "", "temperaturskala når funfacts skal vises")
}

func main() {
	flag.Parse()

	if isFlagPassed("F") && (isFlagPassed("out") || isFlagPassed("funfacts")) {
		if out == "C" {
			celsius := (fahr - 32) * 5 / 9
			fmt.Printf("%.2f°F er %.2f°C\n", fahr, celsius)
		} else if out == "K" {
			kelvin := (fahr + 459.67) * 5 / 9
			fmt.Printf("%.2f°F er %.2f K\n", fahr, kelvin)
		} else if out == "F" {
			celsius := (fahr - 32) * 5 / 9
			fahrenheit := (celsius * 9 / 5) + 32
			fmt.Printf("%.2f°C er %.2f°F\n", celsius, fahrenheit)
		}
	} else if isFlagPassed("funfacts") && isFlagPassed("t") && t == "C" {
		switch funfacts {
		case "sun":
			fmt.Println("The Sun is over 4.5 billion years old.",
				"The Sun's light reaches the Earth in eight minutes.",
				"The Sun is made from Hydrogen and Helium.")
		case "luna":
			fmt.Println("The Moon always shows Earth the same face.",
				"The Moon's surface is actually dark.",
				"The Moon has quakes too.")
		case "terra":
			fmt.Println("Earth has never been perfectly round",
				"Earth's gravity isn't uniform.",
				"In the past, sea levels were very different.")
		default:
			fmt.Println("Ukjent argument for funfacts. Bruk sun, luna eller terra.")
		}
	} else {
		fmt.Println("Feil i argumentene. Bruk -h for hjelp.")
		flag.PrintDefaults()
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
