package main

import "fmt"

//Sigle responsability principle

// A class should have one, and only one, reason to change.
//–Robert C Martin

type robot struct {
	wheels  int
	sensors int
}

func (r robot) move() {
	//do staff
	fmt.Println("Go go go")
}

func (r robot) askWeather() string {
	//do staff
	return "Raining"
}

/////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////

type engine struct{}

func (e engine) move(wheels int) {
	//do staff
	fmt.Println("Go go go")
}

type assistant struct{}

func (a assistant) askWeather() string {
	//do staff
	return "Raining"
}

type robotSrp struct {
	wheels  int
	sensors int
	e       engine
	ia      assistant
}

func (r robotSrp) move() {
	r.e.move(r.wheels)
}

func (r robotSrp) askWeather() {
	fmt.Println(r.ia.askWeather())
}

/////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////

/*func main() {
	r := robotSrp{
		wheels:  3,
		sensors: 30,
		e:       engine{},
		ia:      assistant{},
	}
	r.move()
	r.askWeather()
}*/
