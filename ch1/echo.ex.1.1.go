// echo.ex.1.1 prints CL cmd & args
package main

import "fmt"
import "os"
import "strings"

// go run echo.ex.1.1.go [args] prints full path of go run obj
// go build echo.ex.1.1.go; ./echo.ex.1.1.go [args] prints just cmd name
func main() {
        // concat array slices by separator and print
	fmt.Println(strings.Join(os.Args[0:], " "))
}
