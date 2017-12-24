package service

import (
	"os"
	"sync"

	"github.com/smith-30/qiita-adv-calendar/domain/model"
	"go.uber.org/zap"
)

// refs　http://blog.kaneshin.co/entry/2016/08/18/190435
type (
	Dispatcher struct {
		pool     chan *fetcher
		queue    chan interface{}
		fetchers []*fetcher
		wg       sync.WaitGroup
		quit     chan struct{}
	}
)

const (
	// worker数
	maxFetchers = 10

	maxQueues = 100
)

func NewDispatcher(aggregateCh chan *model.Grid, l *zap.SugaredLogger) *Dispatcher {
	d := &Dispatcher{
		pool:  make(chan *fetcher, maxFetchers),
		queue: make(chan interface{}, maxQueues),
		quit:  make(chan struct{}),
	}

	d.fetchers = make([]*fetcher, cap(d.pool))
	for i := 0; i < cap(d.pool); i++ {
		w := fetcher{
			dispatcher:  d,
			data:        make(chan interface{}),
			quit:        make(chan struct{}),
			aggregateCh: aggregateCh,
			token:       os.Getenv("AUTH_TOKEN"),

			logger: l,
		}
		d.fetchers[i] = &w
	}
	return d
}

func (d *Dispatcher) Start() {
	for _, f := range d.fetchers {
		f.start()
	}

	go func() {
		for {
			select {
			case v := <-d.queue:
				fetcher := <-d.pool
				fetcher.data <- v

			case <-d.quit:
				return
			}
		}
	}()
}

func (d *Dispatcher) Add(v interface{}) {
	d.wg.Add(1)
	d.queue <- v
}

func (d *Dispatcher) Wait() {
	d.wg.Wait()
}
