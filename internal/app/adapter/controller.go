package adapter

import (
	"first-go-app/internal/app/adapter/repository"
	"first-go-app/internal/app/adapter/serializer"
	"first-go-app/internal/app/application/usecase"

	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	EventRepository = repository.Event{}
	EventArgs       = usecase.EventArgs{}
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
	r.GET("/event/count", ctrl.cout_event)
	r.POST("/event", ctrl.create_event)
	return r
}

func (ctrl Controller) index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
}

func (ctrl Controller) event(c *gin.Context) {
	args := usecase.EventArgs{
		R:    EventRepository, // Dependency Injection
		FROM: c.Query("from"),
		TO:   c.Query("to"),
	}
	events, error := usecase.Event(args)

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
	} else {
		if c.Query("groupByType") == "true" && c.Query("groupByOs") == "true" {
			c.JSON(http.StatusOK, serializer.GroupByTypeAndOs(events))
		} else if c.Query("groupByType") == "true" {
			c.JSON(http.StatusOK, serializer.GroupByType(events))
		} else if c.Query("groupByOs") == "true" {
			c.JSON(http.StatusOK, serializer.GroupByOs(events))
		} else {
			c.JSON(http.StatusOK, events)
		}
	}
}

func (ctrl Controller) cout_event(c *gin.Context) {
	count, error := usecase.CountEvent(EventRepository)
	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"count": count})
	}
}

func (ctrl Controller) create_event(c *gin.Context) {
	// TODO Check si tout va bien
	args := usecase.EventArgsCreate{
		R:          EventRepository, // Dependency Injection
		Type_enum:  c.PostForm("type"),
		User_agent: c.GetHeader("User-Agent"),
		Ip:         c.ClientIP(),
	}

	event, error := usecase.CreateEvent(args)

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"newEvent": event})
	}
}
