package egs

import (
	"fmt"
	"strings"
)

const (
	ZERO WordSize = 6 * iota
	SMALL
	MEDIUM
	LARGE
	XLARGE
	XXLARGE   WordSize = 50
	SEPARATOR          = ", "
)

const constants = `
** Constants ***
ZERO: %v
SMALL: %d
MEDIUM: %d
LARGE: %d
XLARGE: %d
XXLARGE: %d
`

type WordSize int

type ChainLink struct {
	Data []string
}

func (v *ChainLink) Value() []string {
	return v.Data
}

type stringFunc func(s string) (result string)

func (v *ChainLink) Map(fn stringFunc) *ChainLink {
	var mapped []string
	orig := *v
	for _, s := range orig.Data {
		mapped = append(mapped, fn(s)) // first-class function
	}
	v.Data = mapped
	return v
}

func (v *ChainLink) Filter(max WordSize) *ChainLink {
	filtered := []string{}
	orig := *v
	for _, s := range orig.Data {
		if len(s) <= int(max) { // embedded logic
			filtered = append(filtered, s)
		}
	}
	v.Data = filtered
	return v
}

func ChainTest() {
	fmt.Printf(constants, ZERO, SMALL, MEDIUM, LARGE, XLARGE, XXLARGE)

	words := []string{
		"tiny",
		"marathon",
		"philanthropinist",
		"supercalifragilisticexpialidocious"}

	data := ChainLink{words}
	fmt.Printf("unfiltered: %#v\n", data.Value())

	filtered := data.Filter(SMALL)
	fmt.Printf("filtered: %#v\n", filtered)

	fmt.Printf("filtered and mapped (<= SMALL sized words): %#v\n",
		filtered.Map(strings.ToUpper).Value())

	data = ChainLink{words}
	fmt.Printf("filtered and mapped (<= Up to MEDIUM sized words): %#v\n",
		data.Filter(MEDIUM).Map(strings.ToUpper).Value())

	data = ChainLink{words}
	fmt.Printf("filtered twice and mapped (<= Up to LARGE sized words): %#v\n",
		data.Filter(XLARGE).Map(strings.ToUpper).Filter(LARGE).Value())
}
