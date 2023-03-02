package handlers

import (
	"bytes"
	"container/list"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	httpHandler "music-player/internal/music-player/handlers/http"
	"music-player/internal/music-player/models"
	"music-player/internal/music-player/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlaylistHandler_AddSong(t *testing.T) {
	playlist := &models.Playlist{
		SongList: list.New(),
	}
	service := services.NewPlaylistService(playlist)
	handler := &httpHandler.PlaylistHandler{Service: service}

	song := &models.Song{
		Title:    "test title",
		Author:   "test author",
		Duration: 180,
	}
	songJSON, err := json.Marshal(song)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/add", bytes.NewBuffer(songJSON))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	handler.AddSong(rr, req, nil)
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, 1, playlist.SongList.Len())
}

func TestPlaylistHandler_GetSong(t *testing.T) {
	playlist := &models.Playlist{
		SongList: list.New(),
	}
	service := services.NewPlaylistService(playlist)
	handler := &httpHandler.PlaylistHandler{Service: service}

	song := &models.Song{
		Title:    "test title",
		Author:   "test author",
		Duration: 180,
	}
	handler.Service.AddSong(song)

	req, err := http.NewRequest("GET", "/songs/0", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	ps := httprouter.Params{httprouter.Param{Key: "index", Value: "0"}}
	handler.GetSong(rr, req, ps)
	assert.Equal(t, http.StatusOK, rr.Code)

	responseSong := &models.Song{}
	err = json.Unmarshal(rr.Body.Bytes(), responseSong)
	assert.NoError(t, err)

	assert.Equal(t, song, responseSong)
}

func TestPlaylistHandler_UpdateSong(t *testing.T) {
	playlist := &models.Playlist{
		SongList: list.New(),
	}
	service := services.NewPlaylistService(playlist)
	handler := &httpHandler.PlaylistHandler{Service: service}

	song := &models.Song{
		Title:    "test title",
		Author:   "test author",
		Duration: 180,
	}
	handler.Service.AddSong(song)

	newSong := &models.Song{
		Title:    "new title",
		Author:   "new author",
		Duration: 200,
	}
	newSongJSON, err := json.Marshal(newSong)
	assert.NoError(t, err)

	req, err := http.NewRequest("PUT", "/songs/0", bytes.NewBuffer(newSongJSON))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	ps := httprouter.Params{httprouter.Param{Key: "index", Value: "0"}}
	handler.UpdateSong(rr, req, ps)
	assert.Equal(t, http.StatusNoContent, rr.Code)

	responseSong, err := service.GetSong(0)
	assert.NoError(t, err)

	assert.Equal(t, newSong, responseSong)
}

func TestPlaylistHandler_DeleteSong(t *testing.T) {
	playlist := &models.Playlist{
		SongList: list.New(),
	}
	service := services.NewPlaylistService(playlist)
	handler := &httpHandler.PlaylistHandler{Service: service}

	song := &models.Song{
		Title:    "test title",
		Author:   "test author",
		Duration: 180,
	}
	handler.Service.AddSong(song)

	req, err := http.NewRequest("DELETE", "/songs/0", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	ps := httprouter.Params{httprouter.Param{Key: "index", Value: "0"}}
	handler.DeleteSong(rr, req, ps)
	assert.Equal(t, http.StatusNoContent, rr.Code)

	assert.Equal(t, 0, playlist.SongList.Len())
}

func TestPlaylistHandler_GetPlaylist(t *testing.T) {
	playlist := &models.Playlist{
		SongList: list.New(),
	}
	service := services.NewPlaylistService(playlist)
	handler := &httpHandler.PlaylistHandler{Service: service}

	song1 := &models.Song{
		Title:    "test title1",
		Author:   "test author1",
		Duration: 180,
	}
	song2 := &models.Song{
		Title:    "test title2",
		Author:   "test author2",
		Duration: 1800,
	}
	handler.Service.AddSong(song1)
	handler.Service.AddSong(song2)

	req, err := http.NewRequest("GET", "/playlist", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	handler.GetPlaylist(rr, req, nil)
	assert.Equal(t, http.StatusOK, rr.Code)

	var playlistData map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &playlistData)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(playlistData["songs"].([]interface{})))
}
