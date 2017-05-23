package main

import (
	"fmt"
)

type myMap map[string]string

func main() {

	mymap := make(myMap, 1)
	mymap["firstKey"] = "firstValue"
	fmt.Printf("Init method nop: Address = %p Len = %d\n", &mymap, len(mymap))
	mymap.grow()
	fmt.Printf("Growed method nop: Address = %p Len = %d\n", &mymap, len(mymap))

	mymap = make(myMap, 1)
	mymap["firstKey"] = "firstValue"
	fmt.Printf("Init method p: Address = %p Len = %d\n", &mymap, len(mymap))
	(&mymap).growp()
	fmt.Printf("Growed method p: Address = %p Len = %d\n", &mymap, len(mymap))

	mymap = make(myMap, 1)
	mymap["firstKey"] = "firstValue"
	fmt.Printf("Init func nop: Address = %p Len = %d\n", &mymap, len(mymap))
	fgrow(mymap)
	fmt.Printf("Growed func nop: Address = %p Len = %d\n", &mymap, len(mymap))

	mymap = make(myMap, 1)
	mymap["firstKey"] = "firstValue"
	fmt.Printf("Init func p: Address = %p Len = %d\n", &mymap, len(mymap))
	fgrowp(&mymap)
	fmt.Printf("Growed func p: Address = %p Len = %d\n", &mymap, len(mymap))

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

func fgrow(m myMap) {
	for i := 1; i < 1000000; i++ {
		m[fmt.Sprintf("nopAddKey%d", i)] = fmt.Sprintf("%d", i)
	}
}

func fgrowp(m *myMap) {
	for i := 1; i < 1000000; i++ {
		(*m)[fmt.Sprintf("pAddKey%d", i)] = fmt.Sprintf("%d", i)
	}
}
