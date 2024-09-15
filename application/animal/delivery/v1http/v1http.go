package v1http

import (
	"github.com/gin-gonic/gin"
	"zoo/application/animal/delivery/v1http/handler"
)

type (
	AnimalV1HttpDelivery struct {
		AnimalHandler handler.IAnimalHandler
	}
)

func NewAnimalV1HttpDelivery(animalHandler handler.IAnimalHandler) *AnimalV1HttpDelivery {
	return &AnimalV1HttpDelivery{
		AnimalHandler: animalHandler,
	}
}
func (d *AnimalV1HttpDelivery) SetupRoute(group *gin.RouterGroup) {
	group.POST("/", d.AnimalHandler.Create)
	group.PUT("/:id", d.AnimalHandler.Update)
	group.GET("/:id", d.AnimalHandler.GetByID)
	group.GET("/", d.AnimalHandler.GetList)
	group.DELETE("/:id", d.AnimalHandler.Delete)
}
