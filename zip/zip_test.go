package zip

import "testing"

func TestZipOneFile(t *testing.T) {
	err := ZipOneFile("../README.md")
	t.Log(err)
}
