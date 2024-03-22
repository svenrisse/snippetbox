package main

import (
	"net/http"
	"testing"

	"github.com/svenrisse/snippetbox/internal/assert"
)

func TestHealthcheck(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/healthcheck")

	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, "OK")
}
