package therapy

import (
	"fmt"

	"github.com/2beens/anas-api/internal/errors"
)

type InMemApi struct {
	Therapies map[int]map[int]*Day // user ID <-> day Index <-> TherapyDay
}

func NewInMemApi() *InMemApi {
	api := &InMemApi{
		Therapies: map[int]map[int]*Day{},
	}

	userId := 1
	api.Therapies[userId] = map[int]*Day{}
	api.Therapies[userId][1] = &Day{
		Index:         1,
		Title:         "Be calm and relax",
		AudioFilePath: "/home/serj/anas-api-data/audio/user1/inu-user1-session1.mp3",
		AudioURL:      fmt.Sprintf("/therapy/%d/%d/audio", userId, 1),
	}

	return api
}

func (api *InMemApi) GetDay(userId int, id int) (*Day, error) {
	days, ok := api.Therapies[userId]
	if !ok {
		return nil, errors.ErrNotFound
	}

	day, ok := days[id]
	if !ok {
		return nil, errors.ErrNotFound
	}

	return day, nil
}
