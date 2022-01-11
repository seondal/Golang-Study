package main

import (
	"container/ring"
	"fmt"
)

func main() {
	//링 생성
	r := ring.New(5)

	//링 길이 반환
	n := r.Len()

	for i := 0; i < n; i++ {
		r.Value = 'A' + i
		r = r.Next()
	}

	//출력
	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value)
		r = r.Next()
	}

	fmt.Println()

	//역순출력
	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value)
		r = r.Prev()
	}
}
