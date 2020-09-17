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

**template.Must**

```
func Must(t *Template, err error) *Template
```
`.` represents current piece of data 
`{{range pipeline}} T1 {{end}}`
	* The value of the pipeline must be an array, slice, map, or channel.
	* If the value of the pipeline has length zero, nothing is output; otherwise, dot is set to the successive elements of the array, slice, or map and T1 is executed.
        * If the value is a map and the keys are of basic type with a defined order, the elements will be visited in sorted key order.
	

**type FuncMap**

* FuncMap is the type of the map defining the mapping from names to functions.
	
```
type FuncMap map[string]interface{}
```
* It takes string for key empty interface for value.ie,here we can pass a string for key and a function for value
* Empty interface is ineterface with no methods and every type has atleat no methods
* So every type implements the empty interface

**func (*Template) Funcs**

```
func (t *Template) Funcs(funcMap FuncMap) *Template
```
* Funcs adds the elements of the argument map to the template's function map. 
* It must be called before the template is parsed.


**Nested template definitions**

* When parsing a template, another template may be defined and associated with the template being parsed.
* Template definitions must appear at the top level of the template, much like global variables in a Go program.
* The syntax of such definitions is to surround each template declaration with a "define" and "end" action.
* The define action names the template being created by providing a string constant. 

`{{define "T1"}}ONE{{end}}
{{define "T2"}}TWO{{end}}
{{define "T3"}}{{template "T1"}} {{template "T2"}}{{end}}
{{template "T3"}}`

* This defines two templates, T1 and T2, and a third T3 that invokes the other two when it is executed. 

**Message Format**

* All HTTP/1.1 messages consist of a start-line followed by a sequence of octets in a format similar to the Internet Message Format [RFC5322]: zero or more header fields           (collectively referred to as the "headers" or the "header section"), an empty line indicating the end of the header section, and an optional message body.

 **Start Line**

   * An HTTP message can be either a request from client to server or a response from server to client. 
   * Syntactically, the two types of message differ only in the start-line, which is either a request-line (for requests) or a status-line (for responses), and in the algorithm
     for determining the length of the message body 
   * In theory, a client could receive requests and a server could receive responses, distinguishing them by their different start-line formats, but, in practice, servers are        implemented to only expect a request (a response is interpreted as an unknown or invalid request method) and clients are implemented to only expect a response.

     `start-line     = request-line / status-line`
     
     `request-line   = method SP request-target SP HTTP-version CRLF`
     
     ` status-line = HTTP-version SP status-code SP reason-phrase CRLF`
     
 **func Listen**
   
   ```
func Listen(network, address string) (Listener, error)
```
* Listen announces on the local network address.
* The network must be "tcp", "tcp4", "tcp6", "unix" or "unixpacket".

`func Listen will return Listener and an error`

**type Listener**

```
type Listener interface {
    // Accept waits for and returns the next connection to the listener.
    Accept() (Conn, error)

    // Close closes the listener.
    // Any blocked Accept operations will be unblocked and return errors.
    Close() error

    // Addr returns the listener's network address.
    Addr() Addr
}
```

```
// Listen on TCP port 2000 on all available unicast and
// anycast IP addresses of the local system.
l, err := net.Listen("tcp", ":2000")
if err != nil {
    log.Fatal(err)
}
defer l.Close()
for {
    // Wait for a connection.
    conn, err := l.Accept()
    if err != nil {
        log.Fatal(err)
    }
 
```
```
type Conn interface {
    // Read reads data from the connection.
    // Read can be made to time out and return an error after a fixed
    // time limit; see SetDeadline and SetReadDeadline.
    Read(b []byte) (n int, err error)

    // Write writes data to the connection.
    // Write can be made to time out and return an error after a fixed
    // time limit; see SetDeadline and SetWriteDeadline.
    Write(b []byte) (n int, err error)

    // Close closes the connection.
    // Any blocked Read or Write operations will be unblocked and return errors.
    Close() error

    // LocalAddr returns the local network address.
    LocalAddr() Addr

    // RemoteAddr returns the remote network address.
    RemoteAddr() Addr
 ```
 
 **func WriteString**
 
 ```
func WriteString(w Writer, s string) (n int, err error)
```
* WriteString writes the contents of the string s to w, which accepts a slice of bytes.
    
    
