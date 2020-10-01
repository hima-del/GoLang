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

**Type assertion**

* Type assertion is used to extract the underlying value of the interface.
* `i.(T)` is the syntax which is used to get the underlying value of interface i whose concrete type is T


```
package main

import (  
    "fmt"
)

func assert(i interface{}) {  
    v, ok := i.(int)
    fmt.Println(v, ok)
}
func main() {  
    var s interface{} = 56
    assert(s)
    var i interface{} = "Steven Paul"
    assert(i)
}
//output
56 true  
0 false  
```

**Type switch**

* A type switch is used to compare the concrete type of an interface against multiple types specified in various case statements. 
* It is similar to switch case. 
* The only difference being the cases specify types and not values as in normal switch.

```
package main

import (  
    "fmt"
)

func findType(i interface{}) {  
    switch i.(type) {
    case string:
        fmt.Printf("I am a string and my value is %s\n", i.(string))
    case int:
        fmt.Printf("I am an int and my value is %d\n", i.(int))
    default:
        fmt.Printf("Unknown type\n")
    }
}
func main() {  
    findType("Naveen")
    findType(77)
    findType(89.98)
}
//output
I am a string and my value is Naveen  
I am an int and my value is 77  
Unknown type  
```

* It is also possible to compare a type to an interface.
* If we have a type and if that type implements an interface, it is possible to compare this type with the interface it implements.

```
package main

import "fmt"

type Describer interface {  
    Describe()
}
type Person struct {  
    name string
    age  int
}

func (p Person) Describe() {  
    fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType(i interface{}) {  
    switch v := i.(type) {
    case Describer:
        v.Describe()
    default:
        fmt.Printf("unknown type\n")
    }
}

func main() {  
    findType("Naveen")
    p := Person{
        name: "Naveen R",
        age:  25,
    }
    findType(p)
}
//outputunknown type  
Naveen R is 25 years old  
```

