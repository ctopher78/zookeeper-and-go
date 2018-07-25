package main

import (
	"fmt"

	"github.com/ctopher78/zookeeper-and-go/internal/zookeeper"
	"github.com/samuel/go-zookeeper/zk"
)

const (
	ZKQUEUE = "/_QUEUE_/"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	conn2 := zookeeper.Connect()
	defer conn2.Close()

	flags := int32(zk.FlagEphemeral | zk.FlagSequence)
	acl := zk.WorldACL(zk.PermAll)

	fmt.Println("creating znode")
	for i := 0; i < 5; i++ {
		_, err := conn2.Create(ZKQUEUE+"ID-", []byte("some important data"), flags, acl)
		must(err)
	}

}
