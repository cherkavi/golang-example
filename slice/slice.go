package main

import "fmt"

func sliceInfo(data []int) {
	fmt.Printf("address=%p  len=%v  capacity=%v  data=%v  representation:%#v  \n", data, len(data), cap(data), data, data)
}

func closureWithSlice(data []int) func() {
	if data != nil && len(data) > 0 {
		data[0] = 3
	}
	var innerData = data
	return func() {
		fmt.Printf("closure data: address=%p  representation:%#v  len=%v  capacity=%v  data=%v \n", innerData, innerData, len(innerData), cap(innerData), innerData)
	}
}

func changeFirstElement(data []int, newValue int) {
	data[0] = newValue
}

func main() {
	var originalData = []int{7, 8, 9, 10, 11}
	// var originalData = make([]int, 5, 10)
	originalData = append(originalData, 7, 8, 9, 10, 11)

	fmt.Printf("\n---- part of slices ---- \n")
	sliceInfo(originalData)

	sliceInfo(originalData[3:])
	sliceInfo(originalData[3:])

	sliceInfo(originalData[4:5])

	sliceInfo(originalData[:1])

	sliceInfo(originalData[:0])

	fmt.Printf("\n---- closure with slice ---- \n")
	closureWithSlice(originalData)()

	closureWithSlice(originalData[3:])()
	closureWithSlice(originalData[3:])()

	closureWithSlice(originalData[4:5])()

	closureWithSlice(originalData[:1])()

	closureWithSlice(originalData[:0])()

	fmt.Printf("\n---- empty slice ---- \n")
	sliceInfo(make([]int, 5, 10))

	fmt.Printf("\n---- slice in slice = doubleSlice ---- \n")
	var doubleSlice = [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
	}
	fmt.Println(doubleSlice)

	fmt.Printf("\n---- slice access via 'for' ---- \n")
	for index, value := range originalData[0:5] {
		fmt.Println(index, value)
	}

	fmt.Printf("\n---- change elements into sub-slice NOT via pointer ---- \n")
	changeFirstElement(originalData, -99)
	sliceInfo(originalData)
	changeFirstElement(originalData[1:], -98)
	sliceInfo(originalData)
	changeFirstElement(originalData[2:], -97)
	sliceInfo(originalData)

	fmt.Printf("\n---- slice elements over 'len' under 'cap' ---- \n")
	originalData = make([]int, 0, 10)
	originalData = append(originalData, 1, 2, 3, 4, 5)
	sliceInfo(originalData)
	sliceInfo(originalData[0:cap(originalData)])

	fmt.Printf("\n---- concantenate slices ---- \n")
	anotherData := []int{9, 8, 7, 6, 5}
	sliceInfo(append(originalData, anotherData[:2]...))
}
