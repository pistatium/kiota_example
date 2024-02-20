package main

import "context"

type Post struct {
	ID      int    `json:"id" binding:"-"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

var FakePosts = []*Post{
	{ID: 1, Title: "First Post", Content: "This is first test post."},
	{ID: 2, Title: "Second Post", Content: "This is second test post."},
	{ID: 3, Title: "Third Post", Content: "This is third test post."},
}

type PostRepository struct{}

func (r *PostRepository) GetPosts(ctx context.Context, offset int, limit int) ([]*Post, error) {
	return getSliceWithinRange(FakePosts, offset, limit), nil
}

func (r *PostRepository) AddPost(ctx context.Context, post *Post) error {
	post.ID = len(FakePosts) + 1
	FakePosts = append(FakePosts, post)
	return nil
}

func getSliceWithinRange(posts []*Post, offset int, limit int) []*Post {
	if offset >= len(posts) {
		return []*Post{}
	}
	if limit == 0 {
		return posts[offset:]
	}
	end := offset + limit
	if end > len(posts) {
		end = len(posts)
	}
	return posts[offset:end]
}
