package roles

import (
	"context"
	"encoding/json"
	"friction/loggers"
	"net/http"
	"net/http/httptest"
	"testing"
)

var roles = []Role{
	{Id: 1, Name: "a", Users: nil, Permissions: nil},
	{Id: 2, Name: "b", Users: nil, Permissions: nil},
}

type roleServiceMock struct {
}

func (r roleServiceMock) RetrieveAllRoles(ctx context.Context) ([]Role, error) {
	return roles, nil
}

func (r roleServiceMock) SaveRole(ctx context.Context, role Role) (Role, error) {
	return role, nil
}

func (r roleServiceMock) RetrieveRole(ctx context.Context, id int64) (Role, error) {
	return Role{}, nil
}

func TestGetRoles(t *testing.T) {
	logger := loggers.NewLogger()
	roleService := roleServiceMock{}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/roles", nil)

	controller := ControllerImpl{db: nil, logger: logger, roleService: roleService}
	controller.GetRoles(w, req)

	var got []Role
	decoder := json.NewDecoder(w.Body)
	decoder.Decode(&got)

	want := roles

	for i, r := range want {
		gottenId := got[i].Id
		if gottenId != r.Id {
			t.Errorf("got %q, want %q", gottenId, r.Id)
		}

		gottenName := got[i].Name
		if gottenName != r.Name {
			t.Errorf("got %q, want %q", gottenName, r.Name)
		}
	}
}
