package service

import (
	"gvb_server/service/image_ser"
	"gvb_server/service/user_ser"
)

type ServiceGruop struct {
	ImageService image_ser.ImageService
	UserService  user_ser.UserService
}

var ServiceApp = new(ServiceGruop)
