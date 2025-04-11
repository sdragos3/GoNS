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

func NewClient(host string, port uint16, protocol models.Protocol) Client {
	return Client{
		connection:    nil,
		host:          host,
		port:          port,
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

func (client *Client) Send(packet *Packet) error {
	encodedPacket := packet.Encode()
	n, err := client.connection.Write(encodedPacket)
	if err != nil {
		return err
	}
	client.lastWriteSize = uint(n)
	return nil
}

func (client *Client) Receive() error {
	encodedAnswer := make([]byte, client.lastWriteSize)
	fmt.Printf("zedzedz")
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
