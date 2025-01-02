// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello World")
// }

// /*
// >>> go run main4.go
// Hello World
// */

// // Printing a variable
// package main

// import "fmt"

// func main() {
// 	var city string = "Kolkata"
// 	fmt.Print(city)
// }

// /*
// >>> go run main4.go
// Kolkata
// */

//Printing a variable and string together

// package main

// import "fmt"

// func main() {
// 	var name string = "KodeKloud"
// 	var user string = "Harry"

// 	fmt.Print("Welcome to ", name, ", ", user)
// }

// /*
// >>> go run main4.go
// Welcome to KodeKloud, Harry
// */

// package main

// import "fmt"

// func main() {
// 	var name string = "KodeKloud"
// 	var user string = "Harry"
// 	fmt.Print(name)
// 	fmt.Print(user)
// }

// /*
// >>> go run main4.go
// KodeKloudHarry
// */

// newline character	\n

// package main

// import "fmt"

// func main() {
// 	var name string = "KodeKloud"
// 	var user string = "Harry"
// 	fmt.Print(name, "\n")
// 	fmt.Print(user)
// }

// /*
// >>> go run main4.go
// KodeKloud
// Harry
// */

// Println

// package main

// import "fmt"

// func main() {
// 	var name string = "KodeKloud"
// 	var user string = "Harry"
// 	fmt.Println(name)
// 	fmt.Println(user)
// }

// /*
// >>> go run main4.go
// KodeKloud
// Harry
// */

// Printf	Formatted input output

// fmt.Printf("Template string %s", Object args(s))

package main

import "fmt"

func main() {
	var name string = "Joe"
	var i int = 78
	fmt.Printf("Hey, %v! You have score %d/100 in Physics", name, i)
}

/*
>>> go run main.go
    Hey, Joe! You have score 78/100 in Physics
*/
