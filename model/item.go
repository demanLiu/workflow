package model

type Context struct {
	flagSet       map[string]interface{}
	parentContext *Context
}

//元素和任务一对一
type Element struct {
	Info BaseInfo
	Task *Task
}

//todo evnet
type Event struct {
}

func (e *Element) Init(id int32, name string) {
	e.Info.Init(id, name)
}

func (e *Element) AddTask(t *Task) {
	e.Task = t
}
