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
