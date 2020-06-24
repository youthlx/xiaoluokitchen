package service

import (
	"context"
	"xiaoluokitchen/model/po"
	"xiaoluokitchen/model/vo"
)

type UserService interface {
	Register(ctx context.Context, info *vo.UserInfo) error
	Login(ctx context.Context, loginType po.LoginType, account string, passw string) (*po.User, error)
}
