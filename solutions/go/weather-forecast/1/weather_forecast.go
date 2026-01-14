// Package weather: This package is about weather forecasting.
package weather

var (
    // CurrentCondition: The current condition.
	CurrentCondition string

    // CurrentLocation: The current location.
	CurrentLocation  string
)

// Forecast: Generating the conditions for a specific location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
