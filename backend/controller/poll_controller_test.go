package controller

import (
	"PollerApplication/model"
	"PollerApplication/service"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type PollControllerTestSuite struct {
	suite.Suite
	recorder       *httptest.ResponseRecorder
	context        *gin.Context
	mockController *gomock.Controller
	service        *service.MockPollService
	pollController PollController
}

func TestPollControllerController(t *testing.T) {
	suite.Run(t, new(PollControllerTestSuite))
}

func (suite *PollControllerTestSuite) SetupTest() {
	suite.recorder = httptest.NewRecorder()
	suite.context, _ = gin.CreateTestContext(suite.recorder)
	suite.mockController = gomock.NewController(suite.T())
	suite.service = service.NewMockPollService(suite.mockController)
	suite.pollController = NewPollController(suite.service)
}

func (suite *PollControllerTestSuite) TearDownTest() {
	suite.mockController.Finish()
}

func (suite *PollControllerTestSuite) TestSaveUserToMap() {
	suite.context.Request, _ = http.NewRequest(http.MethodPost, "/some-url", nil)
	suite.context.Params = gin.Params{gin.Param{Key: "name", Value: "Aparna"}}
	response := model.User{
		Id:             "74824",
		Name:           "Aparna",
		Poll:           model.Poll{},
		SelectedOption: "",
	}
	suite.service.EXPECT().SaveUserInMap("Aparna").Return(response, nil).Times(1)
	suite.pollController.SaveUserToMap(suite.context)
	suite.Equal(http.StatusOK, suite.recorder.Code)
	expectedResponseBody, err := json.Marshal(response)
	suite.Nil(err)
	suite.Equal(string(expectedResponseBody), suite.recorder.Body.String())
}

func (suite *PollControllerTestSuite) TestSaveUserToMap_ShouldReturn400Error_WhenNameIsNotPresent() {
	suite.context.Request, _ = http.NewRequest(http.MethodPost, "/some-url", nil)
	suite.pollController.SaveUserToMap(suite.context)
	suite.Equal(http.StatusBadRequest, suite.recorder.Code)
	suite.Equal("{\"error\":\"Name is required\"}", suite.recorder.Body.String())
}

func (suite *PollControllerTestSuite) TestSaveUserToMap_ShouldReturn500_WhenServiceReturnsError() {
	suite.context.Request, _ = http.NewRequest(http.MethodPost, "/some-url", nil)
	suite.context.Params = gin.Params{gin.Param{Key: "name", Value: "Aparna"}}
	response := model.User{}
	serviceErr := errors.New("some err")
	suite.service.EXPECT().SaveUserInMap("Aparna").Return(response, serviceErr).Times(1)
	suite.pollController.SaveUserToMap(suite.context)
	suite.Equal(http.StatusInternalServerError, suite.recorder.Code)
	suite.Equal("\"Failed to save user. Error: some err\"", suite.recorder.Body.String())
}

func (suite *PollControllerTestSuite) TestGetPollById() {
	suite.context.Request, _ = http.NewRequest(http.MethodPost, "/some-url", nil)
	suite.context.Params = gin.Params{gin.Param{Key: "id", Value: "8765"}}
	response := model.Poll{
		Id:       "8765",
		Name:     "Poll1",
		Question: "",
		Options:  nil,
		Result:   nil,
	}
	suite.service.EXPECT().GetPollByID("8765").Return(response, nil).Times(1)
	suite.pollController.GetPollById(suite.context)
	suite.Equal(http.StatusOK, suite.recorder.Code)
	expectedResponseBody, err := json.Marshal(response)
	suite.Nil(err)
	suite.Equal(string(expectedResponseBody), suite.recorder.Body.String())
}

func (suite *PollControllerTestSuite) TestGetPollById_ShouldReturn400Error_WhenIdIsNotPresent() {
	suite.context.Request, _ = http.NewRequest(http.MethodPost, "/some-url", nil)
	suite.pollController.GetPollById(suite.context)
	suite.Equal(http.StatusBadRequest, suite.recorder.Code)
	suite.Equal("{\"error\":\"ID is required\"}", suite.recorder.Body.String())
}

func (suite *PollControllerTestSuite) TestGetPollById_ShouldReturn500_WhenServiceReturnsError() {
	suite.context.Request, _ = http.NewRequest(http.MethodPost, "/some-url", nil)
	suite.context.Params = gin.Params{gin.Param{Key: "id", Value: "348743"}}
	response := model.Poll{}
	serviceErr := errors.New("some err")
	suite.service.EXPECT().GetPollByID("348743").Return(response, serviceErr).Times(1)
	suite.pollController.GetPollById(suite.context)
	suite.Equal(http.StatusInternalServerError, suite.recorder.Code)
	suite.Equal("\"Failed to get poll. Error: some err\"", suite.recorder.Body.String())
}

func (suite *PollControllerTestSuite) TestSavePollToMap() {
	request := model.Poll{
		Name:     "Poll1",
		Question: "How are you?",
		Options:  []string{"Good", "Great"},
		Result:   nil,
	}
	requestAsString, err := json.Marshal(request)
	suite.Nil(err)
	suite.context.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(string(requestAsString)))

	suite.service.EXPECT().SavePollToMap(request).Return("47635", nil).Times(1)
	suite.pollController.SavePollToMap(suite.context)
	suite.Equal(http.StatusOK, suite.recorder.Code)
	suite.Equal("\"47635\"", suite.recorder.Body.String())
}

