package main

import "fmt"

//Interface segregation principle

//Clients should not be forced to depend on methods they do not use.
//–Robert C. Martin

type cloth struct {
	name  string
	value int
}

func (c cloth) getName() string {
	return c.name
}

func (c cloth) getPrice() int {
	return c.value
}

func (c cloth) wear() {
	//do staff
}

/////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////

type food struct {
	name       string
	valuePerUd int
	unit       int
	cookable   bool
}

func (f food) getName() string {
	return f.name
}

func (f food) getPrice() int {
	return f.valuePerUd * f.unit
}

func (f food) cook() {
	if f.cookable {
		fmt.Println("Cooking ", f.name)
	}
}

type forSale interface {
	getName() string
	getPrice() int
}

/////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////

//func Sell(s []cloth) {
func Sell(s []forSale) {
	for _, i := range s {
		fmt.Printf("Selling %s by %d\n", i.getName(), i.getPrice())
	}
}

/*func main() {

	fs := []forSale{
		cloth{name: "jacket", value: 45},
		food{name: "apple", valuePerUd: 1, unit: 2, cookable: false},
	}

	Sell(fs)
}*/
