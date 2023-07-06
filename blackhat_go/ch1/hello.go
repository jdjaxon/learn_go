package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"time"
)

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

func checkType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("INTEGER!!!")
	case string:
		fmt.Println("STRING!!!")
	default:
		fmt.Println("UNKNOWN!!!")
	}
}

// For concurrency example
func f() {
	fmt.Println("f function concurrency example")
}

func strlen(s string, c chan int) {
	c <- len(s)
}

// Error handling example
type MyError string

func (e MyError) Error() string {
	return string(e)
}

func foo() error {
	return errors.New("Some error Occurred")
}

// For Handling Structured Data example
type FooJSON struct {
	Bar string
	Baz string
}

type FooXML struct {
	Bar string `xml:"id,attr"`
	Baz string `xml:"parent>child"`
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

	fmt.Print("count is an ")
	checkType(count)
	fmt.Print("str is an ")
	checkType(str)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var nums = []int{2, 4, 6, 8, 10}
	for idx, val := range nums {
		fmt.Println(idx, val)
	}
	///// Control Structures /////

	///// Concurrency /////
	go f()
	// Needed so that main doesn't exit before f() is done
	time.Sleep(1 * time.Second)
	fmt.Println("main function")

	// Go provides a type called channels which allow goroutines to sync and
	// communicate.
	var c = make(chan int)
	go strlen("Salutations", c)
	go strlen("World", c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
	///// Concurrency /////

	///// Error Handling /////
	if err := foo(); err != nil {
		fmt.Println("ERROR!!!")
	}
	///// Error Handling /////

	///// Handling Structured Data /////
	jsonData := FooJSON{"Joe Junior", "Hello Shabado"}
	b1, _ := json.Marshal(jsonData)
	fmt.Println(string(b1))
	json.Unmarshal(b1, &jsonData)

	xmlData := FooXML{"Joe Junior", "Hello Shabado"}
	b2, _ := xml.Marshal(xmlData)
	fmt.Println(string(b2))
	json.Unmarshal(b2, &xmlData)
	///// Handling Structured Data /////
}
