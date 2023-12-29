package interaction

import (
	"bufio"
	"os"
	"strings"
)

type InputReader struct {
	reader *bufio.Reader
}

func New() *InputReader {
	return &InputReader{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (ir *InputReader) GatherInput() string {
	input, _ := ir.reader.ReadString('\n')
	return strings.TrimSpace(strings.Replace(input, "\n", "", -1))
}
