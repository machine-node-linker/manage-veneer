package github

import (
	"fmt"
	"io"
)

type ErrorWriter struct {
	io.Writer
}

func (ErrorWriter) Write(p []byte) (n int, err error) {
	n, err = fmt.Printf("::error::%s", p)
	if n < 9 {
		n = 0
	} else {
		n = n - 9
	}
	return n, err
}
