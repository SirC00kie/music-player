package main

import (
	"music-player/internal/music-player/config"
	"music-player/internal/music-player/server"
	"music-player/internal/music-player/services"
)

func main() {

	cfg := config.GetConfig()
	playerService := services.NewPlayerService()
	playlistService := services.NewPlaylistService(playerService.Playlist)

	newServer := server.NewServer(playerService, playlistService, cfg)
	newServer.RunServer()

}
