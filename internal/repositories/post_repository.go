package repositories

import (
	"local-eyes/internal/models"
	"local-eyes/utils"
)

type PostRepository struct {
	FileRepository
}

func NewPostRepository() *PostRepository {
	return &PostRepository{
		FileRepository: *NewFileRepository(utils.PostFile),
	}
}

func (pr *PostRepository) Save(post *models.Post) error {
	posts, err := pr.Load()
	if err != nil {
		return err
	}
	posts = append(posts, post)
	return pr.SaveAll(posts)
}

func (pr *PostRepository) Load() ([]*models.Post, error) {
	var posts []*models.Post
	err := pr.FileRepository.Load(&posts)
	return posts, err
}

func (pr *PostRepository) Delete(postID string) error {
	posts, err := pr.Load()
	if err != nil {
		return err
	}
	for i, post := range posts {
		if post.ID == postID {
			posts = append(posts[:i], posts[i+1:]...)
			return pr.SaveAll(posts)
		}
	}
	return nil
}

func (pr *PostRepository) Update(post *models.Post) error {
	posts, err := pr.Load()
	if err != nil {
		return err
	}
	for i, p := range posts {
		if p.ID == post.ID {
			posts[i] = post
			return pr.SaveAll(posts)
		}
	}
	return nil
}

func (pr *PostRepository) Like(postID string) error {
	posts, err := pr.Load()
	if err != nil {
		return err
	}
	for _, post := range posts {
		if post.ID == postID {
			post.LikeCount++
			return pr.SaveAll(posts)
		}
	}
	return nil
}

func (pr *PostRepository) SaveAll(posts []*models.Post) error {
	return pr.FileRepository.Save(posts)
}
