package Jobs

import (
	"cricradio-go-svc/db/kafka"
	_ "cricradio-go-svc/db/kafka"
	"cricradio-go-svc/logger"
	"context"
	_ "context"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

type Commentary struct {
	TeamA      string `json:"teamA"`
	TeamB      string `json:"teamB"`
	ScoreA     string `json:"scoreA"`
	ScoreB     string `json:"scoreB"`
	Commentary string `json:"commentary"`
	Ball       string `json:"ball"`
}

func CommentaryScraper() {
	for matchId, det := range matchUrls {
		go ScrapeLatestScore(matchId, det["url"], det["previous_ball"])
	}
}

func ScrapeLatestScore(matchId string, url string, prevball string) {
	response, err := http.Get(url)
	if err != nil {
		log.Printf("Making GET request to match - %v : %v\n", matchId, err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode > 400 {
		log.Printf("Status Code for match - %v : %v\n", matchId, response.StatusCode)
		return
	}

	doc, err = goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Printf("Error reading response body of match - %v : %v\n", matchId, err)
		return
	}

	comm := &Commentary{}

	doc.Find("div.ds-text-tight-m.ds-font-regular.ds-flex.ds-px-3.ds-py-2.ds-items-baseline.ds-relative").
		EachWithBreak(func(i int, item *goquery.Selection) bool {
			ball := item.Find("span.ds-text-tight-s.ds-font-regular.ds-mb-1.ds-block.ds-text-center").Text()
			comm.Ball = strings.TrimSpace(ball)

			commentary := ""

			texter := item.Find("div.ds-ml-4")
			texter.Find("span").EachWithBreak(func(i int, selection *goquery.Selection) bool {
				commentary += selection.Text()
				return false
			})
			commentary += fmt.Sprintf(". %v", strings.TrimSpace(texter.Find("p.ci-html-content").Text()))

			comm.Commentary = commentary

			return false
		})

	if comm.Ball == prevball {
		return
	}

	var scores [][]string
	doc.Find("div.ds-text-compact-xxs.ds-p-2.ds-px-4").
		Find("div.ci-team-score.ds-flex.ds-justify-between.ds-items-center.ds-text-typo-title.ds-mb-2").
		EachWithBreak(func(i int, item *goquery.Selection) bool {
			if i > 1 {
				return false
			}

			var details []string
			details = append(details, item.Find("div.ds-flex.ds-items-center").Text())
			details = append(details, item.Find("div.ds-text-compact-m.ds-text-typo-title").Text())
			scores = append(scores, details)

			return true
		})

	if len(scores) == 2 {
		comm.TeamA = scores[0][0]
		comm.ScoreA = scores[0][1]
		comm.TeamB = scores[1][0]
		comm.ScoreB = scores[1][1]
	}

	matchUrls[matchId]["previous_ball"] = comm.Ball

	msg, err := json.Marshal(comm)
	if err != nil {
		logger.Info("Error marshalling commentary object to json !!")
		return
	}

	kafka.ProduceComm(string(msg), matchId, context.Background())

}
