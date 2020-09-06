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
* bool:                    %t
* int, int8 etc.:          %d
* uint, uint8 etc.:        %d, %#x if printed with %#v
* float32, complex64, etc: %g
* string:                  %s
* chan:                    %p
* pointer:                 %p


**String and slice of bytes (treated equivalently with these verbs):**

* %s	the uninterpreted bytes of the string or slice
* %q	a double-quoted string safely escaped with Go syntax
* %x	base 16, lower-case, two characters per byte
* %X	base 16, upper-case, two characters per byte


* %f     default width, default precision
* %9f    width 9, default precision
* %.2f   default width, precision 2
* %9.2f  width 9, precision 2
* %9.f   width 9, precision 0

* width is the minimum number of runes to output, padding the formatted form with spaces if necessary.
* For strings, byte slices and byte arrays, however, precision limits the length of the input to be formatted (not the size of the output), truncating if necessary

```
fmt.Printf("|%6d|%6d|\n", 12, 345)
fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
```

* Given 12.345 the format %6.3f prints 12.345 while %.3g prints 12.3.

**Iota basic example**

* The iota keyword represents successive integer constants 0, 1, 2,…
* It resets to 0 whenever the word const appears in the source code, and increments after each const specification.

```
const (
    C0 = iota
    C1 = iota
    C2 = iota
)
fmt.Println(C0, C1, C2) // "0 1 2"
```
* This can be simplified to

```
const (
	C0 = iota
	C1
	C2
)
```
* Here we rely on the fact that expressions are implicitly repeated in a parenthesized const declaration – this indicates a repetition of the preceding expression and its type.

**Start from one**

* To start a list of constants at 1 instead of 0, you can use iota in an arithmetic expression.

```
const (
    C1 = iota + 1
    C2
    C3
)
fmt.Println(C1, C2, C3) // "1 2 3"
```
**Skip value**

You can use the blank identifier to skip a value in a list of constants.

```
const (
    C1 = iota + 1
    _
    C3
    C4
)
fmt.Println(C1, C3, C4) // "1 3 4"
```

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

**Underlying types**

* Each type T has an underlying type
* If T is one of the predeclared boolean, numeric, or string types, or a type literal, the corresponding underlying type is T itself.
* Otherwise, T's underlying type is the underlying type of the type to which T refers in its type declaration.

```
type (
	A1 = string
	A2 = A1
)

type (
	B1 string
	B2 B1
	B3 []B1
	B4 B3
)
```
* The underlying type of string, A1, A2, B1, and B2 is string. 
* The underlying type of []B1, B3, and B4 is []B1.

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
* We define our function with parameters (if any)
* We call our function and pass in arguments (if any)

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

**Iterating arrays using range**

* The `for` loop can be used to iterate over elements of an array.

```
package main

import "fmt"

func main() {  
    a := [...]float64{67.7, 89.8, 21, 78}
    for i := 0; i < len(a); i++ { //looping from 0 to the length of the array
        fmt.Printf("%d th element of a is %.2f\n", i, a[i])
    }
}
```

* Go provides a better and concise way to iterate over an array by using the `range` form of the `for` loop. 
* `range` returns both the **index and the value at that index.**

```
package main

import "fmt"

func main() {  
    a := [...]float64{67.7, 89.8, 21, 78}
    sum := float64(0)
    for i, v := range a {//range returns both the index and value
        fmt.Printf("%d the element of a is %.2f\n", i, v)
        sum += v
    }
    fmt.Println("\nsum of all elements of a",sum)
}
//output
0 the element of a is 67.70  
1 the element of a is 89.80  
2 the element of a is 21.00  
3 the element of a is 78.00

sum of all elements of a 256.5  
```

* line no. 8 for i, v := range a of the above program is the range form of the for loop. 
* It will return both the index and the value at that index
* In case you want only the value and want to ignore the index, you can do this by replacing the index with the` _ `blank identifier.

```
for _, v := range a { //ignores index  
}
```

