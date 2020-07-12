package main

import "fmt"

/*
* variables at the package level have to be declared like this
* when declared with lowercase the variable is exposed to the package scope. i.e anything
* file within the same package can access it. But when declared with uppercase it is
* available globally.
 */
var i int = 45

// I is available globally camel case and pascal case
var I int = 4

// declaring and initializing multiple variables
var (
	firstName string = "paul"
	lastName  string = "arah"
	age       int    = 23
)

/*
* enumerated constants
* iota is scoped to the constants blocks and resets anytime a new constant is declared
* _ is the only write-only value in go
 */
const (
	_ = iota
	a
	b
	c
)

func main() {

	// variable shadowing: The variable with the innermost scope takes the highest precedence.
	var firstName string = "paul"
	fmt.Println(firstName)
	fmt.Printf("%v, %T\n", I, I)

	/*
	* Data Types in Go
	* All data types have Zero Values
	* signed and unsigned integers int and uint
	* working across types must be explicit
	* Numeric primitive types:  int 8, 16, 32, 64, uint 32, 62, complex 64, 128, float 32, 64
	 */
	var n bool = true
	fmt.Println(n)

	x := 30
	y := int(3.0)
	fmt.Println(x / y)

	s := "this is a string"
	// converts a string to a collection byte slices
	q := []byte(s)
	// a rune is declared with a single ''
	var v rune = 'a'
	fmt.Println(v, q)
	const r int = 78
	/**
	* constants cannot be set to values that are determined by runtime.
	* collection types cannot be constant
	* constans can be shadowed
	 */
	//  untyped constant, the type is implicitly determined.
	const t = 45
	fmt.Println(r, t)
	// implicit conversion of constants
	const a = 2
	const b int = 5
	fmt.Println(a + b)

	/*
	* Collection data types: Arrays and slices
	*the elements in the array are laid out contiguous in memory
	* arrays are more like static arrays while slices are dynamic arrays
	* Hence a slices might have different length and capacities
	* Slices point to the underlying data while arrays makes a copy of the data when
	* assigned
	* slicing syntax works on slices & arrays
	* len - length, cap - capacity
	 */

	var grades [4]int
	test := [3]int{1, 2, 3}
	test = [...]int{3, 5, 6}

	/*
	* Maps and struct
	* The order of a map is guaranteed on returns
	* the key still sorta persist in a map ven after it has been deleted hence checking
	* for presence results in zero. to solve this add an extra variable that will. This does not
	* apply to a struct tho.
	* a boolean to confirm id the key is indeed present.
	* avoid positional syntax for struct and dot notation.
	* a struct is the closests to a javscript object in form and behavior
	* The values of struct remain independent when moving things around
	* composition of struct can be done using embeding. This possible form outside or from the
	* the literal syntax.
	 */

	dict := map[string]int{
		"paul": 23,
		"kemi": 24,
	}
	dict["paul"] = 50
	delete(dict, "kemi")
	_, ok := dict["kemi"]

	type Person struct {
		age  int
		head string
	}

	type Doctor struct {
		Person
		number     int
		actorName  string
		companions []string
	}

	aDoctor := Doctor{
		number:    3,
		actorName: "paul",
		companions: []string{
			"ahmed",
			"wens",
			"eric",
		},
	}

	fmt.Println(grades, test, len(test), test[1:], ok, aDoctor)
}
