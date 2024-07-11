package exceptions

import (
	"fmt"
	"messenger_service/internal/shared/logger"
	"net/http"
	"strconv"
)

type HttpException struct{}

type IHttpException interface {
	Handle(message string, resp *http.Response, err error)
}

func (e *HttpException) Handle(message string, resp *http.Response, err error) {
	log := &logger.Logger{}

	toLog := fmt.Sprintf("Context: HTTP\nURL: %s\nSTATUS: %s STATUS CODE: %s\nError: %s", resp.Request.URL, resp.Status, strconv.Itoa(resp.StatusCode), err)

	log.Read("HTTP EXCEPTION", toLog)
}
