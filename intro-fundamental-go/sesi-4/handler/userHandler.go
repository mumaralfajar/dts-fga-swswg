package handler

type UserHandler struct {
	Name     string
	Email    string
	password string
}

var SecretKey = "rahasia"

func (userHandler *UserHandler) SetPassword(newPassword string) {
	userHandler.password = newPassword
}

func (userHandler UserHandler) GetPassword() string {
	return userHandler.password
}

func Greet() string {
	return "Halo Semuanya"
}
