package main

import (
	"github.com/Bangseungjae/social/internal/auth"
	"github.com/Bangseungjae/social/internal/store"
	"github.com/Bangseungjae/social/internal/store/cache"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newTestApplication(t *testing.T, cfg config) *application {
	t.Helper()

	logger := zap.NewNop().Sugar()
	mockStore := store.NewMockStorage()
	mockCache := cache.NewMockStore()
	testAuth := &auth.TestAuthenticator{}

	return &application{
		logger:        logger,
		store:         mockStore,
		cacheStorage:  mockCache,
		authenticator: testAuth,
		config:        cfg,
	}
}

func executeRequest(req *http.Request, mux http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d", expected, actual)
	}
}
