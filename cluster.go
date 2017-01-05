package sebarcore

type Cluster struct {
	ID, Title string
	Nodes     map[string]*Node
}

type Node struct {
	ID, Title                   string
	Capacity, Used, ObjectCount float64
	Role                        string
	Status                      string
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
