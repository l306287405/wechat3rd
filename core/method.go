package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func PostJson(incompleteURL string, request interface{}, response interface{}) error {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(&request); err != nil {
		return err
	}
	httpResp, err := http.Post(incompleteURL,"application/json; charset=utf-8", &buf)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}
	return json.NewDecoder(httpResp.Body).Decode(response)
}
