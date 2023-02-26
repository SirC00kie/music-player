package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"music-player/internal/music-player/handlers"
	"music-player/internal/music-player/services"
	"net/http"
)

func main() {

	playlistService := services.NewPlaylistService()

	handler := handlers.PlaylistHandler{Service: playlistService}

	go handler.Service.StartListener()
	router := httprouter.New()

	router.POST("/add", handler.AddSong)
	router.POST("/next", handler.NextSong)
	router.POST("/prev", handler.PrevSong)
	router.GET("/play", handler.PlaySong)
	router.GET("/pause", handler.PauseSong)
	router.GET("/playlist", handler.ShowPlaylist)

	log.Fatal(http.ListenAndServe(":8080", router))

}
