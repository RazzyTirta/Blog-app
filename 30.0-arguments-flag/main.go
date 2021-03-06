package main

import (
	"fmt"
	// "os"
	"flag"
)

// arguments
// func main() {
// 	var argsRaw = os.Args
// 	fmt.Printf("-> %v\n", argsRaw)

// 	var args = argsRaw[1:]
// 	fmt.Printf("-> %v\n", args)
// }

// flag
func main() {
	name := flag.String("name", "anonymous", "type your name")
	age := flag.Int64("age", 25, "type your age")

	flag.Parse()
	fmt.Printf("name\t: %s\n", *name)
	fmt.Printf("age\t: %d\n", *age)
}