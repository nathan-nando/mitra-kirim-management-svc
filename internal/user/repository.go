package user

type UserRepository interface {
	Save(user User) error
	FindById(id int64) (*User, error)
}

type userRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

func (u userRepositoryImpl) Save(user User) error {
	return nil
}

func (u userRepositoryImpl) FindById(id int64) (*User, error) {
	return &User{}, nil
}
