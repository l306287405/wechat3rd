package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

func PostJson(incompleteURL string, request interface{}, response interface{}) error {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(request); err != nil {
		return err
	}
	httpResp, err := http.Post(incompleteURL,"application/json; charset=utf-8", &buf)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return errors.New("http.Status:"+httpResp.Status)
	}
	return json.NewDecoder(httpResp.Body).Decode(response)
}

func GetRequest(u string, request url.Values, response interface{}) error {
	if !strings.HasSuffix(u,"?"){
		u+="?"
	}
	httpResp, err := http.Get(u+request.Encode())
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return errors.New("http.Status:"+httpResp.Status)
	}
	return json.NewDecoder(httpResp.Body).Decode(response)
}
