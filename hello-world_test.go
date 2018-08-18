package main

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
    router := setupRouter()

    recorder := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/ping", nil)
    router.ServeHTTP(recorder, req)

    assert.Equal(t, 200, recorder.Code)
    assert.Equal(t, "{\"message\":\"pong\"}", recorder.Body.String())
}
