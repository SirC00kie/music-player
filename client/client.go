package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	api "music-player/pkg/api/playlist"
)

func main() {
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := api.NewPlaylistServiceClient(conn)

	// Add a new song
	song := &api.Song{
		Title:    "New Song",
		Author:   "New Author",
		Duration: 180,
	}
	addResp, err := client.AddSong(context.Background(), song)
	if err != nil {
		log.Fatalf("AddSong request failed: %v", err)
	}
	log.Printf("AddSong response: %v", addResp)

	// Get an existing song
	getResp, err := client.GetSong(context.Background(), &api.GetSongRequest{Id: 0})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			log.Fatalf("Song not found: %v", err)
		}
		log.Fatalf("GetSong request failed: %v", err)
	}
	log.Printf("GetSong response: %v", getResp)

}
