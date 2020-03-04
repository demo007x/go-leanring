package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data, omitempty"`
}

type H struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data, omitempty"`
	Rows  interface{} `json:"rows, omitempty"`
	Total interface{} `json:"total, omitempty"`
}

// 失败的返回结果
func RespFail(w http.ResponseWriter, msg string) {
	Resp(w, -1, nil, msg)
}

// 返回成功
func RespOk(w http.ResponseWriter, data interface{}, msg string) {
	Resp(w, 0, data, msg)
}

// 返回列表数据
func RespList(w http.ResponseWriter, code int, data interface{}, total interface{}) {
	w.Header().Set("Content-Type", "application/json")
	// 设置200状态
	w.WriteHeader(http.StatusOK)
	h := H{
		Code:  code,
		Msg:   "",
		Data:  data,
		Rows:  nil,
		Total: total,
	}
	// 将结构体转换为json数据
	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}
	w.Write(ret)
}
func Resp(w http.ResponseWriter, code int, data interface{}, msg string) {
	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	rep := ResponseData{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	// 转换json数据
	ret, err := json.Marshal(rep)
	if err != nil {
		log.Println(err.Error())
	}
	// 返回json ok
	w.Write(ret)
}
