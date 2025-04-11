package resolver

import (
	log "github.com/sirupsen/logrus"
	"main/models"
	"main/models/dns"
)

func Run(query *dns.Query, server dns.Server) error {
	client := dns.NewClient(server, models.UDP)
	log.Info("Connecting to server ", server.Host, ":", server.Port)
	err := client.Connect()
	if err != nil {
		return err
	}
	log.Info("Connected to server.")
	packet, err := query.ToPacket()
	if err != nil {
		return err
	}
	log.Info("Sending packet...")
	writtenBytes, err := client.Send(&packet)
	if err != nil {
		return err
	}
	log.Info("Packet sent. Written ", writtenBytes, " bytes")
	log.Info("Receiving packet...")
	if err = client.Receive(); err != nil {
		return err
	}
	if err = client.Close(); err != nil {
		return err
	}
	return nil
}
