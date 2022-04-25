package client

import (
	"bufio"
	"log"
	"net"
)

type Client struct {
	addr           string
	conn           net.Conn
	commandReader  *CommandReader
	responseWriter *ResponseWriter
}

func NewClient(addr string, commandReader *CommandReader, responseWriter *ResponseWriter) (*Client, error) {
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		return nil, err
	}

	return &Client{
		addr:           addr,
		conn:           conn,
		commandReader:  commandReader,
		responseWriter: responseWriter,
	}, nil
}

func (c *Client) Run() {
	defer c.conn.Close()

	connReader := bufio.NewReader(c.conn)

	for {
		c.responseWriter.Write("> ")
		command, err := c.commandReader.ReadCommand()

		if err != nil {
			log.Fatal(err)
		}

		if len(command) == 0 {
			continue
		}

		if _, err := c.conn.Write([]byte(command + "\n")); err != nil {
			log.Fatal(err)
		}

		response, err := connReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		c.responseWriter.Write(response)
		c.responseWriter.Write("\n")
	}
}
