package model

const (
	Parallel  = 0
	Exclusive = -1
	//内置变量
	PreResult = "$preResult"
	PreStatus = "$preStatus"
)

type Gateway struct {
	Info      BaseInfo
	Condition func() []int32
}

func (g *Gateway) Init(id int32, name string) {
	g.Info.Init(id, name)
}
func (g *Gateway) SetCondition() {

}
