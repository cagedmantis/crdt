package main

import (
	"fmt"
	"sync"

	"github.com/cagedmantis/crdt/cluster"
)

func main() {
	fmt.Println("Hola... serverca starting up!")

	s, err := cluster.NewCluster("127.0.0.1:7373")
	if err != nil {
		fmt.Println("Serf connection failure")
		return
	}

	members, err := s.Members()
	if err != nil {
		fmt.Println("error querying for members: ", err)
	}

	fmt.Println(members)

	ch := make(chan map[string]interface{}, 10)

	go s.MonitorHandler(ch)

	s.Monitor("member", ch)

	wg := sync.WaitGroup
	wg.Add(1)
}
