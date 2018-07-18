package zookeeper

import (
	"log"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

var servers = []string{"localhost:2181", "localhost:2182", "localhost:2183"}

func Connect() *zk.Conn {
	conn, _, err := zk.Connect(servers, time.Second)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
