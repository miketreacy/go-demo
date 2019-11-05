# GOLANG NOTES
## ENV SETUP
### CLI
 - `go build <go_files>`: compiles module to an executable binary file
   - built executable will be named after first file in args
 - `go run <go_file>`: builds AND executes program
 - `go run -race <go_file`>: builds and executes program and detects race conditions!
 - `./<executable_file>`: runs the built executable for your go package
 - `go test`: runs any tests in directory
 - 

 ### PACKAGES
 Packages may contain multiple files.
 Each file in a package must declare what package they belong to on the first line:
 ```go
 package main
 ```

 #### Types of Packages
 1. Executable: generates a file that we can run
    - the name `main` indicates an executable. `main` is the only package name that builds to an executable
    - must also have a function inside called `main`
 2. Reusable: Code used as helpers. Good place to put reusable logic

 Files in the same package fo not have to be imported into each other

 #### fmt package
 Standard lib package for formatting I/O 
  - `Println`: To print a string to a line:
  - `Printf`: To do string interpolation (must explicitly add newlines as needed)
   - `fmt.Printf("%v", anyPrimitiveType)`: `%v` interpolates any type into string
   - `fmt.Printf("%+v", struct)`: `%+v` prints string representation of struct

### GO WORKSPACE
- one folder, any name, any location
   - bin:
      - compiled binary lives here
   - pkg
      - archives of compiled binary dependencies 
   - src
      - package code

### PKG MANAGEMENT

#### GO MODULES

## VARIABLES
### VAR
types are automatically inferred from the assignment value
```go
// These two declarations are equivalent - the type string is inferred by the value
var name string = "Mike"
name := "Mike"
// The := operator is only used for initialization.
// To reassign an existing variable you must use =
```

### CONST and IOTA
iota: a universal numeric incrementer starting at 0
- use only with `const` declarations
- can be used for making an enum
- you can do bit-shifting with iota

```go
// << and >> are bit-shifting operators
// they move bytes within binary numbers
// using bit-shifting with iota to model byte size
const (
   // dispense the initial iota value of zero
   _ = iota
   // kb = 1024
   kb = 1 << (iota * 10)
   mb = 1 << (iota * 10)
   gb = 1 << (iota * 10)
)

func main () {
   fmt.Printf("\n%d\t\t\t%b\n", kb, kb)
   fmt.Printf("\n%d\t\t\t%b\n", mb, mb)
   fmt.Printf("\n%d\t\t%b\n", gb, gb)
}
// making an enum with const and iota
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
### FUNCS
- receiver functions are go's version approximation of class methods()
- when the instance is not mutated in the receiver function, only the struct type is declared in the receiver
- basic syntax is same as js
- all return value types must be declared before the open bracket
 ```go
 func getNote() string {
    return "C#"
 }
 ```

#### defer
the `defer` directive defers the modified functions execution until the surrounding function exits

```go
// the defer keyword
func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

type struct motif{
   name string 
}
func (m motif) changeName(n string) {
   m.name = n
}

func (motif) getInfo() string {
   return "A motif is a sequence of notes"   
}

// variadic functions

func main() {
   total := add(4, 76, 3, 8, 45)
   fmt.Println(total)
   
   xi := []int{2,3,4,5,6,7}
   // the trailing ellipses spreads a slice into func args
   sliceSpread := add(xi...)
   fmt.Println(sliceSpread)
	
}
// the leading ellipses allow an unlimited number of params of a given type
func add(nums ...int) int {
   result := 0
   // nums is a slice of ints
	for _, num := range nums {
		result += num
	}
	return result
}
```
### ARRAYS v SLICES
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
### STRUCTS
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
### MAPS
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

#### Maps vs Structs
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
### INTERFACES
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
#### Reader Interface
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

#### Writer Interface
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
### SORTS
- to sort slices of primitive types:
   - `sort.Ints()`
   - `sort.Strings()`
```go
import (
	"fmt"
	"sort"
)

func main() {
   xi := []int{6, 3, 9, 2, 11, 79, 1}
   fmt.Println(xi)
   // sort package sorts a slice of values in place (no return value)
	sort.Ints(xi)
	fmt.Println(xi)

	xs := []string{"Larry", "Moe", "Curley"}
   fmt.Println(xs)
   // sort package sorts a slice of values in place (no return value)
	sort.Strings(xs)
	fmt.Println(xs)

}

