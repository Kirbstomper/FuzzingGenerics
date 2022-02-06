package genericmath

import "testing"

func TestAddInt(t *testing.T) {
	ans := Add(int(3), int(5))

	if ans != 8 {
		t.Fail()
	}

}

func TestAddFloat64(t *testing.T) {
	ans := Add(float64(2.6), float64(1.5))

	if ans != 4.1 {
		t.Fail()
	}
}

func FuzzAddInt(f *testing.F) {
	testcases := []int{1, 2, 3, 4, 5, 67, 1231}

	for _, tc := range testcases {
		f.Add(tc, tc)
	}

	f.Fuzz(func(t *testing.T, a, b int) {
		ans := a + b
		if Add(a, b) != ans {
			t.Errorf("A: %v + B: %v should be equal to: %v", a, b, ans)
		}
	})
}
