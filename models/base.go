package models

import (
	"github.com/kataras/iris/v12"
	"time"
)

// CommonReturn 通用返回结构体
type CommonReturn struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(data interface{}) *CommonReturn {
	return SuccessWithMsg("ok", data)
}

func SuccessWithMsg(msg string, data interface{}) *CommonReturn {
	return &CommonReturn{
		Code: 200,
		Msg:  msg,
		Data: data,
	}
}

func Failure(code int, msg string) *CommonReturn {
	return &CommonReturn{
		Code: code,
		Msg:  msg,
		Data: iris.Map{},
	}
}

// CommonPageReturn 分页通用返回结构体
type CommonPageReturn struct {
	Page  int64       `json:"page"`
	Size  int64       `json:"size"`
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}

func CommonPageQuery2Return(query CommonPageQuery, total int64, data interface{}) CommonPageReturn {
	return CommonPageReturn{
		Page:  query.Page,
		Size:  query.Size,
		Total: total,
		Data:  data,
	}
}

// CommonPageQuery 分页查询通用结构体
type CommonPageQuery struct {
	Page int64 `json:"page"`
	Size int64 `json:"size"`
}

type IDBase struct {
	ID uint64 `json:"id"`
}

type Base struct {
	IDBase
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NameBase struct {
	Name string `json:"name"`
}
