package delivery

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"enigmacamp.com/be-enigma-laundry/config"
	"enigmacamp.com/be-enigma-laundry/delivery/controller"
	"enigmacamp.com/be-enigma-laundry/manager"
)

type Server struct {
	uc     manager.UseCaseManager
	engine *gin.Engine
	host   string
}

func (s *Server) setupControllers() {
	rg := s.engine.Group("/api/v1")
	controller.NewBillController(s.uc.BillUseCase(), rg).Route()
}

func (s *Server) Run() {
	s.setupControllers()
	if err := s.engine.Run(s.host); err != nil {
		log.Fatal("masalah di server")
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	infra, err := manager.NewInfraManager(cfg)
	if err != nil {
		log.Fatal(err)
	}
	repo := manager.NewRepoManager(infra)
	uc := manager.NewUseCaseManager(repo)
	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)
	return &Server{
		uc:     uc,
		engine: engine,
		host:   host,
	}
}
