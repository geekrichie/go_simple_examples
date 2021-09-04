package zerocopy

import (
	"testing"
)

func Test_sendfile(t *testing.T) {
	infile := "a.txt"
	outfile := "b.txt"
	err:= sendfile(infile, outfile)
	if err != nil {
		t.Log(err)
	}else {
		t.Log("send file success")
	}
}

func Test_mmap(t *testing.T) {
	err := mmap("a.txt")
	if err != nil {
		t.Log(err)
	}else {
		t.Log("send file success")
	}
}