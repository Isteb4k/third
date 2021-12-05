package server

import (
	"auth/internal/db"
	"auth/transport"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Run() error
}

type server struct {
	router *gin.Engine
	users  db.Users
	amqper transport.Client
}

func New(users db.Users, amqper transport.Client) Server {
	router := gin.Default()

	s := server{
		router: router,
		users:  users,
		amqper: amqper,
	}

	router.GET("/try_amqp", s.pingAMQPClient)
	router.GET("/status", s.statusHandler)
	router.POST("/create_user", s.createUserHandler)
	router.DELETE("/delete_user/:id", s.deleteUserHandler)
	router.GET("/get_user/:id", s.getUserHandler)

	return &s
}

func (s *server) Run() error {
	return s.router.Run(":8081")
}
