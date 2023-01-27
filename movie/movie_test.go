package movie_test

import (
	"testing"

	"github.com/thirathawat/rental/movie"
	"github.com/thirathawat/rental/price"
)

func TestTitle(t *testing.T) {
	movie := movie.New("Toy Story", new(price.Childrens))
	want := "Toy Story"

	title := movie.Title()

	if title != want {
		t.Errorf("Expect %v, got %v", want, title)
	}
}

func TestChange(t *testing.T) {
	movie := movie.New("Toy Story", new(price.Childrens))
	want := 1.5

	change := movie.Change(3)

	if change != want {
		t.Errorf("Expect %v, got %v", want, change)
	}
}

func TestPriceCode(t *testing.T) {
	movie := movie.New("Toy Story", new(price.Childrens))
	want := price.ChildrensCode

	priceCode := movie.PriceCode()

	if priceCode != want {
		t.Errorf("Expect %v, got %v", want, priceCode)
	}
}
