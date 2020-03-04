#vim go ide setup, bash code below
export GOPATH=$HOME/.go
export PATH=$PATH:$GOPATH/bin

#install go
#u might wanna find newest version link and then replace anmebelow
wget https://dl.google.com/go/go1.11.5.linux-amd64.tar.gz
tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
export PATH=$PATH:/usr/local/go/bin


#if mac, ubuntu use apt-get install go
brew install go
#optional?
"go get" the basics
#optional?
go get golang.org/x/tools/cmd/godoc

#install pathogen
mkdir -p ~/.vim/autoload ~/.vim/bundle
curl -LSso ~/.vim/autoload/pathogen.vim https://tpo.pe/pathogen.vim
#install vim-go-ide into vim, pathogen will automatically load the module
git clone https://github.com/farazdagi/vim-go-ide.git ~/.vim_go_runtime
sh ~/.vim_go_runtime/bin/install
#run vim in ide mode for go
vim -u ~/.vimrc.go
#install necessary binaries for go ide to work(important), if run into problem of not finding the import path for modules, check the GOROOT, try unset it
in vim run ":GoInstallBinaries"
#put below line into ~/.vimrc.go
set t_Co=256


go run helloworld.go

godoc fmt println //this cmd shows you the println function under package godoc


//remove new line from string
read_line = read_line[:len(read_line)-1]
//Perhaps a better approach is to use the strings library:
read_line = strings.TrimSuffix(read_line, "\n")


//string manipulation
strings.HasPrefix(the_str, str_pattern)
strings.Split(the_str, separator)

//initialize map[string][]string
mapOfSlices := map[string][]string{
    "first": {},
    "second": []string{"one", "two", "three", "four", "five"},
    "third": []string{"quarter", "half"},
}

//golang datetime formatting
package main

import (
    "fmt"
    "time"
)

func main() {
    layout1 := "03:04:05PM"
    layout2 := "15:04:05"
    t, err := time.Parse(layout1, "07:05:45PM")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(t.Format(layout1))
    fmt.Println(t.Format(layout2))
}


//testing utility
func equals(t *testing.T, actual interface{}, expected interface{}) {
    if actual != expected {
        t.Error("Expected >%s<, actual >%s<", expected, actual)
    }
}
func fail(t *testing.T, err error, expected_err_string string) {
    if err.Error() != expected_err_string {
        t.Error("Expected >%s<, actual >%s<", expected_err_string, err.Error())
    }
}
func ok(t *testing.T, actual_error error) {
    if actual_error != nil {
        t.Error("unexpected error raised %s", actual_error)
    }
}


example code:


package main
import (
    "fmt"
    "math"
    "errors"
)

type person struct{
    name string
    age int

}

func main(){
    fmt.Println("hello world")
    //x:=5
    var x int = 5
    fmt.Println(x)
    if x <6 {
        fmt.Println("false")
    }

    //var a [5]int
    var a []int
    a = append(a,5)

    //when no number assigned to array, it's a slice, variable length
    //var a [5]int{5,4,3,2,1}
    //a[2] = 7
    //a = append(a,110)
    fmt.Println(a)


    key_value_pair := make(map[string]int)
    key_value_pair["key1"] = 2
    key_value_pair["key2"] = 3
    delete(key_value_pair, "key2")
    fmt.Println(key_value_pair)

    for i := 0; i < 5; i++{
        fmt.Println(i)
    }
    
        //while loop
    //i := 0
    //for i < 5 {
    //    fmt.Println(i)
    //}

    m := make(map[string]string)
    m["a"] = "alpha"
    m["b"] = "alpha"
    for key, value := range m{
        fmt.Println("key:",key, "value:",value)
    }

    result := sum(2,3)
    fmt.Println(result)

    sqrt_res, err := sqrt(16)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(sqrt_res)
    p := person{name: "", age: 23}
    fmt.Println(p)


    ab :=7
    fmt.Println(&ab)
    //pass in the address of the pointer
    increase(&ab)
    fmt.Println(ab)

}

