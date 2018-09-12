package main

import (
	"fmt"
	"process/model"
	"process/task"
)

func main() {
	p := new(model.BusinessProcess)
	p.Init(1, "my process1")
	t := new(model.Task)
	t.Init(1, "my task1")
	//todo object pool
	var pf interface{}
	pf = new(task.Print)
	t.Bind(pf.(model.Operation))

	e := new(model.Element)
	e.AddTask(t)

	g := new(model.Gateway)
	g.Init(1, "my gateway1")

	node1 := new(model.Node)
	node1.Init(1, "my node1")
	node1.AddGateway(g)
	node1.AddElement(e)

	p.AddNode(node1)
	p.Process(model.Context{})
	fmt.Println(node1)
	fmt.Println(e)
	fmt.Println(p)
	fmt.Println(t)
	// t.Run(model.Context{})

}
