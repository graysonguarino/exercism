// Package weather: A package for weather forecasting.
package weather

// CurrentCondition: The current weather condition.
var CurrentCondition string

// CurrentLocation: The current location whose weather we're checking.
var CurrentLocation string

// Forecast: Returns the forecast for the given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
