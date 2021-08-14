package main

import "fmt"

func main() {
	fmt.Printf("%-20v $%5v\n", "space", 99)
	fmt.Printf("%-20v $%5v", "Hello", 1000)
}
