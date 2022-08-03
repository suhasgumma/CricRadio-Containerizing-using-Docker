package app

import "cricradio-go-svc/controllers/matches"

func mapUrls() {
	router.GET("/matches/list", matches.ListMatches)

}
