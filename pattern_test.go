// Copyright (c) 2015, xrstf | MIT licensed

package iputils

import (
	"net"
	"testing"
)

type patternTestcase struct {
	Expression string
	Address    string
	Expected   bool
}

func TestPatternMatch(t *testing.T) {
	testcases := []patternTestcase{
		{"0.0.0.0", "0.0.0.0", true},
		{"0.0.0.*", "0.0.0.0", true},
		{"0.0.0.**", "0.0.0.0", true},
		{"0.0.0.*****", "0.0.0.0", true},
		{"0.0.*.*", "0.0.0.0", true},
		{"0.0.*.*", "0.0.0.1", true},
		{"0.0.*.*", "0.0.1.0", true},
		{"0.0.*.*", "0.0.12.13", true},
		{"0.0.*.*", "0.0.0.255", true},
		{"0.0.*.*", "0.5.0.255", false},
		{"0.0.*.*", "255.5.0.255", false},
		{"0.0.*7.0", "0.0.17.0", true},
		{"0.0.*7.0", "0.0.117.0", true},
		{"0.0.*7*.0", "0.0.17.0", false},
		{"0.0.*7**.0", "0.0.17.0", false},
		{"0.0.1**.0", "0.0.1.0", false},
		{"0.0.1**.0", "0.0.17.0", true},
		{"0.0.1**.0", "0.0.174.0", true},
		{"0.0.1**.0", "0.0.41.0", false},
		{"0.0.1**.0", "0.0.211.0", false},
		{"0.0.1**1.0", "0.0.121.0", true},
		{"0.0.1**1.0", "0.0.11.0", false},
		{"0.0.1**1.0", "0.0.122.0", false},
		{"0.0.*.255", "0.0.0.255", true},
		{"0.0.*.255", "0.0.0.255", true},
		{"0.0.0.1*1", "0.0.0.101", true},
		{"0.0.0.1*1", "0.0.0.11", false},
		{"0.0.0.1*1", "0.0.0.110", false},
		{"0.0.0.1*1", "0.0.0.1", false},

		{"2001:db8:85a3:0:0:8a2e:370:7334", "2001:db8:85a3::8a2e:370:7334", true},
		{"2001:db8:85a3:0:0:*:370:7334", "2001:db8:85a3::8a2e:370:7334", true},
		{"2001:db8:8*a3:0:0:8a*e:370:7334", "2001:db8:85a3::8a2e:370:7334", true},
		{"2001:**:8*a3:0:0:8a*e:370:7334", "2001:db8:85a3::8a2e:370:7334", true},
		{"2001:**:8*a3:0:0:8a*e:370:*3*4", "2001:db8:85a3::8a2e:370:7334", true},

		{"2001:db8:85a3:0:0:9*:370:7334", "2001:db8:85a3::8a2e:370:7334", false},
		{"2001:db8:85a3:0:0:8a2e:370*:7334", "2001:db8:85a3::8a2e:370:7334", false},
		{"2001:db8:*85a3:0:0:8a2e:370:7334", "2001:db8:85a3::8a2e:370:7334", false},
		{"2001:d*b8:85a3:0:0:8a2e:370:7334", "2001:db8:85a3::8a2e:370:7334", false},
	}

	for _, testcase := range testcases {
		expression, err := NewPatternExpression(testcase.Expression)
		if err != nil {
			t.Error(err)
		}

		address := net.ParseIP(testcase.Address)

		if expression.Matches(address) != testcase.Expected {
			t.Errorf("Expected '%s'.match('%s') = %t, but got the opposite.", testcase.Expression, testcase.Address, testcase.Expected)
		}
	}
}
