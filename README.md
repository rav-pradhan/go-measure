# go-measure-distance

## What is this?

`go-measure-distance` is a crude library that measures distance between two points (latitude/longitude). This is currently calculated using the [haversine formula](https://en.wikipedia.org/wiki/Haversine_formula).

## Installation

Install with `go get github.com/rav-pradhan/go-measure-distance`

## Usage

```go
// Initialise a coordinates object by passing in latitude and longitude as decimal coordinates
cardiff := haversine.New(51.471583, -3.179090)
london := haversine.New(51.509865, -0.118092)

// run the CalculateDistance method with the two locations as arguments)
result := haversine.CalculateDistance(cardiff, london)

// result is returned as a floating point up to 6 decimal places
fmt.Println(result.InMiles)
fmt.Println(result.InKilometres)
```
