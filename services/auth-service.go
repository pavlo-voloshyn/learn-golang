package services

type AuthService interface {
	Login(username string, password string) bool
}

type loginInformation struct {
	email    string
	password string
}

func NewAuthService() AuthService {
	return &loginInformation{
		email:    "test@gmail.com",
		password: "Password1!",
	}
}
func (info *loginInformation) Login(email string, password string) bool {
	return info.email == email && info.password == password
}