func increase(x *int){
    //dereference the pointer address
    *x++
}

func sum(x int, y int) int {
    return x+y

}

func sqrt(x float64) (float64, error){
    if x < 0 {
        return 0, errors.New("undefined for negative")
    }
    return math.Sqrt(x), nil
}



//method
func sayHello(){
    fmt.Println("hello")
}

//type struct, think struct as a class

type person struct{
    firstname String //struct variable
    lastname String //struct variable
    //we want a speak method for person struct
}
//this method defined a class method for person
func(p person) speak(){
    fmt.Println("i m human")
}

type secretAgent struct{
    person //this struct inherit person struct, now u can call secretAgentIns.firstname ... etc
    licenceToKill bool //added class variable
}

//overwrite the speak() method for secretAgent
func(s secretAgent) speak(){
    fmt.Println('i m james bond')
}

//interface: with the person and secretAgent struct in the scope, now we want to generalize both person and secretAgent
//to be a human class, in goland, this is called interface
type human interface{
    //this speak() defined that "every struct has this speak() method is now a human abstract class"
    speak()
}

//delete element from map a
delete(aMap, key)

//now we have a human interface, we can take any human interface and call its speak function)
func saySomething(h human){
    h.speak()
}

//a pointer receiver, this will enable this method to have the real person struct, so you can modify its value;
//if not a pointer receiver, the p(person struct) will be a copy of original person struct, any modify is to the copy of the struct not itself.
func(p *person) changeFirstName(newFirstName string) {
    p.firstname = newFirstName
}

//golang type conversion
p=person{"joe","underwood"}
secAgen := p.(secretAgent) //this converts the person type to secret agent


//empty interface to deal with multiple type problem
//type switch
var emptyInterface interface{} = 0 //this time we defined an empty interface and assign 0(int) to it
switch emptyInterface.(type) {
    case int: fmt.Println("i m an int")
    case float64: fmt.Println("i m an float64")
    default:
        fmt.Println()

}

    
//when in terms of defining a new var for interface, if you have defined a pointer receiver interface, it would work with either pointer receiver method implementation, or value method implementation
var i = *someInterface{}
//but when defined a value interface
var i = someInterface{}//this will require underlying method implementation to be all value . func(i someInterface) someMethod(){}

    
    
//best way to concatenate string, bytes
package main

import (
    "bytes"
    "fmt"
)

func main() {
    var buffer bytes.Buffer

    for i := 0; i < 1000; i++ {
        buffer.WriteString("a")
        //buffer.Wirte([]bytes) this also works for bytes
    }

    fmt.Println(buffer.String())
}
    
    
//read file

package main

import (
    "compress/gzip"
    "encoding/csv"
    "fmt"
    "log"
    "os"
)

func main() {
    f, err := os.Open("data.csv.gz")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    gr, err := gzip.NewReader(f)
    if err != nil {
        log.Fatal(err)
    }
    defer gr.Close()

    cr := csv.NewReader(gr)
    rec, err := cr.Read()
    if err != nil {
        log.Fatal(err)
    }
    for _, v := range rec {
        fmt.Println(v)
    }
}



//go testing 

//benchmark

func BenchmarkXxx(*testing.B)
are considered benchmarks, and are executed by the "go test" command when its -bench flag is provided. Benchmarks are run sequentially.

For a description of the testing flags, see https://golang.org/cmd/go/#hdr-Description_of_testing_flags.

A sample benchmark function looks like this:

func BenchmarkHello(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprintf("hello")
    }


//reset benchmark timer
b.ResetTimer()



//vim-go fast add qutoe to a word
ciw"Ctrl+r" //this cmd is type ciw in normal mode will delete word, and manually type ", and then press ctrl+r" to paste what is last deleted
//outside vim-go fast add quote
ciw'Ctrl+r"'

//check if map has this element
if _, ok := groupSet["sip"]; ok {
    //do stuff
}



//go routine sync waiting
package main

import "fmt"
import "sync"


