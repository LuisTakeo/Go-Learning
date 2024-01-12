package main

import (
	"fmt"

	"rsc.io/quote"
)
// go mod init <module name> -> creates a go.mod file
// go mod tidy -> adds the required dependencies to the go.mod file

func main() {
	fmt.Println(quote.Go())
}
