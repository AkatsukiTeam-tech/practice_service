package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"practive_service/domain/dto"
	"practive_service/domain/service"
)

type PersonController struct {
	PersonAppService *service.PersonAppService
}

func (p *PersonController) AddPerson(c *gin.Context) {
	var personDto dto.Person
	err := c.BindJSON(&personDto)
	if err != nil {
		c.JSON(400, gin.H{"msg": "Error while BindingJSON notification template"})
		return
	}

	err = p.PersonAppService.AddPerson(personDto)
	if err != nil {
		log.Printf("error when save to db: %s", err)
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
}
