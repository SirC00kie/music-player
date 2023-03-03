package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"music-player/internal/music-player/models"
	"music-player/internal/music-player/services"
	api "music-player/pkg/api/playlist"
	"time"
)

type HandlerPlaylistGRPC struct {
	api.PlaylistServiceServer
	Service *services.PlaylistService
}

func (h *HandlerPlaylistGRPC) AddSong(ctx context.Context, req *api.Song) (*empty.Empty, error) {
	song := models.Song{
		Title:    req.Title,
		Author:   req.Author,
		Duration: time.Duration(req.Duration),
	}
	h.Service.AddSong(&song)
	return &empty.Empty{}, nil
}
func (h *HandlerPlaylistGRPC) GetSong(ctx context.Context, req *api.GetSongRequest) (*api.Song, error) {
	id := req.Id
	song, err := h.Service.GetSong(int(id))
	if err != nil {
		return nil, err
	}
	return &api.Song{
		Title:    song.Title,
		Author:   song.Author,
		Duration: int64(song.Duration),
	}, nil
}
func (h *HandlerPlaylistGRPC) UpdateSong(ctx context.Context, req *api.UpdateSongRequest) (*empty.Empty, error) {
	id := req.Id

	song := models.Song{
		Title:    req.Song.Title,
		Author:   req.Song.Author,
		Duration: time.Duration(req.Song.Duration),
	}

	err := h.Service.UpdateSong(int(id), song)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
func (h *HandlerPlaylistGRPC) DeleteSong(ctx context.Context, req *api.DeleteSongRequest) (*empty.Empty, error) {
	id := req.Id

	err := h.Service.DeleteSong(int(id))
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
func (h *HandlerPlaylistGRPC) GetPlaylist(ctx context.Context, req *empty.Empty) (*api.GetPlaylistResponse, error) {
	playlistData, err := h.Service.GetPlaylist()
	if err != nil {
		return nil, err
	}

	response := &api.GetPlaylistResponse{
		SongList: &api.SongList{},
	}

	for _, song := range playlistData["songs"].([]map[string]interface{}) {
		response.SongList.Song = append(response.SongList.Song, &api.Song{
			Title:    song["title"].(string),
			Author:   song["author"].(string),
			Duration: song["duration"].(int64),
		})
	}

	return response, nil
}
