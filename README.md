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
        * Returning a part of the slice:
        ```
        fruits[startIndex:endNotIncluded]
        // if LHS of the : is ommitted the elements are returned from the start
        // if RHS of the : is ommitted the elements till the end are returned
        ```

- Creating new types:
    - Similar to inheritance in OOP.
    - Example: 
    ```
    type basket []string  // basket will have the properties of slice of strings.
    ```
    - Functions can be associated with the new types, for example:
    ```
    func(b basket) print() {            // receiver function. 'b' is the argument with type basket
        for i, fruit = range b {        // first letter from type declared is used as a convention.
            fmt.Println(i, fruit)
        }
    }

    fruits := basket{"apple","banana","kiwi"}
    fruits.print()           // prints out the list of fruits initialised above
    ```

- _Tip: If a variable is not being used, just replace it with an  _(underscore)_ .

- Functions can return multiple values.
```
func funcName() (type1,type2) {
    return valueOfType1, valueOfType2
}

value1, value2 := funcName()
```

- Functions setup with receivers, receive only the copy of the value that they are called with.
    - they do not mutate the actual value.
    - if the value is to be changed, they should be called with pointers

- **Slice Bytes** are how Golang saves various types of data to a file.
    - Type conversion is done as follows:
    ```
    []byte("Hi there")
    //resulting type(value to be converted)
    ```

- `strings` is a package in golang that provides helper methods to play around with strings.

- `ioutils` is a package used for input output functions in golang

- `shuffle()` functionality
    - time.Now().UnixNano() generates a different int64 number everytime the function is run
    - this value is used to create the "seed" for create a NewSource object
    - the NewSource is used as the basis for the random number generator