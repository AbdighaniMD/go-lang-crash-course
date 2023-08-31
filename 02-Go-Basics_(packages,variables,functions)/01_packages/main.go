package main // Package Declaration

//Import Packages
import (
	"fmt"
	"math"
	"math/rand"
)

/*
Every Go program is made up of packages.

Programs start running in package main.

This program is using the packages with import paths "fmt", "math/rand" & math.
*/

func main() {
	//Exported Name needs to start with a Capital Letter: Pi, Pizza, John, Test
	pi := math.Pi
	fmt.Println("Exported Name: =", pi)
	fmt.Println("My Favourite number is: =", rand.Intn(10))
}