* The above for loop ignores the index. 
* Similarly the value can also be ignored.

**Multidimensional arrays**

```
package main

import (  
    "fmt"
)

func printarray(a [3][2]string) {  
    for _, v1 := range a {
        for _, v2 := range v1 {
            fmt.Printf("%s ", v2)
        }
        fmt.Printf("\n")
    }
}

func main() {  
    a := [3][2]string{
        {"lion", "tiger"},
        {"cat", "dog"},
        {"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
    }
    printarray(a)
    var b [3][2]string
    b[0][0] = "apple"
    b[0][1] = "samsung"
    b[1][0] = "microsoft"
    b[1][1] = "google"
    b[2][0] = "AT&T"
    b[2][1] = "T-Mobile"
    fmt.Printf("\n")
    printarray(b)
//output
lion tiger  
cat dog  
pigeon peacock 

apple samsung  
microsoft google  
AT&T T-Mobile  
```

* Although arrays seem to be flexible enough, they come with the restriction that they are of fixed length. 
* It is not possible to increase the length of an array. 
* This is were slices come into picture. 
* In fact in Go, slices are more common than conventional arrays.

**slices**

```
package main

import (  
    "fmt"
)

func main() {  
    a := [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4] //creates a slice from a[1] to a[3]
    fmt.Println(b)
}
```

* The syntax `a[start:end]` creates a slice from array `a` starting from index `start` to index `end - 1`.
* So in line no. 9 of the above program `a[1:4]` creates a `slice` representation of the array a starting from indexes 1 through 3. 
* Hence the slice b has values `[77 78 79].

```
package main

import (  
    "fmt"
)

func main() {  
    c := []int{6, 7, 8} //creates and array and returns a slice reference
    fmt.Println(c)
}
```

* In the above program in line no. 9, `c := []int{6, 7, 8}` creates an array with 3 integers and returns a slice reference which is stored in c.

**modifying a slice**

* A slice does not own any data of its own. 
* It is just a representation of the underlying array. 
* Any modifications done to the slice will be reflected in the underlying array.

```
package main

import (  
    "fmt"
)

func main() {  
    darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
    dslice := darr[2:5]
    fmt.Println("array before",darr)
    for i := range dslice {
        dslice[i]++
    }
    fmt.Println("array after",darr) 
}
//output
array before [57 89 90 82 100 78 67 69 59]  
array after [57 89 91 83 101 78 67 69 59]  
```

* When a number of slices share the same underlying array, the changes that each one makes will be reflected in the array.

```
package main

import (  
    "fmt"
)

func main() {  
    numa := [3]int{78, 79 ,80}
    nums1 := numa[:] //creates a slice which contains all elements of the array
    nums2 := numa[:]
    fmt.Println("array before change 1",numa)
    nums1[0] = 100
    fmt.Println("array after modification to slice nums1", numa)
    nums2[1] = 101
    fmt.Println("array after modification to slice nums2", numa)
}
//output
array before change 1 [78 79 80]  
array after modification to slice nums1 [100 79 80]  
array after modification to slice nums2 [100 101 80]  
```

* From the output it's clear that **when slices share the same array, the modifications which each one makes are reflected in the array**.

**length and capacity of a slice**

* The length of the slice is the **number of elements in the slice**.
* The capacity of the slice is the **number of elements in the underlying array starting from the index from which the slice is created**.

```
package main

import (  
    "fmt"
)

func main() {  
    fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
    fruitslice := fruitarray[1:3]
    fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) //length of fruitslice is 2 and capacity is 6
}
```

* A slice can be re-sliced upto its capacity. 
* Anything beyond that will cause the program to throw a run time error.

```
package main

import (  
    "fmt"
)

