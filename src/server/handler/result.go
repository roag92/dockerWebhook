package handler

import (
	"dockerWebhook/src/registry"
	"net/http"
	"regexp"
	"strconv"
)

func NewResultHandler(reg registry.Registry) func(http.ResponseWriter, *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		var validPath = regexp.MustCompile("^/(result)/([a-zA-Z0-9]+)$")

		m := validPath.FindStringSubmatch(r.URL.Path)

		if m == nil {
			http.NotFound(w, r)
			return
		}

		id , err := strconv.Atoi(m[2])
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		var jsonResponse string

		jsonResponse, err = reg.Read(int64(id))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(jsonResponse))
	}
}
