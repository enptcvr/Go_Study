package main

import "fmt"

func main() {
	fmt.Print("My weight on the surface of mars is:")
	fmt.Print(149 * 0.3843)
	fmt.Print(" lbs, and i would be ")
	fmt.Print(41 * 365 / 687)
	fmt.Print(" years old ")
	fmt.Println("")
	fmt.Print("My weight on surface of Mars is ", 149*0.3843, " lbs, and i would be ", 41*365/687, " years old \n")

	fmt.Printf("My weight on the surface of Mars is %v lbs.", 149*0.3843)
	fmt.Printf("And i would be %v years old.\n", 41*365/687)
	fmt.Printf("My weight on the surface of Mars is %v And i would be %v years old.\n", 149*0.3843, 41*365/687)

}
