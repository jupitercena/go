package weather // This package contents weather forcasting

var CurrentCondition string // To hold the current condition
var CurrentLocation string // To hold the current location

// This function to execute the weather forcasted based on the current condition and current location
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
