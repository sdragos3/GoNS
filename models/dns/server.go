package dns

type Server struct {
	Host string
	Port uint16
}

func DefaultServer() Server {
	return Server{
		Host: "8.8.8.8",
		Port: 53,
	}
}
