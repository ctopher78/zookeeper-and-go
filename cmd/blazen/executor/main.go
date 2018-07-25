package main

import (
	"fmt"

	"github.com/ctopher78/zookeeper-and-go/internal/zookeeper"
	"github.com/samuel/go-zookeeper/zk"
)

const (
	ZKQUEUE = "/_QUEUE_"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	conn1 := zookeeper.Connect()
	defer conn1.Close()

	flags := int32(0)
	acl := zk.WorldACL(zk.PermAll)
	// conn1.Delete("/watch/child0000000028", int32(0))
	conn1.Create(ZKQUEUE, []byte(""), flags, acl)

	found, _, ech, err := conn1.ChildrenW(ZKQUEUE)
	must(err)
	fmt.Printf("found: %v\n", found)

	evt := <-ech
	fmt.Println("watch fired")
	must(evt.Err)

	found, _, err = conn1.Children(ZKQUEUE)
	must(err)
	fmt.Printf("found: %v\n", found)
	data, _, err := conn1.Get(ZKQUEUE + "/" + found[0])
	must(err)
	fmt.Printf("Data: %v\n", string(data))
}
