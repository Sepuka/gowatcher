package transports

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sepuka/gowatcher/definition/logger"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strings"
)

type loggedRequestSender struct {
	httpClient *http.Client
	log        logger.Logger
	answer     interface{}
}

func (obj loggedRequestSender) sendRequest(req *http.Request) error {
	var response *http.Response
	var err error
	var requestBodyBuffer []byte

	requestBodyBuffer, _ = ioutil.ReadAll(req.Body)
	req.Body = ioutil.NopCloser(bytes.NewBuffer(requestBodyBuffer))

	response, err = obj.httpClient.Do(req)
	if err != nil {
		obj.
			log.
			With(
				zap.String("req", formatRequest(*req)),
				zap.Error(err),
			).
			Error("Request failed")

		return err
	}

	req.Body = ioutil.NopCloser(bytes.NewBuffer(requestBodyBuffer))
	err = parseResponse(*response, obj.answer)
	if err != nil {
		obj.log.Error("Cannot parse body")

		return err
	}

	obj.
		log.
		With(
			zap.String("request",  formatRequest(*req)),
			zap.String("response", formatResponse(*response)),
			zap.Any("answer",   obj.answer),
			zap.String("status",   response.Status),
			zap.String("host", req.Host),
		).
		Debug("Outcoming request")

	defer response.Body.Close()

	return nil
}

func formatRequest(r http.Request) string {
	var request []string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	request = append(request, fmt.Sprintf("Host: %v", r.Host))

	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}

	return strings.Join(request, "\n")
}

func formatResponse(r http.Response) string {
	var request []string

	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	request = append(request, "\n")
	request = append(request, "Body:")
	body, _ := ioutil.ReadAll(r.Body)
	request = append(request, string(body))

	return strings.Join(request, "\n")
}

func parseResponse(r http.Response, response interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json; charset=utf-8", "application/json":
		err := json.Unmarshal(body, response)
		if err != nil {
			return err
		}
		return nil
	default:
		response = &body
		return nil
	}
}
