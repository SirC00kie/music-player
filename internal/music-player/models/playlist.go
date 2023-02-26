package models

import (
	"container/list"
	"time"
)

type Playlist struct {
	SongList    *list.List    `json:"songList"`
	CurrentSong *list.Element `json:"currentSong"`
	StartTime   time.Time     `json:"startTime"`
	CurrentTIme time.Duration `json:"currentTIme"`
	PausedTime  time.Duration `json:"pausedTime"`
	Playing     bool          `json:"playing"`
}
