package model

type BaseInfo struct {
	Id   int32
	Name string
}

func (b *BaseInfo) Init(id int32, name string) {
	b.Id = id
	b.Name = name
}
