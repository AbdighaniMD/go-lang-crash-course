package main

import "fmt"

//var name = "AbdighaniMD" //Global Variable

func main() {

	// var name = "AbdighaniMD"
	var myAge int32 = 23
	const isCool = true
	//isCool = false   // now allowed
	name, email := "AbdighaniMD", "abdighanimd@email.com"
	const size float32 = 5.9 // Constants

	fmt.Println(name, "\n", myAge, "\n", isCool, "\n", size, "\n", email)
	fmt.Printf("%T\n", name) // take the type from name

}

//ALL TYPES:
// string
// bool
// int
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte - alias for uint8
// rune - alias for int32
// float32 float64
// complex64 complex128
