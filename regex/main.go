package main

import (
	"crypto/md5"
	"fmt"
	"regexp"
)

func main() {
	name := ""
	url := ""

	re := regexp.MustCompile(`(?m)https.*` + name + `.*v=(1|0|2|3|4|5)$`)
	fmt.Printf("%q\n", re.Find([]byte(url)))

	data := []byte("5f176ead")
	fmt.Printf("%x", md5.Sum(data))
}
