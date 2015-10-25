// Copyright (c) 2015, xrstf | MIT licensed

package iputils

import "net"

// A subnet expression uses the CIDR code existing in Go to match IPs against
// subnets.
type SubnetExpression struct {
	ip      net.IP
	network *net.IPNet
}

// Parses subnet and extracts the base IP and network. Errors out if Go could
// not understand the expression.
func NewSubnetExpression(subnet string) (*SubnetExpression, error) {
	ip, network, err := net.ParseCIDR(subnet)
	if err != nil {
		return nil, err
	}

	return &SubnetExpression{ip, network}, nil
}

// Check if the expression matches a given IP address.
func (self *SubnetExpression) Matches(ip net.IP) bool {
	return self.network.Contains(ip)
}
