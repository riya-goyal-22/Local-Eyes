package models

type Post struct {
	ID        int
	Title     string
	Content   string
	Type      string
	LikeCount int
	UserId    int
}
