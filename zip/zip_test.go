package zip

import "testing"

func TestZipOneFile(t *testing.T) {
	err := ZipOneFile("../README.md")
	t.Log(err)
}

func TestZipMultiFile(t *testing.T) {
	err := ZipMultiFile("zip_test.zip", []string{"../go.mod", "zip.go"})
	t.Log(err)
}

func TestUnzipFile(t *testing.T) {
	err := UnzipFile("zip_test.zip",".")
	t.Log(err)
}