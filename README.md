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

**Short variable declarations**

* Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
* Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

 
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
* %v	the value in a default format
* %q	a double-quoted string safely escaped with Go syntax
* bool:                    %t
* int, int8 etc.:          %d
* uint, uint8 etc.:        %d, %#x if printed with %#v
* float32, complex64, etc: %g
* string:                  %s
* chan:                    %p
* pointer:                 %p


**Basic types**

**Numbers**

**Integers**

**Unsigned integers[no negatives]**

```
1. uint8[unsigned integer with 8 bits]/byte (0-255)
2. uint16 (0-65535)
3. uint32 (0-4294967295)
4. uint64 (0-18446744073709551615)
```

**Signed integers**

```
1. int8 (-128-127)
2. int16 (-32768-32767)
3. int32/rune (-21477483648 to 2147483647)
4. int64 (-9223372036854775808 to 9223372036854775807)
```

**3 Machine dependent types**

```
1. uint (32 or 64 bits)
2. int (same as unit)
3. unitptr
```

**Floating point numbers**

**Floats**

```
1. float32
2. float64
```

**Complex (Imaginary Parts)**

```
1. complex64
2. complex128
```

**Strings**

```
* "hello world"
```

**Booleans**

```
1. true
2. false
```

**Variables**

* The var statement declares a list of variables; as in function argument lists, the type is last.
* A var statement can be at package or function level. We see both in this example.

```
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
```

**Variables with initializers**

* A var declaration can include initializers, one per variable.
* If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

```
package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
```

**Zero values**

* Variables declared without an explicit initial value are given their zero value.
* 0 for numeric types, false for the boolean type, and "" (the empty string) for strings.

**Constants**

* Constants are declared like variables, but with the const keyword.
* Constants can be character, string, boolean, or numeric values.
* Constants cannot be declared using the := syntax.

**Go >> and << operators**

* `<<` : (left shift): Takes two numbers, left shifts the bits of the first operand, the second operand decides the number of places to shift.
* `>>` :(right shift): Takes two numbers, right shifts the bits of the first operand, the second operand decides the number of places to shift.
* For example if A=60[0011 1100] and B =13[0000 1101] then,
  A << 2 will give 240 which is 1111 0000 and A >> 2 will give 15 which is 0000 1111

**Functions**

**What is a function?**

* A function is a block of code that performs a specific task. 
* A function takes a input, performs some calculations on the input and generates a output.

**Function declaration**

* The general syntax for declaring a function in go is

```
func functionname(parametername type) returntype {  
 //function body
}
```

* The function declaration starts with a func keyword followed by the functionname. 
* The parameters are specified between `(` and `)` followed by the returntype of the function. 
* The syntax for specifying a parameter is `parameter name` followed by the `type`. 
* Any number of parameters can be specified like `(parameter1 type, parameter2 type)`. 
* Then there is a block of code between `{` and `}` which is the body of the function.
* The parameters and return type are optional in a function.
* Hence the following syntax is also a valid function declaration.

```
func functionname() {  
}
```

**Sample function**

```
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

* If consecutive parameters are of the same type, we can avoid writing the type each time and it is enough to be written once at the end

```
package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

**Multiple return values**

* It is possible to return multiple values from a function
* If a function returns multiple return values then they must be specified between `(` and `)`

```
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```

**Named return values**

* It is possible to return named values from a function. 
* If a return value is named, it can be considered as being declared as a variable in the first line of the function.

```
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
```

* Note that the return statement in the function does not explicitly return any value
* x and y are specified in the function declaration as return values, they are automatically returned from the function when a return statement in encountered.

**Blank Identifier**

* _ is know as the blank identifier in Go.
* It can be used in place of any value of any type.

```
package main

import (  
    "fmt"
)

func rectProps(length, width float64) (float64, float64) {  
    var area = length * width
    var perimeter = (length + width) * 2
    return area, perimeter
}
func main() {  
    area, _ := rectProps(10.8, 5.6) // perimeter is discarded
    fmt.Printf("Area %f ", area)
}
```

**Loops**

* A loop statement is used to execute a block of code repeatedly.
* for is the only loop available in Go.
* Go doesn't have while or do while loops 

**for loop syntax**

```
for initialisation; condition; post {  
}
```

