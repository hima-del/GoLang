**What is an interface?**

* In Go, an interface is a set of method signatures. 
* When a type provides definition for all the methods in the interface, it is said to implement the interface.
* Interface specifies what methods a type should have and the type decides how to implement these methods.
* For example WashingMachine can be an interface with method signatures Cleaning() and Drying().
* Any type which provides definition for Cleaning() and Drying() methods is said to implement the WashingMachine interface.

```
package main

import (  
    "fmt"
)

type Worker interface {  
    Work()
}

type Person struct {  
    name string
}

func (p Person) Work() {  
    fmt.Println(p.name, "is working")
}

func describe(w Worker) {  
    fmt.Printf("Interface type %T value %v\n", w, w)
}

func main() {  
    p := Person{
        name: "Naveen",
    }
    var w Worker = p
    describe(w)
    w.Work()
}
//output
Interface type main.Person value {Naveen}  
Naveen is working  
```
* An interface can be thought of as being represented internally by a tuple `(type, value)`.
* `type` is the underlying concrete type of the interface and `value` holds the value of the concrete type.

**Empty interface**

* An interface that has zero methods is called an empty interface. 
* It is represented as interface{}. 
* Since the empty interface has zero methods, all types implement the empty interface.

```
package main

import (  
    "fmt"
)

func describe(i interface{}) {  
    fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {  
    s := "Hello World"
    describe(s)
    i := 55
    describe(i)
    strt := struct {
        name string
    }{
        name: "Naveen R",
    }
    describe(strt)
}
//output
Type = string, value = Hello World  
Type = int, value = 55  
Type = struct { name string }, value = {Naveen R}  
```
