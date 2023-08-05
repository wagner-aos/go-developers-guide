package main

import "fmt"

/*In Golang, the defer keyword is used to delay the execution of a function until the surrounding function completes.
The deferred function calls are executed in Last-In-First-Out (LIFO) order.
That means the most recently deferred function is executed first, followed by the second most recently deferred function, and so on.

The defer keyword is useful when we want to ensure that certain operations are performed before a function returns,
regardless of whether an error occurs or not. This can help simplify error handling and make code more readable.*/

/*Multiple Deferred Function Calls
In Golang, we can defer multiple function calls in a single function.
When we defer multiple functions, they are executed in reverse order. Here's an example
*/

func main() {

	defer fmt.Println("Third")
	defer fmt.Println("Second")
	defer fmt.Println("First")
	fmt.Println("Hello")

	//It will print from bottom to top sequence

}
