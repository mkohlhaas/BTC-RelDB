// Package main is the main entry point for the application.
package main

import (
	"fmt"
	"io"
	"net"
	"os"

	p "github.com/mkohlhaas/btcreldb/protocol"
)

// Client struct
type Client struct {
	Host string
	Port int
}

// TODO: - send version messages to DNS seeds
//       - add wtxidrelay messages
//       - add sendaddrv2 messages
//       - add ping and pong messages
//       - add addr and getaddr messages??? (possibly obsolete)
//       - check checksums
//       - check format of IPv6 address

func main() {
	client := &Client{
		Host: os.Args[1],
		Port: 8333,
	}
	client.Start()
}

// Start TCPClient
func (c *Client) Start() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.Host, c.Port)) // establish TCP connection
	p.Chk(err)                                                         // check error
	defer conn.Close()                                                 // close connection when we are done

	m2 := p.NewMsgVersion(p.IPAddr(c.Host)) // create version message
	e := p.Encode(m2)                       // encode to Bitcoin's network protocol
	conn.Write(e)                           // write out the message

	for {
		msgHeaderBuf := make([]byte, p.MsgHeaderSize)
		_, err := io.ReadFull(conn, msgHeaderBuf)                      // read header
		p.Chk(err)                                                     // check errror
		msgHeader := new(p.MsgHeader)                                  // create new message header
		p.Decode(msgHeaderBuf, msgHeader)                              // decode into header
		msgPayloadBuf := make([]byte, msgHeader.Length)                //
		_, err = io.ReadFull(conn, msgPayloadBuf)                      // read payload
		p.Chk(err)                                                     // check errror
		msg := append(msgHeaderBuf, msgPayloadBuf...)                  // concat header and payload
		command := string(p.RemoveTrailingZeros(msgHeader.Command[:])) // switch on message command
		switch command {
		case "version":
			fmt.Println("Version message")
			msgVersion := new(p.MsgVersion)
			p.Decode(msg, msgVersion)
			fmt.Println(msgVersion)
		case "verack":
			fmt.Println("Version Acknowledge message")
			msgVersionAck := new(p.MsgVersionAck)
			p.Decode(msg, msgVersionAck)
			fmt.Println(msgVersionAck)
		default:
			fmt.Println("Don't know this command:", command)
		}
	}
}
