package iputils

import (
	"net"
	"strings"
)

type Expression interface {
	Matches(net.IP) bool
}

func ParseExpression(expr string) (Expression, error) {
	if strings.Contains(expr, "/") == false {
		if strings.Contains(expr, "*") == false {
			return NewLiteralExpression(expr)
		}

		return NewPatternExpression(expr)
	}

	return NewSubnetExpression(expr)
}

func IsIPv4(ip net.IP) bool {
	return ip.To4() != nil
}

func IsIPv6(ip net.IP) bool {
	return !IsIPv4(ip)
}
