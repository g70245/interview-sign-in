package service

import (
	"app/helper/status"
	"app/model"
	"errors"
	"testing"
)

func TestSignin(t *testing.T) {
	type Check func(error, error) bool

	checkSuccess := func(got, want error) bool {
		return got == nil && want == nil
	}
	checkErrMsg := func(got, want error) bool {
		if got == nil {
			return false
		}
		return got.Error() == want.Error()
	}

	tests := []struct {
		name     string
		username string
		password string
		want     error
		check    Check
	}{{
		name:     "TestSigninSuccess",
		username: model.Username,
		password: model.Password,
		want:     nil,
		check:    checkSuccess,
	}, {
		name:     "TestSigninPasswordIncorrect",
		username: model.Username,
		password: "12345678",
		want:     errors.New(status.UserPasswordIncorrect),
		check:    checkErrMsg,
	}, {
		name:     "TestSigninUserNotFound",
		username: "test123",
		password: "12345678",
		want:     errors.New(status.UserNotFound),
		check:    checkErrMsg,
	}}

	var userStore model.UserStore = model.UserStoreMock{}
	userManager := UserManager{UserStore: &userStore}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := userManager.Signin(test.username, test.password)

			if !test.check(err, test.want) {
				t.Errorf("Signin() = %v, want %v", err, test.want)
			}
		})
	}
}
