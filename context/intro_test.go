package context

import (
	"context"
	"fmt"
	"testing"
)

func TestContextIntroduction(t *testing.T) {
	bg := context.Background()
	fmt.Println(bg)

	todo := context.TODO()
	fmt.Println(todo)
}