func main() {  
    fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
    fruitslice := fruitarray[1:3]
    fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
    fruitslice = fruitslice[:cap(fruitslice)] //re-slicing furitslice till its capacity
    fmt.Println("After re-slicing length is",len(fruitslice), "and capacity is",cap(fruitslice))
}
//output
length of slice 2 capacity 6  
After re-slicing length is 6 and capacity is 6 
```

**creating a slice using make**

* `func make([]T, len, cap) []T` can be used to create a slice by passing the type, length and capacity. 
* The capacity parameter is optional and defaults to the length. 
* The make function creates an array and returns a slice reference to it.

```
package main

import (  
    "fmt"
)

func main() {  
    i := make([]int, 5, 5)
    fmt.Println(i)
}
//output
[0 0 0 0 0]
```

**Appending to a slice**

* As we already know arrays are restricted to fixed length and their length cannot be increased. 
* Slices are dynamic and new elements can be appended to the slice using append function.

```
package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("length is %d and capacity is %d\n", len(a), cap(a))
	a = append(a, 7)
	fmt.Printf("length is %d and capacity is %d\n", len(a), cap(a))
	a = append(a, 8)
	fmt.Printf("length is %d and capacity is %d\n", len(a), cap(a))
	a = append(a, 8)
	fmt.Printf("length is %d and capacity is %d\n", len(a), cap(a))
	a = append(a, 10)
	fmt.Printf("length is %d and capacity is %d\n", len(a), cap(a))
	a = append(a, 11)
	fmt.Printf("length is %d and capacity is %d\n", len(a), cap(a))
	a = append(a, 12)
	fmt.Printf("length is %d and capacity is %d\n", len(a), cap(a))
	a = append(a, 13)
	fmt.Printf("length is %d and capacity is %d\n", len(a), cap(a))
}

//output
length is 6 and capacity is 6
length is 7 and capacity is 12
length is 8 and capacity is 12
length is 9 and capacity is 12
length is 10 and capacity is 12
length is 11 and capacity is 12
length is 12 and capacity is 12
length is 13 and capacity is 24
```

**Passing a slice to a function**

* Slices can be thought of as being represented internally by a structure type.

```
type slice struct {  
    Length        int
    Capacity      int
    ZerothElement *byte
}
```

* A slice contains the length, capacity and a pointer to the zeroth element of the array. 
* When a slice is passed to a function, even though it's passed by value, the pointer variable will refer to the same underlying array.
* Hence when a slice is passed to a function as parameter, changes made inside the function are visible outside the function too.

```
package main

import "fmt"

func changeval(a []int) {
	for i := range a {
		a[i]++
	}
}

func main() {
	b := []int{1, 2, 3, 4, 5}
	fmt.Printf("before function call is %d\n", b)
	changeval(b)
	fmt.Printf("after function call is %d", b)
}
//output
before function call is [1 2 3 4 5]
after function call is [2 3 4 5 6]
```

**multidimensional slices**

```
package main

import "fmt"

func printslice(a [][]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s", v2)
		}
		fmt.Printf("\n")
	}
}

func main() {
	b := [][]string{
		{"cat", "dog"},
		{"rat", "cow"},
		{"lion", "tiger"},
		{"pig", "fox"},
	}
	printslice(b)
}
//output
cat dog
rat cow
lion tiger
pig fox
```

**What is a variadic function?**

* Functions in general accept only a fixed number of arguments.
* A variadic function is a function that accepts a variable number of arguments. 
* If the last parameter of a function definition is prefixed by ellipsis `...`, then the function can accept any number of arguments for that parameter.
* Only the last parameter of a function can be variadic.

**Syntax**

```
func hello(a int, b ...int) {  
}
```

* In the above function, the parameter b is variadic since it's prefixed by ellipsis and it can accept any number of arguments. 
* This function can be called by using the following codes

```
hello(1, 2) //passing one argument "2" to b  
hello(5, 6, 7, 8, 9) //passing arguments "6, 7, 8 and 9" to b  
```

* It is also possible to pass zero arguments to a variadic function.

```
package main

import (  
    "fmt"
)

