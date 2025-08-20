package main

import "fmt"

// Step 1: Define an interface
// An interface is just a list of methods that something can do
type Animal interface {
	Speak() string
}

// Step 2: Create types that implement the interface
// Any type with a Speak() method automatically satisfies Animal interface **
//   The interface is satisfied structurally (by shape), not nominally (by name declaration).

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

type Cow struct {
	Name string
}

func (c Cow) Speak() string {
	return "Moo!"
}

// Step 3: Use the interface
// This function works with ANY type that has a Speak() method
func MakeSound(a Animal) {
	fmt.Printf("The animal says: %s\n", a.Speak())
}

func main() {
	// Create different animals
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}
	cow := Cow{Name: "Bessie"}

	// They all satisfy the Animal interface because they have Speak() method
	fmt.Println("=== Interface Demo ===")

	MakeSound(dog) // Works!
	MakeSound(cat) // Works!
	MakeSound(cow) // Works!

	// You can also store them in a slice of Animal interface
	animals := []Animal{dog, cat, cow}

	fmt.Println("\n=== All animals speaking ===")
	for _, animal := range animals {
		MakeSound(animal)
	}

	fmt.Println("\n🎯 Key Point: MakeSound() doesn't care what type of animal it is!")
	fmt.Println("It just knows each animal can Speak() - that's the power of interfaces!")
}
