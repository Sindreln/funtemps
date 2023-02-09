package conv

var Celsius float64
var Kelvin float64
var Fahrenheit float64

// Konverterer Farhenheit til Celsius
func FahrenheitToCelsius(value float64) float64 {
	Celsius = (Fahrenheit - 32) * (1.8)

	return (value - 32) * (1.8)
}

func CelsiusToFahrenheit(value float64) float64 {
	Fahrenheit = Celsius*(1.8) + 32

	return value*(1.8) + 32
}

func CelsiusToKelvin(value float64) float64 {
	Kelvin = Celsius + 273.15

	return value + 273.15
}
func KelvinToCelsius(value float64) float64 {
	Celsius = Kelvin - 273.15

	return value - 273.15
}
func KelvinToFahrenheit(value float64) float64 {
	Fahrenheit = (Kelvin-273.15)*(1.8) + 32

	return (value-273.15)*(1.8) + 32
}
func FahrenheitToKelvin(value float64) float64 {
	Kelvin = (Fahrenheit-32)*(1.8) + 273.15

	return (value-32)*(1.8) + 273.15
}
