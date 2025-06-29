package stdlib

import (
	"container/ring"
	"fmt"
	"strconv"
)

func Stdlib_ContainerRing() {
	var rings *ring.Ring = ring.New(5)

	for r := 1; r <= rings.Len(); r++ {
		rings.Value = "Rings of #" + strconv.Itoa(r)
		rings = rings.Next()
	}

	rings.Do(func(r any) {
		fmt.Println(r)
	})
}
