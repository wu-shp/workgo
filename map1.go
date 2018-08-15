package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)

	// 往这个map里插入几条数据
	personDB["a"] = PersonInfo{"123231", "张大山", "天津市滨海新区"}
	personDB["b"] = PersonInfo{"213123124", "王小海", "沈阳市沈北新区"}

	person, ok := personDB["c"]
	if ok {
		fmt.Println("找到这个人", person.Name, "他的ID是a")
	} else {
		fmt.Println("没找到这个人")
	}
	sum := 0
	for {
		sum++
		if sum > 100 {
			break
		}
		fmt.Println(sum)
	}
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE
	}
	a := []int{1, 2, 3, 4, 5, 6}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

}
