package main

import "fmt"

// Animal is parent structure
type Animal struct {
	Name string
	mean bool
}

// AnimalSounder interface
type AnimalSounder interface {
	MakeNoise()
}

// Cat is the child of the Animal structure
type Cat struct {
	Basics       Animal
	NoiseType    string
	MeowStrength int
}

// Dog is the child of the Animal structure
type Dog struct {
	Animal
	NoiseType    string
	BarkStrength int
}

// MakeNoise for the dog
func (dog *Dog) MakeNoise() {
	dog.PerformNoise(dog.BarkStrength, dog.NoiseType)
}

// MakeNoise for the cat
func (cat *Cat) MakeNoise() {
	cat.Basics.PerformNoise(cat.MeowStrength, cat.NoiseType)
}

// MakeSomeNoise ...
func MakeSomeNoise(animalSounder AnimalSounder) {
	animalSounder.MakeNoise()
}

// PerformNoise ...
func (animal *Animal) PerformNoise(strength int, sound string) {
	if animal.mean == true {
		strength = strength * 5
	}

	for voice := 0; voice < strength; voice++ {
		fmt.Printf("%s ", sound)
	}

	fmt.Println("")
}

func main() {
	cat := &Cat{
		Basics: Animal{
			"Cat",
			true,
		},
		NoiseType:    "MEOW",
		MeowStrength: 2,
	}

	dog := &Dog{
		Animal: Animal{
			"Dog",
			true,
		},
		NoiseType:    "BARK",
		BarkStrength: 2,
	}

	fmt.Printf("\n")
	MakeSomeNoise(cat)
	fmt.Printf("\n")
	MakeSomeNoise(dog)
	fmt.Printf("\n")
}
