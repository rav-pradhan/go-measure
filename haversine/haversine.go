package haversine

import (
	"math"
)

const (
	RadiusOfEarthInKilometres float64 = 6371
	RadiusOfEarthInMiles float64 = 3958.8
)

type DecimalCoordinates struct {
	Latitude  float64
	Longitude float64
}

type Result struct {
	InKilometres float64
	InMiles      float64
}

func New(latitudeInDecimalDegrees, longitudeInDecimalDegrees float64) DecimalCoordinates {
	return DecimalCoordinates{
		Latitude:  latitudeInDecimalDegrees,
		Longitude: longitudeInDecimalDegrees,
	}
}

func CalculateDistance(startPoint, endPoint DecimalCoordinates) Result {
	startPointLatInRadians, startPointLongInRadians := calculateRadiansFromDecimalDegrees(startPoint)
	endPointLatInRadians, endPointLongInRadians := calculateRadiansFromDecimalDegrees(endPoint)

	deltaLatitude, deltaLongitude := calculateDeltaFromRadians(endPointLatInRadians, startPointLatInRadians, endPointLongInRadians, startPointLongInRadians)

	haversine := math.Pow(math.Sin(deltaLatitude/2), 2) + math.Cos(startPointLatInRadians)*math.Cos(endPointLatInRadians)*
		math.Pow(math.Sin(deltaLongitude/2), 2)

	calculation := 2 * math.Atan2(math.Sqrt(haversine), math.Sqrt(1-haversine))

	return buildResult(calculation)
}

func calculateDeltaFromRadians(endPointLatInRadians float64, startPointLatInRadians float64, endPointLongInRadians float64, startPointLongInRadians float64) (float64, float64) {
	deltaLatitude := endPointLatInRadians - startPointLatInRadians
	deltaLongitude := endPointLongInRadians - startPointLongInRadians
	return deltaLatitude, deltaLongitude
}

func calculateRadiansFromDecimalDegrees(point DecimalCoordinates) (float64, float64) {
	latitudeInRadians := convertDegreesToRadians(point.Latitude)
	longitudeInRadians := convertDegreesToRadians(point.Longitude)
	return latitudeInRadians, longitudeInRadians
}


func convertDegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func buildResult(haversineCalculation float64) Result {
	return Result{
		InKilometres: haversineCalculation * RadiusOfEarthInKilometres,
		InMiles:      haversineCalculation * RadiusOfEarthInMiles,
	}
}
