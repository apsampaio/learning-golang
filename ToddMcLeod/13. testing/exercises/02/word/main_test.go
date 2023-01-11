package word

import (
	"exercises/02/quote"
	"fmt"
	"testing"
)

func TestUseCount(t *testing.T) {
	m := UseCount("one two three three")

	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("expected", 1, "got", v)
			}
		case "two":
			if v != 1 {
				t.Error("expected", 1, "got", v)
			}
		case "three":
			if v != 2 {
				t.Error("expected", 2, "got", v)
			}

		}
	}
}

func ExampleCount() {
	fmt.Println(Count("one two three"))
	// Output:
	// 3
}

func TestCount(t *testing.T) {
	i := Count("one two three")
	if i != 3 {
		t.Error("expected", 3, "got", i)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}
