package service

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewAuthService, NewJWTService, NewUserService)
