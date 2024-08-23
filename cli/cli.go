package cli

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/manifoldco/promptui"
	"local-eyes/UserCreation"
	"local-eyes/constants"
	"local-eyes/internal/controllers"
	"local-eyes/internal/repositories"
	"local-eyes/utils"
)

func StartCLI(userRepo *repositories.UserRepository, postRepo *repositories.PostRepository, notificationCtrl *controllers.NotificationController) {
	for {
		fmt.Println(constants.Magenta + "\n=====================================================")
		fmt.Println("Welcome to Local Eyes!")
		fmt.Println("=====================================================" + constants.Reset)
		fmt.Println(constants.Blue + "1. Sign Up")
		fmt.Println("2. Log In")
		fmt.Println("3. Exit" + constants.Reset)

		choice := getChoice()
		switch choice {
		case 1:
			signUp(userRepo)
		case 2:
			login(userRepo, postRepo, notificationCtrl)
		case 3:
			return
		default:
			fmt.Println(constants.Red + "Invalid choice, please try again." + constants.Reset)
		}
	}
}

func signUp(userRepo *repositories.UserRepository) {
	username := utils.PromptInput("Enter username: ")
	password := promptPassword("Enter password: ")
	userType := utils.PromptInput("Enter user type (newbie/resident): ")
	userExist, err := userRepo.UserNameExists(username)
	if err != nil {
		fmt.Println(err)
	}
	if userExist {
		fmt.Println("User already exist with this username")
		return
	}
	if !utils.ValidateUsername(username) || !utils.ValidatePassword(password) || !utils.ValidateUserType(userType) {
		fmt.Println("Invalid input. Please try again.")
		return
	}

	hashedPassword := hashPassword(password)
	user := userCreate.CreateUser(utils.GenerateUserID(), username, hashedPassword, userType)
	if err := userRepo.Save(user); err != nil {
		fmt.Println("Error signing up:", err)
	} else {
		fmt.Println("User signed up successfully.")
	}
}

func login(userRepo *repositories.UserRepository, postRepo *repositories.PostRepository, notificationCtrl *controllers.NotificationController) {
	username := utils.PromptInput("Enter username: ")
	password := promptPassword("Enter password: ")

	hashedPassword := hashPassword(password)
	user, err := userRepo.FindByUsernameAndPassword(username, hashedPassword)
	if err != nil {
		fmt.Println("Login failed:", err)
		return
	}
	if user != nil {
		switch user.Type {
		case "admin":
			adminCtrl := controllers.NewAdminController(user, userRepo, postRepo)
			adminCtrl.HandleAdminActions()
		case "newbie":
			newbieCtrl := controllers.NewNewbieController(user, postRepo, notificationCtrl)
			newbieCtrl.HandleNewbieActions()
		case "resident":
			residentCtrl := controllers.NewResidentController(user, postRepo, notificationCtrl)
			residentCtrl.HandleResidentActions()
		default:
			fmt.Println("Unknown user type.")
		}
	} else {
		fmt.Println("No user")
	}

}

func getChoice() int {
	fmt.Print("Enter choice: ")
	var choice int
	fmt.Scanln(&choice)
	return choice
}

func promptPassword(prompt string) string {
	prompt1 := promptui.Prompt{
		Label:     prompt,
		Mask:      '*',
		IsConfirm: false,
	}
	result, err := prompt1.Run()
	if err != nil {
		fmt.Println("Prompt failed:", err)
		return ""
	}
	return result
}

func hashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}
