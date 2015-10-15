package iputils

import (
	"net"
	"regexp"
	"strconv"
	"strings"
)

type PatternExpression struct {
	pattern string
}

func NewPatternExpression(pattern string) (*PatternExpression, error) {
	pattern = strings.ToLower(pattern)

	// turn consecutive "*" into single "*"
	re := regexp.MustCompile("\\*+")
	pattern = re.ReplaceAllLiteralString(pattern, "*")

	// replace all invalid characters
	re = regexp.MustCompile("[^0-9a-f.:*]")
	pattern = re.ReplaceAllLiteralString(pattern, "*")

	return &PatternExpression{pattern}, nil
}

func (self *PatternExpression) Matches(ip net.IP) bool {
	splitter := regexp.MustCompile("[.:]")

	ipString := ip.String()
	ipChunks := splitter.Split(ipString, -1)
	exprChunks := splitter.Split(self.pattern, -1)

	// the user probably mixed IPv4 and IPv6
	if len(ipChunks) != len(exprChunks) {
		return false
	}

	for idx, exprChunk := range exprChunks {
		ipChunk := ipChunks[idx]

		if strings.Contains(exprChunk, "*") == false { // we have a literal value to compare against
			// It's okay if the expression contains '.0.' and the IP contains '.000.',
			// we just care for the numerical value (and it's also okay to interprete
			// IPv4 chunks as hex values, as long as we interprete both as hex).
			i, err := strconv.ParseInt(ipChunk, 16, 0)
			if err != nil {
				return false
			}

			e, err := strconv.ParseInt(exprChunk, 16, 0)
			if err != nil {
				return false
			}

			if i != e {
				return false
			}
		} else { // we have an expression with a placeholder, like "12*"
			// turn the abstract placeholder into a regular expression
			exprChunk = strings.Replace(exprChunk, "*", "[0-9a-f]+?", -1)
			exprChunk = "^" + exprChunk + "$"

			expr, err := regexp.Compile(exprChunk)
			if err != nil {
				return false
			}

			if !expr.MatchString(ipChunk) {
				return false
			}
		}
	}

	return true
}
