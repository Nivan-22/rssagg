package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Nivan-22/rssagg/internal/database"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error in parsing json:%v", err))
		return
	}
	feed, err := apiCfg.DB.Createfeed(r.Context(), database.CreatefeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error in creating feed:%v and %v", err, user))
		return
	}

	respondWithJSON(w, 201, databasefeedtoFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	
	feeds, err := apiCfg.DB.Getfeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error in creating feed:%v ", err))
		return
	}

	respondWithJSON(w, 201, databasefeedstoFeeds(feeds))
}
