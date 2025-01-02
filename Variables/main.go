// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var user string
// 	user = "Harry"
// 	fmt.Println(user)
// }

// >>> go run main.go
//     Harry

// Incorrect Values :-
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var s string
// 	s = 123
// 	fmt.Println(s)
// }

// >>> go run main.go
// .\main.go:24:6: cannot use 123 (untyped int constant) as string value in assignment

// Shorthand :-
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var s, t string = "foo", "bar"
// 	fmt.Println(s)
// 	fmt.Println(t)
// }

// >>> go run main.go
//     foo
//	   bar

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var (
// 		s string = "foo"
// 		i int    = 5
// 	)
// 	fmt.Println(s)
// 	fmt.Println(i)
// }

// >>> go run main.go
//     foo
//	   5

// Short Variable Declaration :-

// s := "Hello World"

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	name := "Lisa"
// 	name = "Peter"
// 	fmt.Println(name)
// }

// >>> go run main.go
//     Peter

package main

import (
	"fmt"
)

func main() {
	name := "Lisa"
	name = 12
	fmt.Println(name)
}

// >>> go run main.go
//main.go:95:9: cannot use 12 (untyped int constant) as string value in assignment
