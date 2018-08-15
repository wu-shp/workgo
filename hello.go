package main

import "fmt"

// 我们需要使用fmt包中的Println()函数


func main() {
	/*str := "Hello,世界"
	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i] // 依据下标取字符串中的字符，类型为byte
		fmt.Println(i, ch)
	}*/
	str1 := "Hello,世界"
	for i, ch := range str1 {
		fmt.Println(i, ch)//ch的类型为rune
	}
}
