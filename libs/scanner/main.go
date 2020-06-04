package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	// const src = `created > "2019-11-20 14:34:58-06" and created < "2019-11-27 14:34:58-06"`
	// const src = `created > "2019-11-20 14:34:58-06" and created < "2019-11-27 14:34:58-06" and tags in ["abc",1]`
	const src = `[1111, "bbb"]`

	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	s.Filename = "example"
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
	}

}
