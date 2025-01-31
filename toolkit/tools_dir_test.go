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
