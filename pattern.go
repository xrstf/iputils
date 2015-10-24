package iputils

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
)

var shrinker = regexp.MustCompile("\\*+")
var cleaner = regexp.MustCompile("[^0-9a-f.:*]")
var splitter = regexp.MustCompile("[.:]")

type PatternExpression struct {
	chunks []chunk
}

func NewPatternExpression(pattern string) (*PatternExpression, error) {
	pattern = strings.ToLower(pattern)

	// turn consecutive "*" into single "*"
	pattern = shrinker.ReplaceAllLiteralString(pattern, "*")

	// replace all invalid characters
	pattern = cleaner.ReplaceAllLiteralString(pattern, "*")

	// pre-parse the patterns
	chunks := splitter.Split(pattern, -1)
	list := make([]chunk, 0)

	for idx, chunk := range chunks {
		if len(chunk) == 0 {
			list = append(list, literalChunk{0})
		} else if strings.Contains(chunk, "*") == false { // we have a literal value to compare against
			// It's okay if the expression contains '.0.' and the IP contains '.000.',
			// we just care for the numerical value (and it's also okay to interprete
			// IPv4 chunks as hex values, as long as we interprete both as hex).
			i, err := strconv.ParseInt(chunk, 16, 0)
			if err != nil {
				return nil, fmt.Errorf("Pattern '%s' contains invalid characters in chunk %d.", pattern, idx+1)
			}

			list = append(list, literalChunk{i})
		} else { // we have an expression with a placeholder, like "12*"
			// turn the abstract placeholder into a regular expression
			chunk = strings.Replace(chunk, "*", "[0-9a-f]+?", -1)
			chunk = "^" + chunk + "$"

			expr, err := regexp.Compile(chunk)
			if err != nil {
				return nil, fmt.Errorf("Pattern '%s' contains invalid characters in chunk %d.", pattern, idx+1)
			}

			list = append(list, patternChunk{expr})
		}
	}

	return &PatternExpression{list}, nil
}

func (self *PatternExpression) Matches(ip net.IP) bool {
	ipString := ip.String()
	ipChunks := splitter.Split(ipString, -1)

	// the user probably mixed IPv4 and IPv6
	if len(ipChunks) != len(self.chunks) {
		return false
	}

	for idx, exprChunk := range self.chunks {
		if !exprChunk.Matches(ipChunks[idx]) {
			return false
		}
	}

	return true
}

type chunk interface {
	Matches(string) bool
}

type literalChunk struct {
	literal int64
}

func (self literalChunk) Matches(chunk string) bool {
	dec, err := strconv.ParseInt(chunk, 16, 0)

	return err == nil && dec == self.literal
}

type patternChunk struct {
	pattern *regexp.Regexp
}

func (self patternChunk) Matches(chunk string) bool {
	return self.pattern.MatchString(chunk)
}
