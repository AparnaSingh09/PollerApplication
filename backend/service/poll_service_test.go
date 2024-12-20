package service

import (
	"PollerApplication/model"
	"errors"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PollServiceTestSuite struct {
	suite.Suite
	userMap map[string]model.User
	pollMap map[string]model.Poll
	service PollService
}

func TestPollService(t *testing.T) {
	suite.Run(t, new(PollServiceTestSuite))
}

func (suite *PollServiceTestSuite) SetupTest() {
	suite.userMap = map[string]model.User{
		"985": {
			Id:   "985",
			Name: "User1",
			Poll: model.Poll{
				Id:       "1094238",
				Name:     "Poll1",
				Question: "how are you?",
				Options:  []string{"Good", "Great", "fine"},
				Result: map[string]int{
					"Good":  0,
					"Great": 2,
					"Fine":  1,
				},
			},
			SelectedOption: "Good",
		},
		"0989": {
			Id:   "0989",
			Name: "User2",
			Poll: model.Poll{
				Id:       "1094238",
				Name:     "Poll1",
				Question: "how are you?",
				Options:  []string{"Good", "Great", "fine"},
				Result: map[string]int{
					"Good":  0,
					"Great": 2,
					"Fine":  1,
				},
			},
			SelectedOption: "Great",
		},
		"0981": {
			Id:   "0981",
			Name: "User3",
			Poll: model.Poll{
				Id:       "94239",
				Name:     "Poll2",
				Question: "What day is today?",
				Options:  []string{"Sunday", "Monday"},
				Result: map[string]int{
					"Sunday": 0,
					"Monday": 2,
				},
			},
			SelectedOption: "Sunday",
		},
	}
	suite.pollMap = map[string]model.Poll{
		"1094238": {
			Id:       "1094238",
			Name:     "Poll1",
			Question: "how are you?",
			Options:  []string{"Good", "Great", "fine"},
			Result: map[string]int{
				"Good":  0,
				"Great": 2,
				"Fine":  1,
			},
		},
		"94239": {
			Id:       "94239",
			Name:     "Poll2",
			Question: "What day is today?",
			Options:  []string{"Sunday", "Monday"},
			Result: map[string]int{
				"Sunday": 0,
				"Monday": 2,
			},
		},
	}
	suite.service = NewPollService(suite.userMap, suite.pollMap)
}

func (suite *PollServiceTestSuite) TearDownTest() {
	suite.userMap = nil
	suite.pollMap = nil
}

func (suite *PollServiceTestSuite) TestSaveUserInMap_ShouldCreateAndReturnUser_WhenValidNameIsPassed() {
	RandomInt = func() int {
		return 1234
	}
	expectedUser := model.User{
		Id:             "1234",
		Name:           "Aparna",
		Poll:           model.Poll{},
		SelectedOption: "",
	}
	user, err := suite.service.SaveUserInMap("Aparna")
	suite.Nil(err)
	suite.Equal(expectedUser, user)
}

func (suite *PollServiceTestSuite) TestSaveUserInMap_ShouldThrowError_WhenUnableToSaveUserInMap() {
	RandomInt = func() int {
		return 985
	}
	expectedError := errors.New("Id already exists in the map")
	user, err := suite.service.SaveUserInMap("Aparna")
	suite.NotNil(err)
	suite.Empty(user)
	suite.Equal(expectedError, err)
}

func (suite *PollServiceTestSuite) TestGetPollByID_ShouldReturnPoll_WheValidIdIsPassed() {
	expectedPoll := model.Poll{
		Id:       "1094238",
		Name:     "Poll1",
		Question: "how are you?",
		Options:  []string{"Good", "Great", "fine"},
		Result: map[string]int{
			"Good":  0,
			"Great": 2,
			"Fine":  1,
		},
	}
	poll, err := suite.service.GetPollByID("1094238")
	suite.Nil(err)
	suite.Equal(expectedPoll, poll)
}

func (suite *PollServiceTestSuite) TestGetPollByID_ShouldReturnError_WheInValidIdIsPassed() {
	expectedErr := errors.New("could not get poll key in map")
	poll, err := suite.service.GetPollByID("942")
	suite.NotNil(err)
	suite.Empty(poll)
	suite.Equal(expectedErr, err)
}

func (suite *PollServiceTestSuite) TestSavePollInMap_ShouldSavePoll_WheValidDataIsPassed() {
	RandomIntn = func(n int) int {
		return 42
	}
	poll := model.Poll{
		Name:     "Poll3",
		Question: "What date is it?",
		Options:  nil,
		Result:   nil,
	}
	id, err := suite.service.SavePollToMap(poll)
	suite.Nil(err)
	suite.Equal("1000042", id)
}

func (suite *PollServiceTestSuite) TestSavePollToMap_ShouldThrowError_WhenUnableToSavePollToMap() {
	RandomIntn = func(n int) int {
		return 94238
	}
	poll := model.Poll{
		Name:     "Poll3",
		Question: "What date is it?",
		Options:  nil,
		Result:   nil,
	}
	expectedErr := errors.New("Poll Id already exists in the map")
	id, err := suite.service.SavePollToMap(poll)
	suite.NotNil(err)
	suite.Empty(id)
	suite.Equal(expectedErr, err)
}

func (suite *PollServiceTestSuite) TestUpdatePollResult_ShouldUpdateAndReturnPoll_WheValidIdIsPassed() {
	expectedPoll := model.Poll{
		Id:       "1094238",
		Name:     "Poll1",
		Question: "how are you?",
		Options:  []string{"Good", "Great", "fine"},
		Result: map[string]int{
			"Good":  0,
			"Great": 3,
			"Fine":  1,
		},
	}
	poll, err := suite.service.UpdatePollResult("1094238", "Great")
	suite.Nil(err)
	suite.Equal(expectedPoll, poll)
}

func (suite *PollServiceTestSuite) TestUpdatePollResult_ShouldReturnError_WheInValidIdIsPassed() {
	expecetdErr := errors.New("could not get poll key in map")
	poll, err := suite.service.UpdatePollResult("1498574", "Great")
	suite.NotNil(err)
	suite.Empty(poll)
	suite.Equal(expecetdErr, err)
}

func (suite *PollServiceTestSuite) TestUpdatePollResult_ShouldReturnError_WheInValidOptionIsPassed() {
	expecetdErr := errors.New("option key not found")
	poll, err := suite.service.UpdatePollResult("1094238", "okayish")
	suite.NotNil(err)
	suite.Empty(poll)
	suite.Equal(expecetdErr, err)
}
