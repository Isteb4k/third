package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"polls/internal/entities"
	"strconv"
)

type createPayload struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (s *server) createPollHandler(c *gin.Context) {
	var payload createPayload

	err := c.BindJSON(&payload)
	if err != nil {
		panic(err)
	}

	poll, err := s.polls.Create(context.TODO(), entities.Poll{
		Title:       payload.Title,
		Description: payload.Description,
	})
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, poll)
}

func (s *server) deletePollHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	err = s.polls.DeleteByID(context.TODO(), id)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, nil)
}

func (s *server) getPollHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	poll, err := s.polls.GetByID(context.TODO(), id)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, poll)
}

func (s *server) statusHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
