package matches

import (
	"cricradio-go-svc/db/mysql/matches_db"
	"cricradio-go-svc/logger"
	"cricradio-go-svc/utils/errors"
	"fmt"
)

const (
	indexUniqueEmail   = "email_UNIQUE"
	errorNoRows        = "no rows in result set"
	queryInsertUser    = "INSERT IGNORE INTO matches(matchId, seriesId, teams, details, url) VALUES(?,?,?,?,?);"
	queryListMatches   = "SELECT matchId,seriesId,teams,details,url FROM matches;"
	queryDeleteMatches = "DELETE FROM matches WHERE matchId=?;"
)

func (m *Match) Insert() (int64, *errors.RestErr) {
	stmt, err := matches_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Info("error when trying to prepare insert match statement!")
		return -1, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	result, err := stmt.Exec(m.MatchId, m.SeriesId, m.Teams, m.Details, m.URL)
	if err != nil {
		logger.Info("error when trying to insert match!")
		return -1, errors.NewInternalServerError("database error")
	}

	rows, _ := result.RowsAffected()
	return rows, nil
}

func (m *Match) ListMatches() ([]Match, *errors.RestErr) {
	stmt, err := matches_db.Client.Prepare(queryListMatches)
	if err != nil {
		logger.Error("error when trying to prepare list matches statement!", err)
		return nil, errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		logger.Error("error when trying to list matches!", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()
	results := make([]Match, 0)
	for rows.Next() {
		var match Match
		if err := rows.Scan(&match.MatchId, &match.SeriesId, &match.Teams, &match.Details, &match.URL); err != nil {
			logger.Error("error when trying to scan row to user struct!", err)
			return nil, errors.NewInternalServerError("database error")
		}

		results = append(results, match)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no matches to list"))
	}
	return results, nil

}

func (m *Match) Delete(matchId string) (int64, *errors.RestErr) {
	stmt, err := matches_db.Client.Prepare(queryDeleteMatches)
	if err != nil {
		logger.Info("error when trying to prepare delete match statement!")
		return -1, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	result, err := stmt.Exec(matchId)
	if err != nil {
		logger.Info("error when trying to delete match!")
		return -1, errors.NewInternalServerError("database error")
	}

	rows, _ := result.RowsAffected()

	return rows, nil
}
