package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	t.Run("Sum 2 + 2", func(t *testing.T) {
		res := Sum(2, 2)
		assert.Equal(t, 4, res, "2 + 2 = 4")

		fmt.Println("Sum 1 is done.")
	})
	t.Run("Sum 1000 + 5", func(t *testing.T) {
		res := Sum(1000, 5)
		require.Equal(t, 105, res, "1000 + 5 = 105")

		fmt.Println("Sum 2 is done.")
	})
}
