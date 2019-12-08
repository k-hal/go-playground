package main

import (
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("GET", "http://ascii.jp/", nil)
	req.Write(conn)
	//io.WriteString(conn, "GET / HTTP/1.1\r\nHost: ascii.jp\r\n\r\n")
	io.Copy(os.Stdout, conn)
	conn.Close()
}
