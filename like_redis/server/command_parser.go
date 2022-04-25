package server

import (
	"bufio"
	"errors"
	"net"
	"strings"
)

type CommandParser struct {
	conn   *net.Conn
	reader *bufio.Reader

	commands []interface{}
	isEOF    bool
}

func NewCommandParser(conn *net.Conn) *CommandParser {
	reader := bufio.NewReader(*conn)

	return &CommandParser{
		conn:   conn,
		reader: reader,
	}
}

func (c *CommandParser) Read() (interface{}, error) {
	command, err := c.reader.ReadString('\n')

	if err != nil {
		return nil, err
	}

	command = strings.Trim(strings.Trim(command, "\n"), " ")

	parsedCommand, err := parse(command)
	if err != nil {
		return nil, err
	}

	return parsedCommand, nil
}

func parse(cmd string) (interface{}, error) {
	tokens := strings.Split(cmd, " ")

	if len(tokens) == 0 {
		return nil, errors.New("Empty command")
	}

	name := tokens[0]
	rawArgs := []string{}

	if len(tokens) > 1 {
		rawArgs = tokens[1:]
	}

	switch name {
	case "SET":
		return NewSetCommandDefinition(rawArgs)
	case "GET":
		return NewGetCommandDefinition(rawArgs)
	case "INC":
		return NewIncCommandDefinition(rawArgs)
	case "DEC":
		return NewDecCommandDefinition(rawArgs)
	}

	return NewNotFoundCommandDefinition(rawArgs)
}
