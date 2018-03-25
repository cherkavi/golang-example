package main

import (
	"fmt"
	"sort"
)

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

	fmt.Printf("\n---- remove element  ---- \n")
	var indexForRemove = 3
	sliceInfo(originalData)
	var newOriginalData = append(originalData[0:indexForRemove], originalData[indexForRemove+1:]...)
	sliceInfo(newOriginalData)
	sliceInfo(originalData)

	fmt.Printf("\n---- remove element  ---- \n")
	intSlice := []int{5, 4, 3, 1, 2}
	fmt.Printf("before sorting: %v \n", intSlice)
	sort.Ints(intSlice)
	fmt.Printf("after sorting: %v \n", intSlice)

	strSlice := []string{"b", "c", "a", "e"}
	fmt.Printf("before sorting: %v \n", strSlice)
	sort.Strings(strSlice)
	fmt.Printf("after sorting: %v \n", strSlice)

	dataArray := DataArray{[]Data{Data{"d"}, Data{"b"}, Data{"a"}, Data{"c"}}}
	fmt.Printf("before sorting: %v \n", dataArray)
	sort.Sort(dataArray)
	fmt.Printf("after sorting: %v \n", dataArray)
}

type Data struct {
	name string
}
type DataArray struct {
	elements []Data
}

// Len is the number of elements in the collection.
func (d DataArray) Len() int {
	return len(d.elements)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (d DataArray) Less(i, j int) bool {
	return d.elements[i].name < d.elements[j].name
}

// Swap swaps the elements with indexes i and j.
func (d DataArray) Swap(i, j int) {
	d.elements[i].name, d.elements[j].name = d.elements[j].name, d.elements[i].name
}
