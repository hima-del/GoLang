**Understanding templates**

* A template allows us to create one document and then merge data with it.
* We are learning about templates so that we can create one document, a web page, and then merge customized data to that page.
* Web templates allow us to serve personalized results to users.
* Think of Facebook - you log into your main page and see results tailored for you. 
* That main page was created once. It is a template. However, for each user, that template gets populated with data specific to that user.
* Another common exposure to templates that most of us get every day - junk mail.
* A company creates a piece of mail to send to everyone, and then they merge data with that template to customize the mailing for each individual.

**Template Example - Merged With Data**

```
Dear Mr. Jones,

Are you tired of high electric bills?

We have noticed that your house at .....
```
```
Dear Mr. Smith,

Are you tired of high electric bills?

We have noticed that your house at .....
```

**Template Example - The Template**

```
Dear {{Name}},

Are you tired of high electric bills?

We have noticed that your house at .....
```

**Command-line arguments**

* Command-line arguments are options and data that are passed to programs. 
* We usually pass arguments to console programs, but sometimes we pass arguments to GUI programs as well.
* The `os.Args` holds the command-line arguments.
* The first value in this slice is the name of the program, while the `os.Args[1:]` holds the arguments to the program.
* The individual arguments are accessed with indexing operation.

```
package main

import (
    "fmt"
    "os"
    "reflect"
)

func main() {

    prg_name := os.Args[0]
    fmt.Printf("The program name is %s\n", prg_name)//We get and print the first argument, which is the program name.

    names := os.Args[1:] //We get all the received arguments.
    fmt.Println(reflect.TypeOf(names)) //We print the type which holds the arguments (slice).

    for _, name := range names {

        fmt.Printf("Hello, %s!\n", name) //We go through the arguments and build a message from each of them.
    }
}
$ go build read_args.go 
$ ./read_args Jan Peter Lucia
The program name is ./read_args
[]string
Hello, Jan!
Hello, Peter!
Hello, Lucia!
```

**os.Create**

* This allows us to create a file.

```
func Create(name string) (*File, error)
```

**io.Copy**

* This allows us to copy from from a source to a destination.

```
func Copy(dst Writer, src Reader) (written int64, err error)
```

**strings.NewReader**

* NewReader returns a new Reader reading from s.

```
func NewReader(s string) *Reader
```

**os.Args**

* Args is a variable from package os. 
* Args hold the command-line arguments, starting with the program name.

```
var Args []string
```

**Parsing templates**

**template.ParseFiles**

```
func ParseFiles(filenames ...string) (*Template, error)
```

**template.ParseGlob**

```
func ParseGlob(pattern string) (*Template, error)
```

**template.Parse**

```
func (t *Template) Parse(text string) (*Template, error)
```

**template.ParseFiles**

```
func (t *Template) ParseFiles(filenames ...string) (*Template, error)
```

**template.ParseGlob**

```
func (t *Template) ParseGlob(pattern string) (*Template, error)
```

**Executing templates**

**template.Execute**

```
func (t *Template) Execute(wr io.Writer, data interface{}) error
```

**template.ExecuteTemplate**

```
func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
```


