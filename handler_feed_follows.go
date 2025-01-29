package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Nivan-22/rssagg/internal/database"
	"github.com/go-chi/chi"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error in parsing json:%v", err))
		return
	}
	feed_follow, err := apiCfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error in creating feed follow :%v and %v", err, user))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feed_follow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollows, err := apiCfg.DB.GetFeedfollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error in getting feed follow :%v and %v", err, user))
		return
	}

	respondWithJSON(w, 201, databasefeedfollowsstoFeedsfollows(feedFollows))
}
func (apiCfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedfollowIDstr := chi.URLParam(r, "feedFollowID")
	feedfollowID, err := uuid.Parse(feedfollowIDstr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error in parsing feed follow id :%v and %v", err, user))
		return
	}
	err = apiCfg.DB.DeletefeedFollows(r.Context(), database.DeletefeedFollowsParams{
		ID:     feedfollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error in deleting feed follow :%v and %v", err, user))
		return
	}
	respondWithJSON(w, 200, struct {
		Message string `json:"message"`
	}{
		Message: "Deletion sucess",
	})

}
