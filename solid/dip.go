package main

import "fmt"

//Dependency inversion

//High-level modules should not depend on low-level modules. Both should depend on abstractions.
//Abstractions should not depend on details. Details should depend on abstractions.
//–Robert C. Martin

type book struct {
	pages  int
	weight int
}

func (book) read() {
	fmt.Println("Reading book")
}

type EPUB struct {
	pages int
}

func (EPUB) read() {
	fmt.Println("Reading EPUB")
}

type kindle struct {
	pages int
}

func (kindle) read() {
	fmt.Println("Reading kindle")
}

type readable interface {
	read()
}

/////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////

func Read(r []readable) {
	for _, i := range r {
		i.read()
	}
}

func main() {

	r := []readable{
		book{pages: 256, weight: 102},
		EPUB{pages: 133},
		kindle{pages: 322},
	}

	Read(r)
}
