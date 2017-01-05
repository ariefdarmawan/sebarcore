package sebarcore

import (
	"fmt"
)

type UseSizeCountEnum int

const (
	UseSize  UseSizeCountEnum = 0
	UseCount                  = 1
)

type Cluster struct {
	ID, Title string
	Nodes     map[string]*Node
}

type Node struct {
	ID, Title                  string
	Capacity, Used             float64
	ObjectCount, CapacityCount int
	Role                       string
	Status                     string
}

func NewCluster(id string) *Cluster {
	c := new(Cluster)
	c.ID = id
	c.Title = id
	c.Nodes = map[string]*Node{}
	return c
}

func (c *Cluster) AddNode(n *Node) {
	c.Nodes[n.ID] = n
}

func (c *Cluster) RemoveNode(id string) {
	delete(c.Nodes, id)
}

func (c *Cluster) FindAvailableNode(sizeorcount UseSizeCountEnum, size float64) (node *Node, err error) {
	lowestavail := float64(0)
	highestcount := 0
	idx := 0
	for _, n := range c.Nodes {
		avail := n.Capacity - n.Used
		if sizeorcount == UseSize {
			if avail >= size {
				if lowestavail < avail {
					node = n
					lowestavail = avail
				}
			}
		} else if n.CapacityCount > n.ObjectCount || n.CapacityCount == 0 {
			if node == nil {
				node = n
				highestcount = n.ObjectCount
			} else if highestcount > n.ObjectCount {
				node = n
				highestcount = n.ObjectCount
			}
		}
		idx++
	}
	if node == nil {
		err = ThrowErr("sebarcore", "Cluster",
			fmt.Sprintf("No node has enough available space to store data. Required size is: %f", size))
	}
	return
}
