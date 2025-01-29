package main

import (
	"time"

	"github.com/Nivan-22/rssagg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}

}

type feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databasefeedtoFeed(dbFeed database.Feed) feed {
	return feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}
func databasefeedstoFeeds(dbFeeds []database.Feed) []feed {
	feeds := []feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databasefeedtoFeed(dbFeed))
	}
	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedsFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		FeedID:    dbFeedFollow.FeedID,
		UserID:    dbFeedFollow.ID,
	}
}
func databasefeedfollowsstoFeedsfollows(dbFeedFollows []database.FeedsFollow) []FeedFollow {
	feedfollows := []FeedFollow{}
	for _, dbFeedFollow := range dbFeedFollows {
		feedfollows = append(feedfollows, databaseFeedFollowToFeedFollow(dbFeedFollow))
	}
	return feedfollows
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"descrption"`
	PublishedAt time.Time `json:"publised_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func databasePoststoPost(dbpost database.Post) Post {
	var description *string
	if dbpost.Description.Valid {
		description = &dbpost.Description.String
	}
	return Post{
		ID:          dbpost.ID,
		CreatedAt:   dbpost.CreatedAt,
		UpdatedAt:   dbpost.UpdatedAt,
		Title:       dbpost.Title,
		Description: description,
		PublishedAt: dbpost.PublishedAt,
		Url:         dbpost.Url,
		FeedID:      dbpost.FeedID,
	}
}


func databasePoststoPosts(dbPosts []database.Post) []Post{
	posts := []Post{}
	for _, dbPost := range dbPosts {
		posts = append(posts,databasePoststoPost(dbPost) )
	}
	return posts

}