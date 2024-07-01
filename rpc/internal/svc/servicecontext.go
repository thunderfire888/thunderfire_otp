package svc

import (
	"github.com/thunderfire888/thunderfire_otp/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,
	}
}
