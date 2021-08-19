package recommendations

import (
	"time"

	"github.com/2beens/anas-api/internal/errors"
)

type InMemApi struct {
	// userId <-> rId <-> recommendation
	recommendations map[int]map[int]*DBRecommendation
}

func NewInMemApi() *InMemApi {
	api := &InMemApi{
		recommendations: map[int]map[int]*DBRecommendation{},
	}

	userId := 1
	api.recommendations[userId] = map[int]*DBRecommendation{}
	api.recommendations[userId][0] = &DBRecommendation{
		ID:           0,
		Title:        "Healthy Gluten Free Tacos",
		PrepDuration: 35 * time.Minute,
		Difficulty:   DifficultyEasy,
		PhotoPath:    "/todo",
	}
	api.recommendations[userId][1] = &DBRecommendation{
		ID:           1,
		Title:        "Vegan Pasta",
		PrepDuration: 25 * time.Minute,
		Difficulty:   DifficultyEasy,
		PhotoPath:    "/todo",
	}
	api.recommendations[userId][2] = &DBRecommendation{
		ID:           2,
		Title:        "Burek",
		PrepDuration: 50 * time.Minute,
		Difficulty:   DifficultyNormal,
		PhotoPath:    "/todo",
	}

	return api
}

func (api *InMemApi) Get(userId, rId int) (*DBRecommendation, error) {
	recommendations, ok := api.recommendations[userId]
	if !ok {
		return nil, errors.ErrNotFound
	}
	r, found := recommendations[rId]
	if !found {
		return nil, errors.ErrNotFound
	}
	return r, nil
}

func (api *InMemApi) GetAll(userId int) ([]*DBRecommendation, error) {
	recommendations, ok := api.recommendations[userId]
	if !ok {
		return nil, errors.ErrNotFound
	}
	var foundRecommendations []*DBRecommendation
	for _, r := range recommendations {
		foundRecommendations = append(foundRecommendations, r)
	}
	return foundRecommendations, nil
}

func (api *InMemApi) Add(userId int, r *DBRecommendation) error {
	panic("implement me")
}

func (api *InMemApi) Remove(userId, rId int) error {
	panic("implement me")
}
