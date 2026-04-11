// Package weather stores the current location and weather condition, and offers
// a function to format this information into a human-readable forecast.
package weather

var (
    // CurrentCondition describes the current weather condition.
	CurrentCondition string
	// CurrentLocation represents the current location.
	CurrentLocation  string
)
// Forecast return the weather forecast for a given city or codition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
