package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"music-player/internal/music-player/config"
	httpHandler "music-player/internal/music-player/handlers/http"
	"music-player/internal/music-player/services"
	"net/http"
)

const (
	songsPOST   = "/api/v1/songs"
	songsURL    = "/api/v1/songs/:index"
	playlistURL = "/api/v1/playlist"
	playURL     = "/api/v1/play"
	pauseURL    = "/api/v1/pause"
	nextURL     = "/api/v1/next"
	prevURL     = "/api/v1/prev"
)

func main() {

	playerService := services.NewPlayerService()
	playlistService := services.NewPlaylistService(playerService.Playlist)

	handlerPlayer := httpHandler.PlayerHandler{Service: playerService}
	handlerPlaylist := httpHandler.PlaylistHandler{Service: playlistService}

	cfg := config.GetConfig()

	router := httprouter.New()

	router.POST(songsPOST, handlerPlaylist.AddSong)
	router.GET(songsURL, handlerPlaylist.GetSong)
	router.PUT(songsURL, handlerPlaylist.UpdateSong)
	router.DELETE(songsURL, handlerPlaylist.DeleteSong)
	router.GET(playlistURL, handlerPlaylist.GetPlaylist)
	router.POST(nextURL, handlerPlayer.NextSong)
	router.POST(prevURL, handlerPlayer.PrevSong)
	router.GET(playURL, handlerPlayer.PlaySong)
	router.GET(pauseURL, handlerPlayer.PauseSong)

	log.Fatal(http.ListenAndServe(cfg.Listen.Port, router))

}
