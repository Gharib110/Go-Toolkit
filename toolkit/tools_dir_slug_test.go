package toolkit

import "testing"

func TestTools_CreateDirIfNotExist(t *testing.T) {
	tk := &Tools{}

	err := tk.CreateDirIfNotExist("./test_dir_func")
	if err != nil {
		t.Error(err)
		return
	}

	err = tk.CreateDirIfNotExist("./test_data")
	if err != nil {
		t.Error(err)
		return
	}

}

var slugifyTests = []struct {
	name        string
	s           string
	expected    string
	errExpected bool
}{
	{name: "valid string", s: "now is the time", expected: "now-is-the-time", errExpected: false},
	{name: "empty string", s: "", expected: "", errExpected: true},
	{name: "complex string", s: "Now Good ! + fish & such &^123", expected: "now-good-fish-such-123", errExpected: false},
}

func TestTools_Slugify(t *testing.T) {
	testTools := &Tools{}
	for _, e := range slugifyTests {
		slug, err := testTools.Slugify(e.s)
		if !e.errExpected && err != nil {
			t.Error(err, "unexpected error", e.name)
			return
		}

		if !e.errExpected && slug != e.expected {
			t.Errorf("Expected %s, got %s", e.expected, slug)
			return
		}
	}
}
