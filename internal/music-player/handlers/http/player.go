package http

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"music-player/internal/music-player/models"
	"music-player/internal/music-player/services"
	"net/http"
)

type PlayerHandler struct {
	Service *services.PlayerService
}

func (ph *PlayerHandler) NextSong(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ph.Service.SendNextCommand()
	if ph.Service.Playlist.CurrentSong == nil {
		http.Error(w, "No song is currently playing", http.StatusBadRequest)
		return
	}

	song := ph.Service.Playlist.CurrentSong.Value.(*models.Song)
	songJSON, err := json.Marshal(song)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(songJSON)
}

func (ph *PlayerHandler) PrevSong(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ph.Service.SendPrevCommand()
	if ph.Service.Playlist.CurrentSong == nil {
		http.Error(w, "No song is currently playing", http.StatusBadRequest)
		return
	}

	song := ph.Service.Playlist.CurrentSong.Value.(*models.Song)
	songJSON, err := json.Marshal(song)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(songJSON)
}

func (ph *PlayerHandler) PlaySong(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ph.Service.Play()
	if ph.Service.Playlist.CurrentSong == nil {
		http.Error(w, "No song is currently playing", http.StatusBadRequest)
		return
	}

	song := ph.Service.Playlist.CurrentSong.Value.(*models.Song)
	songJSON, err := json.Marshal(song)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(songJSON)
}

func (ph *PlayerHandler) PauseSong(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ph.Service.SendPauseCommand()
	if ph.Service.Playlist.CurrentSong == nil {
		http.Error(w, "No song is currently playing", http.StatusBadRequest)
		return
	}

	song := ph.Service.Playlist.CurrentSong.Value.(*models.Song)
	songJSON, err := json.Marshal(song)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(songJSON)
}
