package model

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

const (
	baseUrl    = "https://qiita.com/advent-calendar/2017/"
	apiBaseUrl = "https://qiita.com/api/v2/items/"
)

type (
	Calendars struct {
		C  []*Calendar
		wg *sync.WaitGroup
	}

	Calendar struct {
		URL string
	}

	Grid struct {
		URL      string
		QiitaURL string
		Title    string
		Like     int
	}
)

func NewCalendars(name string, count int) *Calendars {
	cs := &Calendars{
		wg: new(sync.WaitGroup),
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

func (cs *Calendars) addCalendar(name string) {
	url := baseUrl + name
	c := &Calendar{
		URL: url,
	}
	cs.C = append(cs.C, c)
}

func (cs *Calendars) Wait() {
	cs.wg.Wait()
}

func (cs *Calendars) FetchGrids(gridUpdateCh chan *Grid) {
	for _, ca := range cs.C {
		cs.wg.Add(1)
		go func(c *Calendar) {
			gridCh := c.SetExecuteURLs()

			for g := range gridCh {
				gridUpdateCh <- g
			}
			cs.wg.Done()
		}(ca)
	}
}

func (c *Calendar) SetExecuteURLs() <-chan *Grid {
	gridCh := make(chan *Grid, 25)

	go func() {
		defer close(gridCh)

		doc, err := goquery.NewDocument(c.URL)
		if err != nil {
			log.Fatal(err)
		}

		doc.Find(".adventCalendarCalendar_comment").Each(func(i int, s *goquery.Selection) {
			a := s.Find("a")
			u, _ := a.Attr("href")

			result, err := url.Parse(u)
			if err != nil {
				log.Fatal(err)
			}

			if result.Host == "qiita.com" {
				i := getItemID(u)
				g := &Grid{
					URL:      apiBaseUrl + i,
					QiitaURL: u,
					Title:    a.Text(),
				}
				gridCh <- g
			}
		})
	}()

	return gridCh
}

func getItemID(url string) string {
	spr := strings.Split(url, "/")
	id := spr[len(spr)-1]
	return id
}
