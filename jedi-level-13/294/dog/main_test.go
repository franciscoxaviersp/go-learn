package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {

	type test struct {
		data   int
		answer int
	}
	tests := []test{
		{1, 7},
		{2, 14},
		{3, 21},
		{4, 28},
		{5, 35},
	}

	for _, v := range tests {
		res := Years(v.data)
		if res != v.answer {
			t.Error("Expected: ", v.answer, ", got: ", res)
		}
	}

}

func TestYearsTwo(t *testing.T) {

	type test struct {
		data   int
		answer int
	}
	tests := []test{
		{1, 7},
		{2, 14},
		{3, 21},
		{4, 28},
		{5, 35},
	}

	for _, v := range tests {
		res := YearsTwo(v.data)
		if res != v.answer {
			t.Error("Expected: ", v.answer, ", got: ", res)
		}
	}

}

func ExampleYears() {
	fmt.Println(Years(2))
	// Output:
	// 14
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(2))
	// Output:
	// 14
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < 100; i++ {
		Years(12)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < 100; i++ {
		YearsTwo(12)
	}
}
