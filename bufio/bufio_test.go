package bufio

import (
	"os"
	"testing"
)

func TestSplitByPad(t *testing.T) {
	SplitByPad("212131  45saffs afdafa")
}

func TestBufferedIO(t *testing.T) {
	t.Log(os.Executable())
}