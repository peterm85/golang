package main

import "fmt"

//Open/closed principle

//Software entities should be open for extension, but closed for modification.
//–Bertrand Meyer, Object-Oriented Software Construction

type sword struct {
	name string
}

func (sword) Damage() int {
	return 30
}

//String implements fmt.Stringer interface
func (s sword) String() string {
	return fmt.Sprintf("%s is a sword that can deal %d points of damage to opponents", s.name, s.Damage())
}

/////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////

type enchantedSword struct {
	sword
}

func (es enchantedSword) Damage() int {
	return 55
}

func (es *enchantedSword) rename(n string) {
	fmt.Println("Renaming sword")
	es.name = n
}

type weapon interface {
	Damage() int
}

/////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////

/*func main() {
	sw := sword{name: "WidowMaker"}
	es := enchantedSword{sword{name: "ThunderBird"}}

	w := [2]weapon{sw, es}
	for i, s := range w {
		fmt.Printf("[%d] This weapon can deal %d points of damage to opponents\n", i, s.Damage())
	}

	st := [2]fmt.Stringer{sw, es}
	for _, s := range st {
		fmt.Println(s.String())
	}

	es.rename("DragonSlayer")
	fmt.Println(es.String())
}*/
