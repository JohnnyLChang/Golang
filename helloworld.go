package main

import "fmt"
import "math"

const limit int = 10

var a [limit]int64

//BasicTypeConversions : for Basic Type conversion example
func BasicTypeConversions() {
	var i int32
	i = int32(math.Sqrt(25)) //explicit conversion
	fmt.Println("value of i:", i)
}

//ArrayAndSlice : for Basic Array and slice example
func ArrayAndSlice() {
	for i := 0; i < limit; i++ {
		a[i] = int64(i * i)
	}
	var b = a[3:8] // Get slice array from array a
	for j := 0; j < 5; j++ {
		fmt.Printf("%d : %d\n", j, b[j])
	}
	c := make([]byte, 10, 15)
	fmt.Printf("Slice Array c, size:%d, cap:%d", len(c), cap(c))

	d := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println(d)
}

func main() {
	fmt.Println("hello world")
	fmt.Println("hello Johnny")
	BasicTypeConversions()
	//countdown()
	ArrayAndSlice()
	testSlice()
	fmt.Println("fibo recrusive:", fibonacciRecursive(10))
	fmt.Println("fibo recrusive:", fibonacci(10))
	fmt.Println("test finished")
	return
}
