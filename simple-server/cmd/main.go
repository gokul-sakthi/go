package main

import (
	"fmt"
	"log"
	"net"
)

type Config struct {
	Port string
	Host string
}

type Message struct {
	from    string
	payload []byte
}

type Server struct {
	lnAddr  string
	ln      net.Listener
	quitch  chan struct{}
	msgchan chan Message
}

func NewServer(lnAddr string) *Server {

	return &Server{
		lnAddr:  lnAddr,
		quitch:  make(chan struct{}),
		msgchan: make(chan Message, 10),
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 2048)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Error read buf (%s): %v", conn.RemoteAddr(), err)
			continue
		}

		s.msgchan <- Message{
			from:    conn.RemoteAddr().String(),
			payload: buf[:n],
		}
	}

}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("Error reading conn", err)
			continue
		}

		go s.readLoop(conn)

	}

}

func (s *Server) Start() {
	fmt.Println("Server addr", s.lnAddr)
	ln, err := net.Listen("tcp", s.lnAddr)

	if err != nil {
		log.Fatalln("Error binding port", err)
	}

	s.ln = ln
	defer s.ln.Close()

	go s.acceptLoop()

	go func() {
		for msg := range s.msgchan {
			fmt.Printf("New message received (%v): %v\n", msg.from, string(msg.payload))
		}
	}()

	<-s.quitch
	close(s.msgchan)

}

func main() {
	fmt.Println("Initializing!")
	run(Config{Host: "localhost", Port: "3000"})
}

func run(cfg Config) {
	addr := net.JoinHostPort(cfg.Host, string(cfg.Port))
	server := NewServer(addr)
	server.Start()

}
