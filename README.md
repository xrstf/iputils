iputils - IP-based expressions in Go
====================================
[![Build Status](https://travis-ci.org/xrstf/iputils.svg?branch=master)](https://travis-ci.org/xrstf/iputils)
[![GoDoc](https://godoc.org/github.com/xrstf/iputils?status.svg)](https://godoc.org/github.com/xrstf/iputils)

This Go package implements a simple way to let users configure IP ranges and
networks to control access to services. It can handle *literals* (``127.0.0.1``),
*patterns* (``127.0.*.*``) and *subnets* (``127.0.0.1/16``), both for IPv4
and IPv6. The IP parsing builds upon Go's [net.IP](https://golang.org/pkg/net/)
package.

Installation
------------

```
go get github.com/xrstf/iputils
```

Usage
-----

```go
// we want to check this IP address
addr := net.ParseIP("127.0.0.8")

// ... against a pattern
expr, err := iputils.NewPatternExpression("127.0.0.*")
if err != nil {
	panic(err)
}

if expr.Matches(addr) {
	fmt.Println("Address matches")
} else {
	fmt.Println("Address does not match.")
}

// ... or against a subnet
expr, err = iputils.NewSubnetExpression("127.0.0.1/16")
if err != nil {
	panic(err)
}

if expr.Matches(addr) {
	fmt.Println("Address matches")
} else {
	fmt.Println("Address does not match.")
}

// ... or a literal address
expr, err = iputils.NewLiteralExpression("127.0.0.1")
if err != nil {
	panic(err)
}

if expr.Matches(addr) {
	fmt.Println("Address matches")
} else {
	fmt.Println("Address does not match.")
}

// Try to auto-detect the type of expression:
expr, err = iputils.ParseExpression("127.0.0.1")
if err != nil {
	panic(err)
}
```

License
-------

This code is licensed under the MIT license.
