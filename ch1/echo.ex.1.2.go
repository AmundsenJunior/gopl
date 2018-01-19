// echo.ex.1.2 prints CL indexes and args
package main

import "fmt"
import "os"

func main() {
	// better for loop standard for Go programs
        // get each index and value from the Args array
        // range iterates implicitly - no need to eval length beforehand
	for idx, arg := range os.Args[1:] {
                fmt.Println(idx, ":", arg) 
        }
}
