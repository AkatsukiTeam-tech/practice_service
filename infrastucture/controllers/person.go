package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"practive_service/domain/dto"
	"practive_service/domain/service"
	"practive_service/infrastucture/helpers"
	"strconv"
)

type PersonController struct {
	PersonAppService *service.PersonAppService `inject:""`
}

func (p *PersonController) AddPerson(c *gin.Context) {
	var personDto dto.Person
	err := c.BindJSON(&personDto)
	if err != nil {
		c.JSON(400, gin.H{"msg": "Error while BindingJSON person"})
		return
	}

	err = p.PersonAppService.AddPerson(personDto)
	if err != nil {
		log.Printf("error when save to db: %s", err)
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"msg": " Person created"})
}

func (p *PersonController) FindPersonById(c *gin.Context) {
	id, _ := strconv.Atoi(c.GetHeader("id"))

	person, err := p.PersonAppService.FindPersonById(id)
	if err != nil {
		log.Printf("error when save to db: %s", err)
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, person)
}

func (p *PersonController) UpdatePerson(c *gin.Context) {
	var personDto dto.Person
	err := c.BindJSON(&personDto)
	if err != nil {
		c.JSON(400, gin.H{"msg": "Error while BindingJSON person"})
		return
	}

	person, err := p.PersonAppService.UpdatePerson(personDto)
	if err != nil {
		log.Printf("error when update the person with id: %s", personDto.ID)
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, person)
}

func (p *PersonController) DeletePerson(c *gin.Context) {
	id, _ := strconv.Atoi(c.GetHeader("id"))

	err := p.PersonAppService.DeletePerson(id)
	if err != nil {
		log.Printf("error when update the person with id: %s", id)
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(201, gin.H{"msg": "Success deleted"})
}

func (p *PersonController) GetAllPerson(c *gin.Context) {
	paging := helpers.PaginateRequest(c)

	res, err := p.PersonAppService.GetAllPerson(paging)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(200, res)
}

func (p *PersonController) GetAllPersonByQuery(c *gin.Context) {
	name, _ := c.GetQuery("name")
	a, _ := c.GetQuery("age")

	age, _ := strconv.Atoi(a)

	res, err := p.PersonAppService.GetAllPersonByQuery(name, age)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(200, res)
}
