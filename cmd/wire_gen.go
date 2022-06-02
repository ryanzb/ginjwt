// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"ginjwt/conf"
	"ginjwt/db"
	"ginjwt/handler"
	"ginjwt/repo"
	"ginjwt/route"
	"ginjwt/service"
)

// Injectors from wire.go:

func initApp(cfg *conf.DB) *route.Route {
	jwtService := service.NewJWTService()
	gormDB := db.NewGormDB(cfg)
	userRepo := repo.NewUserRepo(gormDB)
	authService := service.NewAuthService(userRepo)
	userService := service.NewUserService(userRepo)
	authHandler := handler.NewAuthHandler(authService, jwtService, userService)
	userHandler := handler.NewUserHandler(userService, jwtService)
	routeRoute := route.NewRoute(jwtService, authHandler, userHandler)
	return routeRoute
}
