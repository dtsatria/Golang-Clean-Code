package delivery

import (

	// "final-project/delivery/middleware"

	"final-project-booking-room/config"
	"final-project-booking-room/delivery/controller"
	"final-project-booking-room/delivery/middleware"
	"final-project-booking-room/manager"
	"final-project-booking-room/usecase"
	"final-project-booking-room/utils/common"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	uc         manager.UseCaseManager
	auth       usecase.AuthUseCase
	engine     *gin.Engine
	host       string
	logService common.MyLogger

	jwtService common.JwtToken
}

func (s *Server) setupControllers() {
	s.engine.Use(middleware.NewLogMiddleware(s.logService).LogRequest())
	authMiddlerware := middleware.NewAuthMiddleware(s.jwtService)
	rg := s.engine.Group("/api/v1")
	controller.NewUserController(s.uc.UserUseCase(), rg, authMiddlerware).Route()
	controller.NewBookingController(s.uc.BookingUsecase(), rg, authMiddlerware).Route()
	controller.NewAuthController(s.auth, rg, s.jwtService).Route()
	controller.NewRoomController(s.uc.RoomUsecase(), rg, authMiddlerware).Route()
}

func (s *Server) Run() {
	s.setupControllers()
	if err := s.engine.Run(s.host); err != nil {
		log.Fatal("server can't run")
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
	uc := manager.NewUseCaseManager(repo, common.NewEmailService(cfg))
	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)
	logService := common.NewMyLogger(cfg.LogFileConfig)
	jwtService := common.NewJwtToken(cfg.TokenConfig)
	return &Server{
		uc:         uc,
		engine:     engine,
		host:       host,
		logService: logService,
		auth: usecase.NewAuthUseCase(uc.UserUseCase(),
			jwtService),
		jwtService: jwtService,
	}
}
