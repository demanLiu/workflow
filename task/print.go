package task

import (
	"fmt"
	"process/model"
)

type Print struct {
}

func (p *Print) Excute(ctx model.Context) {
	fmt.Println("hello word,print is runnint")
}
