package main

import  (
	"fmt"
	"math"
)

type person struct{
	first	string
	last	string
	age 	int
}

type square struct{
	side	float32
}

type circle struct{
	radius	float32
}

type shape interface{
	area() float32
}

func main() {
	/*
		Exercise 1
        1. Create a func with the identifier foo that returns an int
        2. Create a func with the identifier bar that returns an int and a string
        3. Call both funcs
        4. Print out their results
	*/
	x := returnInt()
	y, z := returnIntAndString()
	fmt.Printf("foo returns: %v\t\tbar returns: %v, %v\n", x, y, z)

	/*
		Exercise 2
		1. Create a func with the identifier foo that:
			a. Takes in a variadic parameter of type int
			b. Pass in a value of type []int into your func (unfurl the []int)
			c. Returns the sum of all values of type int passed in
		2. Create a func with the identifier bar that: 
			a. Takes in a parameter of type []int
			b. Returns the sum of all values of type int passed in
	*/
	xi := []int{ 1, 10, 100, 1000, 10000 }
	total1 := foo(xi...)
	fmt.Printf("total1: %v\n", total1)

	ii := []int{ 1, 11, 111, 1111, 11111 }
	total2 := bar(ii)
	fmt.Printf("total2: %v\n", total2)

	/*
		Exercise 3
		Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
	*/
	defer testDefer()	

	/*
		Exercise 4
		1. Create a user defined struct with:
			a. Identifier “person”
			b. fields: first, last, age
		2. Attach a method to type person with:
			a. Identifier: “speak”
			b. The method should have the person say their name and age
		3. Create a value of type person
		4. Call the method from the value of type person
	*/
	p1 := person{
		first: "Bob",
		last: "Dylan",
		age: 100,
	}
	p1.speak()

	/*
		Exercise 5
		1. Create a type SQUARE
		2. Create a type CIRCLE
		3. Attach a method to each that calculates AREA and returns it
			a. Circle area= π r 2
			b. Square area = L * W
		4. Create a type SHAPE that defines an interface as anything that has the AREA method
		5. Create a func INFO which takes type shape and then prints the area
		6. Create a value of type square
		7. Create a value of type circle
		8. Use func info to print the area of square
		9. Use func info to print the area of circle
	*/
	sqr := square {
		side: 5.0,
	}	

	crcl := circle {
		radius: 10.0,
	}
	info(sqr)
	info(crcl)
	
	/*
		Exercise 6
		Build and use an anonymous func 
	*/
	func (s string) {
		fmt.Printf("This is an anonymous func that takes a string and prints it here: %v\n", s)
	}(`"Do you, Mr. Jones?"`)
	
	/*
		Exercise 7
		Assign a func to a variable, then call that func
	*/
	f1 := func() {
		fmt.Printf("This is a func expression\n")
	}
	f1()

	/*
		Exercise 8
		1. Create a func which returns a func
		2. assign the returned func to a variable
		3. call the returned func
	*/
	f2 := makeFunc()
	f2()
	
	/*
		Exercise 9
		Pass a func into a func as an argument 
	*/
	f3 := func() {
		fmt.Println("This is a callback func")
	}
	takesCallback(f3)

	/*
		Exercise 10
		Create a func which “encloses” the scope of a variable:
	*/
	f4 := approachZero()
	for i := 0; i<10; i++ {
		f4()
	}
}
	
func returnInt() int {
	return 42
}
	
func returnIntAndString() (int, string) {
	return 1984, "Big brother"
}

func foo(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func bar(xi []int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func testDefer() {
	fmt.Println("This will print last")
}

func (p person) speak() {
	fmt.Printf("My name is %v %v. I am %v years old.\n", p.first, p.last, p.age)
}

func (s square) area() float32 {
	return s.side * s.side
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Printf("area of %T: %v\n", s, s.area())
}

func makeFunc() func() {
	return func() {
		fmt.Println("This func was returned by another func")
	}
}

func takesCallback(callback func()) {
	callback()
}

func approachZero() func() {
	x := 1.0
	return func() {
		x /= 2.0
		fmt.Println(x)
	}
}