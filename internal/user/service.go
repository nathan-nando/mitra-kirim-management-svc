package user

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo}
}

func (r *UserService) GetUserByID(id int64) error {
	return r.repo.FindById(id)
}
