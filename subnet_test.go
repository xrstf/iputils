package iputils

import (
	"net"
	"testing"
)

type subnetTestcase struct {
	Expression string
	Address    string
	Expected   bool
}

func TestSubnetMatch(t *testing.T) {
	testcases := []subnetTestcase{
		{"1.0.0.0/1", "1.0.0.0", true},
		{"1.0.0.0/8", "1.0.0.0", true},
		{"1.0.0.0/8", "1.1.0.0", true},
		{"1.0.0.0/8", "1.255.255.255", true},
		{"1.0.0.0/8", "2.0.0.0", false},
		{"2.0.0.0/7", "2.0.0.0", true},
		{"2.0.0.0/7", "2.0.255.0", true},
		{"2.0.0.0/7", "3.0.0.0", true},
		{"1.0.0.0/32", "1.0.0.0", true},
		{"1.0.0.0/32", "1.0.0.1", false},
		{"1.0.0.0/32", "2.0.0.0", false},

		{"2a01:198:603:0::/65", "2a01:198:603:0:396e:4789:8e99:890f", true},
		{"2a01:198:603:0::/65", "2a00:198:603:0:396e:4789:8e99:890f", false},
		{"2001::/16", "2000::1", false},
	}

	for _, testcase := range testcases {
		expression, err := NewSubnetExpression(testcase.Expression)
		if err != nil {
			t.Error(err)
		}

		address := net.ParseIP(testcase.Address)

		if expression.Matches(address) != testcase.Expected {
			t.Errorf("Expected '%s'.match('%s') = %t, but got the opposite.", testcase.Expression, testcase.Address, testcase.Expected)
		}
	}
}

func TestInvalidSubnets(t *testing.T) {
	testcases := []string{
		"1.0.0.0/",
		"1.0.0.0/null",
		"1.0.0.0/-2",
		"1.0.0.0/33",
		"1.0.0.0/1.2.3.4",
		"1.0.0.0/1.",
		"/1",
		"foo/1",
		"1.2.500.1/1",
		"2001:dH8::1:2/1",
		"::1/-1",
		"::1/200",
		"1.2.*.3/1",
	}

	for _, testcase := range testcases {
		_, err := NewSubnetExpression(testcase)
		if err == nil {
			t.Errorf("Expected '%s' to be an invalid subnet expression, but apparently it was valid.", testcase)
		}
	}
}
