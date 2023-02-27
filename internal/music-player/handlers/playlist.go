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

	e := ph.Service.Playlist.SongList.Front()
	for i := 0; i < index; i++ {
		e = e.Next()
		if e == nil {
			http.NotFound(w, r)
			return
		}
	}
	song := e.Value.(*models.Song)

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

	e := ph.Service.Playlist.SongList.Front()
	for i := 0; i < index; i++ {
		e = e.Next()
		if e == nil {
			http.NotFound(w, r)
			return
		}
	}

	if ph.Service.Playlist.CurrentSong == e.Value && ph.Service.Playlist.Playing {
		http.Error(w, "Cannot update current song while playlist is playing", http.StatusConflict)
		return
	}

	var song models.Song
	err = json.NewDecoder(r.Body).Decode(&song)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	e.Value = &song

	w.WriteHeader(http.StatusNoContent)
}

func (ph *PlaylistHandler) DeleteSong(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	index, err := strconv.Atoi(ps.ByName("index"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	e := ph.Service.Playlist.SongList.Front()
	for i := 0; i < index; i++ {
		e = e.Next()
		if e == nil {
			http.NotFound(w, r)
			return
		}
	}

	if ph.Service.Playlist.CurrentSong == e.Value && ph.Service.Playlist.Playing {
		http.Error(w, "Cannot update current song while playlist is playing", http.StatusConflict)
		return
	}

	ph.Service.Playlist.SongList.Remove(e)

	w.WriteHeader(http.StatusNoContent)
}

func (ph *PlaylistHandler) GetPlaylist(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var songs []map[string]interface{}

	for e := ph.Service.Playlist.SongList.Front(); e != nil; e = e.Next() {
		song := e.Value.(*models.Song)
		songs = append(songs, map[string]interface{}{
			"title":    song.Title,
			"author":   song.Author,
			"duration": song.Duration,
		})
	}

	playlistData := map[string]interface{}{
		"songs":       songs,
		"startTime":   ph.Service.Playlist.StartTime,
		"currentTime": ph.Service.Playlist.CurrentTime,
		"pausedTime":  ph.Service.Playlist.PausedTime,
		"playing":     ph.Service.Playlist.Playing,
	}

	playlistJSON, err := json.Marshal(playlistData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(playlistJSON)
}
