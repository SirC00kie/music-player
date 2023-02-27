package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"music-player/internal/music-player/handlers"
	"music-player/internal/music-player/services"
	"net/http"
)

func main() {

	playerService := services.NewPlayerService()
	playlistService := services.NewPlaylistService(playerService.Playlist)

	handlerPlayer := handlers.PlayerHandler{Service: playerService}
	handlerPlaylist := handlers.PlaylistHandler{Service: playlistService}

	go handlerPlayer.Service.StartListener()
	router := httprouter.New()

	router.POST("/add", handlerPlaylist.AddSong)
	router.GET("/get/:index", handlerPlaylist.GetSong)
	router.GET("/playlist", handlerPlaylist.GetPlaylist)
	router.PUT("/update/:index", handlerPlaylist.UpdateSong)
	router.DELETE("/delete/:index", handlerPlaylist.DeleteSong)
	router.POST("/next", handlerPlayer.NextSong)
	router.POST("/prev", handlerPlayer.PrevSong)
	router.GET("/play", handlerPlayer.PlaySong)
	router.GET("/pause", handlerPlayer.PauseSong)

	log.Fatal(http.ListenAndServe(":8080", router))

}
