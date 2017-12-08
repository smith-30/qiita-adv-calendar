package main

import (
	"github.com/smith-30/qiita-adv-calendar/domain/service"
	"github.com/smith-30/qiita-adv-calendar/helper/env"
	"go.uber.org/zap"
)

var (
	name  = "go"
	count = 4
)

func main() {
	env.LoadEnv()
	cap := 25 * count
	logger, _ := zap.NewDevelopment()

	// ready grid.
	ag := service.NewAggregater(cap)
	gridUpdateCh := ag.UpdateGrid(cap)

	// fetch grids each calendar.
	ga := service.NewGridAggregater(name, count, logger.Sugar())
	ga.FetchGrids(gridUpdateCh)

	// wait to send grid.
	ga.Wait()
	close(gridUpdateCh)
	// wait aggregate.
	ag.Wait()
}
