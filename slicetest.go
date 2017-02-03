package main

import "fmt"

var g_storage []string

//add_grocery: using varadic slice parameter for function
func addGrocery(items ...string) {
	fmt.Println("g_storage capacity:", cap(g_storage), " size:", len(g_storage))
	for _, d := range items {
		g_storage = append(g_storage, d)
	}
}

func listGrocery() {
	fmt.Println("Print all items in grocery:")
	/*for i := 0; i < len(g_storage); i++ {
		fmt.Println("item ", i, ":", g_storage[i])
	}*/

	for i, d := range g_storage {
		fmt.Println(i, d)
	}
}

func testSlice() {
	addGrocery("johnny", "張立帆", "life")
	listGrocery()
	addGrocery("Annie", "謝志萍", "teacher")
	listGrocery()
}
