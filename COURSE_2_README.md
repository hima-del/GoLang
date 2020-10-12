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

**`func (*Template) Funcs`**

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
     
  **func Dial**
  
  ```
func Dial(network, address string) (Conn, error)
```
 * Dial connects to the address on the named network.
     
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
    conn, err := l.Accept() //once we accept we get a connection and we can read or write from that connection
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
    
    
**func NewScanner**

```
func NewScanner(r io.Reader) *Scanner
```
* NewScanner returns a new Scanner to read from r. 
* The split function defaults to ScanLines.
* Returnes a pointer to scanner 
* when we have pointer to scanner we have scan() method attched to it.

**`func (*Scanner) Scan`**

```
func (s *Scanner) Scan() bool
```

* Scan advances the Scanner to the next token, which will then be available through the Bytes or Text method. 
* It returns false when the scan stops, either by reaching the end of the input or an error. 
* After Scan returns false, the Err method will return any error that occurred during scanning, except that if it was io.EOF, Err will return nil. 
* Scan panics if the split function returns too many empty tokens without advancing the input. 

**func ReadAll**

```
func ReadAll(r io.Reader) ([]byte, error)
```

* ReadAll reads from r until an error or EOF and returns the data it read.
* A successful call returns err == nil, not err == EOF. 


```
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)

}
```

**func Fields**

```
func Fields(s string) []string
```
* Fields splits the string s around each instance of one or more consecutive white space characters, as defined by unicode.
* IsSpace, returning a slice of substrings of s or an empty slice if s contains only white space.

```
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
}
//output
Fields are: ["foo" "bar" "baz"]
```

**type Handler**

```
type Handler interface {
    ServeHTTP(ResponseWriter, *Request) //when we have ServeHTTP you have a request and the server responds to that request.
}
```

* A Handler responds to an HTTP request.
* A server handles request by writing responses.
* ServeHTTP should write reply headers and data to the ResponseWriter and then return. 
* Returning signals that the request is finished

**func ListenAndServe**

```
func ListenAndServe(addr string, handler Handler) error
```

* ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections.
* Accepted connections are configured to enable TCP keep-alives.

**Request**

```
type Request struct {
    Method string
    URL *url.URL
	//	Header = map[string][]string{
	//		"Accept-Encoding": {"gzip, deflate"},
	//		"Accept-Language": {"en-us"},
	//		"Foo": {"Bar", "two"},
	//	}
    Header Header
    Body io.ReadCloser
    ContentLength int64
    Host string
    // This field is only available after ParseForm is called.
    Form url.Values
    // This field is only available after ParseForm is called.
    PostForm url.Values
    MultipartForm *multipart.Form
    // RemoteAddr allows HTTP servers and other software to record
	// the network address that sent the request, usually for
	// logging. 
    RemoteAddr string
}
```

* `form` and `postform` allows us to get data from a form
* These fields are only available after ParseForm is called.
* form give data from URL, query string and form body(form data)
* postform only give us data from body

**`func (*Request) ParseForm`**

```
func (r *Request) ParseForm() error
```
* ParseForm populates r.Form and r.PostForm.
* For all requests, ParseForm parses the raw query from the URL and updates r.Form.
* For POST, PUT, and PATCH requests, it also reads the request body, parses it as a form and puts the results into both r.PostForm and r.Form. 
* Request body parameters take precedence over URL query string values in r.Form.


**Request**

```
type Request struct {
    Method string
    URL *url.URL
	//	Header = map[string][]string{
	//		"Accept-Encoding": {"gzip, deflate"},
	//		"Accept-Language": {"en-us"},
	//		"Foo": {"Bar", "two"},
	//	}
    Header Header
    Body io.ReadCloser
    ContentLength int64
    Host string
    // This field is only available after ParseForm is called.
    Form url.Values
    // This field is only available after ParseForm is called.
    PostForm url.Values
    MultipartForm *multipart.Form
    // RemoteAddr allows HTTP servers and other software to record
	// the network address that sent the request, usually for
	// logging. 
    RemoteAddr string
}
```

**Retrieve URL & Form data**

* http.Request is a struct. It has the fields Form & PostForm. 

