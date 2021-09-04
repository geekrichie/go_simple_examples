package bufio

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SplitByPad(input string) (int, error){
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return -1 , err
	}
	return count, nil
}

func BufferedIO() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "Hello, ")
	fmt.Fprint(w, "world!")
	w.Flush() // Don't forget to flush!
}