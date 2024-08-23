package utils

var PostID int
var UserID int

func GeneratePostId() int {
	PostID = PostID + 1
	return PostID
}
func GenerateUserID() int {
	UserID = UserID + 1
	return UserID
}