```
    // Form contains the parsed form data, including both the URL
    // field's query parameters and the POST or PUT form data.
    // This field is only available after **ParseForm** is called.
    // The HTTP client ignores Form and uses Body instead.
    Form url.Values

    // PostForm contains the parsed form data from POST, PATCH,
    // or PUT body parameters.
    // This field is only available after **ParseForm** is called.
    // The HTTP client ignores PostForm and uses Body instead.
    PostForm url.Values
```
* If we look at ParseForm,

`go func (r *Request) ParseForm() error`

* we see that this is a method attached to a *http.Request.

* If we look at FormValue*

`go func (r *Request) FormValue(key string) string`

* we see that this is a method attached to a *http.Request.
* FormValue returns the first value for the named component of the query.
* POST and PUT body parameters take precedence over URL query string values.
* FormValue calls ParseMultipartForm and ParseForm if necessary and ignores any errors returned by these functions. 
* If key is not present, FormValue returns the empty string. 
* To access multiple values of the same key, call ParseForm and then inspect Request.Form directly.


```
type URL struct {
    Scheme     string
    Opaque     string    // encoded opaque data
    User       *Userinfo // username and password information
    Host       string    // host or host:port
    Path       string
    RawPath    string // encoded path hint (Go 1.5 and later only; see EscapedPath method)
    ForceQuery bool   // append a query ('?') even if RawQuery is empty
    RawQuery   string // encoded query values, without '?'
    Fragment   string // fragment for references, without '#'
}
```

**Work with the HTTP header**

* The http.Request type is a struct which has a Header field.

```
type Header map[string][]string
```

**Response**

```
type ResponseWriter interface {
    // Header returns the header map that will be sent by
    // WriteHeader. Changing the header after a call to
    // WriteHeader (or Write) has no effect 
    Header() Header

    // Write writes the data to the connection as part of an HTTP reply.
    //
    // If WriteHeader has not yet been called, Write calls
    // WriteHeader(http.StatusOK) before writing the data. If the Header
    // does not contain a Content-Type line, Write adds a Content-Type set
    // to the result of passing the initial 512 bytes of written data to
    // DetectContentType.
    Write([]byte) (int, error)

    // WriteHeader sends an HTTP response header with status code.
    // If WriteHeader is not called explicitly, the first call to Write
    // will trigger an implicit WriteHeader(http.StatusOK).
    // Thus explicit calls to WriteHeader are mainly used to
    // send error codes.
    WriteHeader(int)
}
```

**Setting a response header**

* An http.ResponseWriter has a method Header() which returns a http.Header.

`type Header map[string][]string`

* Look at the methods which are attached to type http.Header
``` 
type Header
func (h Header) Add(key, value string)
func (h Header) Del(key string)
func (h Header) Get(key string) string
func (h Header) Set(key, value string)
func (h Header) Write(w io.Writer) error
func (h Header) WriteSubset(w io.Writer, exclude map[string]bool) error
```
* We can set headers for a response like this:
```
res.Header().Set("Content-Type", "text/html; charset=utf-8")
```

**type ServeMux**

* ServeMux is an HTTP request multiplexer.
* It matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL.
* // NewServeMux allocates and returns a new ServeMux.

```
func NewServeMux() *ServeMux { return new(ServeMux) }
```

**`func (*ServeMux) Handle`**

```
func (mux *ServeMux) Handle(pattern string, handler Handler)
```

* Handle registers the handler for the given pattern. 
* If a handler already exists for pattern, Handle panics.

**func HandleFunc**

```
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

* HandleFunc registers the handler function for the given pattern in the DefaultServeMux. 
* The `func` can be any function with responsewriter and pointer to request as parameters

```
package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

**type HandlerFunc**

* The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers. 
* If f is a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f.

```
type HandlerFunc func(ResponseWriter, *Request)
```

**func (HandlerFunc) ServeHTTP**

```
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
```

* ServeHTTP calls f(w, r).

* `Handle` wants a `handler`
* A `hander` is anything that has  `ServeHTTP(w ResponseWriter, r *Request)` method
* `Handlerfunc` has this method attched to it


**HTTP request methods**

* HTTP defines a set of request methods to indicate the desired action to be performed for a given resource.

`GET`

* The GET method requests a representation of the specified resource. Requests using GET should only retrieve data.

`HEAD`

