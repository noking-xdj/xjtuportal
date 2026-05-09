package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"xjtuportal/component/basic"
)

func TestSendRequestHandlesErrorStatusWithoutPanic(t *testing.T) {
	server := httptest.NewServer(http.NotFoundHandler())
	defer server.Close()

	requestHelper := &RequestHelper{
		loggerHelper: basic.LoggerTemp,
		requestSettings: &basic.ProgramRequestSettings{
			Header: map[string]string{},
		},
	}
	requestHelper.requestSettings.Connect.Timeout = 1

	_, _, statusCode, err := requestHelper.SendRequest(
		server.URL,
		http.MethodGet,
		nil,
		nil,
		nil,
	)

	if err == nil {
		t.Fatal("expected error for 404 response")
	}
	if statusCode != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, statusCode)
	}
}
