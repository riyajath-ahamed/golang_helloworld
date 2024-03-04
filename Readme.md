
# Go Lang

``` 
$ go mod init example/hello
go: creating new go.mod: module example/hello 
```
To run The code
```
$ go run .

```
To build the code
```
$ go build
$ ls
go.mod  go.sum  hello
$ ./hello
Hello, World!
```
Command to get a list of the others.
```
$ go help
```

In your Go code, import the rsc.io/quote package and add a call to its Go function.
After adding the highlighted lines, your code should include the following:

```go
package main

import (
    "fmt"
    "rsc.io/quote"
)

func main() {
    fmt.Println(quote.Go())
}
```
To run the code
Go will add the quote module as a requirement, as well as a go.sum file for use in authenticating the module
```
$ go mod tidy
go: finding module for package rsc.io/quote
go: found rsc.io/quote in rsc.io/quote v1.5.2
```


# Go Modules
```
$ go mod init example/hello
go: creating new go.mod: module example/hello
$ cat go.mod
module example/hello   

go 1.14
```
# Go Modules
```
$ go mod init example/hello
go: creating new go.mod: module example/hello
$ cat go.mod
module example/hello