* The HEAD method asks for a response identical to that of a GET request, but without the response body.

`POST`

* The POST method is used to submit an entity to the specified resource, often causing a change in state or side effects on the server.

`PUT`

* The PUT method replaces all current representations of the target resource with the request payload.

`DELETE`

* The DELETE method deletes the specified resource.

`CONNECT`

* The CONNECT method establishes a tunnel to the server identified by the target resource.

`OPTIONS`

* The OPTIONS method is used to describe the communication options for the target resource.

`TRACE`

* The TRACE method performs a message loop-back test along the path to the target resource.

`PATCH`

* The PATCH method is used to apply partial modifications to a resource.



**func New**

```
func New() *Router
```
* New returns a new initialized Router. 

**`func (*Router) GET`**

```
func (r *Router) GET(path string, handle Handle)
```

**`func (*Router) POST`**

```
func (r *Router) POST(path string, handle Handle)
```

**`func (*Router) Handle`**

```
func (r *Router) Handle(method, path string, mc MiddlewareChain, h Handler)
```

**type Params**

```
type Params []Param
```

* Params is a Param-slice, as returned by the router. 
* The slice is ordered, the first URL parameter is also the first slice value.

**func (Params) ByName**

```
func (ps Params) ByName(name string) string
```

* ByName returns the value of the first Param which key matches the given name. 
* If no matching Param is found, an empty string is returned.

**`func (*File) Stat`**

```
func (f *File) Stat() (FileInfo, error)
```

* Stat returns the FileInfo structure describing file

**type FileInfo**

* A FileInfo describes a file and is returned by Stat and Lstat.

```
type FileInfo interface {
    Name() string       // base name of the file
    Size() int64        // length in bytes for regular files; system-dependent for others
    Mode() FileMode     // file mode bits
    ModTime() time.Time // modification time
    IsDir() bool        // abbreviation for Mode().IsDir()
    Sys() interface{}   // underlying data source (can return nil)
}
```

**func ServeContent**

```
func ServeContent(w ResponseWriter, req *Request, name string, modtime time.Time, content io.ReadSeeker)
```

* ServeContent replies to the request using the content in the provided ReadSeeker. 
* The main benefit of ServeContent over io.Copy is that it handles Range requests properly, sets the MIME type, and handles If-Match, If-Unmodified-Since, If-None-Match, If-       Modified-Since, and If-Range requests.

**func ServeFile**

```
func ServeFile(w ResponseWriter, r *Request, name string)
```

* ServeFile replies to the request with the contents of the named file or directory.


**func FileServer**

```
func FileServer(root FileSystem) Handler
```

* FileServer returns a handler that serves HTTP requests with the contents of the file system rooted at root.
* To use the operating system's file system implementation, use `http.Dir`
```
http.Handle("/", http.FileServer(http.Dir("/tmp")))
```

**Passing values**

**Form submission**

* We can pass values from the client to the server through the URL or through the body of the request.
* When you submit a form, you can use either the "POST" or "GET" method. 
* The "POST" method sends the form submission through the body of the request. 
* The "GET" method for a form submission sends the form submission values through the url.

```
post
body

get
url
```

* Post has four letters and so does form.
* Get has three letters and so does url.

**URL values**

* You can always append values to a URL.
* Anything after the `?` is the query string - the area where values are stored.
* The values are stord in a `identifier=value` fashion.
* You can have multiple `identifier=value` by separating them with the & ampersand.


**`func (*Request) FormValue`**

```
func (r *Request) FormValue(key string) string
```

* FormValue returns the first value for the named component of the query.

**`func (*Request) FormFile`**

```
func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
```

* FormFile returns the first file for the provided form key.

**type File**

* File is an interface to access the file part of a multipart message. 
* Its contents may be either stored in memory or on disk. 
* If stored on disk, the File's underlying concrete type will be an *os.File.

```
type File interface {
    io.Reader
    io.ReaderAt
    io.Seeker
    io.Closer
}
```

**type FileHeader**

* A FileHeader describes a file part of a multipart request.

```
type FileHeader struct {
    Filename string
    Header   textproto.MIMEHeader
    Size     int64 // Go 1.9
    // contains filtered or unexported fields
}
```

**func Create**

```
func Create(name string) (*File, error)
```