func find(num int, nums ...int) {  
    fmt.Printf("type of nums is %T\n", nums)
    found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            found = true
        }
    }
    if !found {
        fmt.Println(num, "not found in ", nums)
    }
    fmt.Printf("\n")
}
func main() {  
    find(89, 89, 90, 95)
    find(45, 56, 67, 45, 90, 109)
    find(78, 38, 56, 98)
    find(87)
}
```

**What is a map?**

* A map is a builtin type in Go that is used to store key-value pairs.

**How to create a map?**

* A map can be created by passing the type of key and value to the make function. 

```
make(map[type of key]type of value)  
```

```
package main

import (  
    "fmt"
)

func main() {  
    employeeSalary := make(map[string]int)
    fmt.Println(employeeSalary)
}
//output
map[]  
```

* The program above creates a `map` named employeeSalary with `string` `key` and `int` `value`.

**Adding items to a map**

```
package main

import (  
    "fmt"
)

func main() {  
    employeeSalary := make(map[string]int)
    employeeSalary["steve"] = 12000
    employeeSalary["jamie"] = 15000
    employeeSalary["mike"] = 9000
    fmt.Println("employeeSalary map contents:", employeeSalary)
}
//output
employeeSalary map contents: map[steve:12000 jamie:15000 mike:9000] 
```

**Retrieving value for a key from a map**

```
package main

import (  
    "fmt"
)

func main() {  
    employeeSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
        "mike": 9000,
    }
    employee := "jamie"
    salary := employeeSalary[employee]
    fmt.Println("Salary of", employee, "is", salary)
}
//output
Salary of jamie is 15000  
```

* What will happen if an element is not present? The map will return the zero value of the type of that element

**Checking if a key exists**

* In the above section we learned that when a key is not present, the zero value of the type will be returned. 
* This doesn't help when we want to find out whether the key actually exists in the map.

```
value, ok := map[key]  
```
* The above is the syntax to find out whether a particular key is present in a map. 
* If ok is true, then the key is present and its value is present in the variable value, else the key is absent.

```
package main

import (  
    "fmt"
)

func main() {  
    employeeSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    newEmp := "joe"
    value, ok := employeeSalary[newEmp]
    if ok == true {
        fmt.Println("Salary of", newEmp, "is", value)
        return
    }
    fmt.Println(newEmp, "not found")

}
//output
joe not found  
```

**Iterate over all elements in a map**

* The range form of the for loop is used to iterate over all elements of a map.

```
package main

import (  
    "fmt"
)

func main() {  
    employeeSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
        "mike":  9000,
    }
    fmt.Println("Contents of the map")
    for key, value := range employeeSalary {
        fmt.Printf("employeeSalary[%s] = %d\n", key, value)
    }

}
//output
Contents of the map  
employeeSalary[mike] = 9000  
employeeSalary[steve] = 12000  
employeeSalary[jamie] = 15000  
```

* One important fact is that the order of the retrieval of values from a map when using for range is not guaranteed to be the same for each execution of the program.
* It is also not the same as the order in which the elements were added to the map

**Deleting items from a map**

* `delete(map, key)` is the syntax to delete key from a map. 
* The delete function does not return any value.

```
package main

import (  
    "fmt"
)

func main() {  
    employeeSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,     
        "mike": 9000,
    }
    fmt.Println("map before deletion", employeeSalary)
    delete(employeeSalary, "steve")
    fmt.Println("map after deletion", employeeSalary)

}
```

**Length of the map**

* Length of the map can be determined using the len function.

```
package main

import (  
    "fmt"
)

func main() {  
    employeeSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    fmt.Println("length is", len(employeeSalary))

}

//output
length is 2 
```

**Maps are reference types**

* Similar to slices, maps are reference types. 
* When a map is assigned to a new variable, they both point to the same internal data structure. 
* Hence changes made in one will reflect in the other.

```
package main

import (  
    "fmt"
)