* All the three components namely initialisation, condition and post are optional in Go. 
* The initialisation statement will be executed only once.
* After the loop is initialised, the condition will be checked. 
* If the condition evaluates to true, the body of loop inside the { } will be executed followed by the post statement. 
* The post statement will be executed after each successful iteration of the loop.
* After the post statement is executed, the condition will be rechecked. 
* If it's true, the loop will continue executing, else the for loop terminates.
* The variables declared in a for loop are only available within the scope of the loop. Hence i cannot be accessed outside the body for loop.

```
package main

import (  
    "fmt"
)

func main() {  
    for i := 1; i <= 10; i++ {
        fmt.Printf(" %d",i)
    }
}
```

**break**

* The break statement is used to terminate the for loop abruptly before it finishes its normal execution and move the control to the line of code just after the for loop.

```
package main

import (  
    "fmt"
)

func main() {  
    for i := 1; i <= 10; i++ {
        if i > 5 {
            break //loop is terminated if i > 5
        }
        fmt.Printf("%d ", i)
    }
    fmt.Printf("\nline after for loop")
}
```

**continue**

* The continue statement is used to skip the current iteration of the for loop.
* All code present in a for loop after the continue statement will not be executed for the current iteration. 
* The loop will move on to the next iteration.

```
package main

import (  
    "fmt"
)

func main() {  
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            continue
        }
        fmt.Printf("%d ", i)
    }
}
```


**switch statement**

* A switch is a conditional statement that evaluates an expression and compares it against a list of possible matches and executes the corresponding block of code.

```
package main

import (  
    "fmt"
)

func main() {  
    finger := 4
    fmt.Printf("Finger %d is ", finger)
    switch finger {
    case 1:
        fmt.Println("Thumb")
    case 2:
        fmt.Println("Index")
    case 3:
        fmt.Println("Middle")
    case 4:
        fmt.Println("Ring")
    case 5:
        fmt.Println("Pinky")

    }
}
```

* In the above program `switch finger` in line no. 10, compares the value of `finger` with each of the `case` statements. 
* The cases are evaluated from top to bottom and the first `case` which matches the expression is executed.
* In this case, `finger` has a value of 4 and hence

**multiple expressions case**

```
package main

import (  
    "fmt"
)

func main() {  
    letter := "i"
    fmt.Printf("Letter %s is a ", letter)
    switch letter {
    case "a", "e", "i", "o", "u": //multiple expressions in case
        fmt.Println("vowel")
    default:
        fmt.Println("not a vowel")
    }
}
```

**expressionless switch**

* The expression in a switch is optional and it can be omitted.
* If the expression is omitted, the switch is considered to be switch true and each of the case expression is evaluated for truth and the corresponding block of code is           executed.

```
package main

import (  
    "fmt"
)

func main() {  
    num := 75
    switch { // expression is omitted
    case num >= 0 && num <= 50:
        fmt.Printf("%d is greater than 0 and less than 50", num)
    case num >= 51 && num <= 100:
        fmt.Printf("%d is greater than 51 and less than 100", num)
    case num >= 101:
        fmt.Printf("%d is greater than 100", num)
    }

}
```

**fallthrough**

* In Go, the control comes out of the switch statement immediately after a case is executed. 
* A fallthrough statement is used to transfer control to the first statement of the case that is present immediately after the case which has been executed.

```
package main

import (  
    "fmt"
)

func number() int {  
        num := 15 * 5 
        return num
}

func main() {

    switch num := number(); { //num is not a constant
    case num < 50:
        fmt.Printf("%d is lesser than 50\n", num)
        fallthrough
    case num < 100:
        fmt.Printf("%d is lesser than 100\n", num)
        fallthrough
    case num < 200:
        fmt.Printf("%d is lesser than 200", num)
    }

}
//output
75 is lesser than 100  
75 is lesser than 200  
```

* Switch and case expressions need not be only constants. 
* They can be evaluated at runtime too.
* In the program above `num` is initialized to the return value of the function `number()` in line no. 14.
* The control comes inside the switch and the cases are evaluated. 
* `case num < 100:` in line no. 18 is true and the program prints 75 is lesser than 100.
* The next statement is `fallthrough`.
* When `fallthrough` is encountered the control moves to the first statement of the next case and also prints 75 is lesser than 200.

