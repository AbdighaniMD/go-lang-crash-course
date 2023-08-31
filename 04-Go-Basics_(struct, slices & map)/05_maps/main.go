package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// Map literals are like struct literals, but the keys are required.
var Map_literals = map[string]Vertex{
	//If the top-level type is just a type name, you can omit it from the elements of the literal.
	"Google": { // Vertex
		37.42202, -12.08408,
	},
}

func main() {

	// The zero value of a map is nil. A nil map has no keys, nor can keys be added.
	var createMap map[int]string
	fmt.Println(createMap)

	//Mutating Maps
	M := make(map[int]string)
	//Insert or update an element in map m:
	M[1] = "key1"
	M[2] = "key2"
	M[3] = "key3"
	fmt.Println(M)
	// Retrieve an element:
	elem := M[3]
	fmt.Println("Retrieve an element:", elem)
	delete(M, 2)
	fmt.Println("delete key2:", M[2], M)

	// Test that a key is present with a two-value assignment:
	value, ok := M[2]
	fmt.Println("The value:", value, "Present?", ok)

	// Map with Structs
	structMap := make(map[string]Vertex)
	structMap["Computer leb"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(structMap["Computer leb"])

	fmt.Println("map literals", Map_literals["Google"])

}
