package rental

import (
	"github.com/thirathawat/rental/movie"
	"github.com/thirathawat/rental/price"
)

type R struct {
	movie      movie.M
	daysRented int
}

func New(movie movie.M, daysRented int) R {
	return R{
		movie:      movie,
		daysRented: daysRented,
	}
}

func (r R) Movie() movie.M {
	return r.movie
}

func (r R) DaysRented() int {
	return r.daysRented
}

func (r R) Charge() float64 {
	return r.movie.Change(r.daysRented)
}

func (r R) Points() int {
	if r.Movie().PriceCode() == price.NewReleaseCode && r.DaysRented() > 1 {
		return 2
	}
	return 1
}
