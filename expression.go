// Copyright (c) 2015, xrstf | MIT licensed

// Package iputils can be used to check whether a given IP matches a given pattern
// or is contained in a given subnet.
//
// There are three types of expressions:
//
// Literal expressions are IP addresses and always match exactly one address. This
// boils down to a simple equality comparision.
//
// Pattern expressions are IP address strings with * as a placeholder, for example
// 127.0.0.* or fe*:0:1:2:*b*:de:*0:8. The asterisk can match multiple characters
// in the string and there can be as many asterisks as needed.
//
// Subnet expressions wrap the existing CIDR logic in Go, allowing expressions
// like "127.0.0.1/8".
package iputils

import (
	"net"
	"strings"
)

// An expression is something that can describe an amount of IP addresses.
type Expression interface {
	Matches(net.IP) bool
}

// This function tries to auto-detect the type of expression based on the
// given expr. If neither / nor * appear in expr, it's assumed to be a literal,
// if no / is present, it's assumed to be a pattern; otherwise a subnet
// expression is used.
//
// If the chosen expression chokes, an error is returned.
func ParseExpression(expr string) (Expression, error) {
	if strings.Contains(expr, "/") == false {
		if strings.Contains(expr, "*") == false {
			return NewLiteralExpression(expr)
		}

		return NewPatternExpression(expr)
	}

	return NewSubnetExpression(expr)
}

// Determines whether or not the given IP is an IPv4 address.
func IsIPv4(ip net.IP) bool {
	return ip.To4() != nil
}

// Determines whether or not the given IP is an IPv6 address.
func IsIPv6(ip net.IP) bool {
	return !IsIPv4(ip)
}
