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
