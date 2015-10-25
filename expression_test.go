// Copyright (c) 2015, xrstf | MIT licensed

package iputils

import "testing"

type parseExpressionTestcase struct {
	Input    string
	Expected string
}

func TestParseExpression(t *testing.T) {
	testcases := []parseExpressionTestcase{
		{"0.0.0.0", "literal"},
		{"::1", "literal"},
		{"fe80::1", "literal"},

		{"0.0.0.0/8", "subnet"},
		{"::1/8", "subnet"},
		{"fe80::/128", "subnet"},

		{"fe*:0:0:0:0:0:0:0", "pattern"},
		{"127.*.*.0", "pattern"},
	}

	for _, testcase := range testcases {
		result, err := ParseExpression(testcase.Input)
		if err != nil {
			t.Error(err)
		}

		var okay bool

		switch testcase.Expected {
		case "literal":
			_, okay = result.(*LiteralExpression)
		case "subnet":
			_, okay = result.(*SubnetExpression)
		case "pattern":
			_, okay = result.(*PatternExpression)
		}

		if !okay {
			t.Errorf("Expected a '%s' expression, but did not get one for '%s'.", testcase.Expected, testcase.Input)
		}
	}
}
