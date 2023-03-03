package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"music-player/internal/music-player/models"
	"music-player/internal/music-player/services"
	apiPl "music-player/pkg/api/player"
)

type HandlerPlayerGRPC struct {
	apiPl.PlayerServiceServer
	Service *services.PlayerService
}

func (h *HandlerPlayerGRPC) NextSong(ctx context.Context, req *empty.Empty) (*apiPl.SongResponse, error) {
	h.Service.SendNextCommand()
	if h.Service.Playlist.CurrentSong == nil {
		return nil, status.Errorf(codes.NotFound, "No song is currently playing")
	}
	song := h.Service.Playlist.CurrentSong.Value.(*models.Song)
	return &apiPl.SongResponse{Title: song.Title, Author: song.Author, Duration: int64(song.Duration)}, nil
}
func (h *HandlerPlayerGRPC) PrevSong(ctx context.Context, req *empty.Empty) (*apiPl.SongResponse, error) {
	h.Service.SendPrevCommand()
	if h.Service.Playlist.CurrentSong == nil {
		return nil, status.Errorf(codes.NotFound, "No song is currently playing")
	}
	song := h.Service.Playlist.CurrentSong.Value.(*models.Song)
	return &apiPl.SongResponse{Title: song.Title, Author: song.Author, Duration: int64(song.Duration)}, nil
}
func (h *HandlerPlayerGRPC) PlaySong(ctx context.Context, req *empty.Empty) (*apiPl.SongResponse, error) {
	err := h.Service.Play()
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "No song is currently playing %s", err)
	}
	if h.Service.Playlist.CurrentSong == nil {
		return nil, status.Errorf(codes.NotFound, "No song is currently playing")
	}
	song := h.Service.Playlist.CurrentSong.Value.(*models.Song)
	return &apiPl.SongResponse{Title: song.Title, Author: song.Author, Duration: int64(song.Duration)}, nil
}
func (h *HandlerPlayerGRPC) PauseSong(ctx context.Context, req *empty.Empty) (*apiPl.SongResponse, error) {
	h.Service.SendPauseCommand()
	if h.Service.Playlist.CurrentSong == nil {
		return nil, status.Errorf(codes.NotFound, "No song is currently playing")
	}
	song := h.Service.Playlist.CurrentSong.Value.(*models.Song)
	return &apiPl.SongResponse{Title: song.Title, Author: song.Author, Duration: int64(song.Duration)}, nil
}
