package main

import "fmt"

func main() {
	var slice []string = []string{"t1", "t2"}
	var m map[string]int = map[string]int{"test1": 1, "test2": 2}
	fmt.Printf("slice : %v\n", slice)

	m["test3"] = 3

	fmt.Printf("map : %v\n", m)
	delete(m, "test3") // remove item from map
	for k, v := range m {
		fmt.Printf("m[%s] : %d\n", k, v)
	}

	k, ok := m["xxx"]
	fmt.Println("The value m[xxx]:", k, "is exist:", ok) // zero value of int is zero
	m["xxx"] = 0
	k, ok = m["xxx"]
	fmt.Println("The value m[xxx]:", k, "is exist:", ok)

	var nilMap map[string]int
	if nilMap == nil {
		fmt.Println("nilMap is nil")
	}

	emptyMap := make(map[string]int) // book of memory
	if emptyMap == nil {
		fmt.Println("emptyMap is nil")
	} else {
		fmt.Println("emptyMap is not nil")
	}
	fmt.Println("emptyMap: ", emptyMap)
	fmt.Printf("emptyMap: %#v\n", emptyMap)
}
