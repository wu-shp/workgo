package main

import "fmt"

func main() {
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var mySlice []int = myArray[:5] //基于数组的前五个元素创建切片
	fmt.Println("数组的元素：")
	for _, v := range myArray {
		fmt.Print(v, " ")

	}
	fmt.Println("\n数组切片的元素：")
	for _, v := range mySlice {
		fmt.Print(v, " ")

	}
	fmt.Println()

}
