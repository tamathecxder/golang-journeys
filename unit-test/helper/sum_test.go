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

type Assignment struct {
	Name               string
	RequestX, RequestY int
	Expected           int
	Description        string
}

func TestTableTest(t *testing.T) {
	list := []Assignment{
		{
			Name:        "20 + 20",
			RequestX:    20,
			RequestY:    20,
			Expected:    40,
			Description: "The result must be '40'",
		},
		{

			Name:        "100 + 90",
			RequestX:    100,
			RequestY:    90,
			Expected:    190,
			Description: "The result must be '190'",
		},
		{
			Name:        "1 + 1",
			RequestX:    1,
			RequestY:    1,
			Expected:    2,
			Description: "The result must be '2'",
		},
	}

	for _, v := range list {
		t.Run(v.Name, func(t *testing.T) {
			assert.Equal(t, v.Expected, Sum(v.RequestX, v.RequestY), v.Description)
			fmt.Printf("Test with name '%s' is success.\n", v.Name)
		})
	}
}
