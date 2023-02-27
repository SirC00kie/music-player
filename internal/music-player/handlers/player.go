package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"music-player/internal/music-player/models"
	"music-player/internal/music-player/services"
	"net/http"
)

type PlayerHandler struct {
	Service *services.PlaylistService
}

func (ph *PlayerHandler) NextSong(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ph.Service.SendNextCommand()
	if ph.Service.Playlist.CurrentSong == nil {
		http.Error(w, "No song is currently playing", http.StatusBadRequest)
		return
	}
	song := ph.Service.Playlist.CurrentSong.Value.(*models.Song)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("NextSong %s", song)))
}

func (ph *PlayerHandler) PrevSong(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ph.Service.SendPrevCommand()
	if ph.Service.Playlist.CurrentSong == nil {
		http.Error(w, "No song is currently playing", http.StatusBadRequest)
		return
	}
	song := ph.Service.Playlist.CurrentSong.Value.(*models.Song)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Prev song %s", song)))
}

func (ph *PlayerHandler) PlaySong(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ph.Service.SendPlayCommand()
	if ph.Service.Playlist.CurrentSong == nil {
		http.Error(w, "No song is currently playing", http.StatusBadRequest)
		return
	}
	song := ph.Service.Playlist.CurrentSong.Value.(*models.Song)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Playing song %s with duration %s", song.Title, song.Duration)))
}

func (ph *PlayerHandler) PauseSong(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ph.Service.SendPauseCommand()
	if ph.Service.Playlist.CurrentSong == nil {
		http.Error(w, "No song is currently playing", http.StatusBadRequest)
		return
	}
	song := ph.Service.Playlist.CurrentSong.Value.(*models.Song)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Playing song %s pause time %s", song.Title, ph.Service.Playlist.PausedTime)))
}
