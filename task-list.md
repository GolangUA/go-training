#Task List

## Short declaration syntax

- Follow [ this link ](https://play.golang.org/p/3Vl75w72JO) and press the *Run* button.
- Rewrite the program using the short declaration syntax; there should be no `var` declarations, only `:=`.

## Slice exercises

- Follow [this link](https://play.golang.org/p/AJk1Jgp1iE) for instructions.
- Declare a variable called `i` which is a slice of 5 `int`.
- Declare a variable called `f` which is a slice of 9 `float64`.
- Declare a variable called `s` which is a slice of 4 `string`.
- Does your program print the expected result, `5`9`4`?

## Slice initialisation 

- Follow [ this link ](https://play.golang.org/p/P-eBqzPCWh) and complete the exercise.

## Subslices

- Follow [this link](https://play.golang.org/p/d1-jl42aTF) and complete the exercise.

## Inserting values into a map

- Follow [ this link ](https://play.golang.org/p/a-V5I0nZ5l) and complete the exercise.

## Range over slices

- Follow [ this link ](https://play.golang.org/p/AmQW-OrPC1) and complete the exercise.
- If you cannot figure it out, don't worry, there is an answer on the next slide.

## What time is it (exercise)

- `cd`$GOPATH/src/whattimeisit`
- Build the program with, `go`build` (if you make an error, go back and edit `main.go`)
- Run your program `./whattimeisit`, it should print something like this

```
 The current time is 2016-12-05 12:33:18.222821474 +0900 JST
```

## Writing tests

- The code for this exercise is in `$GOPATH/src/simplestrings/`
- Read the source code for `simplestrings.go`

## Methods

- Complete the exercise in `$GOPATH/src/exercises/methods` by making all the tests pass.

## Pointers

- Complete the exercise in `$GOPATH/src/exercises/pointers` by making the test pass.

##  Reading input

This exercise asks you to write a function that reads lines from an io.Reader and returns a string containing all the lines read.

The code for this exercise is in `$GOPATH/src/exercises/input`

##  Readers

To familarise you with the `io.Reader` implementations available in the `io` package, this exercise is all about Readers.

- Complete the exercise in `$GOPATH/src/exercises/readers` by making the test pass.
- If you get stuck, consult the documentation in the [`io`](https://golang.org/pkg/io/ ) package.

##  Counting the number of lines in a file

Now we know about `io.Reader`'s, `error`'s, we can write some more useful programs.

The code for this exercise is in `$GOPATH/src/exercises/countlines`

_Note_:

- `go`test` always executes from the package's source directory, this makes it simple to include fixtures for your tests.
- The go tool ignores any directory called `testdata`, or starts with a `.` or `_`.

## Error handling (cont.)

In the previous counting example, if an error happened, the program would exit.

In this exercise, we'll handle errors by returning them to the caller.

The code for this exercise is in `$GOPATH/src/exercises/errorhandling`

## Passing in a reader

Let's turn out `Countlines` function into a program.

The code for this exercise is in `$GOPATH/src/exercises/linecount`

Complete the program so it reads the number of lines sent to it via stdin.

```
 cat testdata/moby.txt | ./linecount 
 22659
```

## Handling multiple files

Let's extend our `linecount` program to handle files passed on the command line.

The code for this exercise is in `$GOPATH/src/exercises/countmanyfiles`

Complete the program so it counts the lines in files passed via the command line.

```
 ./countmanyfiles testdata/*.txt
 testdata/dracula.txt    15973
 testdata/pride-and-prejudice.txt        13427
 testdata/sherlock.txt   13052
```

## Reading all the *.txt files in a directory

In the previous example we used the shell to list files to process.

In this exercise, let's extend our `countmanyfiles` program to walk the directory listing itself.

To do this we use the [`ReadDir`](https://golang.org/pkg/os/#File.Readdir) method on [`os.File`](https://golang.org/pkg/os/#File).

_Note_: Be careful to only read _files_, not directories, and do not read files that don't end in `.txt`

The code for this exercise is in `$GOPATH/src/exercises/countdir`

Complete the program so it counts the lines in files passed via the command line.

```
 ./countdir testdata/
 testdata/christmas-carol.txt    4236
 testdata/tom-sawyer.txt 9209
```

##  HTTP request

The Go standard library supports writing HTTP clients and servers with the [`net/http`](https://golang.org/pkg/net/http/) package.

Using the `net/http` package is very straight forward:

```
 resp, err := http.Get("http://example.com/")
 if err != nil {
         // handle error
 }
```
 
`resp` is a [`http.Response`](https://golang.org/pkg/net/http/#Response) structure, which has lots of interesting fields.

Let's write a simple HTTP client that can fetch HTTP URLs.

```
 ./httpget http://httpbin.org/ip
 {
   "origin": "125.203.122.114"
 }
```

The code for this exercise is in `$GOPATH/src/exercises/httpget`

## JSON parsing

The service at `http://httpbin.org/` returns JSON bodies.

The [`encoding/json`](https://golang.org/pkg/encoding/json/) package can decode JSON data into a map.

```
 result := make(map[string]string)
 dec := json.NewDecoder(resp.Body)
 err := dec.Decode(&result)
 if err != nil {
         // handle error
 }
```

Let's use this to write a program that will tell us our public IP address.

```
 ./whatismyip 
 My IP address is: 125.203.122.114
```

The code for this exercise is in `$GOPATH/src/exercises/whatismyip`

## Controlling JSON encoding

The `encoding/json` package requires the fields of a struct to be public (start with an upper case letter), this means the keys in your JSON document will be upper case.

We can fix this and control the output of the JSON with a _tag_.

The format of the JSON tag is documented on the [`json.Encode`](https://golang.org/pkg/encoding/json/#Marshal) method.

The code for this exercise is in `$GOPATH/src/exercises/jsonenc`

## Line counting HTTP service

Let's write a HTTP service that counts the lines in a book via http.

Features:

- When the client requests `/books/{book}` we look up the book and return the number of lines counted.
- The response to the client should be in JSON format and include the number of lines and the title of the book.
- Book directory is configurable.

```
 ./httplinecount ../../../books/ &
 [1] 17554
 curl http://localhost:8080/books/moby.txt
 {"title":"moby.txt","lines":"22659"}
```

The code for this exercise is in `$GOPATH/src/exercises/httplinecount`

