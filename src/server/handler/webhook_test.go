package handler_test

import (
	"bytes"
	"dockerWebhook/src/server"
	"dockerWebhook/src/server/handler"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestWebhookHandlerInternalServerError(t *testing.T) {
	_ = os.Setenv(handler.TagKey, "latest")

	invalidJson := `invalid json`
	req, err := http.NewRequest("POST", "/webhook/", bytes.NewBuffer([]byte(invalidJson)))

	if err != nil {
		t.Fatal(err)
	}

	d := NewDockerMock()

	res := httptest.NewRecorder()
	h := http.HandlerFunc(handler.NewWebhookHandler(d, NewRegistryMock()))

	h.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf(
			"handler returned wrong status code: got %v expected %v",
			status,
			http.StatusInternalServerError,
		)
	}
}

func TestWebhookHandlerBadRequest(t *testing.T) {
	_ = os.Setenv(handler.TagKey, "latest")
	_ = os.Setenv(server.HostPortKey, "8000")
	_ = os.Setenv(server.HostUrlKey, "http://localhost")

	jsonPayload := `{"push_data": { "tag": "latest", "pushed_at": 1584400136}}`
	req, err := http.NewRequest("POST", "/webhook/", bytes.NewBuffer([]byte(jsonPayload)))

	if err != nil {
		t.Fatal(err)
	}

	d := NewDockerMock()

	res := httptest.NewRecorder()
	h := http.HandlerFunc(handler.NewWebhookHandler(d, NewRegistryMock()))

	h.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf(
			"handler returned wrong status code: got %v expected %v",
			status,
			http.StatusBadRequest,
		)
	}
}

func TestWebhookHandlerWithoutDockerStart(t *testing.T) {
	_ = os.Setenv(handler.TagKey, "latest")
	_ = os.Setenv(server.HostPortKey, "8000")
	_ = os.Setenv(server.HostUrlKey, "http://localhost")

	jsonPayload := `{"push_data": { "tag": "dev", "pushed_at": 1584400136}, "repository": { "repo_name": "foo-repo"}}`
	req, err := http.NewRequest("POST", "/webhook/", bytes.NewBuffer([]byte(jsonPayload)))

	if err != nil {
		t.Fatal(err)
	}

	d := NewDockerMock()

	res := httptest.NewRecorder()
	h := http.HandlerFunc(handler.NewWebhookHandler(d, NewRegistryMock()))

	h.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned wrong status code: got %v expected %v",
			status,
			http.StatusOK,
		)
	}

	var body = res.Body.String()

	response := server.Response{
		"success",
		"Tag not deployed",
		"CD by foo-repo",
		"http://localhost:8000/result/1584400136",
	}

	expected, err := json.Marshal(response)
	if err != nil {
		t.Error(err)
	}

	if body != string(expected) {
		t.Errorf(
			"not contains `%s` into the body `%s`",
			expected,
			body,
		)
	}
}

func TestWebhookHandlerWithDockerStart(t *testing.T) {
	_ = os.Setenv(handler.TagKey, "latest")
	_ = os.Setenv(server.HostPortKey, "8000")
	_ = os.Setenv(server.HostUrlKey, "http://localhost")

	jsonPayload := `{"push_data": { "tag": "latest", "pushed_at": 1584400136}, "repository": { "repo_name": "foo-repo"}}`
	req, err := http.NewRequest("POST", "/webhook/", bytes.NewBuffer([]byte(jsonPayload)))

	if err != nil {
		t.Fatal(err)
	}

	d := NewDockerMock()

	res := httptest.NewRecorder()
	h := http.HandlerFunc(handler.NewWebhookHandler(d, NewRegistryMock()))

	h.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned wrong status code: got %v expected %v",
			status,
			http.StatusOK,
		)
	}

	var body = res.Body.String()

	response := server.Response{
		"success",
		"Service deployed",
		"CD by foo-repo",
		"http://localhost:8000/result/1584400136",
	}

	expected, err := json.Marshal(response)
	if err != nil {
		t.Error(err)
	}

	if body != string(expected) {
		t.Errorf(
			"not contains `%s` into the body `%s`",
			expected,
			body,
		)
	}
}
