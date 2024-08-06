package repository

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (userRepository *UserRepository) GetInfoUser() string {
	return "minhbear"
}
