package iputils

import "net"

type SubnetExpression struct {
	ip      net.IP
	network *net.IPNet
}

func NewSubnetExpression(subnet string) (*SubnetExpression, error) {
	ip, network, err := net.ParseCIDR(subnet)
	if err != nil {
		return nil, err
	}

	return &SubnetExpression{ip, network}, nil
}

func (self *SubnetExpression) Matches(ip net.IP) bool {
	return self.network.Contains(ip)
}
