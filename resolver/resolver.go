package resolver

import (
	"main/models"
	"main/models/dns"
)

func Run(query *dns.Query) error {
	client := dns.NewClient("8.8.8.8", 85, models.UDP)
	err := client.Connect()
	if err != nil {
		return err
	}
	packet, err := query.ToPacket()
	if err != nil {
		return err
	}
	err = client.Send(&packet)
	if err != nil {
		return err
	}
	client.Receive()
	client.Close()
	return nil
}
