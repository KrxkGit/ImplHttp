package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
)

func main() {
	u := &url.URL{
		Scheme: "http",
		Host:   "rz.protect-file.com",
		Path:   "/2019/u3492432002/rzbfy.asp",
	}
	query := u.Query()
	query.Set("uid", "A144205")
	query.Set("pwd", "23966237")
	u.RawQuery = query.Encode()

	var address, Method, Port string
	Method = "GET"
	Port = "80"
	address = fmt.Sprintf("%s:%s", u.Host, Port)
	fmt.Println(address)

	req := fmt.Sprintf("%s %s HTTP/1.1\r\nHost: %s\r\nConnection:close\r\n\r\n", Method,
		u.Path+"?"+u.RawQuery, u.Host) /*注意关闭 长连接，在HTTP/1.1是默认开启的，长连接对方不会主动关闭请求*/
	fmt.Println(req)

	client, err := net.Dial("tcp", address)
	if err != nil {
		log.Println(err.Error())
	}
	nw, err := client.Write([]byte(req))
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("written: ", nw)

	//reader := bufio.NewReader(client)
	for {
		buf := make([]byte, 1024)
		n, err := client.Read(buf)
		if err != nil {
			if err == io.EOF { /*对方可能没有实现该方法，默认采用 \r\n\r\n 标志或对方断开连接表示 结束*/
				break
			}
		}
		fmt.Println("Read: ", n)
		fmt.Println("content:", string(buf))
	}
}
