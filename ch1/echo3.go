// echo3 prints CL args
package main

import "fmt"
import "os"
import "strings"

func main() {
        // concat array slices by separator and print
	fmt.Println(strings.Join(os.Args[1:], " "))
}
