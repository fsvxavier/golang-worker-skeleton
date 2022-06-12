package udp

import "net"

type MessageType int

const (
	ControlMessage MessageType = iota
	TextMessage
	VoiceMessage
)

//some bytes associated with an address
type Packet struct {
	bytes         []byte
	returnAddress *net.UDPAddr
}

type Message struct {
	Type    MessageType
	Message []byte
}

type Server struct {
	Connection *net.UDPConn
	Client     *net.UDPAddr //or use map with an uuid

	Messages chan Message
	Packets  chan Packet
	Kill     chan bool
}

type Client struct {
	Connection *net.UDPConn
	Port       int

	Messages chan Message
	Packets  chan Packet
	Kill     chan bool
}

//create a new server.
func NewServer() *Server {
	return &Server{
		Packets:  make(chan Packet),
		Messages: make(chan Message),
		Kill:     make(chan bool),
	}
}

//create a new client.
func NewClient() *Client {
	return &Client{
		Packets:  make(chan Packet),
		Messages: make(chan Message),
		Kill:     make(chan bool),
	}
}