func main() {  
    employeeSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,     
        "mike": 9000,
    }
    fmt.Println("Original employee salary", employeeSalary)
    modified := employeeSalary
    modified["mike"] = 18000
    fmt.Println("Employee salary changed", employeeSalary)

}
//output

Original employee salary map[jamie:15000 mike:9000 steve:12000]  
Employee salary changed map[jamie:15000 mike:18000 steve:12000]  
```

**What is a String?**

* A string is a slice of bytes in Go. 
* Strings can be created by enclosing a set of characters inside double quotes `" "`.

**Accessing individual bytes of a string**

* Since a string is a slice of bytes, it's possible to access each byte of a string.

```
package main

import (  
    "fmt"
)

func printBytes(s string) {  
    fmt.Printf("Bytes: ")
    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
}

func main() {  
    name := "Hello World"
    fmt.Printf("String: %s\n", name)
    printBytes(name)
}
//output
String: Hello World  
Bytes: 48 65 6c 6c 6f 20 57 6f 72 6c 64  
```

* %s is the format specifier to print a string. 
* In line no. 16, the input string is printed. 
* In line no. 9 of the program above, `len(s)` returns the number of bytes in the string and we use a for loop to print those bytes in hexadecimal notation.
* %x is the format specifier for hexadecimal. 

**Rune**

* A rune is a builtin type in Go and it's the alias of int32. 
* Rune represents a Unicode code point in Go. 

```
package main

import (  
    "fmt"
)

func printBytes(s string) {  
    fmt.Printf("Bytes: ")
    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
}

func printChars(s string) {  
    fmt.Printf("Characters: ")
    runes := []rune(s)
    for i := 0; i < len(runes); i++ {
        fmt.Printf("%c ", runes[i])
    }
}

func main() {  
    name := "Hello World"
    fmt.Printf("String: %s\n", name)
    printChars(name)
    fmt.Printf("\n")
    printBytes(name)
    fmt.Printf("\n\n")
    name = "Señor"
    fmt.Printf("String: %s\n", name)
    printChars(name)
    fmt.Printf("\n")
    printBytes(name)
}
//output
String: Hello World  
Characters: H e l l o   W o r l d  
Bytes: 48 65 6c 6c 6f 20 57 6f 72 6c 64 

String: Señor  
Characters: S e ñ o r  
Bytes: 53 65 c3 b1 6f 72  
```

**Accessing individual runes using for range loop**

```
package main

import (  
    "fmt"
)

func charsAndBytePosition(s string) {  
    for index, rune := range s {
        fmt.Printf("%c starts at byte %d\n", rune, index)
    }
}

func main() {  
    name := "Señor"
    charsAndBytePosition(name)
}
//output
S starts at byte 0  
e starts at byte 1  
ñ starts at byte 2
o starts at byte 4  
r starts at byte 5 
```

**Creating a string from a slice of bytes**

```
package main

import (  
    "fmt"
)

func main() {  
    byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
    str := string(byteSlice)
    fmt.Println(str)
}

//output
Café 
```

**Creating a string from a slice of runes**

```
package main

import (  
    "fmt"
)

func main() {  
    runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
    str := string(runeSlice)
    fmt.Println(str)
}
//output
Señor 
```

**String length**

* The `RuneCountInString(s string) (n int)` function of the utf8 package can be used to find the length of the string. 
* This method takes a string as an argument and returns the number of runes in it.
* `len(s)` is used to find the number of bytes in the string and it doesn't return the string length. 
* Some Unicode characters have code points that occupy more than 1 byte.
* Using `len` to find out the length of those strings will return the incorrect string length.

```
package main

import (  
    "fmt"
    "unicode/utf8"
)

func main() {  
    word1 := "Señor"
    fmt.Printf("String: %s\n", word1)
    fmt.Printf("Length: %d\n", utf8.RuneCountInString(word1))
    fmt.Printf("Number of bytes: %d\n", len(word1))

    fmt.Printf("\n")
    word2 := "Pets"
    fmt.Printf("String: %s\n", word2)
    fmt.Printf("Length: %d\n", utf8.RuneCountInString(word2))
    fmt.Printf("Number of bytes: %d\n", len(word2))
}

