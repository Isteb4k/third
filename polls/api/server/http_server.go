package server

import (
	"github.com/gin-gonic/gin"
	"polls/internal/db"
)

type Server interface {
	Run() error
}

type server struct {
	router *gin.Engine
	polls  db.Polls
}

func New(polls db.Polls) Server {
	router := gin.Default()

	s := server{
		router: router,
		polls:  polls,
	}

	router.GET("/status", s.statusHandler)
	router.POST("/create_poll", s.createPollHandler)
	router.DELETE("/delete_poll/:id", s.deletePollHandler)
	router.GET("/get_poll/:id", s.getPollHandler)

	return &s
}

func (s *server) Run() error {
	return s.router.Run(":8080")
}
