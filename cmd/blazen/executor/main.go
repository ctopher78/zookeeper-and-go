package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/contiv/executor"
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
	// conn := zookeeper.Connect()
	// defer conn.Close()

	e := executor.NewCapture(exec.Command("cmd/blazen/fixtures/test-script.py"))
	var buff bytes.Buffer
	_, err := buff.Write([]byte("test input\n"))
	if err != nil {
		log.Fatal(err)
	}
	e.Stdin = &buff
	e.Start() // start
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	er, err := e.Wait(ctx) // wait for only 10 seconds
	if err != nil {
		fmt.Println("Err: ", err)
	}

	fmt.Println("Out: ", er.Stdout)
	fmt.Println("Err: ", er.Stderr)

	// flags := int32(0)
	// acl := zk.WorldACL(zk.PermAll)
	// // conn1.Delete("/watch/child0000000028", int32(0))
	// conn.Create(ZKQUEUE, []byte(""), flags, acl)

	// found, _, ech, err := conn.ChildrenW(ZKQUEUE)
	// must(err)
	// fmt.Printf("found: %v\n", found)

	// evt := <-ech
	// fmt.Println("watch fired")
	// must(evt.Err)

	// found, _, err = conn.Children(ZKQUEUE)
	// must(err)
	// fmt.Printf("found: %v\n", found)
	// data, _, err := conn.Get(ZKQUEUE + "/" + found[0])
	// must(err)
	// fmt.Printf("Data: %v\n", string(data))
}
