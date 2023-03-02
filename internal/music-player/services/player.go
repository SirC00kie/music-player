package services

import (
	"container/list"
	"errors"
	"fmt"
	"music-player/internal/music-player/models"
	"time"
)

type PlayerService struct {
	Playlist  *models.Playlist
	playChan  chan bool
	pauseChan chan bool
	nextChan  chan bool
	prevChan  chan bool
	stopChan  chan bool
}

func NewPlayerService() *PlayerService {
	playlist := &models.Playlist{
		SongList: list.New(),
	}
	return &PlayerService{
		Playlist:  playlist,
		playChan:  make(chan bool, 1),
		pauseChan: make(chan bool, 1),
		nextChan:  make(chan bool, 1),
		prevChan:  make(chan bool, 1),
		stopChan:  make(chan bool, 1),
	}
}

func (ps *PlayerService) Play() error {
	if ps.Playlist.Playing {
		return errors.New("already playing")
	}
	if ps.Playlist.SongList.Len() == 0 {
		return errors.New("list empty")
	}

	if ps.Playlist.CurrentSong == nil {
		ps.Playlist.CurrentSong = ps.Playlist.SongList.Front()
	}

	ps.Playlist.Playing = true

	if ps.Playlist.PausedTime > 0 {
		ps.Playlist.StartTime = time.Now().Add(-ps.Playlist.PausedTime)
	} else {
		ps.Playlist.StartTime = time.Now()
	}

	song := ps.Playlist.CurrentSong.Value.(*models.Song)
	// for trigger event when song finished
	duration := ps.Playlist.CurrentSong.Value.(*models.Song).Duration - ps.Playlist.PausedTime
	timer := time.NewTimer(duration)

	fmt.Printf("Now playing: %s by %s. Time left: %v\n", song.Title, song.Author, timer)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	go func() {
		fmt.Print("go func")
		select {
		case <-timer.C:
			fmt.Print("timer c")
			timer.Stop()
			err := ps.NextSong()
			if err != nil {
				return
			}
		case <-ps.pauseChan:
			timer.Stop()
			ps.Playlist.CurrentTime = time.Since(ps.Playlist.StartTime)
			err := ps.Pause()
			if err != nil {
				return
			}
		case <-ps.nextChan:
			timer.Stop()
			err := ps.NextSong()
			if err != nil {
				return
			}
		case <-ps.prevChan:
			timer.Stop()
			err := ps.PrevSong()
			if err != nil {
				return
			}
		}
	}()

	return nil
}

func (ps *PlayerService) Pause() error {
	if ps.Playlist.Playing {
		ps.Playlist.Playing = false
		ps.Playlist.PausedTime = ps.Playlist.CurrentTime
		fmt.Println("pause")
	} else {
		return errors.New("no playing song")
	}

	return nil
}

func (ps *PlayerService) NextSong() error {
	ps.Playlist.CurrentTime = 0
	ps.Playlist.PausedTime = 0
	ps.Playlist.Playing = false

	if ps.Playlist.SongList.Len() == 0 {
		return errors.New("list empty")
	}

	if ps.Playlist.CurrentSong == nil {
		ps.Playlist.CurrentSong = ps.Playlist.SongList.Front()
	} else if ps.Playlist.CurrentSong.Next() != nil {
		ps.Playlist.CurrentSong = ps.Playlist.CurrentSong.Next()
	} else {
		ps.Playlist.CurrentSong = ps.Playlist.SongList.Front()
	}

	err := ps.Play()
	if err != nil {
		return err
	}
	fmt.Println("Next Song")
	return nil
}

func (ps *PlayerService) PrevSong() error {
	ps.Playlist.CurrentTime = 0
	ps.Playlist.PausedTime = 0
	ps.Playlist.Playing = false

	if ps.Playlist.SongList.Len() == 0 {
		return errors.New("list empty")
	}
	if ps.Playlist.CurrentSong == nil {
		ps.Playlist.CurrentSong = ps.Playlist.SongList.Front()
	} else if ps.Playlist.CurrentSong.Prev() != nil {
		ps.Playlist.CurrentSong = ps.Playlist.CurrentSong.Prev()
	} else {
		ps.Playlist.CurrentSong = ps.Playlist.SongList.Front()
	}

	err := ps.Play()
	if err != nil {
		return err
	}
	fmt.Println("Prev Song")
	return nil
}

func (ps *PlayerService) SendPauseCommand() {
	ps.pauseChan <- true
}

func (ps *PlayerService) SendNextCommand() {
	ps.nextChan <- true
}

func (ps *PlayerService) SendPrevCommand() {
	ps.prevChan <- true
}
