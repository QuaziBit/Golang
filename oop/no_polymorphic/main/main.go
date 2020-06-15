package main

import "fmt"

// Animal is parent structure
type Animal struct {
	Name string
	mean bool
}

// Cat is the child of the Animal structure
type Cat struct {
	Basics       Animal
	MeowStrength int
}

// Dog is the child of the Animal structure
type Dog struct {
	Animal
	BarkStrength int
}

// MakeNoise for the dog
func (dog *Dog) MakeNoise() {
	barkStrength := dog.BarkStrength

	if dog.mean == true {
		barkStrength *= 5
	}

	for bark := 0; bark < barkStrength; bark++ {
		fmt.Printf("BARK ")
	}
}

// MakeNoise for the cat
func (cat *Cat) MakeNoise() {
	meowStrength := cat.MeowStrength

	if cat.Basics.mean == true {
		meowStrength *= 5
	}

	for meow := 0; meow < meowStrength; meow++ {
		fmt.Printf("MEOW ")
	}
}

func main() {
	cat := Cat{
		Basics: Animal{
			"Cat",
			true,
		},
		MeowStrength: 2,
	}

	dog := Dog{
		Animal: Animal{
			"Dog",
			true,
		},
		BarkStrength: 2,
	}

	fmt.Printf("\n")
	cat.MakeNoise()
	fmt.Printf("\n")

	fmt.Printf("\n")
	dog.MakeNoise()
	fmt.Printf("\n\n")
}
