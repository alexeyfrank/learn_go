package client

import (
	"bufio"
	"os"
)

type ResponseWriter struct {
	io     *os.File
	writer bufio.Reader
}

func NewResponseWriter(io *os.File) *ResponseWriter {
	return &ResponseWriter{
		io: io,
	}
}

func (w *ResponseWriter) Write(str string) (int, error) {
	return w.io.Write([]byte(str))
}
