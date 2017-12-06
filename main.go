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
	cap := 25 * count

	gridWg := new(sync.WaitGroup)
	aggregateWg := new(sync.WaitGroup)

	ag := service.NewAggregater(aggregateWg, cap)
	gridUpdateCh := ag.UpdateGrid(cap)

	cs := model.NewCalendars(name, count, gridWg)
	cs.FetchGrids(gridUpdateCh)

	gridWg.Wait()
	close(gridUpdateCh)
	aggregateWg.Wait()
}
