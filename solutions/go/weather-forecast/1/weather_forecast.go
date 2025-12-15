//Package weather attempts to forcast the impossible via the Forecast function.
package weather

var (
	// CurrentCondition describes the current conditions.
    CurrentCondition string
    // CurrentLocation describes the current location.
	CurrentLocation  string
)

//Forecast only returns what you put into it.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
