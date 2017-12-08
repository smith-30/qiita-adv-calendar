package model

import (
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	parser "github.com/smith-30/go-goo.gl-parser"
	"go.uber.org/zap"
)

const (
	apiBaseUrl = "https://qiita.com/api/v2/items/"
)

type (
	Calendar struct {
		URL    string
		parser parser.Parser

		logger *zap.SugaredLogger
	}

	Grid struct {
		URL      string
		QiitaURL string
		Title    string
		Like     int
	}
)

func NewCalendar(u string, l *zap.SugaredLogger) *Calendar {
	return &Calendar{
		URL:    u,
		parser: parser.NewParser(os.Getenv("URL_SHORTER_API_KEY")),
		logger: l,
	}
}

func (c *Calendar) SetExecuteURLs() <-chan *Grid {
	gridCh := make(chan *Grid, 25)

	go func() {
		defer close(gridCh)

		doc, err := goquery.NewDocument(c.URL)
		if err != nil {
			c.logger.Fatal(err)
			return
		}

		doc.Find(".adventCalendarCalendar_comment").EachWithBreak(func(i int, s *goquery.Selection) bool {
			a := s.Find("a")
			u, _ := a.Attr("href")

			if u == "" {
				return false
			}

			result, err := url.Parse(u)
			if err != nil {
				c.logger.Fatal(err)
				return false
			}

			if result.Host == "goo.gl" {
				u, err = c.parser.DecodeURL(u)
				if err != nil {
					c.logger.Errorf("DecodeURL failed: %s", err)
				}

				result, err = url.Parse(u)
				if err != nil {
					c.logger.Fatal(err)
					return false
				}
			}

			c.logger.Infof("get grid url: %s", u)

			if result.Host == "qiita.com" {
				i := getItemID(u)
				g := &Grid{
					URL:      apiBaseUrl + i,
					QiitaURL: u,
					Title:    a.Text(),
				}
				gridCh <- g
			}

			return true
		})
	}()

	return gridCh
}

func getItemID(url string) string {
	spr := strings.Split(url, "/")
	id := spr[len(spr)-1]
	return id
}
