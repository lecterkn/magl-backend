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