var wg sync.WaitGroup // 1
func cleanup(){
    if r := recover(); r!=nil {
        fmt.Println("recovered",r)
    }
}
func routine(i int) {
    defer wg.Done() // 3
    defer cleanup()
    fmt.Printf("routine %v finished\n", i)
    if i ==2 {
    panic("2 should not be in here")
    }
}

func main() {
    for i := 0; i < 10; i++ {
        wg.Add(1) // 2
        go routine(i) // *
    }
    wg.Wait() // 4
    fmt.Println("main finished")
}

//go routine channel
package main
import "fmt"
import "sync"

var wg sync.WaitGroup
func foo(c chan int, v int){
    defer wg.Done()
    c <- v*5
}
func main(){
    channelIns := make(chan int, 10)
    for i :=0; i< 10;i++{
        wg.Add(1)
        go foo(channelIns, i)
    }
    //go foo(channelIns, 5)
    //go foo(channelIns, 3)
    //v1 := <- channelIns
    //v2 := <- channelIns
    //fmt.Println(v1, v2)
    wg.Wait()
    //only close channel after nothing will send to channel
    close(channelIns)
    for item := range channelIns{
        fmt.Println(item)
    }
}


//in go only have pass by value, everything is duplicated when passing around, so use pointer
import (
    "fmt"
)

type Point struct {
    x int
    y int
}

