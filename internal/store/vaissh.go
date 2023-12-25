package store

import (
	"database/sql/driver"
	"net"
	"time"

	"golang.org/x/crypto/ssh"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type ViaSSHDialer struct {
	client *ssh.Client
}

func NewDialer(c *ssh.Client) *ViaSSHDialer {
	return &ViaSSHDialer{
		client: c,
	}
}

func (self *ViaSSHDialer) Open(s string) (_ driver.Conn, err error) {
	return pq.DialOpen(self, s)
}

func (self *ViaSSHDialer) Dial(network, address string) (net.Conn, error) {
	return self.client.Dial(network, address)
}

func (self *ViaSSHDialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	return self.client.Dial(network, address)
}
