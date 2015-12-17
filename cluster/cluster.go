package cluster

import (
	"fmt"

	"github.com/hashicorp/serf/client"
)

type Cluster struct {
	serf    *client.RPCClient
	members []client.Members
}

func NewCluster(host string) (Cluster, error) {
	//"127.0.0.1:7373"
	s, err := client.NewRPCClient(host)
	if err != nil {
		fmt.Println(err)
		return Cluster{}, err
	}
	return s, nil
}

func (c *Cluster) Members() ([]client.Members, error) {
	c.members, err = c.serf.Members()
	if err != nil {
		fmt.Println(err)
		return
	}
	return c.members
}

func (c *Cluster) Monitor(filter string, ch chan<- map[string]interface{}) {
	_, err := c.serf.Stream(filter, ch)
	if err != nil {
		fmt.Println(err)
	}
}

func (c *Cluster) MonitorHandler(c chan map[string]interface{}) {

	for {
		d := <-c

		fmt.Printf("%v\n", d)
	}
}
