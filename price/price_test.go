package price_test

import (
	"testing"

	"github.com/thirathawat/rental/price"
)

func TestChildrensCharge(t *testing.T) {
	childrens := new(price.Childrens)

	testCases := []struct {
		daysRented int
		want       float64
	}{
		{1, 1.5},
		{3, 1.5},
		{4, 3.0},
	}

	for _, tc := range testCases {
		charge := childrens.Change(tc.daysRented)

		if charge != tc.want {
			t.Errorf("Childrens Change method did not return the correct amount for %d days rented. Expected: %f, Got: %f", tc.daysRented, tc.want, charge)
		}
	}
}

func TestChildrensCode(t *testing.T) {
	childrens := new(price.Childrens)
	want := price.ChildrensCode

	code := childrens.Code()

	if code != want {
		t.Error("Childrens Code method did not return the correct code.")
	}
}

func TestNewReleaseCharge(t *testing.T) {
	newRelease := new(price.NewRelease)

	testCases := []struct {
		daysRented int
		want       float64
	}{
		{1, 3.0},
		{2, 6.0},
	}

	for _, tc := range testCases {
		charge := newRelease.Change(tc.daysRented)

		if charge != tc.want {
			t.Errorf("NewRelease Change method did not return the correct amount for %d days rented. Expected: %f, Got: %f", tc.daysRented, tc.want, charge)
		}
	}
}

func TestNewReleaseCode(t *testing.T) {
	newRelease := new(price.NewRelease)
	want := price.NewReleaseCode

	code := newRelease.Code()

	if code != want {
		t.Error("NewRelease Code method did not return the correct code.")
	}
}

func TestRegularCharge(t *testing.T) {
	regular := new(price.Regular)

	testCases := []struct {
		daysRented int
		want       float64
	}{
		{1, 0.5},
		{2, 2.0},
		{3, 3.5},
	}

	for _, tc := range testCases {
		charge := regular.Change(tc.daysRented)

		if charge != tc.want {
			t.Errorf("Regular Change method did not return the correct amount for %d days rented. Expected: %f, Got: %f", tc.daysRented, tc.want, charge)
		}
	}
}

func TestRegularCode(t *testing.T) {
	regular := new(price.Regular)
	want := price.RegularCode

	code := regular.Code()

	if code != want {
		t.Error("Regular Code method did not return the correct code.")
	}
}
