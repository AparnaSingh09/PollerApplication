package service

import (
	"PollerApplication/model"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

//go:generate mockgen -destination=mock_poll_service.go -package=service . PollService

var RandomInt = rand.Int
var RandomIntn = func(n int) int {
	return rand.Intn(n)
}

type PollService interface {
	SaveUserInMap(name string) (model.User, error)
	GetPollByID(id string) (model.Poll, error)
	SavePollToMap(poll model.Poll) (string, error)
	UpdatePollResult(id string, option string) (model.Poll, error)
}

type pollService struct {
	userMap map[string]model.User
	pollMap map[string]model.Poll
}

func (s *pollService) SaveUserInMap(name string) (model.User, error) {
	log.Println("Saving user in map")
	user := model.User{
		Id:   fmt.Sprint(RandomInt()),
		Name: name,
	}

	if _, exists := s.userMap[user.Id]; exists {
		log.Println("Id already exists in the map")
		return model.User{}, errors.New("Id already exists in the map")

	}
	s.userMap[user.Id] = user
	log.Println("User successfully saved in memory")
	return user, nil
}

func (s *pollService) GetPollByID(id string) (model.Poll, error) {
	log.Println("Getting poll by id")
	var poll model.Poll
	var exists bool
	if poll, exists = s.pollMap[id]; !exists {
		log.Printf("could not get poll key in map")
		return model.Poll{}, errors.New("could not get poll key in map")
	}
	log.Println("Successfully returned poll")
	return poll, nil
}

func (s *pollService) SavePollToMap(poll model.Poll) (string, error) {
	log.Println("Saving poll in poll map")
	rand.Seed(time.Now().UnixNano())
	// Generate a random 7-digit number
	pollId := RandomIntn(9000000) + 1000000
	poll.Id = fmt.Sprint(pollId)
	poll.Result = make(map[string]int)
	for _, option := range poll.Options {
		poll.Result[option] = 0
	}

	if _, exists := s.pollMap[poll.Id]; exists {
		log.Println("Poll Id already exists in the map")
		return "", errors.New("Poll Id already exists in the map")
	}

	s.pollMap[poll.Id] = poll

	log.Println("Successfully saved poll in memory")
	return poll.Id, nil

}

func (s *pollService) UpdatePollResult(id string, option string) (model.Poll, error) {
	var poll model.Poll
	var exists bool
	if poll, exists = s.pollMap[id]; !exists {
		log.Println("could not get poll key in map")
		return model.Poll{}, errors.New("could not get poll key in map")
	}
	var result int
	var ok bool
	if result, ok = poll.Result[option]; !ok {
		log.Println("option key not found")
		return model.Poll{}, errors.New("option key not found")
	}
	poll.Result[option] = result + 1

	s.pollMap[poll.Id] = poll

	log.Println("Successfully updated poll in memory")
	return poll, nil
}

func NewPollService(userMap map[string]model.User, pollMap map[string]model.Poll) PollService {
	return &pollService{userMap: userMap, pollMap: pollMap}
}
