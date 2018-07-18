package main

import (
	"fmt"

	"github.com/ctopher78/zookeeper-and-go/internal/zookeeper"
	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	conn := zookeeper.Connect()
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

func must(err error) {
	if err != nil {
		panic(err)
	}
}
