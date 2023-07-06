package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Friend interface {
	SayHello()
}

func (p *Person) SayHello() {
	fmt.Println("Hello,", p.Name)
}

func Greet(f Friend) {
	f.SayHello()
}

type Dog struct{}

func (d *Dog) SayHello() {
	fmt.Println("Woof woof")
}

// Testing and demonstrating the concepts above as a refresher.
func main() {
	fmt.Println("Hello, black hat gophers!")

	///// Types /////
	var count = int(42)
	ptr := &count
	fmt.Println(*ptr) // 42
	*ptr = 100
	fmt.Println(count) // 100

	// Using Person example.
	var guy = new(Person) // Inits new Person
	guy.Name = "Dave"
	guy.SayHello() // Calls SayHello method on guy
	Greet(guy)     //

	// Using Dog example.
	var dog = new(Dog)
	Greet(dog) // Demonstrating Friend interface
	///// Types /////

	///// Control Structures /////
	var x = 1
	if x == 1 {
		fmt.Println("x is equal to 1")
	} else {
		fmt.Println("x is not equal to 1")
	}

	// Go switches do NOT fall through.
	var str = "foo"
	switch str {
	case "foo":
		fmt.Println("Found foo")
	case "bar":
		fmt.Println("Found bar")
	default:
		fmt.Println("Default case")
	}
	///// Control Structures /////
}
