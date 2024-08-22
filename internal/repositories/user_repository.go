package repositories

import (
	"local-eyes/internal/models"
	"local-eyes/utils"
)

type UserRepository struct {
	FileRepository
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		FileRepository: *NewFileRepository(utils.UserFile),
	}
}

func (ur *UserRepository) Save(user *models.User) error {
	users, err := ur.Load()
	if err != nil {
		return err
	}
	users = append(users, user)
	return ur.SaveAll(users)
}

func (ur *UserRepository) Load() ([]*models.User, error) {
	var users []*models.User
	err := ur.FileRepository.Load(&users)
	for i, _ := range users {
		users[i].NotificationCh = make(chan models.Notification, 10)
	}
	return users, err
}

func (ur *UserRepository) Delete(userID string) error {
	users, err := ur.Load()
	if err != nil {
		return err
	}
	for i, user := range users {
		if user.ID == userID {
			users = append(users[:i], users[i+1:]...)
			return ur.SaveAll(users)
		}
	}
	return nil
}

func (ur *UserRepository) FindByUsernameAndPassword(username, hashedPassword string) (*models.User, error) {
	users, err := ur.Load()
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.Username == username && user.Password == hashedPassword {
			return user, nil
		}
	}
	return nil, nil
}

func (ur *UserRepository) FindAllNewbies() ([]*models.User, error) {
	var newbies []*models.User
	users, err := ur.Load()
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.Type == "newbie" {
			newbies = append(newbies, user)
		}
	}
	return newbies, nil
}

func (ur *UserRepository) SaveAll(users []*models.User) error {
	return ur.FileRepository.Save(users)
}

func (ur *UserRepository) UserNameExists(name string) (bool, error) {
	users, err := ur.Load()
	if err != nil {
		return false, err
	}
	for _, user := range users {
		if user.Username == name {
			return true, nil
		}
	}
	return false, nil
}
