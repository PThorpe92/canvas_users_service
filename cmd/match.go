package cmd

import "sort"

type MatchScore struct {
	CanvasUser CanvasUser
	Score      float64
}

func calculateMatchScore(canvasUser CanvasUser, unlockEdUser UnlockEdUser) float64 {
	// Define weights for each field
	const (
		nameWeight         = 0.3
		sortableNameWeight = 0.2
		shortNameWeight    = 0.2
		// ... other weights
	)

	// Calculate similarity for each field (you can use a custom similarity function)
	nameSimilarity := similarity(canvasUser.Name, unlockEdUser.NameFirst+unlockEdUser.NameLast)
	sortableNameSimilarity := similarity(canvasUser.SortableName, unlockEdUser.NameFirst+unlockEdUser.NameLast)
	shortNameSimilarity := similarity(canvasUser.ShortName, unlockEdUser.NameFirst+unlockEdUser.NameLast)
	// ... other similarities

	// Calculate total score
	score := nameSimilarity*nameWeight + sortableNameSimilarity*sortableNameWeight + shortNameSimilarity*shortNameWeight // ... add other scores

	return score
}

// Implement a similarity function (e.g., Levenshtein distance, Jaccard index, etc.)
func similarity(a, b string) float64 {
	// Dummy implementation, replace with actual algorithm
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
