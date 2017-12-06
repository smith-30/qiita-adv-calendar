package model

import (
	"fmt"
	"log"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

const (
	baseUrl = "https://qiita.com/advent-calendar/2017/"
)

type (
	Calendars struct {
		C []*Calendar
	}

	Calendar struct {
		URL string
	}

	Grid struct {
		URL   string
		Title string
		Like  int
	}
)

func NewCalendars(name string, count int) *Calendars {
	cs := &Calendars{}

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

func (c *Calendar) SetExecuteURLs() <-chan *Grid {
	gridCh := make(chan *Grid, 25)

	go func() {
		doc, err := goquery.NewDocument(c.URL)
		// log.Print("parse: ", c.URL)
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
				g := &Grid{
					URL:   u,
					Title: a.Text(),
				}
				gridCh <- g
			}
		})
	}()

	return gridCh
}