**Fallthrough happens even when the case evaluates to false**

* Fallthrough will happen even when the case evaluates to false.

```
package main

import (  
    "fmt"
)

func main() {  
    switch num := 25; { 
    case num < 50:
        fmt.Printf("%d is lesser than 50\n", num)
        fallthrough
    case num > 100:
        fmt.Printf("%d is greater than 100\n", num)     
    }

}
//output
25 is lesser than 50  
25 is greater than 100  
```

* So be sure that you understand what you are doing when using fallthrough.

**fallthrough should be the last statement in a case. If it is present somewhere in the middle, the compiler will complain that `fallthrough statement out of place`.**
**One more thing is fallthrough cannot be used in the last case of a switch since there are no more cases to fallthrough. If fallthrough is present in the last case, it will       result in the following compilation error.`cannot fallthrough final case in switch`** 

**Breaking switch**

* The break statement can be used to terminate a switch early before it completes.

```
package main

import (  
    "fmt"
)

func main() {  
    switch num := -5; {
    case num < 50:
        if num < 0 {
            break
        }
        fmt.Printf("%d is lesser than 50\n", num)
        fallthrough
    case num < 100:
        fmt.Printf("%d is lesser than 100\n", num)
        fallthrough
    case num < 200:
        fmt.Printf("%d is lesser than 200", num)
    }

}
```

* In the above program `num` is -5. 
* When the control reaches the `if` statement in line no. 10, the condition is satisfied since `num < 0`. 
* The `break` statement terminates the `switch` before it completes and the program doesn't print anything

**Arrays**

* An array is a collection of elements that belong to the same type. 
* For example the collection of integers 5, 8, 9, 79, 76 form an array.
* Mixing values of different types, for example an array that contains both strings and integers is not allowed in Go.

```
package main

import (  
    "fmt"
)


func main() {  
    var a [3]int //int array with length 3
    fmt.Println(a)
}
```

* `var a [3]int` declares a integer array of length 3.
* All elements in an array are automatically assigned the zero value of the array type.
* In this case a is an integer array and hence all elements of a are assigned to 0, the zero value of `int`. 
* Running the above program will output `[0 0 0]`.
* The index of an array starts from 0 and ends at length - 1.


**short hand declaration for array**

```
package main 

import (  
    "fmt"
)

func main() {  
    a := [3]int{12, 78, 50} // short hand declaration to create array
    fmt.Println(a)
}
```

* You can even ignore the length of the array in the declaration and replace it with `...` and let the compiler find the length for you

```
package main

import (  
    "fmt"
)

func main() {  
    a := [...]int{12, 78, 50} // ... makes the compiler determine the length
    fmt.Println(a)
}
````

* The size of the array is a part of the type. 
* Hence [5]int and [25]int are distinct types. 
* Because of this, arrays cannot be resized

**Arrays are value types**

* Arrays in Go are value types and not reference types. 
* This means that when they are assigned to a new variable, a copy of the original array is assigned to the new variable.
* If changes are made to the new variable, it will not be reflected in the original array.

```
package main

import "fmt"

func main() {  
    a := [...]string{"USA", "China", "India", "Germany", "France"}
    b := a // a copy of a is assigned to b
    b[0] = "Singapore"
    fmt.Println("a is ", a)
    fmt.Println("b is ", b) 
}
//output
a is [USA China India Germany France]  
b is [Singapore China India Germany France]  
```

* Similarly when arrays are passed to functions as parameters, they are passed by value and the original array in unchanged.

```
package main

import "fmt"

func changeLocal(num [5]int) {  
    num[0] = 55
    fmt.Println("inside function ", num)

}
func main() {  
    num := [...]int{5, 6, 7, 8, 8}
    fmt.Println("before passing to function ", num)
    changeLocal(num) //num is passed by value
    fmt.Println("after passing to function ", num)
}
//output
before passing to function  [5 6 7 8 8]  
inside function  [55 6 7 8 8]  
after passing to function  [5 6 7 8 8]  
```

**Length of an array**

* The length of the array is found by passing the array as parameter to the len function.

```
package main

import "fmt"

func main() {  
    a := [...]float64{67.7, 89.8, 21, 78}
    fmt.Println("length of a is",len(a))

}
```







































