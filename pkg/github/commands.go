package github

import (
	"fmt"
	"io"
)

const ErrorCommand = "::error::"

type ErrorWriter struct {
	io.Writer
}

func (ErrorWriter) Write(p []byte) (int, error) {
	n, err := fmt.Printf("$ErrorCommand%s", p)
	if n < len(ErrorCommand) {
		n = 0
	} else {
		n -= len(ErrorCommand)
	}

	return n, err
}
