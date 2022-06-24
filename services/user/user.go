package user

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"

	"oasis/ready/biz/dal/mysql/user_dal"
	"oasis/ready/biz/model/handler_model"
	"oasis/ready/biz/model/user_model"
	"oasis/ready/global"
	"oasis/ready/utils"
)

type userService struct {
}

var Service = new(userService)

func (s *userService) Register(userLoginInfo *user_model.UserLoginInfo, userBaseInfo *user_model.UserBaseInfo) (error, *user_model.UserInfo) {
	rowEffect, userId, err := user_dal.CreateUserLoginInfo(userLoginInfo)
	if err != nil {
		return err, nil
	}
	if rowEffect == 0 || userId == 0 {
		return errors.New(fmt.Sprintf("创建用户失败, rowEffect = %v, userLoginId = %v", rowEffect, userId)), nil
	}

	userInfo := &user_model.UserInfo{
		UserBaseInfo: *userBaseInfo,
		LoginInfoId:  userId,
	}
	rowEffect, err = user_dal.CreateUserInfo(userInfo)
	if err != nil {
		return err, nil
	}
	if rowEffect == 0 {
		return errors.New(fmt.Sprintf("创建用户失败, rowEffect = %v", rowEffect)), nil
	}

	return nil, userInfo
}

func (s *userService) Login(request *handler_model.LoginRequest) (*user_model.UserInfo, error) {
	user, err := user_dal.GetUserLoginInfoByEmail(request.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.Wrap(errors.New("record not found"), "查找用户记录失败")
	}

	if !utils.BcryptMakeCheck([]byte(request.Password), user.Password) {
		return nil, errors.Wrap(errors.New("password [BcryptMakeCheck] fail"), "密码错误")
	}

	return user_dal.GetUserInfo(user.ID.ID)
}

// GetUserInfo 获取用户信息
func (s *userService) GetUserInfo(id string) (err error, user *user_model.UserInfo) {
	intId, err := strconv.Atoi(id)
	err = global.App.DB.Where("id = ?", intId).First(user).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return
}
