package main

import (
	"fmt"
	"time"
)

/*func main() {
	numSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(length1(numSlice))

	addToSlice(&numSlice, 6)

	fmt.Println(length2(numSlice))

	fmt.Println(keys(map[string]int{"key1": 1, "key2": 2, "key3": 3}))
}

func length1(slice []int) int {
	var sliceLength int = 0

	for i := 0; i < len(slice); i++ {
		sliceLength += 1
	}

	return sliceLength
}

func length2(slice []int) int {
	sliceLength := len(slice)

	return sliceLength
}

func addToSlice(slice *[]int, x int) {
	*slice = append(*slice, x)
}

func keys(keysMap map[string]int) []string {
	var keys []string

	for k := range keysMap {
		keys = append(keys, k)
	}

	return keys
}*/

// make() is used for creating slices and maps. is

type person struct {
	name string
	age  int
}

type family struct {
	members []person
}

func main() {
	familys := map[string]family{
		"Canudas": {
			members: []person{
				{name: "Oscar", age: 50},
				{name: "Michelle", age: 50},
				{name: "Emma", age: 16},
				{name: "Michael", age: 15},
				{name: "Juliana", age: 13}}},
		"Capra": {
			members: []person{
				{name: "Tom", age: 50},
				{name: "Jen", age: 50},
				{name: "Emily", age: 18},
				{name: "John", age: 16},
				{name: "Lila", age: 12}}}}

	go func() {
		logfamily(familys["Canudas"])
	}()
	go func() {
		logfamily(familys["Capra"])
	}()
}

func logfamily(f family) {
	for i := 0; i < len(f.members); i++ {
		fmt.Println(f.members[i].name)
		time.Sleep(time.Second)
	}
}
