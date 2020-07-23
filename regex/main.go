package main

import (
	"fmt"
	"regexp"
)

func main() {
	name := ""
	url := ""

	re := regexp.MustCompile(`(?m)https.*` + name + `.*v=(1|0|2|3|4|5)$`)
	fmt.Printf("%q\n", re.Find([]byte(url)))

}
