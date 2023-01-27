package rental_test

import (
	"testing"

	"github.com/thirathawat/rental"
	"github.com/thirathawat/rental/movie"
	"github.com/thirathawat/rental/price"
)

func TestMovie(t *testing.T) {
	rental := rental.New(movie.New("Regular Movie", new(price.Regular)), 3)
	want := "Regular Movie"

	movie := rental.Movie().Title()

	if movie != want {
		t.Errorf("Expect %v, got %v", want, movie)
	}
}

func TestDaysRented(t *testing.T) {
	rental := rental.New(movie.New("Regular Movie", new(price.Regular)), 3)
	want := 3

	daysRented := rental.DaysRented()

	if daysRented != want {
		t.Errorf("Expect %v, got %v", want, daysRented)
	}
}

func TestCharge(t *testing.T) {
	rental := rental.New(movie.New("Regular Movie", new(price.Regular)), 3)
	want := 3.5

	charge := rental.Charge()

	if charge != want {
		t.Errorf("Expect %v, got %v", want, charge)
	}
}

func TestPoints(t *testing.T) {
	regular := rental.New(movie.New("Regular Movie", new(price.Regular)), 3)
	want := 1

	points := regular.Points()

	if points != want {
		t.Errorf("Expect %v, got %v", want, points)
	}
}
