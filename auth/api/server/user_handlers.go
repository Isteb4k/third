package server

import (
	"auth/internal/entities"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type createPayload struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

func (s *server) createUserHandler(c *gin.Context) {
	var payload createPayload

	err := c.BindJSON(&payload)
	if err != nil {
		panic(err)
	}

	user, err := s.users.Create(context.TODO(), entities.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	})
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, user)
}

func (s *server) deleteUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	err = s.users.DeleteByID(context.TODO(), id)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, nil)
}

func (s *server) getUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	user, err := s.users.GetByID(context.TODO(), id)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, user)
}

func (s *server) statusHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (s *server) pingAMQPClient(c *gin.Context) {
	s.amqper.Publish()

	c.JSON(http.StatusOK, gin.H{"amqp": "message published"})
}
