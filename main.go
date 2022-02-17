package main

import "fmt"

var (
	a [4]string //four variable array
)

func main() {

	a[0] = "This is"
	a[1] = "what you earned"
	a[2] = "from"
	a[3] = "this encounter"

	fmt.Println(a[0], a[1], a[2], a[3])         // Prints entire array
	fmt.Printf("%s %s %s \n", a[0], a[1], a[2]) // Prints 0,1,2
	fmt.Println(a)                              // Prints array as a horizontal list with brackets

	lootFish()
}

//Sup Buddy. How you doin'? I'm doin' aight."

func lootFish() {
	a := []string{"ring of protection +1", "scalemail"}
	a = append(a, "potion of cure light wounds") // adds to collection

	fmt.Println(a[0]) // prints ring
	fmt.Println(a[1]) // prints scalemail
	fmt.Println(a[2]) // prints potion

	fmt.Printf("%q", a) // prints collection in brackets
}
