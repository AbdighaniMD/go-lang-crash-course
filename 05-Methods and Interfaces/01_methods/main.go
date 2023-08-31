package main

import "fmt"

/*
Go does not have classes. However, you can define methods on types.
note: a method is just a function with a receiver argument.
*/

type additions struct {
	X, Y int32
}

// A method is a function with a special receiver argument.
func (v additions) add() int32 {
	return v.X + v.Y
}

/*
You can declare a method on non-struct types, too.
note:
cannot declare a method with a receiver whose type is defined in another package
*/
type MyInt int32

func (a MyInt) addInt() int32 {
	if a < 5 {
		return int32(-a)
	}
	return int32(a)
}

/*
You can declare methods with pointer receivers.
Methods with pointer receivers can modify the value to which the receiver points
*/
func (v *additions) Scale(f int32) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	result := additions{3, 4}
	fmt.Println(result.add())
	//pointer receivers
	result.Scale(10)
	fmt.Println(result.add())

	f := MyInt(3)
	fmt.Println(f.addInt())

}
