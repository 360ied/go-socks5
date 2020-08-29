package main

import (
	"go-socks5/keepalive"
	"go-socks5/socks"
)

func main() {
	//var IPAddress, Port string
	//
	//if len(os.Args) == 2 {
	//	IPAddress = "localhost"
	//	Port = os.Args[1]
	//} else if len(os.Args) == 3 {
	//	IPAddress = os.Args[1]
	//	Port = os.Args[2]
	//} else {
	//	fmt.Println("syntax: ")
	//	fmt.Println("  <IPAddress> <port>  Binds to the specified IP and port")
	//	fmt.Println("  <port>              Binds to localhost and the specified port")
	//	return
	//}

	go keepalive.KeepAlive()

	var IPAddress = "127.0.0.1"
	var Port = "6969"

	conf := &socks.Config{}
	server, err := socks.New(conf)
	if err != nil {
		panic(err)
	}

	if err := server.ListenAndServe("tcp", IPAddress+":"+Port); err != nil {
		panic(err)
	}
}
