package stdlib

import (
	"container/list"
	"fmt"
)

func Stdlib_ContainerList() {
	queues := list.New()

	queues.PushBack("A")
	queues.PushBack("B")
	queues.PushBack("C")

	for q := queues.Front(); q != nil; q = q.Next() {
		fmt.Println(q.Value)
	}
}