//output

String: Señor  
Length: 5  
Number of bytes: 6

String: Pets  
Length: 4  
Number of bytes: 4  
```

**String comparison**

* The `==` operator is used to compare two strings for equality.
* If both the strings are equal, then the result is `true` else it's `false`.

```
package main

import (  
    "fmt"
)

func compareStrings(str1 string, str2 string) {  
    if str1 == str2 {
        fmt.Printf("%s and %s are equal\n", str1, str2)
        return
    }
    fmt.Printf("%s and %s are not equal\n", str1, str2)
}

func main() {  
    string1 := "Go"
    string2 := "Go"
    compareStrings(string1, string2)

    string3 := "hello"
    string4 := "world"
    compareStrings(string3, string4)

}
//output
Go and Go are equal  
hello and world are not equal 
```

**String concatenation**

```
package main

import (  
    "fmt"
)

func main() {  
    string1 := "Go"
    string2 := "is awesome"
    result := string1 + " " + string2
    fmt.Println(result)
}
```

```
package main

import (  
    "fmt"
)

func main() {  
    string1 := "Go"
    string2 := "is awesome"
    result := fmt.Sprintf("%s %s", string1, string2)
    fmt.Println(result)
}
```

**Strings are immutable**

* Strings are immutable in Go. 
* Once a string is created it's not possible to change it.
* To workaround this string immutability, strings are converted to a slice of runes. 
* Then that slice is mutated with whatever changes are needed and converted back to a new string.

```
package main

import (  
    "fmt"
)

func mutate(s []rune) string {  
    s[0] = 'a' 
    return string(s)
}
func main() {  
    h := "hello"
    fmt.Println(mutate([]rune(h)))
}
```

**In Go, '⌘' represents a single character (called a Rune), whereas "⌘" represents a string containing the character ⌘.**

**Double quotes**

* It is used to define a string.
* A string defined within double quotes will honor escaping characters.
* For, eg for when printing a string having `\n` there will be a new line printed. Similarly, `\t` will have tab printed.

**Back Quotes**

* It is also used to define a string. 
* A string encoded in back quotes is a raw literal string and doesn’t honor any kind of escaping.

**Single quotes**

* To declare either a byte or a rune we use single quotes.
* While declaring byte we have to specify the type. 
* If we don’t specify the type, then the default type is meant as a rune.
* A single quote will allow only one character.

**defer**

* A defer statement defers the execution of a function until the surrounding function returns.
* The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

```
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
//output
hello
world
```
**Pointers**

* A pointer is a variable which stores the memory address of another variable.
* If a variable b has value 156 and is stored at memory address 0x1040a124 and if the variable a holds the address of b => Now a is said to point to b.

**Declaring pointers**

* `*T` is the type of the pointer variable which points to a value of type T.

```
package main

import (  
    "fmt"
)

func main() {  
    b := 255
    var a *int = &b
    fmt.Printf("Type of a is %T\n", a)
    fmt.Println("address of b is", a)
}
```

* The `&` operator is used to get the address of a variable. 
* In line no. 9 of the above program we are assigning the address of b to a whose type is `*int`.
* Now a is said to point to b. 
* When we print the value in a, the address of b will be printed. 

**Creating pointers using the new function**

* Go also provides a handy function `new` to create pointers.
* The new function takes a type as argument and returns a pointer to a newly allocated zero value of the type passed as argument.

```
package main

import (  
    "fmt"
)

func main() {  
    size := new(int)
    fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
    *size = 85
    fmt.Println("New size value is", *size)
}
//output
Size value is 0, type is *int, address is 0x414020  
New size value is 85  
```

* In the above program, in line no. 8 we use the `new` function to create a pointer of type `int`.
* This function will return a pointer to a newly allocated zero value of the type `int`.
* The zero value of type `int` is 0.
* Hence size will be of type `*int` and will point to 0 i.e `*size` will be 0.

**Dereferencing a pointer**

* Dereferencing a pointer means accessing the value of the variable which the pointer points to. 
* `*a` is the syntax to deference a.

```
package main  
import (  
    "fmt"
)

