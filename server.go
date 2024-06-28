package main

import (
	"database/sql"
	"fmt"
	"log"
	"todo-challange/config"
	"todo-challange/controller"
	"todo-challange/repository"
	"todo-challange/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// struct
type Server struct {
	uS      service.UserService
	tS      service.TaskService
	engine  *gin.Engine
	PortApp string
}

func (s *Server) RouteInit() {
	routeGroup := s.engine.Group("/api")
	controller.NewTaskController(s.tS, routeGroup).Route()
	controller.NewUserController(s.uS, routeGroup).Route()
}

func (s *Server) Start() {
	s.RouteInit()
	s.engine.Run(s.PortApp)
}

// constuctor
func NewServer() *Server {

	cf, err := config.NewConfig()

	urlConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cf.Host, cf.Port, cf.User, cf.Password, cf.DbName)

	db, err := sql.Open(cf.Driver, urlConnection)
	if err != nil {
		fmt.Println("Error koneksi database")
		log.Fatal(err)
	}

	portApp := cf.AppConfig.AppPort
	if portApp == "" {
		portApp = "8080"
	}

	userRepo := repository.NewUserRepository(db)
	taskRepo := repository.NewTaskRepository(db)

	userServ := service.NewUserService(userRepo)
	taskServe := service.NewTaskService(taskRepo, userServ)

	return &Server{
		uS:      userServ,
		tS:      taskServe,
		engine:  gin.Default(),
		PortApp: portApp,
	}
}
