package output

import "time"

type UserLoginOutput struct {
	AccessToken  string
	RefreshToken string
}

type RefreshOutput struct {
	AccessToken string
	ExpiresIn   time.Time
}
