# GOLANG NOTES

## CLI
 - go build <file>: compiles module to an executable binary file
 - go run <file>: builds AND executes program
 - go 

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
types are autmatically inferred from the assignment value
```go
// These two declarations are equivalent - the type string is inferred by the value
var name string = "Mike"
name := "Mike"
// The := operator is only used for initialization.
// To reassign an existing variable you must use =
```
## TYPES
type default zero values in parens
  - int (0)
  - float64 (0)
  - string ("")
  - bool (false)
  - array
   - fixed-length list of values
   - every value in array must be of same type
  - slice
   - an array that can grow or shrink
  - map
  - struct (nil)
  - func
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


## fmt
  - `Println`: To print a string to a line:
  - `Printf`: To do string interpolation (must explicitly add newlines as needed)
   - `fmt.Printf("%v", anyPrimitiveType)`: `%v` interpolates any type into string
   - `fmt.Printf("%+v", struct)`: `%+v` prints string representation of struct


 ##
 How to Access Course Diagrams
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