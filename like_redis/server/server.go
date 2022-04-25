package server

import (
	"fmt"
	"io"
	"log"
	"net"
)

type Server struct {
	addr     string
	listener net.Listener
	storage  *Storage
}

func NewServer(addr string, storage *Storage) *Server {
	server := Server{
		addr:    addr,
		storage: storage,
	}

	return &server
}

func (s *Server) Run() {
	listener, err := net.Listen("tcp", s.addr)

	s.listener = listener

	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Accept new conn")

		go s.Handle(conn)
	}
}

func (s *Server) Handle(conn net.Conn) {
	defer conn.Close()
	parser := NewCommandParser(&conn)
	executor := NewCommandExecutor(s.storage)

	for {
		command, err := parser.Read()

		if err == io.EOF {
			log.Printf("Connection is closed\n")
			return
		}

		if err != nil {
			handleError(conn, err)
			continue
		}

		log.Printf("%v\n", command)

		result, err := executor.Execute(command)

		if err != nil {
			handleError(conn, err)
			continue
		}

		conn.Write([]byte(fmt.Sprintf("%v\n", result)))
	}
}

func handleError(conn net.Conn, err error) {
	conn.Write([]byte(fmt.Sprintf("%v\n", err)))
	log.Printf("Error: %v\n", err)
}