// CUSTOM COLLECTION SORTS
import (
	"fmt"
	"sort"
)

type Note struct {
	Name     string
	Value    int
	Duration int
}
// defines string representation of Note by satisfying native struct interface
func (n Note) String() string {
	return fmt.Sprintf("%d:%d:%s", n.Value, n.Duration, n.Name)
}

// CUSTOM STRUCT SORT
// ByValue implicitly implements the built-in sort.Interface https://godoc.org/sort#Interface
// by receiving methods Len, Swap, Less
// explicitly implementing the sort interface for your collection is more performant because
// the native sort.Sort() has to use reflection to get the length and order of the collection
type ByValue []Note

func (v ByValue) Len() int           { return len(v) }
func (v ByValue) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v ByValue) Less(i, j int) bool { return v[i].Value < v[j].Value }

func main() {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	fmt.Println("int slice: unsorted\t\t", s)
	sort.Ints(s)
	fmt.Println("int slice: sorted\t\t", s)
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println("int slice: reverse sorted\t", s)
	fmt.Println()

	n1 := Note{"D#", 52, 16}
	n2 := Note{"F#", 43, 16}
	n3 := Note{"", 0, 24}
	n4 := Note{"B", 60, 4}
	n5 := Note{"D#", 52, 8}
	n6 := Note{"A#", 46, 4}
	n7 := Note{"D#", 52, 32}

	notes := []Note{n1, n2, n3, n4, n5, n6, n7}
	fmt.Println("notes unsorted\t\t\t", notes)

	sort.Sort(ByValue(notes))
	fmt.Println("notes sorted by value\t\t", notes)

	// stable sort (keeps original order of equal values)
	sort.SliceStable(notes, func(i, j int) bool { return notes[i].Value < notes[j].Value })
	fmt.Println("notes stable sorted by value\t", notes)

}
```
### POINTERS
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

#### pointer examples
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


### ERRORS
   - `Go` does not have exceptions
   - It is `Go` convention to return an error value as the last value in the tuple that a func returns
      - if there is no error, the error value will be `nil`
   - Errors in `Go` are children of a common `Error` interface


#### THE ERROR INTERFACE
   - the `Error` interface is very simple
   - any type that implements `func Error() string` satisfies the error interface
```go
type error interface {
   Error() string
}
```
#### TO RETURN AN ERROR
   - `errors.New()`
   - `fmt.Errorf()`
   - Create your own custom error type
#### CHECKING FOR ERRORS
   - use `if` statement to check if err is not nil
```go
// define custom error
var ErrNotFound = errors.New("Data not found")

func getData(num int)(int, err) {
   if num < 5 {
      // return custom error type
      return nil, ErrNotFound
   } else {
      return num, nil
   }
}

func main () {
   // check for any error
   if result, err := getData(); err != nil {
      fmt.Println("ERROR:", err)
   } else {
      fmt.Println("RESULT:", result)
   }
   // check for instance of specific error
   result, err := getData(3)
   if err == ErrNotFound {
      // handle specific error
      fmt.Println("ERROR:", err)
   } else {
      fmt.Println("RESULT:", result)
   }
}
```
#### PANIC
   - `panic()` is the closest `Go` comes to raising an `exception`
   - `panic()` takes a message and can ONLY be called from a deferred function
   - you can handle a `panic` in your program with `recover`
   - if you don't handle the `panic`, the program will crash and print the call stack
   - the primary use-case for `panic/recover` is handling errors across multiple concurrent routines
```go
func main() {
   defer func() {
      if err := recover(); err != nil {
         fmt.Println("A panic recovered", err)
      }
   }()
}
```

## CONCURRENCY
"Concurrency is not parallelism"
   - CONCURRENCY: multiple threads executing code
      - Go routines are executed concurrently on a single CPU core  
   - PARALLELISM: multiple threads executing code at the exact same time (requires multiple CPU cores)
      - Go routines are executed in parallel on multiple CPU cores   
"Do not communicate by sharing memory; instead, share memory by communicating"
   - Do not communicate by locking variables between threads
   - Communicate by sending values from one concurrent piece of code to another

### GOROUTINES
- go routines:
   - lightweight "threads"
      - in reality they are not real parallel threads
      - they are pieces of code that get scheduled among multiple os threads that makes execution concurrent
      - go routines can share an os thread with other go routines
   - an engine that executes code in a given process.
   - when a goroutine hits a blocking call, then the goroutine has to wait
   - while waiting, it passes control flow back to the main goroutine
   - go routines run concurrently
   - Go routine completion is not deterministic

#### Theory of Go Routines
Hardware Environments:
  1. One CPU Core:
   - Go's default behavior is to run on one CPU Core
   - When running on one Core, Go routines do not run at the same time
   - the Go Scheduler runs one routine until it finishes or makes a blocking call, then executes the next go routine
  2. Multiple CPU Cores:
   -  Go Scheduler runs one thread on each "logical" core

#### Go Routine Gotchas
The main routine is the parent routine that decides when our program exits.
The main routine cursor doesn't wait for child routines to return.
Never reference a pointer variable from multiple go routines!
   - don't close over outer variables from a go routine
```go

