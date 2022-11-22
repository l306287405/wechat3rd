package wechat3rd

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	token := "63_vV-1ucItNdryGAAVjtt9uAi-q3bSXXqwQvudlG6D-Iq2H5pCazfUOHy0xqUJJ9Y8NiuwskHkdivpfWiY86wmsK1eAui_Y4tM"
	req := &GetWxaCodeUnLimitReq{
		Scene: "aaa",
		Page:  "pages/xxxx",
	}
	s := Server{}
	resp := s.GetWxaCodeUnLimit(token, req)
	fileName := "./test.jpeg"
	err := ioutil.WriteFile(fileName, resp.Buffer, 0666)
	if err != nil {
		fmt.Println("write fail")
	}
	log.Print(resp)
}
