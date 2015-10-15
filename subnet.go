package iputils

import "net"

type SubnetExpression struct {
	subnet string
}

func NewSubnetExpression(subnet string) (*SubnetExpression, error) {
	return &SubnetExpression{subnet}, nil
}

func (self *SubnetExpression) Matches(ip net.IP) bool {
	return false
}
