package service

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/smith-30/qiita-adv-calendar/domain/model"
	"go.uber.org/zap"
)

type (
	Aggregater struct {
		wg          *sync.WaitGroup
		dispatcher  *Dispatcher
		aggregateCh chan *model.Grid
		grids       []*model.Grid

		logger *zap.SugaredLogger
	}
)

func NewAggregater(cap int, l *zap.SugaredLogger) *Aggregater {
	aCh := make(chan *model.Grid, cap)
	d := NewDispatcher(aCh, l)
	d.Start()

	return &Aggregater{
		wg:          new(sync.WaitGroup),
		dispatcher:  d,
		aggregateCh: aCh,
		logger:      l,
	}
}

func (a *Aggregater) UpdateGrid(cap int) chan *model.Grid {
	gridUpdateCh := make(chan *model.Grid, cap)
	a.wg.Add(1)

	go func() {
		defer func() {
			a.wg.Done()
			a.dispatcher.Wait()
			close(a.aggregateCh)
		}()
		for g := range gridUpdateCh {
			a.dispatcher.Add(g)
		}
		a.logger.Info("update each grid is finished.")
	}()

	go a.wg.Add(1)

	go func() {
		defer func() {
			a.logger.Info("start sort. output ranking...\n")
			a.Output()
			a.wg.Done()
		}()
		for g := range a.aggregateCh {
			a.grids = append(a.grids, g)
		}
	}()

	return gridUpdateCh
}

// refs https://mattn.kaoriya.net/software/lang/go/20161004092237.htm
func (a *Aggregater) Output() {
	sort.Slice(a.grids, func(i, j int) bool {
		return a.grids[i].Like > a.grids[j].Like
	})

	for _, g := range a.grids {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%v, %v, %v\n\n", g.Like, g.Title, g.QiitaURL)
	}
}

func (a *Aggregater) Wait() {
	a.wg.Wait()
}
