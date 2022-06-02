package main

import (
	"ginjwt/conf"
	"ginjwt/controller"
	"ginjwt/db"
	"ginjwt/repo"
	"ginjwt/route"
	"ginjwt/service"

	"github.com/gin-gonic/gin"
	"github.com/ryanzb/zlog"
)

var (
	log = zlog.New()
)

func main() {
	cfg, err := conf.Load("./config.yaml")
	if err != nil {
		log.Panicf("load config failed: %v", err)
	}
	log.Infof("config: %+v", cfg)

	db, err := db.Init(cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		log.Panicf("init db failed: %v", err)
	}

	userRepo := repo.NewUserRepo(db)
	authService := service.NewAuthService(userRepo)
	jwtService := service.NewJWTService()
	userService := service.NewUserService(userRepo)

	r := route.New(
		controller.NewAuthController(authService, jwtService, userService), 
		controller.NewUserController(userService, jwtService),
	)

	app := gin.Default()
	r.Register(app)

	if err := app.Run(cfg.Address); err != nil {
		log.Fatalf("run server at %s failed: %v", cfg.Address, err)
	}
}