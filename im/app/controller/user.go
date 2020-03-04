package controller

import (
	"im/app/model"
	"im/app/service"
	"im/app/util"
	"net/http"
)

// 注册用户
func UserRegister(w http.ResponseWriter, r *http.Request) {
	var user model.User
	util.Bind(r, &user)
	//user, err := service.UserService.UserRegister(user.Mobile, user.Passwd, user.Nickname, user.Avatar, user.Sex)
	user, err := service.UserService.UserRegister(user.Mobile, user.Passwd, user.Nickname, user.Avatar, user.Sex)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOk(w, user, "")
	}
}

// 用户登录
func UserLogin(w http.ResponseWriter, r *http.Request) {
	// 解析表单数据
	r.ParseForm()
	mobile := r.PostForm.Get("mobile")
	plainpwd := r.PostForm.Get("passwd")
	// 校验参数
	if len(mobile) == 0 || len(plainpwd) == 0 {
		util.RespFail(w, "用户名或者密不正确")
	}
	loginUser, err := service.UserService.Login(mobile, plainpwd)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOk(w, loginUser, "")
	}
}
