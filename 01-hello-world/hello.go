// This is a comment
package main

import (
	"fmt"

	"rsc.io/quote"
)

/*
This is a block comment
Print Hello World to the screen
*/
func main() {
	fmt.Println("Hello World")
	fmt.Println("Quote generated:-", quote.Go())
}
