package handlers

import (
	"testing"
)

func TestMatchPathSuccess(t *testing.T) {
	usersRg := "/api/users/([^/]+)"

	usersPaths := []string{}
	usersPaths = append(usersPaths, "/api/users/abc")
	usersPaths = append(usersPaths, "/api/users/1a")
	usersPaths = append(usersPaths, "/api/users/a1")
	usersPaths = append(usersPaths, "/api/users/11")

	want := []string{"abc", "1a", "a1", "11"}

	for i, v := range usersPaths {
		strSubmatchs := matchPath(usersRg, v)
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
