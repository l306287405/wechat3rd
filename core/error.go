package core

import (
	"strconv"
)

type Error struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (err *Error) Error() string {
	return "errcode: "+strconv.FormatInt(err.ErrCode,10)+", errmsg: "+err.ErrMsg
}

type H map[string]interface{}
