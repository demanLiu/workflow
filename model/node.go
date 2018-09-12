package model

type Node struct {
	Info     BaseInfo
	Gateway  *Gateway
	Status   int32
	Elements []*Element
}

func (n *Node) Init(id int32, name string) {
	n.Info.Init(id, name)
}
func (n *Node) AddGateway(g *Gateway) {
	n.Gateway = g
}
func (n *Node) AddElement(e *Element) {
	n.Elements = append(n.Elements, e)
}

func (n *Node) Process(ctx Context) {
	n.GetNextElemets()
}
func (n *Node) GetNextElemets() {

}
