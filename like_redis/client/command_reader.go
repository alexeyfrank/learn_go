package client

import (
	"bufio"
	"os"
	"strings"
)

type CommandReader struct {
	io     *os.File
	reader bufio.Reader
}

func NewCommandReader(io *os.File) *CommandReader {
	return &CommandReader{
		io:     io,
		reader: *bufio.NewReader(io),
	}
}

func (c *CommandReader) ReadCommand() (string, error) {
	str, err := c.reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.Trim(strings.Trim(str, " "), "\n"), nil
}
