package iyzigo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

type (
	Options struct {
		ApiKey    string
		SecretKey string
		BaseUrl   string
	}
	client struct {
		Options
	}
	iyzicoLog struct {
		Request  interface{} `json:"request,omitempty"`
		Response interface{} `json:"response,omitempty"`
		Error    string      `json:"error,omitempty"`
	}
)

var (
	once    sync.Once
	options Options
	c       *client
)

func getClient() *client {
	once.Do(func() {
		c = &client{
			Options: options,
		}
	})
	return c
}

func IsProd() bool {
	return !strings.HasPrefix(getClient().ApiKey, "sandbox-")
}

func sendToIyzicoApi(method string, endpoint string, object interface{}) (output []byte, err error) {
	b, err := json.Marshal(object)
	if err != nil {
		logError(err)
		return nil, err
	}
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", getClient().BaseUrl, endpoint), bytes.NewReader(b))
	if err != nil {
		logError(err)
		return nil, err
	}
	req.Header = getClient().getHttpHeader(object)
	hc := http.Client{}
	hc.Transport = &http.Transport{DisableCompression: true}
	il := iyzicoLog{Request: maskSensitiveInfo(object)}
	res, err := hc.Do(req)
	if err != nil {
		logError(err)
		return nil, err
	}
	defer res.Body.Close()
	output, err = ioutil.ReadAll(res.Body)
	if err != nil {
		logError(err)
	}
	il.Response = string(output)
	logRequest(il)
	return
}

func doGet(endpoint string) ([]byte, error) {
	return sendToIyzicoApi("GET", endpoint, nil)
}

func doPost(endpoint string, object interface{}) ([]byte, error) {
	return sendToIyzicoApi("POST", endpoint, object)
}

func doPut(endpoint string, object interface{}) ([]byte, error) {
	return sendToIyzicoApi("PUT", endpoint, object)
}

func doDelete(endpoint string, object interface{}) ([]byte, error) {
	return sendToIyzicoApi("DELETE", endpoint, object)
}
