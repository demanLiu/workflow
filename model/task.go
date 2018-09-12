package model

import (
	"time"
)

const (
	DefaultLimitTime = 3 * time.Minute
	Wait             = 0
	Running          = 1
	Failed           = 2
	Finished         = 3
)

type Operation interface {
	Excute(ctx Context)
}
type Task struct {
	Info         BaseInfo
	StartTime    time.Time
	EndTime      time.Time
	LimitTime    time.Duration
	Status       int32
	IsCancel     bool
	IsExcuted    bool
	CustomExcute Operation
	//TOOD 权限
}

func (t *Task) Init(id int32, name string) {
	t.Info.Init(id, name)
	t.LimitTime = DefaultLimitTime
	t.IsExcuted = true
}

func (t *Task) Bind(f interface{}) {
	t.CustomExcute = f.(Operation)
}
func (t *Task) Run(ctx Context) {
	if t.IsExcuted && t.CustomExcute != nil {
		t.CustomExcute.Excute(ctx)
	}
}
