package main

import "fmt"

func main()  {
	//var m []int
	//var k int
	//for k = 0; k < 10; k++{
	//	//fmt.Println(k)
	//	m = append(m,k + 200)
	//	fmt.Printf("Element[%d] = %d\n", k, m[k])
	//}

	// findElementWithSum of k from arr of size
	//var m int
	//var k int
	//for k = 0; k < 10; k++{
	//	//fmt.Println(k)
	//	m = k * 200
	//	fmt.Printf("Element[%d] = %d\n", k, m)
	//}

	// Two dimensional slices
	var rows int
	var cols int
	rows =7
	cols = 9
	var twoSlices =make([][]int, rows)
	var i int
	for i =range twoSlices{
		twoSlices[i] = make([]int, cols)
	}
	fmt.Println(twoSlices)

}