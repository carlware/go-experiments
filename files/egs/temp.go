package egs

import (
	"fmt"
	"io/ioutil"
	"os"
)

func WorkWithTemp() {
	t, err := ioutil.TempDir("", "tmp")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(t)

	tf, err := ioutil.TempFile(t, "tmp")
	if err != nil {
		panic(err)
	}
	fmt.Println(tf.Name())
}
