package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func PostJson(incompleteURL string, request interface{}, response interface{}) error {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(request); err != nil {
		return err
	}
	httpResp, err := http.Post(incompleteURL, "application/json; charset=utf-8", &buf)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return errors.New("http.Status:" + httpResp.Status)
	}
	return json.NewDecoder(httpResp.Body).Decode(response)
}

func PostJsonReturnBuffer(incompleteURL string, request interface{}, response interface{}) error {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(request); err != nil {
		return err
	}
	httpResp, err := http.Post(incompleteURL, "application/json; charset=utf-8", &buf)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return errors.New("http.Status:" + httpResp.Status)
	}
	if err = json.NewDecoder(httpResp.Body).Decode(response); err != nil {
		buffer, err := io.ReadAll(httpResp.Body)
		if err != nil {
			return err
		}
		m := make(map[string]interface{})
		m["errcode"] = 0
		m["errmsg"] = ""
		m["buffer"] = buffer
		jsonBytes, err := json.Marshal(m)
		if err != nil {
			return err
		}
		json.Unmarshal(jsonBytes, &response)
		return nil
	}
	return nil
}

func GetRequest(u string, request url.Values, response interface{}) error {
	if !strings.HasSuffix(u, "?") {
		u += "?"
	}
	httpResp, err := http.Get(u + request.Encode())
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return errors.New("http.Status:" + httpResp.Status)
	}
	return json.NewDecoder(httpResp.Body).Decode(response)
}

func AuthTokenUrlValues(token string) url.Values {
	v := make(url.Values)
	v.Set("access_token", token)
	return v
}

func PostFile(url string, filePath string, fileParameterName string, response interface{}) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile(fileParameterName, filePath)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	//打开文件句柄操作
	fh, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	httpResp, err := http.Post(url, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode != http.StatusOK {
		return errors.New("http.Status:" + httpResp.Status)
	}
	return json.NewDecoder(httpResp.Body).Decode(response)
}

func PostFileDirectly(url string, file io.Reader, fileName, fileParameterName string, response interface{}) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile(fileParameterName, fileName)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}
	//iocopy
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return err
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	httpResp, err := http.Post(url, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode != http.StatusOK {
		return errors.New("http.Status:" + httpResp.Status)
	}
	return json.NewDecoder(httpResp.Body).Decode(response)
}
