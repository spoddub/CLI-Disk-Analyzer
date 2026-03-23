package code

import "testing"

func TestGetPathSize_File(t *testing.T) {
	got, err := GetPathSize("testdata/data.csv", false, true, false)
	if err != nil {
		t.Fatal(err)
	}

	want := "416B"

	if got != want {
		t.Errorf("file data.csv size should be %s, got %s", want, got)
	}
}

func TestGetPathSize_Dir(t *testing.T) {
	got, err := GetPathSize("testdata/dir1", false, true, false)
	if err != nil {
		t.Fatal(err)
	}

	want := "3.8KB"

	if got != want {
		t.Errorf("dir testdata/dir1 size is %s, got %s", want, got)
	}
}
