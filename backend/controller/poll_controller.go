package controller

import (
	"PollerApplication/model"
	"PollerApplication/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type PollController struct {
	pollService service.PollService
}

func (pc *PollController) SaveUserToMap(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}
	user, err := pc.pollService.SaveUserInMap(c.Request.Context(), name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to save user. Error: %v", err))
		return
	}
	log.Printf("Successsfully saved user in memory")
	c.JSON(http.StatusOK, user)
}

func (pc *PollController) GetPollByIdV2(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	poll, err := pc.pollService.GetPollByIDV2(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to get poll. Error: %v", err))
		return
	}
	log.Printf("Successsfully retrived poll by ID")
	c.JSON(http.StatusOK, poll)
}

func (pc *PollController) SavePollToMap(c *gin.Context) {
	var poll model.Poll
	fmt.Println(poll)
	err := c.BindJSON(&poll)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind request"})
		return
	}
	pollId, err := pc.pollService.SavePollInMap(c.Request.Context(), poll)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to save poll. Error: %v", err))
		return
	}
	log.Printf("Successsfully saved poll in memory")
	c.JSON(http.StatusOK, pollId)
}

func (pc *PollController) UpdatePollResultV2(c *gin.Context) {
	id := c.Query("id")
	option := c.Query("option")
	updatedPoll, err := pc.pollService.UpdatePollResultV2(c.Request.Context(), id, option)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to save poll. Error: %v", err))
		return
	}
	log.Printf("Successsfully updated poll in memory")
	c.JSON(http.StatusOK, updatedPoll)
}

func NewPollController(pollService service.PollService) *PollController {
	return &PollController{
		pollService: pollService,
	}
}
