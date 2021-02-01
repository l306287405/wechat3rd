package core

type Error struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (s *Error) Success() bool {
	if s.ErrCode==0{
		return true
	}
	return false
}

type H map[string]interface{}
