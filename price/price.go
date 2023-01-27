package price

type Code int

const (
	_ Code = iota
	ChildrensCode
	NewReleaseCode
	RegularCode
)

type Pricer interface {
	Change(daysRented int) float64
	Code() Code
}

type Childrens struct{}

func (c Childrens) Change(daysRented int) float64 {
	amount := 1.5
	if daysRented > 3 {
		amount += float64(daysRented-3) * 1.5
	}
	return amount
}

func (c Childrens) Code() Code {
	return ChildrensCode
}

type NewRelease struct{}

func (n NewRelease) Change(daysRented int) float64 {
	return float64(daysRented) * 3
}

func (n NewRelease) Code() Code {
	return NewReleaseCode
}

type Regular struct{}

func (r Regular) Change(daysRented int) float64 {
	return 2 + float64(daysRented-2)*1.5
}

func (r Regular) Code() Code {
	return RegularCode
}
