package main

import (
	"fmt"
	"time"

	"github.com/ctopher78/zookeeper-and-go/internal/zookeeper"
	"github.com/samuel/go-zookeeper/zk"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	conn1 := zookeeper.Connect()
	defer conn1.Close()
	conn2 := zookeeper.Connect()
	defer conn2.Close()

	flags := int32(0)
	acl := zk.WorldACL(zk.PermAll)

	conn1.Delete("/watch/child1", int32(0))
	conn1.Create("/watch", []byte(""), flags, acl)

	found, _, ech, err := conn1.ChildrenW("/watch")
	must(err)
	fmt.Printf("found: %v\n", found)

	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("creating znode")
		_, err = conn2.Create("/watch/child1", []byte("here"), flags, acl)
		must(err)
	}()

	evt := <-ech
	fmt.Println("watch fired")
	must(evt.Err)

	found, _, err = conn1.Children("/watch")
	must(err)
	fmt.Printf("found: %v\n", found)
}