* Create creates or truncates the named file.


**func Join**

```
func Join(elem ...string) string
```

* Join joins any number of path elements into a single path, separating them with an OS specific Separator. 
* Empty elements are ignored. 

```
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println("On Unix:")
	fmt.Println(filepath.Join("a", "b", "c"))
	fmt.Println(filepath.Join("a", "b/c"))
	fmt.Println(filepath.Join("a/b", "c"))
	fmt.Println(filepath.Join("a/b", "/c"))
}

//output
a/b/c
a/b/c
a/b/c
a/b/c
```

**Enctype**

* “When you make a POST request, you have to encode the data that forms the body of the request in some way. 
* HTML forms provide three methods of encoding:

●	application/x-www-form-urlencoded (the default) 
●	multipart/form-data 
●	text/plain 

* All you need to know is use multipart/form-data when your form includes any <input type="file"> elements. 

**Redirects**

* A client makes a request to server => request is for certain resource => the server redirects the client to another location

```
StatusMultipleChoices  = 300 // RFC 7231, 6.4.1
    StatusMovedPermanently = 301 // RFC 7231, 6.4.2
    StatusFound            = 302 // RFC 7231, 6.4.3
    StatusSeeOther         = 303 // RFC 7231, 6.4.4
    StatusNotModified      = 304 // RFC 7232, 4.1
    StatusUseProxy         = 305 // RFC 7231, 6.4.5

    StatusTemporaryRedirect = 307 // RFC 7231, 6.4.7
    StatusPermanentRedirect = 308 // RFC 7538, 3
 ```
   
  * On the web, we have a client / server architecture. 
  * Clients make requests, and servers write responses to those clients.
  * The request and response are both just text that must conform to the rules of HTTP.
  * Both the request & response have a start line, headers, and a body. 
  * The request start line is called the “request line”. It consists of 

 `request-line = method SP request-target SP HTTP-version CRLF`

  * The response start line is called the “status line”. It consists of 

  `status-line = HTTP-version SP status-code SP reason-phrase CRLF`


   **301 Moved Permanently**

   * The 301 (Moved Permanently) status code indicates that the target resource has been assigned a new permanent URI and any future references to this resource ought to use one      of the enclosed URIs.

   **303 See Other**

   * The 303 (See Other) status code indicates that the server is redirecting the user agent to a different resource, as indicated by a URI in the Location header field, which        is intended to provide an indirect response to the original request.
   
   **307 Temporary Redirect**

   * The 307 (Temporary Redirect) status code indicates that the target resource resides temporarily under a different URI and the user agent MUST NOT change the request method      if it performs an automatic redirection to that URI. 
   * Since the redirection can change over time, the client ought to continue using the original effective request URI for future requests.


 **func Redirect**
 
 ```
func Redirect(w ResponseWriter, r *Request, url string, code int)
```

* Redirect replies to the request with a redirect to url, which may be a path relative to the request path.
* The provided code should be in the 3xx range and is usually StatusMovedPermanently, StatusFound or StatusSeeOther.


**cookies**

* cookies is a file that can store information in which a server can write to client's machine if the client allows cookies to be written.
* if we are able to write a cookie to a client's machine then anytime the client makes request back to server,it will send that cookie
* A cookie is a file that is stored on the client's machine.
* Cookies are written by domains to store information on the client's machine.
* Cookies are only sent by the browser to the domain which wrote them.
* With every request to a specific domain, the client's web browser looks to see if there is a cookie from that domain on the client's machine.
* If there is a cookie that has been written by that particular domain, then the browser will send the cookie with every request to that domain.
* Cookies are domain specific.

**Expiring a cookie**

* If the Expires or MaxAge field isn't set, then the cookie is deleted when the browser is closed. This is colloquially known as a "session cookie."
* You can expire a cookie by setting one of these two fields: Expires or MaxAge
* Expires sets the exact time when the cookie expires. Expires is Deprecated.
* MaxAge sets how long the cookie should live (in seconds).

**type Cookie**

* A Cookie represents an HTTP cookie as sent in the Set-Cookie header of an HTTP response or the Cookie header of an HTTP request.

