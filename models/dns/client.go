package dns

import (
	"fmt"
	"log"
	"main/models"
	"net"
	"time"
)

type Client struct {
	connection net.Conn
	host       string
	port       uint16
	protocol   string
}

func New(host string, port uint16, protocol models.Protocol) (Client, error) {
	return Client{
		connection: nil,
		host:       host,
		port:       port,
		protocol:   string(protocol),
	}, nil
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

func (client *Client) Close() error {
	if client.connection == nil {
		return nil
	}
	return client.connection.Close()
}
