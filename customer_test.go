package rental_test

import (
	"testing"

	"github.com/thirathawat/rental"
	"github.com/thirathawat/rental/movie"
	"github.com/thirathawat/rental/price"
)

func TestCustomerName(t *testing.T) {
	customer := rental.NewCustomer("Thirathawat")
	want := "Thirathawat"

	name := customer.Name()

	if name != want {
		t.Errorf("Expect %v, got %v", want, name)
	}
}

func TestCustomerPoints(t *testing.T) {
	customer := rental.NewCustomer("Thirathawat")
	customer.AddRental(rental.New(movie.New("Kingsman", new(price.Regular)), 2))
	customer.AddRental(rental.New(movie.New("Iron Man", new(price.Regular)), 3))
	want := 2

	points := customer.Points()

	if points != want {
		t.Errorf("Expect %v, got %v", want, points)
	}
}

func TestCustomerAmount(t *testing.T) {
	customer := rental.NewCustomer("Thirathawat")
	customer.AddRental(rental.New(movie.New("Kingsman", new(price.Regular)), 2))
	customer.AddRental(rental.New(movie.New("Iron Man", new(price.Regular)), 3))
	want := 5.5

	amount := customer.Amount()

	if amount != want {
		t.Errorf("Expect %v, got %v", want, amount)
	}
}

func TestCustomerStatement(t *testing.T) {
	customer := rental.NewCustomer("Thirathawat")
	customer.AddRental(rental.New(movie.New("Kingsman", new(price.Regular)), 2))
	customer.AddRental(rental.New(movie.New("Iron Man", new(price.Regular)), 3))
	customer.AddRental(rental.New(movie.New("The Avengers", new(price.NewRelease)), 1))
	customer.AddRental(rental.New(movie.New("Shang-chi", new(price.NewRelease)), 2))
	customer.AddRental(rental.New(movie.New("Ant-Man", new(price.Childrens)), 3))
	customer.AddRental(rental.New(movie.New("The Batman", new(price.Childrens)), 4))

	want := `Rental Record for Thirathawat
	Kingsman	2.0
	Iron Man	3.5
	The Avengers	3.0
	Shang-chi	6.0
	Ant-Man	1.5
	The Batman	3.0
Amount owed is 19.0
You earned 7 frequent renter points`

	if want != customer.Statement() {
		t.Errorf("Expect \n%v\n, got \n%v", want, customer.Statement())
	}
}
