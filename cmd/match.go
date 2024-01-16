package cmd

import (
	"sort"
	"strings"
)

type MatchScore struct {
	CanvasUser CanvasUser
	Score      float64
}

func calculateMatchScore(canvasUser CanvasUser, unlockEdUser UnlockEdUser) float64 {
	name_first := strings.TrimSpace(strings.Split(canvasUser.SortableName, ",")[1])
	name_last := strings.TrimSpace(strings.Split(canvasUser.SortableName, ",")[0])
	var total float64
	const (
		nameWeight    = 0.5 // last + first both must match
		loginIdWeight = 0.3 // this field is generally the email, so a match on a split("@")[0] -> username is a good sign
	)
	if strings.Contains(canvasUser.LoginId, "@") {
		username := strings.Split(canvasUser.LoginId, "@")[0]
		if username == unlockEdUser.Username {
			total += loginIdWeight
		}
	}
	if unlockEdUser.NameFirst == name_first && unlockEdUser.NameLast == name_last {
		total += nameWeight
	}
	// nameSimilarity := similarity(canvasUser.Name, unlockEdUser.NameFirst+unlockEdUser.NameLast)
	// sortableNameSimilarity := similarity(canvasUser.SortableName, unlockEdUser.NameFirst+unlockEdUser.NameLast)
	// shortNameSimilarity := similarity(canvasUser.ShortName, unlockEdUser.NameFirst+unlockEdUser.NameLast)

	return total
}

func similarity(a, b string) float64 {
	// TODO: implement similarity/matching algo

	if a == b {
		return 1
	}
	return 0
}

func (db *DbHandler) FindBestMatches(unlockEdUser UnlockEdUser) []MatchScore {
	similarUsers := db.QuerySimilarUsers(&unlockEdUser)
	var matches []MatchScore
	for _, user := range similarUsers {
		score := calculateMatchScore(user, unlockEdUser)
		matches = append(matches, MatchScore{CanvasUser: user, Score: score})
	}

	// Sort matches by score, highest first
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].Score > matches[j].Score
	})
	return matches
}