func (p Point) String() string {
    return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func modifyValue(point Point) {
    point.x += 10
}

func modifyPointer(point *Point) {
    point.x = 5
    point.y = 5
}

func modifyReference(point *Point) {
    point = &Point{5, 5}
}

func main() {
    p := Point{0, 0}
    fmt.Println(p) // prints (0, 0)

    modifyValue(p)
    fmt.Println(p) // prints (0, 0)

    modifyPointer(&p)
    fmt.Println(p) // prints (5, 5)

    p = Point{0, 0}
    modifyReference(&p)
    fmt.Println(p) // prints (0, 0)
}

//check if string is in a string
strings.Contains("something", "some") // true



testing, show code coverage:
1.navigate to source code package dir
2.run go tool cover -html=c.out
3. if not sucessful, run touch c.out, then try again, if successful, you should see message like 
HTML output written to /tmp/cover716341004/coverage.html
4.now checkout the file generated above, in this case it is at "/tmp/cover716341004/coverage.html"
5. now color coding is through a class called covX, known one is cov0, this means not covered, cov8 means covered
<span class="cov0" title="0">{
                fmt.Errorf("%s", err.Error())
                return []string{}, 1, err
}</span>
above code means the part code of golang error handling is not covered
// another example go test -coverprofile=c.out && go tool cover -html=c.out

//string to bool
b, err := strconv.ParseBool("true")
//string to float
f, err := strconv.ParseFloat("3.1415", 64)
//string to int
i, err := strconv.ParseInt("-42", 10, 64)
//string to int
u, err := strconv.ParseUint("42", 10, 64)
//float to int
int(42.0)

//unmarshal unstructured json string
  birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":"none"}`
  var result map[string]interface{}
  json.Unmarshal([]byte(birdJson), &result)

  // The object stored in the "birds" key is also stored as 
  // a map[string]interface{} type, and its type is asserted from
  // the interface{} type
  birds := result["birds"].(map[string]interface{})

  //simple for loop
  for key, value := range birds {
    // Each value is an interface{} type, that is type asserted as a string
    fmt.Println(key, value.(string))
  }


//golang time formatting conversion
func main() {
	fmt.Println("Hello, playground")
	input := "2018-11-12T16:18:06.821Z"
	layout := "2006-01-02T15:04:05.000Z"
	t, _ := time.Parse(layout, input)

}



//golang net package
import (
    "net"
)
conn, err := net.Dial("tcp", "golang.org:80")
if err != nil {
	// handle error
}
fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
status, err := bufio.NewReader(conn).ReadString('\n')

//use net to create a server
ln, err := net.Listen("tcp", ":8080")
if err != nil {
	// handle error
}
for {
	conn, err := ln.Accept()
	if err != nil {
		// handle error
	}
	go handleConnection(conn)
}
// choose net dns resolve method, prefer go method since it only consumes a go routine
export GODEBUG=netdns=go    # force pure Go resolver
export GODEBUG=netdns=cgo   # force cgo resolver


// buffered channel
package main
import "fmt"
func main() {
    // by specifying 2 for a channel made it a buffered channel
    messages := make(chan string, 2)
    messages <- "buffered"
    messages <- "channel"
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}


//log module. set date and time for it
log.SetFlags(log.Ldate | log.Ltime)

//flag module is used to parse argument passed into golang cmd
var flagname = flag.String("flagname", "str1", "help message") // this binds the argument flagname to pointer flagname 
flag.Parse() // this line will parse the argument
fmt.Println(*flagname) //this will print out flagname pointer

// directional channel 
sendch chan<- int

// directional chan is subclass of channel

// check if receive message is ok, false when chanmel is cloased
v, ok := <- ch

// format golang string
fmt.Sprintf("(%d, %d)", p.x, p.y)

// define new variable
val1 := "123"
//or
var var1 string;
var1 = "123"
//or
var var1 string = "123"


// net packages
conn, err := net.Dial("tcp", "golang.org:80")
if err != nil {
	// handle error
}
fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
status, err := bufio.NewReader(conn).ReadString('\n')

// a net service 
ln, err := net.Listen("tcp", ":8080")
if err != nil {
	// handle error
}
for {
	conn, err := ln.Accept()
	if err != nil {
		// handle error
	}
	go handleConnection(conn)
}

// swap domain name resolution 
// export GODEBUG=netdns=go    # force pure Go resolver
// export GODEBUG=netdns=cgo   # force cgo resolver

// net.Conn interface
type Conn interface {
    // Read reads data from the connection.
    // Read can be made to time out and return an Error with Timeout() == true
    // after a fixed time limit; see SetDeadline and SetReadDeadline.
    Read(b []byte) (n int, err error)

    // Write writes data to the connection.
    // Write can be made to time out and return an Error with Timeout() == true
    // after a fixed time limit; see SetDeadline and SetWriteDeadline.
    Write(b []byte) (n int, err error)

    // Close closes the connection.
    // Any blocked Read or Write operations will be unblocked and return errors.
    Close() error

    // LocalAddr returns the local network address.
    LocalAddr() Addr

    // RemoteAddr returns the remote network address.
    RemoteAddr() Addr

    // SetDeadline sets the read and write deadlines associated
    // with the connection. It is equivalent to calling both
    // SetReadDeadline and SetWriteDeadline.
    //
    // A deadline is an absolute time after which I/O operations
    // fail with a timeout (see type Error) instead of
    // blocking. The deadline applies to all future and pending
    // I/O, not just the immediately following call to Read or
    // Write. After a deadline has been exceeded, the connection
    // can be refreshed by setting a deadline in the future.
    //
    // An idle timeout can be implemented by repeatedly extending
    // the deadline after successful Read or Write calls.
    //
    // A zero value for t means I/O operations will not time out.
    //
    // Note that if a TCP connection has keep-alive turned on,
    // which is the default unless overridden by Dialer.KeepAlive
    // or ListenConfig.KeepAlive, then a keep-alive failure may
    // also return a timeout error. On Unix systems a keep-alive
    // failure on I/O can be detected using
    // errors.Is(err, syscall.ETIMEDOUT).
    SetDeadline(t time.Time) error

    // SetReadDeadline sets the deadline for future Read calls
    // and any currently-blocked Read call.
    // A zero value for t means Read will not time out.
    SetReadDeadline(t time.Time) error

    // SetWriteDeadline sets the deadline for future Write calls
    // and any currently-blocked Write call.
    // Even if write times out, it may return n > 0, indicating that
    // some of the data was successfully written.
    // A zero value for t means Write will not time out.
    SetWriteDeadline(t time.Time) error
}

// convert string represetation of bytes that like "0xFF" to binary which is 0xFF, treat this function as a type conversion of string to []bytes
// func DecodeString(s string) ([]byte, error)
a, _ := hex.DecodeString("0xFF")

// decode bytes to string
import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	src := []byte("48656c6c6f20476f7068657221")  // this converts string representation of bytes to actual []bytes

	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", dst[:n])
}


