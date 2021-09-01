package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	b, err := HTTPwithCookies("")
	if err != nil {
		panic(err)
	}
	println(string(b))
}

func HTTPwithCookies(url string) (b []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.AddCookie(&http.Cookie{Name: "Id", Value: ""})
	req.AddCookie(&http.Cookie{Name: "u", Value: ""})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(url +
			"\nresp.StatusCode: " + strconv.Itoa(resp.StatusCode))
		return
	}

	return ioutil.ReadAll(resp.Body)
}
