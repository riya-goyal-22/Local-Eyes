package cli

import (
	"fmt"
	"local-eyes/internal/controllers"
)

func HandleAdminActions(adminCtrl *controllers.AdminController) {
	fmt.Println("Admin actions:")
	fmt.Println("1. Delete User")
	fmt.Println("2. Delete Post")
	fmt.Println("3. Exit")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		adminCtrl.DeleteUser()
	case 2:
		adminCtrl.DeletePost()
	case 3:
		return
	default:
		fmt.Println("Invalid choice, please try again.")
	}
}

func HandleNewbieActions(newbieCtrl *controllers.NewbieController) {
	fmt.Println("Newbie actions:")
	fmt.Println("1. View Posts")
	fmt.Println("2. Like Post")
	fmt.Println("3. Exit")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		newbieCtrl.ViewPosts()
	case 2:
		newbieCtrl.LikePost()
	case 3:
		return
	default:
		fmt.Println("Invalid choice, please try again.")
	}
}

func HandleResidentActions(residentCtrl *controllers.ResidentController) {
	fmt.Println("Resident actions:")
	fmt.Println("1. Create Post")
	fmt.Println("2. Update Post")
	fmt.Println("3. Delete Post")
	fmt.Println("4. Exit")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		residentCtrl.CreatePost()
	case 2:
		residentCtrl.UpdatePost()
	case 3:
		residentCtrl.DeletePost()
	case 4:
		return
	default:
		fmt.Println("Invalid choice, please try again.")
	}
}
