package adapter

import (
	"first-go-app/internal/app/adapter/repository"
	"first-go-app/internal/app/application/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	EventRepository = repository.Event{}
)

// Controller is a controller
type Controller struct{}

// Router is routing settings
func Router() *gin.Engine {
	r := gin.Default()
	ctrl := Controller{}
	// NOTICE: following path is from CURRENT directory, so please run Gin from root directory
	r.GET("/", ctrl.index)
	r.GET("/event", ctrl.event)
	return r
}

func (ctrl Controller) index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
}

func (ctrl Controller) event(c *gin.Context) {
	event := usecase.Event(EventRepository) // Dependency Injection
	c.JSON(200, event)
}
