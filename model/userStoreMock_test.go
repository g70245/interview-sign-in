package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetUser(t *testing.T) {
	tests := []struct {
		name     string
		username string
		want     User
	}{{
		name:     "TestGetUserSuccess",
		username: "test",
		want:     Users[0],
	}, {
		name:     "TestGetUserFailure",
		username: "test123",
		want:     User{},
	}}

	userStore := UserStoreMock{}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := userStore.GetUser(test.username); !cmp.Equal(got, test.want) {
				t.Errorf("GetUser() = %v, want %v", got, test.want)
			}
		})
	}
}
