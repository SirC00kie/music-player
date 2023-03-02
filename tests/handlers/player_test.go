package handlers

import (
	"github.com/stretchr/testify/assert"
	"music-player/internal/music-player/models"
	"music-player/internal/music-player/services"
	"testing"
	time "time"
)

func TestPlayerService_Play(t *testing.T) {
	playerService := services.NewPlayerService()
	playlistService := services.NewPlaylistService(playerService.Playlist)

	song1 := &models.Song{
		Title:    "test title1",
		Author:   "test author1",
		Duration: 1 * time.Second,
	}
	song2 := &models.Song{
		Title:    "test title2",
		Author:   "test author2",
		Duration: 3 * time.Second,
	}
	playlistService.AddSong(song1)
	playlistService.AddSong(song2)

	err := playerService.Play()
	if err != nil {
		return
	}
	time.Sleep(1 * time.Second)
	assert.True(t, playerService.Playlist.Playing)
	time.Sleep(1 * time.Second)
	assert.Equal(t, song2.Title, playerService.Playlist.CurrentSong.Value.(*models.Song).Title)
}

func TestPlayerService_Play_EmptyList(t *testing.T) {
	playerService := services.NewPlayerService()
	err := playerService.Play()
	assert.NotNil(t, err)
	assert.Equal(t, "list empty", err.Error())
}

func TestPlayerService_Play_AlreadyPlaying(t *testing.T) {
	playerService := services.NewPlayerService()
	playlistService := services.NewPlaylistService(playerService.Playlist)

	song1 := &models.Song{
		Title:    "test title1",
		Author:   "test author1",
		Duration: 2 * time.Second,
	}
	playlistService.AddSong(song1)

	playerService.Playlist.Playing = true
	err := playerService.Play()
	assert.NotNil(t, err)
	assert.Equal(t, "already playing", err.Error())
}

func TestPlayerService_Pause(t *testing.T) {
	playerService := services.NewPlayerService()
	playlistService := services.NewPlaylistService(playerService.Playlist)

	song1 := &models.Song{
		Title:    "test title1",
		Author:   "test author1",
		Duration: 6 * time.Second,
	}
	playlistService.AddSong(song1)

	err := playerService.Play()
	time.Sleep(2 * time.Second)
	playerService.SendPauseCommand()
	time.Sleep(1 * time.Second)
	assert.Nil(t, err)
	assert.False(t, playerService.Playlist.Playing)
	assert.Equal(t, 2*time.Second, playerService.Playlist.PausedTime.Round(time.Millisecond*100))
}

func TestPlayerService_Pause_NoPlayingSong(t *testing.T) {
	playerService := services.NewPlayerService()
	err := playerService.Pause()
	assert.NotNil(t, err)
	assert.Equal(t, "no playing song", err.Error())
}

func TestPlayerService_NextSong(t *testing.T) {
	playerService := services.NewPlayerService()
	playlistService := services.NewPlaylistService(playerService.Playlist)

	song1 := &models.Song{
		Title:    "test title1",
		Author:   "test author1",
		Duration: 3 * time.Second,
	}
	song2 := &models.Song{
		Title:    "test title2",
		Author:   "test author2",
		Duration: 3 * time.Second,
	}
	playlistService.AddSong(song1)
	playlistService.AddSong(song2)

	err := playerService.Play()
	assert.Nil(t, err)
	time.Sleep(1 * time.Second)
	playerService.SendNextCommand()
	time.Sleep(1 * time.Second)
	assert.Equal(t,
		playerService.Playlist.SongList.Front().Next().Value.(*models.Song),
		playerService.Playlist.CurrentSong.Value.(*models.Song))
	assert.True(t, playerService.Playlist.Playing)
	assert.Equal(t, 0*time.Second, playerService.Playlist.CurrentTime)
	assert.Equal(t, 0*time.Second, playerService.Playlist.PausedTime)
}

func TestPlayerService_PrevSong(t *testing.T) {
	playerService := services.NewPlayerService()
	playlistService := services.NewPlaylistService(playerService.Playlist)

	song1 := &models.Song{
		Title:    "test title1",
		Author:   "test author1",
		Duration: 3 * time.Second,
	}
	song2 := &models.Song{
		Title:    "test title2",
		Author:   "test author2",
		Duration: 3 * time.Second,
	}
	playlistService.AddSong(song1)
	playlistService.AddSong(song2)

	err := playerService.Play()
	assert.Nil(t, err)
	playerService.SendNextCommand()
	time.Sleep(1 * time.Second)
	playerService.SendPrevCommand()
	time.Sleep(1 * time.Second)
	assert.Equal(t,
		playerService.Playlist.SongList.Front().Value.(*models.Song),
		playerService.Playlist.CurrentSong.Value.(*models.Song))
	assert.True(t, playerService.Playlist.Playing)
	assert.Equal(t, 0*time.Second, playerService.Playlist.CurrentTime)
	assert.Equal(t, 0*time.Second, playerService.Playlist.PausedTime)
}
