package main

func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{21, 21}, 42},
		test{[]int{3, 3}, 6},
		test{[]int{1, 1}, 2},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}

}
