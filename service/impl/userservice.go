package impl

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
	"xiaoluokitchen/constant"
	"xiaoluokitchen/mapper"
	"xiaoluokitchen/model/po"
	"xiaoluokitchen/model/vo"
	"xiaoluokitchen/service"
)

type UserServiceImpl struct {
	userDao mapper.UserMapper
}

func NewUserService() service.UserService {
	return &UserServiceImpl{userDao: mapper.NewUserMapper()}
}

func (s *UserServiceImpl) Register(ctx context.Context, info *vo.UserInfo) error {
	userModel := s.toRegisterModel(info)
	if userModel != nil {
		return s.userDao.CreateUser(ctx, userModel)
	}
	return errors.New("空对象错误")
}

func (s *UserServiceImpl) toRegisterModel(info *vo.UserInfo) *po.User {
	if info == nil {
		return nil
	}
	uidGen := uuid.NewV4()
	uid := strings.ReplaceAll(uidGen.String(), "-", "")
	u := po.User{
		UserName: info.UserName,
		Salt:     uid,
		Birthday: info.BirthDay,
		Email:    info.Email,
		Phone:    info.Phone,
		Sex:      info.Sex,
		Address:  info.Address,
		Ctime:    int(time.Now().Unix()),
		Mtime:    int(time.Now().Unix()),
		Taste:    info.Taste,
	}
	u.Passw = s.encryptPassword(info.Password, uid)
	return &u
}

func (s *UserServiceImpl) encryptPassword(password, salt string) string {
	secretObj := md5.New()
	secretObj.Write([]byte(password))
	secretPassword := hex.EncodeToString(secretObj.Sum([]byte(salt)))
	return secretPassword
}

func (s *UserServiceImpl) Login(ctx context.Context, loginType po.LoginType, account string, password string) (*po.User, error) {
	cond := make(map[string]interface{})
	if loginType == constant.UserNameLogin {
		cond["user_name = ?"] = account
	} else if loginType == constant.PhoneLogin {
		cond["phone = ?"] = account
	} else if loginType == constant.EmailLogin {
		cond["email = ?"] = account
	} else {
		return nil, errors.New("不支持的登录类型")
	}
	users, err := s.userDao.GetUser(ctx, cond)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.New("用户未注册")
	}
	userModel := users[0]
	secretPassword := s.encryptPassword(password, userModel.Salt)
	if userModel.Passw != secretPassword {
		return nil, errors.New("密码错误")
	}
	return userModel, nil
}
