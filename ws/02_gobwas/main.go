package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

type msg struct {
	Num int
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println("Could not open file.", err)
	}
	fmt.Fprintf(w, "%s", content)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Origin") != "http://"+r.Host {
		http.Error(w, "Origin not allowed", 403)
		return
	}
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

    fmt.Println("conn", conn)
	go echo(conn)
}

func echo(conn net.Conn) {
	defer conn.Close()
	for {
		// m := msg{}
		// err := conn.ReadJSON(&m)
		msg, op, err := wsutil.ReadClientData(conn)
		if err != nil {
			fmt.Println("error", err)
			break
		}
		fmt.Println("msg", string(msg))
		fmt.Println("op", op)
		err = wsutil.WriteServerMessage(conn, op, msg)
		if err != nil {
			fmt.Println("error", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/", rootHandler)

	panic(http.ListenAndServe(":8080", nil))
}
