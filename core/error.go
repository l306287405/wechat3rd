package core

const UNKNOWN_ERROR = -99

type Error struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg,omitempty"`
}

func (s *Error) Success() bool {
	if s.ErrCode == 0 {
		return true
	}
	return false
}

func (s *Error) Err(err error) {
	if err != nil {
		s.ErrMsg = err.Error()
		s.ErrCode = UNKNOWN_ERROR
	}
}

func NewUnknownError(errMsg string) *Error {
	return &Error{
		ErrCode: UNKNOWN_ERROR,
		ErrMsg:  errMsg,
	}
}

type H map[string]interface{}
