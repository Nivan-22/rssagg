package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string){
	if code > 499{
		log.Println("Respondimg with error:",msg)
	}
	type errResposne struct{
		Error string `json:"error"`
	}
	respondWithJSON(w,code,errResposne{
		Error:msg,
	})
}
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}){

	dat,err := json.Marshal(payload)
    if err != nil{
		log.Printf(
		 "failed to Marshal the : %v",payload,
		)
		w.WriteHeader(500)
		return
	}
    w.Header().Add("Content-Type","application/json")
	w.WriteHeader(code )
	w.Write(dat)


}