func main() {  
    b := 255
    a := &b
    fmt.Println("address of b is", a)
    fmt.Println("value of b is", *a)
}
//output
address of b is 0x1040a124  
value of b is 255  
```

**Passing pointer to a function**

```
package main

import (  
    "fmt"
)

func change(val *int) {  
    *val = 55
}
func main() {  
    a := 58
    fmt.Println("value of a before function call is",a)
    b := &a
    change(b)
    fmt.Println("value of a after function call is", a)
}
//output
value of a before function call is 58  
value of a after function call is 55  
```

**Returning pointer from a function**

```
package main

import (  
    "fmt"
)

func hello() *int {  
    i := 5
    return &i
}
func main() {  
    d := hello()
    fmt.Println("Value of d", *d)
}
//output
Value of d 5  
```

**Structs**

* A struct is a user-defined type that represents a collection of fields.

```
package main

import (  
    "fmt"
)

type Employee struct {  
    firstName string
    lastName  string
    age       int
    salary    int
}

func main() {

    //creating struct specifying field names
    emp1 := Employee{
        firstName: "Sam",
        age:       25,
        salary:    500,
        lastName:  "Anderson",
    }

    //creating struct without specifying field names
    emp2 := Employee{"Thomas", "Paul", 29, 800}

    fmt.Println("Employee 1", emp1)
    fmt.Println("Employee 2", emp2)
}
//output
Employee 1 {Sam Anderson 25 500}  
Employee 2 {Thomas Paul 29 800}  
```

**Creating anonymous structs**

* It is possible to declare structs without creating a new data type. 
* These types of structs are called anonymous structs.

```
package main

import (  
    "fmt"
)

func main() {  
    emp3 := struct {
        firstName string
        lastName  string
        age       int
        salary    int
    }{
        firstName: "Andreah",
        lastName:  "Nikola",
        age:       31,
        salary:    5000,
    }

    fmt.Println("Employee 3", emp3)
}
//output
Employee 3 {Andreah Nikola 31 5000}  
```

* In line no 8. of the above program, an anonymous struct variable emp3 is defined. 
* This struct is called anonymous because it only creates a new struct variable emp3 and does not define any new struct type like named structs.

**Accessing individual fields of a struct**

* The dot `.` operator is used to access the individual fields of a struct.

```
package main

import (  
    "fmt"
)

type Employee struct {  
    firstName string
    lastName  string
    age       int
    salary    int
}

func main() {  
    emp6 := Employee{
        firstName: "Sam",
        lastName:  "Anderson",
        age:       55,
        salary:    6000,
    }
    fmt.Println("First Name:", emp6.firstName)
    fmt.Println("Last Name:", emp6.lastName)
    fmt.Println("Age:", emp6.age)
    fmt.Printf("Salary: $%d\n", emp6.salary)
    emp6.salary = 6500
    fmt.Printf("New Salary: $%d", emp6.salary)
}
```

**Pointers to a struct**

* It is also possible to create pointers to a struct.

```
package main

import (  
    "fmt"
)

type Employee struct {  
    firstName string
    lastName  string
    age       int
    salary    int
}

func main() {  
    emp8 := &Employee{
        firstName: "Sam",
        lastName:  "Anderson",
        age:       55,
        salary:    6000,
    }
    fmt.Println("First Name:", emp8.firstName)
    fmt.Println("Age:", emp8.age)
}
//output
First Name: Sam  
Age: 55  
```

**Anonymous fields**

* It is possible to create structs with fields that contain only a type without the field name. 
* These kinds of fields are called anonymous fields.
* The snippet below creates a struct Person which has two anonymous fields string and int

```
type Person struct {  
    string
    int
}
```

* Even though anonymous fields do not have an explicit name, by default the name of an anonymous field is the name of its type.
* For example in the case of the Person struct above, although the fields are anonymous, by default they take the name of the type of the fields. 
* So Person struct has 2 fields with name string and int.

```
package main

