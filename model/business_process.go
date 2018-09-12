package model

type BusinessProcess struct {
	Info          BaseInfo
	Nodes         []*Node
	CurrentNodeId int32
	Status        int32
}

func (p *BusinessProcess) Init(id int32, name string) {
	p.Info.Init(id, name)
}
func (p *BusinessProcess) AddNode(node *Node) {
	p.Nodes = append(p.Nodes, node)
}
func (p *BusinessProcess) SetCurrentNodeId(node *Node) {
	p.CurrentNodeId = node.Info.Id
}
func (p *BusinessProcess) Process(ctx Context) {
	p.Status = Running
	for _, node := range p.Nodes {
		p.SetCurrentNodeId(node)
		node.Process(ctx)
	}
	p.Status = Finished
}
