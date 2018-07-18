package main

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

var servers = []string{"localhost:2181", "localhost:2182", "localhost:2183"}

func connect() *zk.Conn {
	conn, _, err := zk.Connect(servers, time.Second)
	must(err)
	return conn
}

func main() {
	conn := connect()
	defer conn.Close()

	flags := int32(0)
	acl := zk.WorldACL(zk.PermAll)

	path, err := conn.Create("/01", []byte("data"), flags, acl)
	must(err)
	fmt.Printf("create: %+v\n", path)

	data, stat, err := conn.Get("/01")
	must(err)
	fmt.Printf("get:    %+v %+v\n", string(data), stat)

	stat, err = conn.Set("/01", []byte("newdata"), stat.Version)
	must(err)
	fmt.Printf("set:    %+v\n", stat)

	err = conn.Delete("/01", -1)
	must(err)
	fmt.Printf("delete: ok\n")

	exists, stat, err := conn.Exists("/01")
	must(err)
	fmt.Printf("exists: %+v %+v\n", exists, stat)
}