import (  
    "fmt"
)

type Person struct {  
    string
    int
}

func main() {  
    p1 := Person{
        string: "naveen",
        int:    50,
    }
    fmt.Println(p1.string)
    fmt.Println(p1.int)
}
//output
naveen  
50  
```

**nested structs**

```
package main

import (  
    "fmt"
)

type Address struct {  
    city  string
    state string
}

type Person struct {  
    name    string
    age     int
    address Address
}

func main() {  
    p := Person{
        name: "Naveen",
        age:  50,
        address: Address{
            city:  "Chicago",
            state: "Illinois",
        },
    }

    fmt.Println("Name:", p.name)
    fmt.Println("Age:", p.age)
    fmt.Println("City:", p.address.city)
    fmt.Println("State:", p.address.state)
}

```

**Promoted fields**

* Fields that belong to an anonymous struct field in a struct are called promoted fields since they can be accessed as if they belong to the struct which holds the anonymous     struct field.

```
package main

import (  
    "fmt"
)

type Address struct {  
    city  string
    state string
}
type Person struct {  
    name string
    age  int
    Address
}

func main() {  
    p := Person{
        name: "Naveen",
        age:  50,
        Address: Address{
            city:  "Chicago",
            state: "Illinois",
        },
    }

    fmt.Println("Name:", p.name)
    fmt.Println("Age:", p.age)
    fmt.Println("City:", p.city)   //city is promoted field
    fmt.Println("State:", p.state) //state is promoted field
}
```

**Structs Equality**

* Structs are value types and are comparable if each of their fields are comparable. 
* Two struct variables are considered equal if their corresponding fields are equal.


```
package main

import (  
    "fmt"
)

type name struct {  
    firstName string
    lastName  string
}

func main() {  
    name1 := name{
        firstName: "Steve",
        lastName:  "Jobs",
    }
    name2 := name{
        firstName: "Steve",
        lastName:  "Jobs",
    }
    if name1 == name2 {
        fmt.Println("name1 and name2 are equal")
    } else {
        fmt.Println("name1 and name2 are not equal")
    }

    name3 := name{
        firstName: "Steve",
        lastName:  "Jobs",
    }
    name4 := name{
        firstName: "Steve",
    }

    if name3 == name4 {
        fmt.Println("name3 and name4 are equal")
    } else {
        fmt.Println("name3 and name4 are not equal")
    }
}
//output
name1 and name2 are equal  
name3 and name4 are not equal  
```

**type WaitGroup**

* A WaitGroup waits for a collection of goroutines to finish. 
* The main goroutine calls `Add` to set the number of goroutines to wait for. 
* Then each of the goroutines runs and calls `Done` when finished. 
* `Wait` can be used to block until all goroutines have finished.

**func (*WaitGroup) Add**

```
func (wg *WaitGroup) Add(delta int)
```
* Add adds delta, which may be negative, to the WaitGroup counter. 
* If the counter becomes zero, all goroutines blocked on Wait are released.

**func (*WaitGroup) Done**

```
func (wg *WaitGroup) Done()
```

* Done decrements the WaitGroup counter by one.

**func (*WaitGroup) Wait**

```
func (wg *WaitGroup) Wait()
```

* Wait blocks until the WaitGroup counter is zero.

```
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
//output

OS       windows
ARCH     amd64    
CPUs     4        
Goroutines       1
bar: 0
bar: 1
bar: 2
bar: 3
bar: 4
bar: 5
bar: 6
bar: 7
bar: 8
bar: 9
CPUs     4
Goroutines       2
foo: 0
foo: 1
foo: 2
foo: 3
foo: 4
foo: 5
foo: 6
foo: 7
foo: 8
foo: 9
```


















