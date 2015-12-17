package cluster

import (
	"fmt"

	"github.com/hashicorp/serf/client"
)

type Cluster struct {
	serf    *client.RPCClient
	members []client.Member
}

func NewCluster(host string) (Cluster, error) {
	//"127.0.0.1:7373"
	s, err := client.NewRPCClient(host)
	if err != nil {
		fmt.Println(err)
		return Cluster{}, err
	}
	c := Cluster{
		serf: s,
	}
	return c, nil
}

func (c *Cluster) Members() ([]client.Member, error) {
	members, err := c.serf.Members()
	if err != nil {
		fmt.Println(err)
		return []client.Member{}, err
	}
	c.members = members
	return c.members, nil
}

func (c *Cluster) Monitor(filter string, ch chan<- map[string]interface{}) {
	_, err := c.serf.Stream(filter, ch)
	if err != nil {
		fmt.Println(err)
	}
}

func (c *Cluster) MonitorHandler(ch chan map[string]interface{}) {

	for {
		d := <-ch

		fmt.Printf("%v\n", d)
	}
}
