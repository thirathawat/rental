package rental

import "fmt"

type Customer struct {
	name    string
	rentals []R
}

func NewCustomer(name string) Customer {
	return Customer{
		name:    name,
		rentals: []R{},
	}
}

func (c Customer) Name() string {
	return c.name
}

func (c *Customer) AddRental(r R) {
	c.rentals = append(c.rentals, r)
}

func (c Customer) Points() int {
	points := 0
	for _, r := range c.rentals {
		points += r.Points()
	}
	return points
}

func (c Customer) Amount() float64 {
	amount := 0.0
	for _, r := range c.rentals {
		amount += r.Charge()
	}
	return amount
}

func (c Customer) Statement() string {
	result := fmt.Sprintf("Rental Record for %s\n", c.name)
	for _, r := range c.rentals {
		result += fmt.Sprintf("\t%s\t%.1f\n", r.Movie().Title(), r.Charge())
	}
	result += fmt.Sprintf("Amount owed is %.1f\n", c.Amount())
	result += fmt.Sprintf("You earned %v frequent renter points", c.Points())
	return result
}
