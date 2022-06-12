package udp

import (
	"net"

	log "github.com/sirupsen/logrus"

	"time"

	"github.com/vmihailenco/msgpack"
)

const ALIVE_CHECK_TIME = time.Second * 10

func (c *Client) SetupConnection(address string) {
	addr, err := net.ResolveUDPAddr("udp4", address)

	ErrorCheck(err, "setupConnection", true)
	log.Printf("> server address: %s ... connecting ", addr.String())

	conn, err := net.DialUDP("udp4", nil, addr)
	c.Connection = conn

	//also listen from requests from the server on a random port
	listeningAddress, err := net.ResolveUDPAddr("udp4", ":0")
	ErrorCheck(err, "setupConnection", true)
	log.Printf("...CONNECTED! ")

	conn, err = net.ListenUDP("udp4", listeningAddress)
	ErrorCheck(err, "setupConnection", true)

	log.Printf("listening on: local:%s\n", conn.LocalAddr())

}

func (c *Client) ReadFromSocket(buffersize int) {
	for {
		var b = make([]byte, buffersize)
		n, addr, err := c.Connection.ReadFromUDP(b[0:])
		ErrorCheck(err, "readFromSocket", false)

		b = b[0:n]

		if n > 0 {
			pack := Packet{b, addr}
			select {
			case c.Packets <- pack:
				continue
			case <-c.Kill:
				break
			}
		}

		select {
		case <-c.Kill:
			break
		default:
			continue
		}
	}
}

func (c *Client) ProcessPackets() {
	for pack := range c.Packets {
		var msg Message
		err := msgpack.Unmarshal(pack.bytes, &msg)
		ErrorCheck(err, "processPackets", false)
		c.Messages <- msg
	}
}

func (c *Client) ProcessMessages() {
	for msg := range c.Messages {
		if msg.Type == TextMessage {
			log.Printf("Received TXT : %s", msg.Message)
		}
		if msg.Type == VoiceMessage {
			panic("todo:// voice message :)")
		}
	}
}

func (c *Client) Send(message string) {

	msg := Message{
		Type:    TextMessage,
		Message: []byte(message),
	}

	b, err := msgpack.Marshal(msg)
	ErrorCheck(err, "Send", false)

	_, err = c.Connection.Write(b)
	ErrorCheck(err, "Send", false)

}
