package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	apiPl "music-player/pkg/api/player"
	api "music-player/pkg/api/playlist"
)

func main() {
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	clientPlaylist := api.NewPlaylistServiceClient(conn)
	clientPlayer := apiPl.NewPlayerServiceClient(conn)

	// Add a new song
	song1 := &api.Song{
		Title:    "New Song",
		Author:   "New Author",
		Duration: 180,
	}
	song2 := &api.Song{
		Title:    "New Song",
		Author:   "New Author",
		Duration: 380,
	}
	addResp, err := clientPlaylist.AddSong(context.Background(), song1)
	if err != nil {
		log.Fatalf("AddSong request failed: %v", err)
	}
	log.Printf("AddSong response: %v", addResp)

	addResp, err = clientPlaylist.AddSong(context.Background(), song2)
	if err != nil {
		log.Fatalf("AddSong request failed: %v", err)
	}
	log.Printf("AddSong response: %v", addResp)

	// Get an existing song
	getResp, err := clientPlaylist.GetSong(context.Background(), &api.GetSongRequest{Id: 0})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			log.Fatalf("Song not found: %v", err)
		}
		log.Fatalf("GetSong request failed: %v", err)
	}
	log.Printf("GetSong response: %v", getResp)

	// Update an existing song
	updateReq := &api.UpdateSongRequest{
		Id: 0,
		Song: &api.Song{
			Title:    "Updated Song",
			Author:   "Updated Author",
			Duration: 200,
		},
	}
	updateResp, err := clientPlaylist.UpdateSong(context.Background(), updateReq)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			log.Fatalf("Song not found: %v", err)
		}
		log.Fatalf("UpdateSong request failed: %v", err)
	}
	log.Printf("UpdateSong response: %v", updateResp)

	_, err = clientPlayer.PlaySong(context.Background(), &empty.Empty{})
	if err != nil {
		return
	}

	_, err = clientPlayer.NextSong(context.Background(), &empty.Empty{})
	if err != nil {
		return
	}

	_, err = clientPlayer.PauseSong(context.Background(), &empty.Empty{})
	if err != nil {
		return
	}

	// Delete an existing song
	deleteResp, err := clientPlaylist.DeleteSong(context.Background(), &api.DeleteSongRequest{Id: 0})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			log.Fatalf("Song not found: %v", err)
		}
		log.Fatalf("DeleteSong request failed: %v", err)
	}
	log.Printf("DeleteSong response: %v", deleteResp)

}
