package main

import "fmt"


func sliceInfo( data []int ){
	fmt.Printf("address=%p  len=%v  capacity=%v  data=%v  representation:%#v  \n", data, len(data), cap(data), data, data)
}

func closureWithSlice(data []int) func(){
	if data!=nil && len(data)>0 {
		data[0]=3
	}
	var innerData = data
	return func(){
		fmt.Printf("closure data: address=%p  representation:%#v  len=%v  capacity=%v  data=%v \n", innerData, innerData, len(innerData), cap(innerData), innerData)
	}
}

func main(){
	var originalData = []int{7,8,9,10,11}
	sliceInfo(originalData)

	sliceInfo(originalData[3:])
	sliceInfo(originalData[3:])

	sliceInfo(originalData[4:5])

	sliceInfo(originalData[:1])

	sliceInfo(originalData[:0])


	closureWithSlice(originalData)()

	closureWithSlice(originalData[3:])()
	closureWithSlice(originalData[3:])()

	closureWithSlice(originalData[4:5])()

	closureWithSlice(originalData[:1])()

	closureWithSlice(originalData[:0])()


	sliceInfo(make([]int, 5))
	var doubleSlice = [][]int{
		[]int{1,2,3,4},
		[]int{5,6,7,8},
	}
	fmt.Println(doubleSlice)

	sliceInfo(append(originalData, 44))
	sliceInfo(append(originalData, 55))

	for index, value := range originalData {
		fmt.Println(index, value)
	}
}
