package main

// User ...Layout for each user row
type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Hash      string `json:"hash"`
	Bio       string `json:"bio"`
	AvatarURL string `json:"avatar_url"`
	Followers int    `json:"followers"`
}

// Post ...Post structure in DB
type Post struct {
	ID        string `json:"id"`
	UserID    int    `json:"user_id"`
	ImageURL  string `json:"image_url"`
	Caption   string `json:"caption"`
	Time      string `json:"time"`
	LikeCount int    `json:"like_count"`
	Edited    bool   `json:"edited"`
}

// Comment ... Comment structure in DB
type Comment struct {
	ID      string `json:"id"`
	PostID  string `json:"post_id"`
	UserID  int    `json:"user_id"`
	Comment string `json:"comment"`
	Likes   int    `json:"likes"`
}

// Config ...Defines the layout of our config file
type Config struct {
	DatabaseURI string `json:"databaseURI"`
}
