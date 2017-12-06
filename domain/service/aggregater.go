package service

import (
	"fmt"

	"github.com/smith-30/qiita-adv-calendar/domain/model"
)

type (
	Aggregater struct {
	}
)

func NewAggregater() *Aggregater {
	return &Aggregater{}
}

func (a *Aggregater) StartCalc(cap int) chan *model.Grid {
	ch := make(chan *model.Grid, cap)

	go func() {
		for {
			select {
			case g, ok := <-ch:
				fmt.Println(g)
				fmt.Println(ok)
			}
		}
	}()

	return ch
}
