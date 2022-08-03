package Jobs

import (
	"cricradio-go-svc/db/kafka"
	"cricradio-go-svc/domain/matches"
	"cricradio-go-svc/logger"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var (
	source    = "https://www.espncricinfo.com/live-cricket-score"
	srcDomain = "https://www.espncricinfo.com"
	re        = regexp.MustCompile(`(?m)series/([^//]*)/([^//]*)/(live-cricket-score|full-scorecard)`)
	doc       *goquery.Document

	matchUrls = make(map[string]map[string]string)
)

func LiveScraper() {

	response, err := http.Get(source)
	if err != nil {
		log.Fatalf("Making GET request to src : %v\n", err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode > 400 {
		fmt.Println("Status Code : ", response.StatusCode)
		return
	}

	doc, err = goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatalf("Error reading response body : %v\n", err)
		return
	}

	go UpdateMatchesDB()

	//go FetchScores()

}

func UpdateMatchesDB() {

	doc.Find("div.ds-px-4.ds-py-3").Find("div.ds-text-compact-xxs").Each(func(i int, item *goquery.Selection) {

		status := item.Find("span.ds-text-tight-xs.ds-font-bold.ds-uppercase.ds-leading-5").Text()

		if status == "Live" || status == "Drinks" {

			var match matches.Match
			var teams []string
			item.Find("p.ds-text-tight-m.ds-font-bold.ds-capitalize").Each(func(i int, selection *goquery.Selection) {
				teams = append(teams, selection.Text())
			})
			detail := item.Find("div.ds-text-tight-xs.ds-truncate.ds-text-ui-typo-mid").Text()
			match.Teams = fmt.Sprintf("%v VS %v", teams[0], teams[1])
			match.Details = detail
			id, exist := item.Find("a").Attr("href")
			if exist {
				ids := re.FindStringSubmatch(id)
				match.SeriesId = ids[1]
				match.MatchId = ids[2]

				id = strings.ReplaceAll(id, "live-cricket-score", "ball-by-ball-commentary")
				match.URL = fmt.Sprintf("%v%v", srcDomain, id)
			}

			matchUrls[match.MatchId] = map[string]string{"url": match.URL, "previous_ball": "-1"}
			rows, err := match.Insert()
			if err != nil {
				logger.Info("Error inserting match into DB : ")
			}

			if rows > 0 {
				go kafka.CreateTopic(match.MatchId)
			}

		} else if status == "RESULT" || status == "Stumps" {
			id, exist := item.Find("a").Attr("href")
			if exist {
				ids := re.FindStringSubmatch(id)
				matchId := ids[2]
				m := &matches.Match{}
				rows, err := m.Delete(matchId)
				if err != nil {
					logger.Info("Error deleting match from DB : ")
				}
				delete(matchUrls, matchId)

				if rows > 0 {
					go kafka.DeleteTopic(matchId)
				}

			}
		}
	})
}