func (suite *PollControllerTestSuite) TestSavePollToMap_ShouldReturn400Error_WhenBadRequestPresent() {
	request := model.Poll{
		Name:     "Poll1",
		Question: "How are you?",
		Result:   nil,
	}
	requestAsString, err := json.Marshal(request)
	suite.Nil(err)
	suite.context.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(string(requestAsString)))

	suite.pollController.SavePollToMap(suite.context)
	suite.Equal(http.StatusBadRequest, suite.recorder.Code)
	suite.Equal("{\"error\":\"Failed to bind request\"}", suite.recorder.Body.String())
}

func (suite *PollControllerTestSuite) TestSavePollToMap_ShouldReturn500_WhenServiceReturnsError() {
	request := model.Poll{
		Name:     "Poll1",
		Question: "How are you?",
		Options:  []string{"Good", "Great"},
		Result:   nil,
	}
	requestAsString, err := json.Marshal(request)
	suite.Nil(err)
	suite.context.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(string(requestAsString)))

	suite.service.EXPECT().SavePollToMap(request).Return("", errors.New("some err")).Times(1)
	suite.pollController.SavePollToMap(suite.context)
	suite.Equal(http.StatusInternalServerError, suite.recorder.Code)
	suite.Equal("\"Failed to save poll. Error: some err\"", suite.recorder.Body.String())
}

func (suite *PollControllerTestSuite) TestUpdatePollResult() {
	suite.context.Request, _ = http.NewRequest(http.MethodPost, "/some-url?id=7216&option=Good", nil)
	poll := model.Poll{
		Id:       "7216",
		Name:     "Poll1",
		Question: "How are you?",
		Options:  []string{"Good", "Great"},
		Result:   nil,
	}
	suite.service.EXPECT().UpdatePollResult("7216", "Good").Return(poll, nil).Times(1)
	suite.pollController.UpdatePollResult(suite.context)
	suite.Equal(http.StatusOK, suite.recorder.Code)
	expectedResponseBody, err := json.Marshal(poll)
	suite.Nil(err)
	suite.Equal(string(expectedResponseBody), suite.recorder.Body.String())
}

func (suite *PollControllerTestSuite) TestUpdatePollResult_ShouldReturn400Error_WhenBadRequestPresent() {
	suite.context.Request, _ = http.NewRequest(http.MethodPost, "/some-url?option=Good", nil)
	suite.pollController.UpdatePollResult(suite.context)
	suite.Equal(http.StatusBadRequest, suite.recorder.Code)
	suite.Equal("{\"error\":\"ID and option are required\"}", suite.recorder.Body.String())
}

func (suite *PollControllerTestSuite) TestUpdatePollResult_ShouldReturn500_WhenServiceReturnsError() {
	suite.context.Request, _ = http.NewRequest(http.MethodPost, "/some-url?id=7216&option=Good", nil)
	serviceErr := errors.New("some err")
	suite.service.EXPECT().UpdatePollResult("7216", "Good").Return(model.Poll{}, serviceErr).Times(1)
	suite.pollController.UpdatePollResult(suite.context)
	suite.Equal(http.StatusInternalServerError, suite.recorder.Code)
	suite.Equal("\"Failed to update poll. Error: some err\"", suite.recorder.Body.String())
}
