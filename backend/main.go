package main

import (
	"PollerApplication/controller"
	"PollerApplication/model"
	"PollerApplication/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	userMap := make(map[string]model.User)
	pollMap := make(map[string]model.Poll)
	pollService := service.NewPollService(userMap, pollMap)
	pollController := controller.NewPollController(*pollService)
	router := setupRouter(*pollController)
	// Serve static files from the React app
	fs := http.FileServer(http.Dir("./frontend/build"))
	http.Handle("/", fs)

	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server : %v", err)
	}
}

func setupRouter(controller controller.PollController) *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	router.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "I'm healthy!",
		})
	})

	router.GET("/user/:name", controller.SaveUserToMap)

	router.POST("/poll", controller.SavePollToMap)

	router.GET("/poll/:id", controller.GetPollByIdV2)

	router.GET("/updatePoll", controller.UpdatePollResultV2)

	return router

}
