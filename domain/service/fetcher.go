package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/smith-30/qiita-adv-calendar/domain/model"

	"go.uber.org/zap"
)

type (
	fetcher struct {
		dispatcher  *Dispatcher
		data        chan interface{}
		quit        chan struct{}
		aggregateCh chan *model.Grid
		token       string

		logger *zap.SugaredLogger
	}
)

var UpdateGrid = func() (*model.ItemInfo, error) {
	return &model.ItemInfo{}, nil
}

func (f *fetcher) start() {
	go func() {
		for {
			f.dispatcher.pool <- f

			select {
			case v := <-f.data:
				if g, ok := v.(*model.Grid); ok {
					// update grid.
					info, err := f.fetchGridInfo(g.URL)
					if err != nil {
						f.logger.Errorf("%v\n", err)
					} else {
						f.logger.Infof("success api request %s, %d", g.URL, info.PageViewsCount)
						g.Like = info.LikesCount
						f.aggregateCh <- g
					}
				}

				f.dispatcher.wg.Done()

			case <-f.quit:
				return
			}
		}
	}()
}

func (f *fetcher) fetchGridInfo(url string) (*model.ItemInfo, error) {
	ii := &model.ItemInfo{}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+f.token)

	fmt.Printf("f.token %#v\n", f.token)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return ii, errors.Wrap(err, "failed HTTPClient.Do")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ii, errors.Wrap(err, "resp body read err")
	}

	if err := json.Unmarshal(body, &ii); err != nil {
		return ii, errors.Wrap(err, "failed parse json")
	}

	return ii, nil
}
