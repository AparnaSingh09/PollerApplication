package service

import (
	"PollerApplication/model"
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type PollService struct {
	userMap map[string]model.User
	pollMap map[string]model.Poll
}

func (s *PollService) SaveUserInMap(ctx context.Context, name string) (model.User, error) {
	user := model.User{
		Id:   fmt.Sprint(rand.Int()),
		Name: name,
	}

	s.userMap[user.Id] = user
	log.Println("User successfully saved in memory")
	if _, exists := s.userMap[user.Id]; !exists {
		log.Fatalf("could not get key in map")
	}
	return user, nil
}

func (s *PollService) GetPollByIDV2(ctx context.Context, id string) (model.Poll, error) {
	log.Println("Getting poll by id")
	var poll model.Poll
	var exists bool
	if poll, exists = s.pollMap[id]; !exists {
		log.Printf("could not get poll key in map")
		return model.Poll{}, errors.New("could not get poll key in map")
	}
	return poll, nil
}

func (s *PollService) SavePollInMap(ctx context.Context, poll model.Poll) (string, error) {
	rand.Seed(time.Now().UnixNano())

	// Generate a random 7-digit number
	pollId := rand.Intn(9000000) + 1000000
	poll.Id = fmt.Sprint(pollId)
	poll.Result = make(map[string]int)
	for _, option := range poll.Options {
		poll.Result[option] = 0
	}

	s.pollMap[poll.Id] = poll
	log.Println("Poll successfully saved in memory")

	if _, exists := s.pollMap[poll.Id]; !exists {
		log.Fatalf("could not get poll key in map")
	}

	return poll.Id, nil

}

func (s *PollService) UpdatePollResultV2(ctx context.Context, id string, option string) (model.Poll, error) {
	var poll model.Poll
	var exists bool
	if poll, exists = s.pollMap[id]; !exists {
		log.Fatalf("could not get poll key in map")
	}
	var result int
	var ok bool
	if result, ok = poll.Result[option]; !ok {
		log.Fatalf("Option Key not found")
	}
	poll.Result[option] = result + 1

	s.pollMap[poll.Id] = poll

	log.Println("Poll successfully updated in memory")

	return poll, nil
}

func NewPollService(userMap map[string]model.User, pollMap map[string]model.Poll) *PollService {
	return &PollService{userMap: userMap, pollMap: pollMap}
}
