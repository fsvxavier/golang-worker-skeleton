package udp

import (
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack"
)

func (s *Server) SetupServerConnection(address string) {

	//also listen from requests from the server on a random port
	listeningAddress, err := net.ResolveUDPAddr("udp4", address)
	ErrorCheck(err, "setupConnection", true)
	log.Info("...CONNECTED! ")

	s.Connection, err = net.ListenUDP("udp4", listeningAddress)
	ErrorCheck(err, "setupConnection", true)

	log.Printf("listening on: local:%s\n", s.Connection.LocalAddr())

}

func (s *Server) ReadFromSocket(buffersize int) {
	for {
		var b = make([]byte, buffersize)
		n, addr, err := s.Connection.ReadFromUDP(b[0:])
		ErrorCheck(err, "readFromSocket", false)

		s.Client = addr

		b = b[0:n]
		if n > 0 {
			pack := Packet{b, addr}
			select {
			case s.Packets <- pack:
				continue
			case <-s.Kill:
				break
			}
		}

		select {
		case <-s.Kill:
			break
		default:
			continue
		}
	}
}

func (s *Server) ProcessPackets() {
	for pack := range s.Packets {
		var msg Message
		err := msgpack.Unmarshal(pack.bytes, &msg)
		ErrorCheck(err, "processPackets", false)
		s.Messages <- msg
	}
}

func (s *Server) ReceiveMessages() {
	for msg := range s.Messages {
		if msg.Type == TextMessage {
			log.Printf("Received TXT : %s", msg.Message)
			// return string(msg.Message)
		}
	}
	// return ""
}

func (s *Server) SendMessage(message string) {

	msg := Message{
		Type:    TextMessage,
		Message: []byte(message),
	}

	b, err := msgpack.Marshal(msg)
	ErrorCheck(err, "Send", false)

	_, err = s.Connection.WriteToUDP(b, s.Client)
	ErrorCheck(err, "Send", false)

}
