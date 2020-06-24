package controller

import (
	"xiaoluokitchen/service"
	"xiaoluokitchen/service/impl"
)

type ServiceGroup struct {
	userSrv service.UserService
}

var sg *ServiceGroup

func ServiceInit() {
	if sg == nil {
		sg = new(ServiceGroup)
	}
	sg.userSrv = impl.NewUserService()
}
