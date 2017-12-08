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
	s := logger.Sugar()

	// ready grid.
	ag := service.NewAggregater(cap, s)
	gridUpdateCh := ag.UpdateGrid(cap)

	// fetch grids each calendar.
	ga := service.NewGridAggregater(name, count, s)
	ga.FetchGrids(gridUpdateCh)

	// wait to send grid.
	ga.Wait()
	s.Info("finished FetchGrids.")
	close(gridUpdateCh)
	// wait aggregate.
	ag.Wait()
}
