// Package weather is get weather for Location.
package weather

var (
    // CurrentCondition is get current status.
	CurrentCondition string

    // CurrentLocation is get current place.
	CurrentLocation  string
)

// Forecast returns a sting value.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
