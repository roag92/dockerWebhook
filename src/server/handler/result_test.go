package handler_test

import (
	"bytes"
	"dockerWebhook/src/server/handler"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestResultMissingId(t *testing.T) {
	_ = os.Setenv(handler.TagKey, "latest")

	invalidJson := `invalid json`
	req, err := http.NewRequest("POST", "/result", bytes.NewBuffer([]byte(invalidJson)))

	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	h := http.HandlerFunc(handler.NewResultHandler(NewRegistryMock()))

	h.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusNotFound {
		t.Errorf(
			"handler returned wrong status code: got %v expected %v",
			status,
			http.StatusNotFound,
		)
	}
}

func TestResultInvalidId(t *testing.T) {
	_ = os.Setenv(handler.TagKey, "latest")

	invalidJson := `invalid json`
	req, err := http.NewRequest("POST", "/result/a", bytes.NewBuffer([]byte(invalidJson)))

	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	h := http.HandlerFunc(handler.NewResultHandler(NewRegistryMock()))

	h.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf(
			"handler returned wrong status code: got %v expected %v",
			status,
			http.StatusInternalServerError,
		)
	}
}

func TestResultIdNotFound(t *testing.T) {
	_ = os.Setenv(handler.TagKey, "latest")

	invalidJson := `invalid json`
	req, err := http.NewRequest("POST", "/result/0", bytes.NewBuffer([]byte(invalidJson)))

	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	h := http.HandlerFunc(handler.NewResultHandler(NewRegistryMock()))

	h.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusNotFound {
		t.Errorf(
			"handler returned wrong status code: got %v expected %v",
			status,
			http.StatusNotFound,
		)
	}
}

func TestResultIdFound(t *testing.T) {
	_ = os.Setenv(handler.TagKey, "latest")

	invalidJson := `invalid json`
	req, err := http.NewRequest("POST", "/result/1584400381", bytes.NewBuffer([]byte(invalidJson)))

	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	h := http.HandlerFunc(handler.NewResultHandler(NewRegistryMock()))

	h.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned wrong status code: got %v expected %v",
			status,
			http.StatusOK,
		)
	}

	var body = res.Body.String()

	expected := fmt.Sprintf(`{"id": %d}}`, 1584400381)

	if body != expected {
		t.Errorf(
			"not contains `%s` into the body `%s`",
			expected,
			body,
		)
	}
}
