package main

import (
	"time"

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
	// wg := sync.WaitGroup{}

	ag := service.NewAggregater()
	agCh := ag.StartCalc(25 * count)

	cs := model.NewCalendars(name, count)

	for _, ca := range cs.C {
		go func(c *model.Calendar) {
			gridCh := c.SetExecuteURLs()

			for g := range gridCh {
				agCh <- g
			}
		}(ca)
	}

	// ctx := context.Background()
	// ctx, cancel := context.WithCancel(ctx)

	// go func() {
	// 	time.Sleep(5 * time.Second)
	// 	close(agCh)
	// }()

	time.Sleep(10 * time.Second)
	close(agCh)
}
