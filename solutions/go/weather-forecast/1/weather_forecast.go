// Package weather provides weather implementation for this exercise with test cases.
package weather

var (
    // CurrentCondition bitch bal bla.
	CurrentCondition string
	// CurrentLocation bitch2 bla bla bla.
	CurrentLocation  string
)
// Forecast returns the shit.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
