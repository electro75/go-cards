## Learning Go Basics.

This repo contains the code written while going through [this course](https://www.udemy.com/course/go-the-complete-developers-guide/).

Some of the learnings are jotted down below.

### Package main

- This package is imported into your main file and tells the compiler where the start of the program is.
- needs to contain the `func main() {}`
- Following commands can be used to run code:
```
go run yourfilename/s.go // runs code after compilation
go build yourfilename/s.go // makes an executable file. Does not execute it.
```


### Variable declarations and Types

- Go is a statically typed language. A type of variable, once set, cannot be changed.
- A variable can be initialised in the following ways:

```
var abc string = "Hello"
abc := "Hello"  //  := is used to a declare and initialise a new variable.
```
- Functions that return a variable need to specify the return type at declaration:
```
func returnsAString string () {

    return "Hello"
}
```
- There are two types of data structures for holding values with same data types:
    - Arrays : Fixed length of data.
    - Slices : Arrays that can shrink or grow.
        * Example of slice declaration:
        ```
        fruits := []string{"apple", "banana"}
        ```
        * Adding a new element to a slice:
        ```
        fruits = append(fruits, "kiwi"); // append does not modify the existing array, but returns a new one
        ```
        * Iteration over a slice:
        ```
        for i, fruit := range fruits {
            fmt.Println(i, fruit)
        }
        ```
- Creating new types:
    - Similar to inheritance in OOP.
    - Example: 
    ```
    type basket []string  // basket will have the properties of slice of strings.
    ```
    -
