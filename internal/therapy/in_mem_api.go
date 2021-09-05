package therapy

import (
	"fmt"
	"log"

	"github.com/2beens/anas-api/internal/errors"
)

type InMemApi struct {
	Therapies map[int]map[int]*Day // user ID <-> day Index <-> TherapyDay
}

func NewInMemApi() *InMemApi {
	api := &InMemApi{
		Therapies: map[int]map[int]*Day{},
	}

	// hardcode therapy days for now
	userId := 1
	for i := 1; i <= 8; i++ {
		title := "In progress..."
		if i == 1 {
			title = "Be calm and relax"
		}
		api.Therapies[userId] = map[int]*Day{}
		api.Therapies[userId][i] = &Day{
			Index:         i,
			Title:         title,
			AudioFilePath: fmt.Sprintf("/home/serj/anas-api-data/audio/user1/inu-user1-session%d.mp3", i),
			AudioURL:      fmt.Sprintf("/therapy/%d/%d/audio", userId, i),
		}
	}

	log.Println("therapy days added for user 1:")
	log.Printf("%+v\n", api.Therapies[userId])

	return api
}

func (api *InMemApi) GetDay(userId int, index int) (*Day, error) {
	days, ok := api.Therapies[userId]
	if !ok {
		return nil, errors.ErrNotFound
	}

	day, ok := days[index]
	if !ok {
		return nil, errors.ErrNotFound
	}

	return day, nil
}
