package main

import "carlware/templ/src"


type B int
type A string
type C = B | A

func main() {
	src.TemplateTest()
}
