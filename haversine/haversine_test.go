package haversine_test

import (
	"github.com/rav-pradhan/go-measure-distance/haversine"
	"math"
	"testing"
)

const tolerance = .001

func TestHaversineFormula(t *testing.T) {
	t.Run("Haversine formula is correctly calculated between two locations", func(t *testing.T) {
		cardiff := haversine.New(51.471583, -3.179090)
		london := haversine.New(51.509865, -0.118092)

		got := haversine.CalculateDistance(cardiff, london)
		want := haversine.Result{
			InKilometres: 211.954134,
			InMiles:      131.703661,
		}

		assertMatch(t, got, want)
	})
	t.Run("Haversine formula is correctly calculated between two locations", func(t *testing.T) {
		cardiffMillenniumCentre := haversine.New(51.465232, -3.165094)
		cardiffCastle := haversine.New(51.481725, -3.180801)

		got := haversine.CalculateDistance(cardiffMillenniumCentre, cardiffCastle)
		want := haversine.Result{
			InKilometres: 2.132325,
			InMiles:      1.324980,
		}

		assertMatch(t, got, want)
	})
}

func assertMatch(t *testing.T, got haversine.Result, want haversine.Result) {
	if math.Abs(got.InKilometres-want.InKilometres) > tolerance {
		t.Errorf("got %f, wanted %f", got.InKilometres, want.InKilometres)
	}

	if math.Abs(got.InMiles-want.InMiles) > tolerance {
		t.Errorf("got %f, wanted %f", got.InMiles, want.InMiles)
	}
}