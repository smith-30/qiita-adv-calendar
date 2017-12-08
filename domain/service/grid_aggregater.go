package service

import (
	"fmt"
	"sync"

	"github.com/smith-30/qiita-adv-calendar/domain/model"
	"go.uber.org/zap"
)

const (
	baseUrl = "https://qiita.com/advent-calendar/2017/"
)

type (
	GridAggregater struct {
		C []*model.Calendar

		wg     *sync.WaitGroup
		logger *zap.SugaredLogger
	}
)

func NewGridAggregater(name string, count int, l *zap.SugaredLogger) *GridAggregater {
	cs := &GridAggregater{
		wg:     new(sync.WaitGroup),
		logger: l,
	}

	for i := 1; i <= count; i++ {
		if i == 1 {
			cs.addCalendar(name)
			continue
		}
		n := name + fmt.Sprint(i)
		cs.addCalendar(n)
	}
	return cs
}

func (cs *GridAggregater) addCalendar(name string) {
	url := baseUrl + name
	c := model.NewCalendar(url, cs.logger)
	cs.C = append(cs.C, c)
}

func (cs *GridAggregater) Wait() {
	cs.wg.Wait()
}

func (cs *GridAggregater) FetchGrids(gridUpdateCh chan *model.Grid) {
	for _, ca := range cs.C {
		cs.wg.Add(1)
		go func(c *model.Calendar) {
			cs.logger.Infof("calendar %s is start.", c.URL)
			gridCh := c.SetExecuteURLs()

			for g := range gridCh {
				gridUpdateCh <- g
			}
			cs.wg.Done()
		}(ca)
	}
}
