package controller

import (
	"PollerApplication/model"
	"PollerApplication/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type PollController interface {
	SaveUserToMap(c *gin.Context)
	GetPollById(c *gin.Context)
	SavePollToMap(c *gin.Context)
	UpdatePollResult(c *gin.Context)
}

type pollController struct {
	pollService service.PollService
}

func (pc *pollController) SaveUserToMap(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}
	user, err := pc.pollService.SaveUserInMap(name)
	if err != nil {
		log.Print("Failed to save user. Error:", err)
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to save user. Error: %v", err))
		return
	}
	log.Printf("Successsfully saved user in memory")
	c.JSON(http.StatusOK, user)
}

func (pc *pollController) GetPollById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	poll, err := pc.pollService.GetPollByID(id)
	if err != nil {
		log.Print("Failed to get poll by id")
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to get poll. Error: %v", err))
		return
	}
	log.Printf("Successsfully retrived poll by ID")
	c.JSON(http.StatusOK, poll)
}

func (pc *pollController) SavePollToMap(c *gin.Context) {
	var poll model.Poll
	fmt.Println(poll)
	err := c.BindJSON(&poll)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind request"})
		return
	}
	pollId, err := pc.pollService.SavePollToMap(poll)
	if err != nil {
		log.Print("Failed to save poll to map")
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to save poll. Error: %v", err))
		return
	}
	log.Printf("Successsfully saved poll in memory")
	c.JSON(http.StatusOK, pollId)
}

func (pc *pollController) UpdatePollResult(c *gin.Context) {
	id := c.Query("id")
	option := c.Query("option")
	if id == "" || option == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID and option are required"})
		return
	}
	updatedPoll, err := pc.pollService.UpdatePollResult(id, option)
	if err != nil {
		log.Print("Failed to update poll")
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to update poll. Error: %v", err))
		return
	}
	log.Printf("Successsfully updated poll in memory")
	c.JSON(http.StatusOK, updatedPoll)
}

func NewPollController(pollService service.PollService) PollController {
	return &pollController{
		pollService: pollService,
	}
}
