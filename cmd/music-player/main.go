package main

import (
	"music-player/internal/music-player/models"
	"music-player/internal/music-player/services"
	"time"
)

func main() {

	song1 := &models.Song{
		Title:    "1",
		Author:   "1",
		Duration: 10 * time.Second,
	}

	song2 := &models.Song{
		Title:    "2",
		Author:   "2",
		Duration: 10 * time.Second,
	}

	playlistService := services.NewPlaylistService()

	playlistService.AddSong(song1)
	playlistService.AddSong(song2)

	go playlistService.Play()
	time.Sleep(2 * time.Second)

	playlistService.Pause()
	time.Sleep(2 * time.Second)
	playlistService.Play()

}