```
type Cookie struct {
    Name  string
    Value string

    Path       string    // optional
    Domain     string    // optional
    Expires    time.Time // optional
    RawExpires string    // for reading cookies only

    // MaxAge=0 means no 'Max-Age' attribute specified.
    // MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
    // MaxAge>0 means Max-Age attribute present and given in seconds
    MaxAge   int
    Secure   bool //this is set true when we are using https . https is what secure means
    HttpOnly bool // you can access this cookie with javascript
    SameSite SameSite // Go 1.11
    Raw      string
    Unparsed []string // Raw text of unparsed attribute-value pairs
}
```

**func SetCookie**

```
func SetCookie(w ResponseWriter, cookie *Cookie)
```

* SetCookie adds a Set-Cookie header to the provided ResponseWriter's headers. 
* The provided cookie must have a valid Name.

**`func (*Request) Cookie`**

```
func (r *Request) Cookie(name string) (*Cookie, error)
```

`ErrNoCookie is returned by Request's Cookie method when a cookie is not found.

var ErrNoCookie = errors.New("http: named cookie not present")
`

**func Atoi**

```
func Atoi(s string) (int, error)
```

* Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.

```
package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := "10"
	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T, %v", s, s)
	}

}
//output
int, 10
```

**func Itoa**

```
func Itoa(i int) string
```

* Itoa is equivalent to FormatInt(int64(i), 10).

```
package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 10
	s := strconv.Itoa(i)
	fmt.Printf("%T, %v\n", s, s)

}
//output
string, 10
```

**func GenerateFromPassword**

```
func GenerateFromPassword(password []byte, cost int) ([]byte, error)
```

* GenerateFromPassword returns the bcrypt hash of the password at the given cost. 

**func CompareHashAndPassword**

```
func CompareHashAndPassword(hashedPassword, password []byte) error
```

* CompareHashAndPassword compares a bcrypt hashed password with its possible plaintext equivalent.

**package "crypto/hmac"**

**func New**

`
func New(h func() hash.Hash, key []byte) hash.Hash
`
* New returns a new HMAC hash using the given hash.Hash type and key.

**encoding/base64**

**func NewEncoding**
`
func NewEncoding(encoder string) *Encoding
`
* NewEncoding returns a new padded Encoding defined by the given alphabet


**`func (*Encoding) EncodeToString`**
`
func (enc *Encoding) EncodeToString(src []byte) string
`
* EncodeToString returns the base64 encoding of src.

**`func (*Encoding) DecodeString`**
`
func (enc *Encoding) DecodeString(s string) ([]byte, error)
`
* DecodeString returns the bytes represented by the base64 string s.


**package context**

**func WithValue**
`
func WithValue(parent Context, key, val interface{}) Context
`
* WithValue returns a copy of parent in which the value associated with key is val.

**func WithTimeout**

`
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	return WithDeadline(parent, time.Now().Add(timeout))
`
* WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).
* Canceling this context releases resources associated with it, so code should
* call cancel as soon as the operations running in this Context complete:


**func WithCancel**
`
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
`
* WithCancel returns a copy of parent with a new Done channel.
* The returned context's Done channel is closed when the returned cancel function is called or when the parent context's Done channel is closed, whichever happens first.
* Canceling this context releases resources associated with it, so code should call cancel as soon as the operations running in this Context complete.

**func Background**
`
func Background() Context
`
* Background returns a non-nil, empty Context. 
* It is never canceled, has no values, and has no deadline. 
* It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests.

**func ListenAndServeTLS**
`
func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error
`
* ListenAndServeTLS acts identically to ListenAndServe, except that it expects HTTPS connections. 
* Additionally, files containing a certificate and matching private key for the server must be provided.

**Go & JSON - Marshal & Encode** 

* Now we will see how to use Go with JSON. 
* The most important thing to understand is that you can marshal *OR* encode Go code to JSON.
* Regardless of whether or not you use “marshal” or “encode”, your Go data structures will be turned into JSON. 
* So what’s the difference? Marshal is for turning Go data structures into JSON and then assigning the JSON to a variable. 
* Encode is used to turn Go data structures into JSON and then send it over the wire. 
* Both “marshal” and “encode” have their counterparts: “unmarshal” and “decode”.

**Using Curl**

* Start your server
* go run main.go
* Enter this at the terminal
* curl http://localhost:8080
