# GOLANG NOTES

## CLI
 - go build <go_files>: compiles module to an executable binary file
   - built executable will be named after first file in args
 - go run <go_file>: builds AND executes program
 - ./<executable_file>: runs the built executable for your go package
 - go test: runs any tests in directory
 - 

 ## PACKAGES
 Packages may contain multiple files.
 Each file in a package must declare what package they belong to on the first line:
 ```go
 package main
 ```

 ### Types of Packages
 1. Executable: generates a file that we can run
    - the name `main` indicates an executable. `main` is the only package name that builds to an executable
    - must also have a function inside called `main`
 2. Reusable: Code used as helpers. Good place to put reusable logic

 Files in the same package fo not have to be imported into each other

 ### fmt package
 Standard lib package for formatting I/O

 ### func
 basic syntax is same as js

 All return values must be declared before the open bracket
 ```go
 func getNote() string {
    return "C#"
 }
 ```


## VARIABLES
types are automatically inferred from the assignment value
```go
// These two declarations are equivalent - the type string is inferred by the value
var name string = "Mike"
name := "Mike"
// The := operator is only used for initialization.
// To reassign an existing variable you must use =
```

## TYPES

### VALUE TYPES
use pointers to mutate these things in a function
type default zero values in parens
  - `int` (0)
  - `float64` (0)
  - `string` ("")
  - `bool` (false)
  - `struct` (nil)
### REFERENCE TYPES
  - `array`
   - fixed-length list of values
   - every value in array must be of same type
  - `slice`
   - an array that can grow or shrink
  - `map`
  - `channels`
  - `pointers`  
  - `func`
## ARRAYS v SLICES
-  Arrays:
   - primitive data structure
   - can't be resized
   - rarely used directly
   - value type
- Slices:
   - can grow and shrink
   - used 99% of the time for lists of elements
   - reference type:
      - a slice is a pointer to an underlying array
      - every time you create a slice, go creates two items in memory
         - the slice value itself
         - an internal array 
         - when a slice is passed by value into a func,
         the func makes a new copy of the slice but it still points to the underlying array value

```go
type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

func main() {

	var d Direction = North
	fmt.Print(d)
	switch d {
	case North:
		fmt.Println(" goes up.")
	case South:
		fmt.Println(" goes down.")
	default:
		fmt.Println(" stays put.")
	}

	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
   name := "Bill"
   // cast string => to pointer address `&` => back to string value `*`
	fmt.Println(*&name)
}
// slices are passed by value BUT the copied values points to the same underlying array
// so when you mutate a slice in a func it changes the slice value in the outer scope
func updateSlice(s []string) {
	s[0] = "Bye"

}
```

## STRUCTS
```go
type contactInfo struct {
   email string
   zipCode int
}
// embedded struct with an explicit key
type person struct {
   firstName string
   lastName string
   contact contactInfo
}
// embedded struct with an implicit key (the name of the embedded struct)
type person struct {
   firstName string
   lastName string
   contactInfo
}
// 3 ways to intantiate struct
// relies on struct field sequence (BAD!)
me := person{"Mike", "Treacy"} 
// explicit, readable, independent of field sequence (GOOD!)
me := person{
   firstName: "Mike",
   lastName: "Treacy",
   contact: contactInfo{
      email: "mike@gmail.com",
      zipCode: 99508
      }
      }
// instantiates all struct fields with zero values
var me person
me.firstName = "Mike"
me.lastName = "Treacy"

me.updateFirstName("Michael")
me.print() // prints firstname: Mike unexpectedly

func (p person) updateName(newFirstName string) {
   p.firstName = newFirstName
}
func (p person) print() {
   fmt.Printf("%+p", p)
}
```
## MAPS
- a collection of key-value pairs
- like an Object in JS or Dict in Python
- both the keys and the values are statically typed
   - all keys must be the same type
   - all values must be the same type

```go
// 3 ways to declare a map
// #1 declaration with initialization
colors := map[string]string{
      "red":   "#ff0000",      
		"green": "#4bf745", // map literal requires trailing comma!
   }
// #2 declaration without initialization
var colors map[string]string
// #3 declaration without initialization
colors := make(map[string]string)

// add keys to map with bracket notation
// maps do NOT have dot notation
// this is because all map keys are typed
colors["white"] = "#ffffff"

colors := make(map[int]string)
colors[10] = "#ffffff"

// to delete keys and values from a map
delete(colors, 10)

// iterating over values in a map
func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
}

func printMap(colorMap map[string]string) {
	// range returns a tuple of key, value
	for color, hex := range colorMap {
		fmt.Println(color + ": " + hex)

	}
}
```

### Maps vs Structs
 - Map:
   - used to represent a collection of related properties
   - all keys same type
   - all values same type
   - keys are indexed - can iterate over them   
   - don't need to know all the keys at compile time
      - can dynamically change keys after declaration
   - Reference Type!

- Struct:
   - used to model a "thing" with a lot of different properties
   - values can be different types
   - keys don't support indexing - can't iterate over them
   - need to know all the different fields at compile time
      - can't dynamically change field names after declaration   
   - Value Type!


