// echo2 prints CL args
package main

import "fmt"
import "os"

func main() {
	// short variable declarations
        // w/multiple assignments in list order
        s, sep := "", ""

	// better for loop standard for Go programs
        // _ is a blank identifier, used when indexes need to be assigned
        // but not used, as the elements/values of that array do
        // range iterates implicitly - no need to eval length beforehand
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
