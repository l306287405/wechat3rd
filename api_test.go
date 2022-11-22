package wechat3rd

import (
	"log"
	"testing"
)

func TestName(t *testing.T) {
	token := "63_vKrfyQ39H0f0fFJWSLcsIaTCOZD5KvFBtGrPkl9QxtzkDW5g7Y_R6CvZ4zxkZCqDpC6wfjl6RX4vCYQqiK9WfGprVcJQPS17K-2h-NBB0j_Kr1eHjs9ETpqOtnRurWOWrLpzdVnrgDgN1LksSLJaAIDJCA"
	req := &GetWxaCodeUnLimitReq{
		Scene: "aaa",
	}
	s := Server{}
	limit := s.GetWxaCodeUnLimit(token, req)
	log.Print(limit)
}
