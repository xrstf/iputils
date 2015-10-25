// Copyright (c) 2015, xrstf | MIT licensed

package iputils

import (
	"fmt"
	"net"
)

// A literal expression is the string representation of any IP address and
// therefore only ever matches one specific address.
type LiteralExpression struct {
	ip string
}

// Parses ip and returns a literal expression if ip is a valid IP address.
// Otherwise, an error is returned.
func NewLiteralExpression(ip string) (*LiteralExpression, error) {
	parsed := net.ParseIP(ip)
	if parsed == nil {
		return nil, fmt.Errorf("Could not parse '%s' as an IP address.", ip)
	}

	return &LiteralExpression{parsed.String()}, nil
}

// Check if the expression matches a given IP address.
func (self *LiteralExpression) Matches(ip net.IP) bool {
	return ip.String() == self.ip
}
