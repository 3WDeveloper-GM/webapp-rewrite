package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg/ping"
	"github.com/3WDeveloper-GM/webapp-rewrite/internal/assert"
)

func TestPing(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	ping.Ping(rr, r)

	rs := rr.Result()

	assert.Equal(t, rs.StatusCode, http.StatusOK)

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, string(body), "OK")
}
