package eg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

type RW struct {
	*bufio.Reader // *bufio.Reader
	*bufio.Writer // *bufio.Writer
}

func write(rw ReadWriter, content string) (int, error) {
	return rw.Write([]byte(content))
}

func read(rw ReadWriter, w []byte) (int, error) {
	return rw.Read(w)
}

func TestRW() {
	// w := bytes.NewBufferString("your string")
	r := bufio.NewReader(strings.NewReader("some io.Reader stream to be read\n"))
	w := bufio.NewWriter(os.Stdout)

	rw := &RW{r, w}
	i, e := write(rw, "\nnew line")
	fmt.Println(i, e)
	i, e = read(rw, []byte("hol"))
	fmt.Println(i, e)

}
