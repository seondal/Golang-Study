package main

import (
	"container/list"
	"fmt"
)

//Queue 구조체 정의
type Queue struct {
	v *list.List
}

//요소 추가
func (q *Queue) Push(val interface{}) {
	q.v.PushBack(val)
}

//맨 앞 요소 반환 후 삭제
func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

// 새로운 큐 인스턴스 생성
func NewQueue() *Queue {
	return &Queue{list.New()}
}

func main() {
	queue := NewQueue()

	for i := 1; i < 5; i++ {
		queue.Push(i)
	}
	v := queue.Pop()
	for v != nil {
		fmt.Printf("%v -> ", v)
		v = queue.Pop()
	}
}
