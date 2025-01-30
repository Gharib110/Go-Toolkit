package toolkit

import "testing"

func TestTools_RandomString(t *testing.T) {
	tk := &Tools{}

	s := tk.RandomString(10)
	if len(s) != 10 {
		t.Errorf("Expected string length to be 10, got %d", len(s))
	}
}
