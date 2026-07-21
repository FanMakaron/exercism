// Package weather славный пакет.
package weather

var (
	// CurrentCondition условия.
	CurrentCondition string
	// CurrentLocation локация.
	CurrentLocation string
)

// Forecast Сам прогноз.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