for _, link := range links {
   // adding the "go" keyword in front of a function call 
   // creates a new child go routine process to execute the code in
   go checkLink(link)
}

```
#### CONTEXT
   - good for http request properties
   - in go servers, each incoming request is handled in its own goroutine
   - context is a convenient way to handle request-scoped values
   - when launching multiple goroutines from a request it is possible to leak memory
   - context is good for killing all child goroutines when the parent thread closes to free up memory right away
   
### SYNCHRONICITY PRIMITIVES
#### WAITGROUPS
```go
var wg sync.WaitGroup

func main() {
	fmt.Println("Go Routines\t", runtime.NumGoroutine())
	fmt.Println("CPUs\t\t", runtime.NumCPU())

	// Add one item to WaitGroup
	wg.Add(1)
	go foo()
	bar()
	// this line will block until everything we added to the WaitGroup
	// calls Done()
	wg.Wait()
}

func foo() {
	fmt.Println("foo")
	// WaitGroup can stop waiting now
	wg.Done()
}

func bar() {
	fmt.Println("bar")
}
```
### RACE CONDITIONS AND MUTEXES

#### Race Conditions
   - can happen when you have multiple go routines reading from and writing to a shared variable
   - this is an anti-pattern in go
```go
// NOTE: Anti-Pattern Alert
// create a race condition among go routines
// reading from and writing to a shared variable
func main() {	
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Go Routines:", runtime.NumGoroutine())

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	counter := 0

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			// yields control back to parent routine
			runtime.Gosched()
			// yielding before writing to shared variable
			// creates race condition
			v++
			counter = v
			wg.Done()

		}()
		fmt.Println("Go Routines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Go Routines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)

}
```
#### MUTEX
   - `sync.Mutex` alows you to prevent race conditions when sharing variables across go routines
   - use `mu.Lock()` and `mu.Unlock()` to control access to shared variables in scope of a goroutine
```go
// Using a mutex to fix the race condition above
func main() {	
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Go Routines:", runtime.NumGoroutine())

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	counter := 0
	// creates a mutex
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			// locks access to all shared (closure) variables in
			// go routine
			mu.Lock()
			v := counter
			// yields control back to parent routine
			runtime.Gosched()
			// yielding before writing to shared variable
			// creates race condition
			v++
			counter = v
			mu.Unlock()
			wg.Done()

		}()
		fmt.Println("Go Routines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Go Routines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
```
### CHANNELS
"Share memory by communicating"
Channels are used to communicate between multiple running go routines
- channels are the ONLY WAY to communicate between go routines
- any data sent to the channel is sent to all running go routines
- channels are typed - all the messages sent to a channel must be of the same type

Receiving messages from a channel is a blocking operation for the main routine.
- channel receivers will wait until they receive a value from the channel
- if you have more channel receivers than senders, your program will hang

Directional Channels
   - you can specify if a channel is meant to only send or receive values.
   - this specificity increases the type-safety of the program.

#### Sending Data with Channels
```go
// create a bidirectional channel
c := make(chan int)
// create a send-only channel
sc := make(chan<- int)
// create a receive-only channel
rc := make(<-chan int)

// send the value '5' into this channel
c <- 5

// wait for a value to be sent into the channel
// when we get a value, assign it to 'myNumber'
myNumber <- c

// wait for a value to be sent into the channel
// when we get one, log it out immediately

fmt.Println(<- c)

// The boolean variable ok returned by a receive operator indicates whether the received value was sent on the channel (true) or is a zero value returned because the channel is closed and empty (false).
val, ok := <-c
```

```go
func main() {
	xs := []string{" Hello ", " world ", " come ", " code ", " with ", " me!"}
	c := make(chan bool)

	for i, s := range xs {
      // cast index int to time.Duration to make each successive
      // func call wait one second longer than the last
      // invoke all functions now on their own go routines
      // we'll trigger them later through the channel
		go waitAndSay(c, s, time.Duration(i))
	}

	// we send a signal to c in order to allow waitAndSay to continue
	for i := 0; i < len(xs); i++ {
		c <- true
	}
	// wait for a message from each go routine you created
	// before we exit the main routine	
	for i := 0; i < len(xs); i++ {
      // we don't need to do anything with this value
		<-c
	}
}

func waitAndSay(c chan bool, s string, d time.Duration) {
   // when we get the bool signal from the channel
   // check if signal is true, then proceed
   if b := <-c; b {
		if d != 0 {
			time.Sleep(d * time.Second)
		}
		fmt.Printf(s)
   }
   // send message back through channel to caller in main routine
	c <- true
}
```
#### BUFFERED CHANNELS
   - only use buffered channels if absolutely necessary, favor unbuffered channels
   - a collection of individual stacked channels contained in one buffer
   - senders block ONLY if the buffer is full
   - receivers block ONLY when the buffer is empty
   - to declare a buffered channel:
   ```go
   // this creates a stack of 100 buffered int channels
   ch := make(chan int, 100)
   ```
#### CHANNELS: RANGE AND CLOSE
   - `range` can be used to receive values from a channel inside a `for` loop
   - `close` is used to indicate the channel has retired
   - these keywords are useful for listening to a speficic amount of messages
   ```go
   func main() {
      c := make(chan string)
      go SayHelloMultipleTimes(c, 5)
      
      // do this for every message on channel c
      for s := range c {
         fmt.Println(s)
      }
      
      // channel receiver has a second tuple arg 
      // bool representing channel status      
      v, ok := <-c
      // the above line is blocking as long as the channel is open
      // this line only runs when channel is closed
      fmt.Println("Close channel?", !ok, " value ", v)

   }
   func SayHelloMultipleTimes(c chan string, n int) {
      for i := 0; i <= n; i++ {
         c <- "Hello"
      }
      // close the channel after sending message n times
      close(c)
   }
   ```
#### CHANNELS: SELECT STATEMENTS
   - select statements allow our code to wait on multiple channels at the same time
   - select blocks until one channel is ready
   - if multiple channels are ready, select picks one at random
   - syntax is similar to `switch` statement
   - if `select` has a `default` case, then `select` won't block

   ```go
   select {
      case value1 := <- channel1:
      // do stuff
      case channel 2 <- value2:
      // do stuff
      default:
      fmt.Println("Too slow!")
   }
  
   func main() {
      rand.Seed(time.Now().UnixNano())

      c1 := make(chan int)
      c2 := make(chan int)

      name := "Mike"
      // query two db servers simultaneously for Mike's ID
      go findID(name, "Server 1", c1)
      go findID(name, "Server 2", c2)

      // the select statement blocks until one of the channels returns
      select {
      case id := <-c1:
         fmt.Println(name, "has an id of", id, "found in Server 1")

      case id := <-c2:
         fmt.Println(name, "has an id of", id, "found in Server 2")
      // time.After() creates a channel that returns after given time
      case <-time.After(1 * time.Millisecond):
         fmt.Println("Search timed out!!")
         //default:
         //fmt.Println("Too slow!")
      }

   }

   var idMapping = map[string]int{
      "Mike":  7,
      "Asher": 22,
      "Karen": 13,
   }

   func findID(name, server string, c chan int) {
      // simulate searching
      time.Sleep(time.Duration(rand.Intn(50)) * time.Minute)

      // return security clearance from map
      c <- idMapping[name]
   }
   ```