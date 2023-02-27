package handlers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"music-player/internal/music-player/models"
	"music-player/internal/music-player/services"
	"net/http"
	"strconv"
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

	w.WriteHeader(http.StatusCreated)
}

func (ph *PlaylistHandler) GetSong(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	index, err := strconv.Atoi(ps.ByName("index"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	song, err := ph.Service.GetSong(index)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	songJSON, err := json.Marshal(song)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(songJSON)
}

func (ph *PlaylistHandler) UpdateSong(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	index, err := strconv.Atoi(ps.ByName("index"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var song models.Song
	err = json.NewDecoder(r.Body).Decode(&song)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ph.Service.UpdateSong(index, song)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (ph *PlaylistHandler) DeleteSong(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	index, err := strconv.Atoi(ps.ByName("index"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ph.Service.DeleteSong(index)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (ph *PlaylistHandler) GetPlaylist(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	playlistData, err := ph.Service.GetPlaylist()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	playlistJSON, err := json.Marshal(playlistData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(playlistJSON)
}
