package main

import "fmt"

func main() {

	//맵 생성
	m := make(map[int]string)

	//키와 값 추가
	m[1] = "일"
	m[2] = "이"
	m[3] = "삼"

	delete(m, 3)

	v3, ok3 := m[3]
	fmt.Println(v3)
	fmt.Println(ok3)

	v1, ok1 := m[1]
	fmt.Println(v1)
	fmt.Println(ok1)
}
