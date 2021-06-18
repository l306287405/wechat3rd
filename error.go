package wechat3rd

import (
	"log"
	"net/http"
	"os"
)

var (
	Success = []byte("success")
	Fail    = []byte("Fail")
)

type WechatErrorer interface {
	ServeError(w http.ResponseWriter, r *http.Request, err error)
}

var DefaultErrorHandler WechatErrorer = ErrorHandlerFunc(defaultErrorHandlerFunc)

type ErrorHandlerFunc func(http.ResponseWriter, *http.Request, error)

func (fn ErrorHandlerFunc) ServeError(w http.ResponseWriter, r *http.Request, err error) {
	fn(w, r, err)
}

var errorLogger = log.New(os.Stderr, "[WECHAT_ERROR] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)

func defaultErrorHandlerFunc(w http.ResponseWriter, r *http.Request, err error) {
	errorLogger.Output(3, err.Error())
}
