package egs

const TXT = `
hello
this is a long text
hi there
testing
`

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CreateFileWithText() {

}
