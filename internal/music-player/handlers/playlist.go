package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"music-player/internal/music-player/models"
	"music-player/internal/music-player/services"
	"net/http"
)

type PlaylistHandler struct {
	Service *services.PlaylistService
}

func (ph *PlaylistHandler) AddSong(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var song models.Song
	err := json.NewDecoder(r.Body).Decode(&song)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ph.Service.AddSong(&song)

	w.Write([]byte(fmt.Sprintf("Added song %s", &song)))
	w.WriteHeader(http.StatusOK)
}

func (ph *PlaylistHandler) NextSong(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ph.Service.SendNextCommand()
	if ph.Service.Playlist.CurrentSong == nil {
		http.Error(w, "No song is currently playing", http.StatusBadRequest)
		return
	}
	song := ph.Service.Playlist.CurrentSong.Value.(*models.Song)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("NextSong %s", song)))
}

func (ph *PlaylistHandler) PrevSong(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ph.Service.SendPrevCommand()
	if ph.Service.Playlist.CurrentSong == nil {
		http.Error(w, "No song is currently playing", http.StatusBadRequest)
		return
	}
	song := ph.Service.Playlist.CurrentSong.Value.(*models.Song)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Prev song %s", song)))
}

func (ph *PlaylistHandler) PlaySong(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ph.Service.SendPlayCommand()
	if ph.Service.Playlist.CurrentSong == nil {
		http.Error(w, "No song is currently playing", http.StatusBadRequest)
		return
	}
	song := ph.Service.Playlist.CurrentSong.Value.(*models.Song)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Playing song %s with duration %s", song.Title, song.Duration)))
}

func (ph *PlaylistHandler) PauseSong(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ph.Service.SendPauseCommand()
	if ph.Service.Playlist.CurrentSong == nil {
		http.Error(w, "No song is currently playing", http.StatusBadRequest)
		return
	}
	song := ph.Service.Playlist.CurrentSong.Value.(*models.Song)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Playing song %s pause time %s", song.Title, ph.Service.Playlist.PausedTime)))
}

func (ph *PlaylistHandler) ShowPlaylist(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Playlist %s", ph.Service.Playlist)))
}
