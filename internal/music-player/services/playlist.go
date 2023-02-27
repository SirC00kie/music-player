package services

import (
	"fmt"
	"music-player/internal/music-player/models"
)

type PlaylistService struct {
	playlist *models.Playlist
}

func NewPlaylistService(playlist *models.Playlist) *PlaylistService {
	return &PlaylistService{playlist: playlist}
}

func (ps *PlaylistService) AddSong(song *models.Song) {
	ps.playlist.SongList.PushBack(song)
}

func (ps *PlaylistService) GetSong(index int) (*models.Song, error) {

	if ps.playlist.SongList.Len() == 0 {
		return nil, fmt.Errorf("playlist empty")
	}

	e := ps.playlist.SongList.Front()
	for i := 0; i < index; i++ {
		e = e.Next()
		if e == nil {
			return nil, fmt.Errorf("song not found")
		}
	}

	song := e.Value.(*models.Song)
	return song, nil
}

func (ps *PlaylistService) UpdateSong(index int, song models.Song) error {
	if ps.playlist.SongList.Len() == 0 {
		return fmt.Errorf("playlist empty")
	}

	e := ps.playlist.SongList.Front()
	for i := 0; i < index; i++ {
		e = e.Next()
		if e == nil {
			return fmt.Errorf("song not found")
		}
	}
	if ps.playlist.CurrentSong == e.Value && ps.playlist.Playing {
		return fmt.Errorf("cannot update current song while playlist is playing")
	}

	e.Value = &song

	return nil
}

func (ps *PlaylistService) DeleteSong(index int) error {
	if ps.playlist.SongList.Len() == 0 {
		return fmt.Errorf("playlist empty")
	}

	e := ps.playlist.SongList.Front()

	for i := 0; i < index; i++ {
		e = e.Next()
		if e == nil {
			return fmt.Errorf("song not found")
		}
	}

	if ps.playlist.CurrentSong == e.Value && ps.playlist.Playing {
		return fmt.Errorf("cannot update current song while playlist is playing")
	}

	ps.playlist.SongList.Remove(e)

	return nil
}

func (ps *PlaylistService) GetPlaylist() (map[string]interface{}, error) {
	if ps.playlist.SongList.Len() == 0 {
		return nil, fmt.Errorf("playlist empty")
	}

	var songs []map[string]interface{}

	for e := ps.playlist.SongList.Front(); e != nil; e = e.Next() {
		song := e.Value.(*models.Song)
		songs = append(songs, map[string]interface{}{
			"title":    song.Title,
			"author":   song.Author,
			"duration": song.Duration,
		})
	}

	playlistData := map[string]interface{}{
		"songs":       songs,
		"currentSong": ps.playlist.CurrentSong,
		"currentTime": ps.playlist.CurrentTime,
		"pausedTime":  ps.playlist.PausedTime,
		"playing":     ps.playlist.Playing,
	}

	return playlistData, nil
}
