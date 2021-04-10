package main

import "fmt"

//Liskov substitution principle

//Two types are substitutable if they exhibit behaviour such that the caller is unable to tell the difference.
//-Barbara Liskov

type car interface {
	Run()
}

type mercedes struct{}

func (mercedes) Run() {
	fmt.Println("Running with a Mercedes")
}

type fiat struct{}

func (fiat) Run() {
	fmt.Println("Running with a Fiat")
}

type honda struct{}

func (honda) Run(kmh int) {
	fmt.Println("Running with a Honda at ", kmh)
}

/////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////

/*func main() {

	c := [2]car{mercedes{}, fiat{}}
	//c := [2]car{mercedes{}, fiat{}, honda{}}
	for _, s := range c {
		s.Run()
	}
}*/
