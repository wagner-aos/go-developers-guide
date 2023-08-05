package main

//. "go-developers-guide/function-first-class/practicaluse"
import (
	"fmt"
)

/*
https://golangbot.com/first-class-functions/

What are first class functions?
A language that supports first class functions allows functions to be assigned to variables,
passed as arguments to other functions and returned from other functions. Go has support
for first class functions.
- Anonymous functions
- Anonymous Functions with arguments
- User defined functions (function as a type)
- High-Order Functions
	- takes one or more functions as arguments.
	- returns a function as its result.
*/

//Here we defined a variables of type Function.
type add func(a, b int) int

func main() {

	////Anonymous Functions.
	//anonymous()

	////Anonymous Functions with arguments.
	//anonymousWithArguments()

	////User defined Functions.
	//userDefined()

	////Higher-order functions
	//functionsAsArguments()
	//functionAsReturn()

	////Closures
	//closure()

	//PRATICAL USE OF FIRST CLASS FUNCTIONS
	praticalUseOfFirstClassFunction()

}

func anonymous() {
	a := func() {
		fmt.Println("Hello world first class functions")
	}
	a()

	fmt.Printf("%T", a) //Format directive will print the type of the function, func().
}

func anonymousWithArguments() {
	//It is also possible to pass arguments to anonymous functions just like any other function.
	func(n string) {
		fmt.Println("Anonymous with arguments", n)
	}("for Gophers")
}

func userDefined() {
	//User defined function is declared type, just like this:
	// type add func(a, b int) int
	var a add = func(a, b int) int {
		return a + b
	}
	fmt.Println("Sum", a(3, 5))
}

/*
The definition of Higher-order function from wiki is a function which does at least one of the following.

- takes one or more functions as arguments.
- returns a function as its result.
*/

//Passing functions as arguments to other functions
func simple(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

func functionsAsArguments() {
	f := func(a, b int) int {
		return a + b
	}
	simple(f)
}

//Returning functions from other functions
func simpleReturn() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func functionAsReturn() {
	s := simpleReturn()
	fmt.Println(s(60, 8))
}

/*Closures are a special case of anonymous functions.
Closures are anonymous functions that access the variables defined outside the body of the function.
*/
func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string { //the same signature (siguinatiure) of the return of the function appendStr.
		t = t + " " + b + "\n"
		return t
	}
	return c
}

func closure() {
	a, b := appendStr(), appendStr()
	fmt.Print((a("World")))
	fmt.Print((b("Everyone")))
}
