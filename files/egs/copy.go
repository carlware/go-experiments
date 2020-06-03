package egs

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func Copy(in io.Reader, out io.Writer) error {
	w := io.MultiWriter(out, os.Stdout)
	buf := make([]byte, 64)
	if _, err := io.CopyBuffer(w, in, buf); err != nil {
		return err
	}
	return nil
}

func Scanlines() {
	in := bytes.NewReader([]byte("( a > 1 ) && ( b < 10 ) || x == a"))

	scanner := bufio.NewScanner(in)
	// scanner.Split(bufio.ScanLines)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func CopyBuffer() {
	in := bytes.NewReader([]byte(TXT))
	out := &bytes.Buffer{}

	if err := Copy(in, out); err != nil {
		panic(err)
	}

	fmt.Println("out bytes buffer =", out.String())
}

func PipeExample() {
	r, w := io.Pipe()
	out := &bytes.Buffer{}

	go func() {
		_, _ = w.Write([]byte("test\n"))
		w.Close()
	}()

	if _, err := io.Copy(out, r); err != nil {
		panic(err)
	}

	fmt.Println("out bytes buffer =", out.String())
}
