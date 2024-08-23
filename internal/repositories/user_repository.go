package repositories

import (
	"errors"
	"fmt"
	"local-eyes/internal/models"
	"strings"
)

type UserRepository struct {
	FileRepository
}

func NewUserRepository(filepath string) *UserRepository {
	return &UserRepository{
		FileRepository: *NewFileRepository(filepath),
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

func (ur *UserRepository) Delete(userID int) error {
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
	return errors.New("user not found")
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

func (ur *UserRepository) UserDisplayTable() error {
	// Define column headers
	headers := []string{"Id", "Username", "Password", "Type", "NotificationChannel"}

	// Define column widths
	colWidths := make([]int, len(headers))
	for i, header := range headers {
		colWidths[i] = len(header)
	}
	users, err := ur.Load()
	if err != nil {
		return err
	}

	// Compute the maximum width needed for each column
	for _, userElement := range users {
		values := []string{
			fmt.Sprintf("%v", userElement.ID),
			userElement.Username,
			userElement.Type,
		}
		for i, value := range values {
			if len(value) > colWidths[i] {
				colWidths[i] = len(value)
			}
		}
	}

	// Print the headers
	for i, header := range headers {
		fmt.Printf("%-*s ", colWidths[i], header)
	}
	fmt.Println()

	// Print the separator line
	fmt.Println(strings.Repeat("-", sum(colWidths)+len(colWidths)-1))
	// Print the post details
	for _, userElement := range users {
		values := []string{
			fmt.Sprintf("%d", userElement.ID),
			userElement.Username,
			userElement.Type,
		}
		for i, value := range values {
			fmt.Printf("%-*s ", colWidths[i], value)
		}
		fmt.Println()
	}
	fmt.Println()
	return nil
}

//// sum calculates the sum of integers in a slice
//func sum(numbers []int) int {
//	total := 0
//	for _, num := range numbers {
//		total += num
//	}
//	return total
//}
