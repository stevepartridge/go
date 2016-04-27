package log

import (
	// "fmt"
	"net"
)

type Papertrail struct {
	Destination string
}

func NewPapertrail(destination string) *Papertrail {
	newPapertrail := new(Papertrail)

	newPapertrail.Destination = destination

	return newPapertrail
}

func (p *Papertrail) Send(payload []byte) error {

	conn, err := net.Dial("udp", p.Destination)
	if err == nil {
		_, err = conn.Write(payload)
	}

	return err
}
