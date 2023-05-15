package word

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {

	s := "Testing count function , will it work ?"

	res := Count(s)
	if res != 8 {
		t.Error("Expected: 8, got: ", res)
	}

}

func TestUseCount(t *testing.T) {

	s := "Testing count Testing , will it work Testing ?"

	res := UseCount(s)
	for k, v := range res {
		switch k {
		case ",":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "?":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "Testing":
			if v != 3 {
				t.Error("got", v, "want", 3)
			}
		case "count":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "it":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "will":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "work":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}

		}
	}

}

func ExampleCount() {
	fmt.Println(Count("Testing count function , will it work ?"))
	// Output:
	// 8
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < 100; i++ {
		Count("Testing count function , will it work ?")
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < 100; i++ {
		UseCount("Testing count function , will it work ?")
	}
}
