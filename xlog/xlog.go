package xlog

import (
	"bufio"
	"fmt"
	"os"
)

var out *bufio.Writer

func InitLogger(file *os.File) {
	out = bufio.NewWriter(file)
}

func Debug(msg string, args ...any) {
	fmt.Fprintf(out, msg, args...)
	out.Flush()
}
