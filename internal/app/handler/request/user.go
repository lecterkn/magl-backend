package request

type UserSigninRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserSignupRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserUpdatePermissionRequest struct {
	Permission int `json:"permission" validate:"required"`
}
