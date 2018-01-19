// echo1 prints CL args
package main

import "fmt"
import "os"

func main() {
	// vars s, sep are string types
	// implicitly initialized empty ""
	var s, sep string

	// for initialization; condition; post-action
	for i := 1; i < len(os.Args); i++ {
		// empty sep, w/space added after first arg appended to s
		// ensures that spaces are only put in between the args
		// not before the first arg or after the last arg
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
