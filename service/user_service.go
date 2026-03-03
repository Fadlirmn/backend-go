package service


import(
	"backend-api-belajar/model"
	"backend-api-belajar/repository"
)

type UserService struct{
	repo repository.UserRepository
}
func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetUsers()[]model.User  {
	return  s.repo.FindAll()
}

func (s *UserService) CreateUser(user model.User)  {
	s.repo.Save(user)
}

func (s *UserService) UpdateUser(id string,user model.User)  error{
	return s.repo.Update(id, user)
}
func (s *UserService) DeleteUser(id string) error {
	return s.repo.Delete(id)
}