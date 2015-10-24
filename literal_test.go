package iputils

import (
	"net"
	"testing"
)

type literalTestcase struct {
	Expression string
	Address    string
	Expected   bool
}

func TestLiteralMatch(t *testing.T) {
	testcases := []literalTestcase{
		{"0.0.0.0", "0.0.0.0", true},
		{"12.0.0.0", "12.0.0.0", true},
		{"12.0.0.255", "12.0.0.255", true},
		{"255.254.255.255", "254.255.255.255", false},
		{"12.0.0.0", "1.0.0.0", false},
		{"12.0.0.0", "1.2.0.0", false},
		{"::1", "::1", true},
		{"::1", "0:0:0:0:0:0:0:1", true},
		{"0:0:0:0:0:0:0:1", "::1", true},
		{"1:0:0:0:0:0:0:1", "::1", false},
	}

	for _, testcase := range testcases {
		expression, err := NewLiteralExpression(testcase.Expression)
		if err != nil {
			t.Error(err)
		}

		address := net.ParseIP(testcase.Address)

		if expression.Matches(address) != testcase.Expected {
			t.Errorf("Expected '%s'.match('%s') = %t, but got the opposite.", testcase.Expression, testcase.Address, testcase.Expected)
		}
	}
}
