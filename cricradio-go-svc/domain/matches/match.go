package matches

import (
	"cricradio-go-svc/logger"
	"encoding/json"
	"fmt"
)

type Match struct {
	Teams    string `json:"teams"`
	Details  string `json:"details"`
	SeriesId string `json:"seriesId"`
	MatchId  string `json:"matchId"`
	URL      string `json:"url"`
}

type Matches []Match

func (m *Match) Validate() {
	if m.MatchId == "" {
		logger.Info(fmt.Sprintf("Something wrong with match object : %v\n", m))
	}
	if m.URL == "" {
		logger.Info(fmt.Sprintf("Something wrong with match object : %v\n", m))
	}
}

func (m Matches) Marshall() []interface{} {
	result := make([]interface{}, len(m))
	for index, match := range m {
		result[index] = match.Marshall()
	}

	return result
}

func (m *Match) Marshall() interface{} {

	userJson, _ := json.Marshal(m)

	var match Match
	json.Unmarshal(userJson, &match)
	return match
}
