package main

import (
	"fmt"
)

type myMap map[string]string

func main() {

	mymap := make(myMap, 1)
	mymap["firstKey"] = "firstValue"
	fmt.Printf("Init nop: Address = %p Len = %d\n", &mymap, len(mymap))
	mymap.grow()
	fmt.Printf("Growed nop: Address = %p Len = %d\n", &mymap, len(mymap))

	mymap = make(myMap, 1)
	mymap["firstKey"] = "firstValue"
	fmt.Printf("Init p: Address = %p Len = %d\n", &mymap, len(mymap))
	(&mymap).growp()
	fmt.Printf("Growed p: Address = %p Len = %d\n", &mymap, len(mymap))

}

func (m myMap) grow() {
	for i := 1; i < 1000000; i++ {
		m[fmt.Sprintf("nopAddKey%d", i)] = fmt.Sprintf("%d", i)
	}
}

func (m *myMap) growp() {
	for i := 1; i < 1000000; i++ {
		(*m)[fmt.Sprintf("pAddKey%d", i)] = fmt.Sprintf("%d", i)
	}
}
