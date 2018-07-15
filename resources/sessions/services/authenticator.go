package sessions

import (
	model "github.com/defact/user/resources/sessions/models"
)

type Authenticator struct{}

func (a Authenticator) Authenticate(auth model.Authentication) (string, error) {
	return auth.Token, nil
}