## INTERFACES
Interfaces solves these problems:
- makes it easier to re-use code by declaring a typed struct signature
- funcs that are passed interface values can also be passed values of structs with the interface signature
- interfaces can NOT be function receiver types
- interfaces are never instantiated
- interfaces are NOT generic types but they present a different approach to the same problem
   - go famously does not have generics
- interfaces are satisfied implicitly
   - no need to explicitly declare a link from your custom type to an interface
- interfaces are a contract to help us to manage types and reuse code
   - interfaces do not serve as unit tests for your types
   - if our custom types implementation of a func is broken, interfaces won't help

```go
// interfaces aren't explicitly inherited or extended 
// when an interface is declared, all other types in the program that match
// the signature are implicitly given that interface type
type bot interface {
   // interfaces declare func arguments and return types
   getGreeting(int, string) (int, string)
   // you can declare an interface that requires multiple funcs to satisfy membership
   getBotVersion() float64
   respondToUser(user) string
}
type Reader interface {
    Read(p []byte) (n int, err error)
}
type Closer interface {
    Close() error
}
// interfaces can be created by embedding other interfaces,
// thus requiring types to satisfy all child interfaces
type ReadCloser interface {
   Reader
   Closer
}
type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "Hi there!"
}
func (spanishBot) getGreeting() string {
	return "Hola!"
}
// this becomes available to 
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
```
### Reader Interface
The std lib [Reader](https://golang.org/pkg/io/#Reader) interface provides a common ouput `[]byte` for many disparate forms of input
   - http request body
   - text file
   - image file
   - user CLI input
   - 
```go
// the calling func passes a byte slice to Reader.Read() 
// our byte slice is then mutated by pointer
// we get back an int representing the length of the byte slice
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

### Writer Interface
The std lib [Writer](https://golang.org/pkg/io/#Writer) interface takes a byte slice and
 transforms it into one of many output formats.
   - http request body
   - text file
   - image file
   - user CLI input

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

## POINTERS
Go is a "pass-by-value" language.
When a function runs, it makes a new copy of the values passed in that exist within the scope of 
that function. In order for a function to mutate outer state you must use pointers.
Guess what? mutating outer state is bad practice. In the functional programming paradigm that is 
known as a side-effect. Don't be a jerk, update state via pure functions that take a value, do work, then return a new value. It's a good thing that Go encourages this style.

Pointer syntax:
- `&variable`: a pointer (memory address) to this variable's value
- `*pointer`: the stored value at this pointer address
```bash
# memory heap diagram
|---------|-----------------------------|
| address | value                       |
|---------|-----------------------------|
| 0001    | person{firstName: "Mike"..} |
|---------|-----------------------------|
```
- turn `address` into `value` with `*address`
- turn `value` into `address` with `&value`

### pointer examples
```go
func (p person) print() {
   fmt.Printf("%+p", p)
}
func main () {
   var me person
   me.firstName = "Mike"
   me.lastName = "Treacy"   
   
   // approach #1: unexpected struct instance behavior
   func (p person) updateName(newFirstName string) {
      // this is not updating the original struct of "me" in memory
      p.firstName = newFirstName
   }
   me.updateFirstName("Michael")
   me.print() // prints firstname: Mike unexpectedly
   
   // approach #2: using pointers to reference original instance
   // NOTE: the * symbol means different things in different contexts
   // NOTE: this is a type description:
   // the "*person" receiver type here is a pointer that points to a person
   // NOTE: this "*person" type will take EITHER a pointer to a person OR a person
   // this is a syntactic abstraction in go
   func (p *person) updateName(newFirstName string) {
      // this updates the original struct of "me" in memory
      // NOTE: this "*" is an operator:
      // it means we want to manipulate the value the pointer is referencing
      (*p).firstName = newFirstName
   }
   mePointer := &me
   mePointer.updateFirstName("Michael")
   me.print() // prints firstname: Michael as expected
   
   // approach #3: pointer shortcut syntax (because the print() receiver type is flexible)
    me.updateFirstName("Michael")
    me.print() // prints firstname: Michael as expected   
   }


```

## FUNCS
- receiver functions are go's version approximation of class methods()
- when the instance is not mutated in the receiver function, only the struct type is declared in the receiver
```go
type struct motif{
   name string 
}
func (m motif) changeName(n string) {
   m.name = n
}

func (motif) getInfo() string {
   return "A motif is a sequence of notes"   
}
```
## fmt
  - `Println`: To print a string to a line:
  - `Printf`: To do string interpolation (must explicitly add newlines as needed)
   - `fmt.Printf("%v", anyPrimitiveType)`: `%v` interpolates any type into string
   - `fmt.Printf("%+v", struct)`: `%+v` prints string representation of struct

## CONCURRENCY
structures in go for using concurrency 
 - channels
 - goroutine

### CHANNELS

### GOROUTINES



## How to Access Course Diagrams
All of the diagrams in this course can be downloaded and marked up by you!  Here's how:

Go to https://github.com/StephenGrider/GoCasts/tree/master/diagrams

Open the folder containing the set of diagrams you want to edit

Click on the ‘.xml’ file

Click the ‘raw’ button

Copy the URL

Go to https://www.draw.io/

On the ‘Save Diagrams To…’ window click ‘Decide later’ at the bottom

Click ‘File’ -> ‘Import From’ -> ‘URL’

Paste the link to the XML file

Tada!