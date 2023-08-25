package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name", "wincher", "specify name")
	flag.Parse()
	fmt.Println("Hello", *name)
	fmt.Println(os.Args)

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	for i, c := range *name {
		fmt.Println(i, string(c))
	}

	myArray := [3]int{1, 2, 3}
	for i, item := range myArray {
		fmt.Println(i, item)
	}

	a := []int{1, 2, 3}
	c := append(a, 1)
	fmt.Println(a)
	fmt.Println(c)
}