// make a bytes array with length and capacity
buf := make([]byte, 4, 16)  // 4 is the length, and the 16 is the capacity
// will get [0,0,0,0], but this byte array could expand to size 16

// if else example
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

// appending to slice or add to slice, returns a new copy of slice
s = append(s, 1)

// lock variable, thread safe variable
package main

import (
	"fmt"
	"sync"  // this library provides sync
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex  // we need this Mutex in struct to provide method level lock/unlock
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}



// lock vs rlock, rlock can read many but write one
// we can use defer to make method with locking simpler
r.mux.RLock()
defer r.mux.RUnlock()

//remove element from map
delete(map_var, key1)
//check if element exist in map
if val, ok := dict["foo"]; ok {
    //do something here
}


//function programming in golang
func main() {
    var list = []string{"Orange", "Apple", "Banana", "Grape"}
    // we are passing the array and a function as arguments to mapForEach method.
    var out = mapForEach(list, func(it string) int {
        return len(it)
    })
    fmt.Println(out) // [6, 5, 6, 5]

}

// The higher-order-function takes an array and a function as arguments
func mapForEach(arr []string, fn func(it string) int) []int {
    var newArray = []int{}
    for _, it := range arr {
        // We are executing the method passed
        newArray = append(newArray, fn(it))
    }
    return newArray
}

//partial functions
func add(x int) func(y int) int {
    // A function is returned here as closure
    // variable x is obtained from the outer scope of this method and memorized in the closure
    return func(y int) int {
        return x + y
    }
}

func main() {

    // we are currying the add method to create more variations
    var add10 = add(10)
    var add20 = add(20)
    var add30 = add(30)

    fmt.Println(add10(5)) // 15
    fmt.Println(add20(5)) // 25
    fmt.Println(add30(5)) // 35
}


//implement generic golang functions, any type
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

// Insertion sort
func insertionSort(data Interface, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}



// ...  this means recursively for all files, useful when writing makefile for go




//golang debugger
go get -u github.com/go-delve/delve/cmd/dlv
dlv debug file


// ways to initilize struct
//Struct types
//A struct is a typed collection of fields, useful for grouping data into records.
//way 1
type Student struct {
    Name string
    Age  int
}

var a Student    // a == Student{"", 0}
a.Name = "Alice" // a == Student{"Alice", 0}
//To define a new struct type, you list the names and types of each field.
//The default zero value of a struct has all its fields zeroed.
//You can access individual fields with dot notation.

//way 2
var pa *Student   // pa == nil
pa = new(Student) // pa == &Student{"", 0}
pa.Name = "Alice" // pa == &Student{"Alice", 0}
You can also create and initialize a struct with a struct literal.

//way 3
b := Student{ // b == Student{"Bob", 0}
    Name: "Bob",
}
    
pb := &Student{ // pb == &Student{"Bob", 8}
    Name: "Bob",
    Age:  8,
}

c := Student{"Cecilia", 5} // c == Student{"Cecilia", 5}
d := Student{}             // d == Student{"", 0}


// concatenate two slices
// ... allows you pass any slice, which will loop thru the slice and apply each element in slice
append([]int{1,2}, []int{3,4}...)

//byte array to string
str:=hex.EncodeToString([]byte{01000000})


//uint8 is alias for byte
var a int8 = 12
str2:=hex.EncodeToString([]byte{a})



// string to bytes
b := []byte("ABC€")
fmt.Println(b) // [65 66 67 226 130 172]

//Convert bytes to string
s := string([]byte{65, 66, 67, 226, 130, 172})
fmt.Println(s) // ABC€

// group of variables, mass var

var (
a string = "123"
)

// type of variable
reflect.TypeOf(a)
