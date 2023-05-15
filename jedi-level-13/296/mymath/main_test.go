package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}

	res := CenteredAvg(list)
	if res != 3 {
		t.Error("Expected: 3, got: ", res)
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 3, 4, 5}))
	// Output:
	// 3
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < 100; i++ {
		CenteredAvg([]int{1, 2, 3, 4, 5})
	}
}
