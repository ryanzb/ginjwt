// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"ginjwt/conf"
	"ginjwt/handler"
	"ginjwt/db"
	"ginjwt/repo"
	"ginjwt/route"
	"ginjwt/service"

	"github.com/google/wire"
)

func initApp(cfg *conf.DB) *route.Route {
	panic(wire.Build(
		db.ProviderSet,
		repo.ProviderSet,
		service.ProviderSet,
		handler.ProviderSet,
		route.ProviderSet,
	))
}
