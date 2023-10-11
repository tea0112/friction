package handlers

import (
	"database/sql"
	"friction/loggers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMatchPathSuccess(t *testing.T) {
	usersRg := "/api/users/([^/]+)"

	suite := []struct {
		path string
		want string
	}{
		{"/api/users/abc", "abc"},
		{"/api/users/1a", "1a"},
		{"/api/users/a1", "a1"},
		{"/api/users/11", "11"},
	}

	want := []string{"abc", "1a", "a1", "11"}

	for i, v := range suite {
		strSubmatchs := matchPath(usersRg, v.path)
		if strSubmatchs == nil {
			t.Fatal("path matches are nil")
		}
		if len(strSubmatchs) <= 1 {
			t.Fatal("path matches are not enough")
		}

		got := strSubmatchs[1]
		if got != want[i] {
			t.Errorf("got %q, want %q", got, want)
		}
	}
}

func TestMatchPathPagination(t *testing.T) {
	path := "/api/roles?limit=3&page=2"
	want := []string{"/api/roles?limit=3&page=2", "3", "2"}
	rg := `\/api\/roles\?limit=([0-9]+)&page=([0-9]+)`

	got := matchPath(rg, path)
	if len(got) != 3 {
		t.Fatal("failed")
	}

	for idx, v := range want {
		if v != got[idx] {
			t.Errorf("want %q, got %q", v, got[idx])
		}
	}
}

func TestSetupHandlers(t *testing.T) {
	logger := loggers.NewLogger()

	req := httptest.NewRequest(http.MethodGet, "/api/roles/5312abc", nil)
	w := httptest.NewRecorder()

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	var routes = []Route{
		{Method: http.MethodGet, PathRegex: "/api/roles/([^/]+)", Handler: h},
	}

	mux := SetupHandlers(&sql.DB{}, logger, routes)
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("want %q, got %q", http.StatusOK, w.Code)
	}
}
