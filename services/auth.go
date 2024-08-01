package services

type AuthService struct{}

type User struct {
	UserId     string
	Email      string
	IsVerified bool
}

func (authService AuthService) GetBackendToken() string {
	return "backend_token"
}

func (authService AuthService) ValidateBackendToken(token string) bool {
	return true
}

func (authService AuthService) GetAndValidateUser(userId string) (User, error) {
	user := authService.ReadUser(userId)
	err := authService.ValidateUser(user)

	return user, err
}

func (authService AuthService) ValidateUser(user User) error {
	return nil
}

func (authService AuthService) ReadUser(userId string) User {
	return User{
		UserId:     "UserId",
		Email:      "mock@test.com",
		IsVerified: true,
	}
}
