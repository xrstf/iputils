package iputils

import (
	"fmt"
	"net"
)

type LiteralExpression struct {
	ip string
}

func NewLiteralExpression(ip string) (*LiteralExpression, error) {
	parsed := net.ParseIP(ip)
	if parsed == nil {
		return nil, fmt.Errorf("Could not parse '%s' as an IP address.", ip)
	}

	return &LiteralExpression{parsed.String()}, nil
}

func (self *LiteralExpression) Matches(ip net.IP) bool {
	return ip.String() == self.ip
}
