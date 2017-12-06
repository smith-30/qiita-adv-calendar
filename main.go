package main

import (
	"sync"

	"github.com/smith-30/qiita-adv-calendar/domain/model"
	"github.com/smith-30/qiita-adv-calendar/domain/service"
	"github.com/smith-30/qiita-adv-calendar/helper/env"
)

var (
	name  = "go"
	count = 4
)

func main() {
	env.LoadEnv()
	gridWg := new(sync.WaitGroup)
	aggregateWg := new(sync.WaitGroup)

	ag := service.NewAggregater(aggregateWg, 25*count)
	gridUpdateCh := ag.UpdateGrid(25 * count)

	cs := model.NewCalendars(name, count)

	for _, ca := range cs.C {
		gridWg.Add(1)
		go func(c *model.Calendar) {
			gridCh := c.SetExecuteURLs()

			for g := range gridCh {
				gridUpdateCh <- g
			}
			gridWg.Done()
		}(ca)
	}

	gridWg.Wait()
	close(gridUpdateCh)
	aggregateWg.Wait()
}
