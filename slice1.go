package main

import "fmt"

func main() {
	mySlice := make([]int, 5, 10)
	fmt.Println("a", len(mySlice)) //数组切片中当前储存元素的个数
	fmt.Println("b", cap(mySlice)) //数组切片分配空间的大小

}
