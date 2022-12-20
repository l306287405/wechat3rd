package util

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Sign 微信公众号 url 签名.
func Sign(datas ...string) (signature string) {
	return MsgSign(datas...)
}

// MsgSign 微信公众号/企业号 消息体签名.
func MsgSign(datas ...string) (signature string) {
	sort.Strings(datas)
	h := sha1.New()
	_, _ = io.WriteString(h, strings.Join(datas, ""))
	return fmt.Sprintf("%x", h.Sum(nil))
}
