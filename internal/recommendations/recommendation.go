package recommendations

import (
	"time"
)

type Difficulty string

const (
	DifficultyEasy        Difficulty = "easy"
	DifficultyNormal      Difficulty = "normal"
	DifficultyComplicated Difficulty = "complicated"
)

type DBRecommendation struct {
	ID           int
	Title        string
	PrepDuration time.Duration
	Difficulty   Difficulty
	PhotoPath    string
}

func (r *DBRecommendation) ToRecommendation() *Recommendation {
	return &Recommendation{
		ID:           r.ID,
		Title:        r.Title,
		PrepDuration: r.PrepDuration.String(),
		Difficulty:   r.Difficulty,
		PhotoUrl:     "/todo",
	}
}

type Recommendation struct {
	ID           int        `json:"id"`
	Title        string     `json:"title"`
	PrepDuration string     `json:"prep_duration"`
	Difficulty   Difficulty `json:"difficulty"`
	PhotoUrl     string     `json:"photo_url"`
}
