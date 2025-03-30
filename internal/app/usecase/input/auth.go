package input

type UserCreateInput struct {
	Username string
	Email    string
	Password string
}

type UserLoginInput struct {
	Username string
	Password string
}

type RefreshInput struct {
	RefreshToken string
}
