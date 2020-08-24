**Packages**

* Every Go program is made up of packages.
* Programs start running in package main.

```
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```

* This program is using the packages with import paths "fmt" and "math/rand".
* By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package     rand.

**func Intn**

```
func Intn(n int) int
```

* Intn returns, as an int, a non-negative pseudo-random number in [0,n)from the default Source. 
* It panics if n <= 0.

```
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("my favourite number is ", rand.Intn(10))
	fmt.Println("my favourite number is ", rand.Intn(10))
	fmt.Println("my favourite number is ", rand.Intn(10))
	fmt.Println("my favourite number is ", rand.Intn(100))
	fmt.Println("my favourite number is ", rand.Intn(100))
	fmt.Println("my favourite number is ", rand.Intn(100))
	fmt.Println("my favourite number is ", rand.Intn(10))
	fmt.Println("my favourite number is ", rand.Intn(1000))
	fmt.Println("my favourite number is ", rand.Intn(1000))
}

//output
my favourite number is  1
my favourite number is  7
my favourite number is  7
my favourite number is  59
my favourite number is  81
my favourite number is  18
my favourite number is  5
my favourite number is  540
my favourite number is  456
```
 
 **Println vs Printf in Golang**
 
 * Println and Printf are the functions to print the output in Golang. Both are present inside the package “fmt“. 
 * Both these functions provide output differently
 
 **Println**
 
 * Means “Print Line”. 
 * It helps us to print integers, strings, etc but inserts a new line at the end i.e. a line break is inserted.
 * It formats the string using the default formats for its operands. 
 * Println also inserts spaces between arguments.
 ```
 package main 
  
import "fmt"
  
func main() { 
  
    // The Println prints and 
    // adds a new line 
    s := "Sam"
    age := 25 
    fmt.Println("His name is", s) 
  
    fmt.Println("His age is", age, "years") 
} 

//output
His name is Sam
His age is 25 years
```

**Printf**

* Means “Print Formatter”. 
* It prints formatted strings. 
* It contains symbols in the string which you want to print and then arguments after it will replace those symbol points. 
* It does not insert a new line at the end like Println. For that, you’ll have to add “\n” in the end. 
* It formats the string according to a format specifier.
**Printf formats according to a specified format specifier but Println uses the default formats for its operands**

```
package main 
  
import "fmt"
  
func main() { 
    m, n, p := 15, 25, 40 
  
    fmt.Println( 
        "(m + n = p) :", m, "+", n, "=", p, 
    ) 
  
    fmt.Printf( 
        "(m + n = p) : %d + %d = %d\n", m, n, p, 
    ) 
} 

//Output
(m + n = p) : 15 + 25 = 40
(m + n = p) : 15 + 25 = 40
```

**The verbs**

* %T :	a Go-syntax representation of the type of the value
* %t	the word true or false
* bool:                    %t
* int, int8 etc.:          %d
* uint, uint8 etc.:        %d, %#x if printed with %#v
* float32, complex64, etc: %g
* string:                  %s
* chan:                    %p
* pointer:                 %p
