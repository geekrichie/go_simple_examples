package tar

import (
	"testing"
)

func TestPractice(t *testing.T) {
	err := TarOneFile("../README.md")
	t.Log(err)
}

func TestTarMultiFile(t *testing.T) {
	err := TarMultiFile("mutli.test","../README.md", "../README.zh-CN.md")
	t.Log(err)
}

func TestUnTarFile(t *testing.T) {
	err := UnTarFile("mutli.test.tar","../simpleevent")
	t.Log(err)
}