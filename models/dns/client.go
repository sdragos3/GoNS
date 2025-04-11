package dns

import (
	"bufio"
	"fmt"
	"log"
	"main/models"
	"net"
	"time"
)

type Client struct {
	connection    net.Conn
	host          string
	port          uint16
	protocol      string
	lastWriteSize uint
}

func NewClient(server Server, protocol models.Protocol) Client {
	return Client{
		connection:    nil,
		host:          server.Host,
		port:          server.Port,
		protocol:      string(protocol),
		lastWriteSize: 0,
	}
}

func (client *Client) Connect() error {
	dnsServer := fmt.Sprintf("%s:%d", client.host, client.port)
	conn, err := net.Dial(client.protocol, dnsServer)
	if err != nil {
		return err
	}

	err = conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Fatal("failed to set deadline: ", err)
	}
	client.connection = conn
	return nil
}

func (client *Client) Send(packet *Packet) (int, error) {
	encodedPacket := packet.Encode()
	n, err := client.connection.Write(encodedPacket)
	if err != nil {
		return n, err
	}
	client.lastWriteSize = uint(n)
	return n, nil
}

func (client *Client) Receive() error {
	encodedAnswer := make([]byte, client.lastWriteSize)
	if _, err := bufio.NewReader(client.connection).Read(encodedAnswer); err != nil {
		log.Fatal(err)
	}
	fmt.Printf(">> encodedAnswer: %#x\n", encodedAnswer)
	return nil
}

func (client *Client) Close() error {
	if client.connection == nil {
		return nil
	}
	return client.connection.Close()
}
