package dog

import "testing"

func TestYears(t *testing.T) {
	y := Years(2)
	if y != 14 {
		t.Error("expected", 14, "got", y)
	}
}

func TestYearsTwo(t *testing.T) {
	y := YearsTwo(2)
	if y != 14 {
		t.Error("expected", 14, "got", y)
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(2)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(2)
	}
}
