package services

import (
	"container/list"
	"errors"
	"fmt"
	"music-player/internal/music-player/models"
	"time"
)

type PlayerService struct {
	Playlist         *models.Playlist
	playChan         chan bool
	pauseChan        chan bool
	nextChan         chan bool
	prevChan         chan bool
	listenerRunning  bool
	listenerStopChan chan bool
}

func NewPlayerService() *PlayerService {
	playlist := &models.Playlist{
		SongList: list.New(),
	}
	return &PlayerService{
		Playlist:         playlist,
		playChan:         make(chan bool, 1),
		pauseChan:        make(chan bool, 1),
		nextChan:         make(chan bool, 1),
		prevChan:         make(chan bool, 1),
		listenerStopChan: make(chan bool, 1),
	}
}

func (ps *PlayerService) StartListener() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for range ticker.C {

		select {
		case <-ps.playChan:
			err := ps.Play()
			if err != nil {
				//return err
			}
			continue
		case <-ps.nextChan:
			ps.Playlist.Playing = false
			err := ps.NextSong()
			if err != nil {
				//return err
			}
			continue
		case <-ps.prevChan:
			ps.Playlist.Playing = false
			err := ps.PrevSong()
			if err != nil {
				//return err
			}
			continue
		case <-ps.listenerStopChan:
			return
		}

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

	select {
	case <-ps.pauseChan:
		err := ps.Pause()
		if err != nil {
			return err
		}
		return nil
	case <-ps.nextChan:
		err := ps.NextSong()
		if err != nil {
			return err
		}
		return nil
	case <-ps.prevChan:
		err := ps.PrevSong()
		if err != nil {
			return err
		}
		return nil
	//case <-time.After(durationLeft):
	//	if ps.Playlist.Playing {
	//		ps.SendNextCommand()
	//	}
	default:
		ps.Playlist.CurrentTime = time.Since(ps.Playlist.StartTime)
		song := ps.Playlist.CurrentSong.Value.(*models.Song)

		if ps.Playlist.CurrentTime >= ps.Playlist.CurrentSong.Value.(*models.Song).Duration {
			ps.SendNextCommand()
		} else {
			durationLeft := ps.Playlist.CurrentSong.Value.(*models.Song).Duration - ps.Playlist.CurrentTime
			fmt.Printf("Now playing: %s by %s. Time left: %v\n", song.Title, song.Author, durationLeft)
			time.Sleep(ps.Playlist.CurrentSong.Value.(*models.Song).Duration - ps.Playlist.CurrentTime)
			if ps.Playlist.Playing {
				ps.SendNextCommand()
			}
		}
	}

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
	ps.Playlist.StartTime = time.Now()
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

	ps.SendPlayCommand()
	fmt.Println("Next Song")
	return nil
}

func (ps *PlayerService) PrevSong() error {
	ps.Playlist.CurrentTime = 0
	ps.Playlist.PausedTime = 0
	ps.Playlist.StartTime = time.Now()
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

	ps.SendPlayCommand()
	fmt.Println("Prev Song")
	return nil
}

func (ps *PlayerService) SendPlayCommand() {
	ps.playChan <- true
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
