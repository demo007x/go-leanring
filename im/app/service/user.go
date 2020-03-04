package service

import (
	"errors"
	"fmt"
	"im/app/model"
	"im/app/util"
	"math/rand"
	"time"
)

type UserService struct {
}

// 用户注册
func (us *UserService) UserRegister(mobile, plainPwd, nickname, avatar, sex string) (user model.User, err error) {
	registerUser := model.User{}
	_, err = model.DbEngine.Where("mobile=?", mobile).Get(&registerUser)
	if err != nil {
		return registerUser, err
	}

	if registerUser.Id > 0 {
		return registerUser, errors.New("该手机号码已注册")
	}

	registerUser.Mobile = mobile
	registerUser.Avatar = avatar
	registerUser.Nickname = nickname
	registerUser.Sex = sex
	registerUser.Salt = fmt.Sprintf("%06d", rand.Int31n(100000))
	registerUser.Passwd = util.MakePassword(plainPwd, registerUser.Salt)
	registerUser.Createat = time.Now()
	_, err = model.DbEngine.InsertOne(&registerUser)

	return registerUser, err
}

// 用户登录
func (us *UserService) Login(mobile, plainPwd string) (user model.User, err error) {
	loginUser := model.User{}
	model.DbEngine.Where("mobile = ?", mobile).Get(&loginUser)
	if loginUser.Id == 0 {
		return loginUser, errors.New("用户不存在")
	}

	// 判断密码是否正确
	if !util.ValidatePasswd(plainPwd, loginUser.Salt, loginUser.Passwd) {
		return loginUser, errors.New("密码不正确")
	}
	// 刷新用户登录的token
	token := util.GenRandomStr(32)
	loginUser.Token = token
	model.DbEngine.Id(loginUser.Id).Cols("token").Update(&loginUser)

	return loginUser, nil
}
