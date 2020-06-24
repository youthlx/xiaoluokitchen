package mapper

import (
	"context"
	"github.com/jinzhu/gorm"
	"xiaoluokitchen/config"
	"xiaoluokitchen/model/po"
)

type UserMapper interface {
	GetUser(ctx context.Context, where map[string]interface{}) ([]*po.User, error)
	CreateUser(ctx context.Context, u *po.User) error
}

type userMapper struct {
}

func NewUserMapper() UserMapper {
	return &userMapper{}
}

func (m *userMapper) GetUser(ctx context.Context, where map[string]interface{}) ([]*po.User, error) {
	var ret []*po.User
	db := config.Com.DB
	if len(where) == 0 {
		db = db.Table(po.UserTab.TableName()).Find(&ret)
		if db.Error != nil && !gorm.IsRecordNotFoundError(db.Error) {
			return nil, db.Error
		}
		return ret, nil
	}
	var condLines []string
	var valLines []interface{}
	for k, v := range where {
		condLines = append(condLines, k)
		valLines = append(valLines, v)
	}
	db = db.Table(po.UserTab.TableName()).Where(condLines, valLines...).Find(&ret)
	if db.Error != nil && !gorm.IsRecordNotFoundError(db.Error) {
		return nil, db.Error
	}
	return ret, nil
}

func (m *userMapper) CreateUser(ctx context.Context, u *po.User) error {
	statement := config.Com.DB.Table(po.UserTab.TableName()).Create(u)
	if statement.Error != nil {
		return statement.Error
	}
	return nil
}
