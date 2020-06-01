package eg

import "bufio"

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

func write(rw ReadWriter) (int, error) {
	return rw.Write([]byte("testing"))
}

func TestRW() {

}
