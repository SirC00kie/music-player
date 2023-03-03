package server

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"music-player/internal/music-player/config"
	grpcHandler "music-player/internal/music-player/handlers/grpc"
	httpHandler "music-player/internal/music-player/handlers/http"
	"music-player/internal/music-player/services"
	"net/http"
)

const (
	songsPOST   = "/api/v1/songs"
	songsURL    = "/api/v1/songs/:index"
	playlistURL = "/api/v1/playlist"
	playURL     = "/api/v1/play"
	pauseURL    = "/api/v1/pause"
	nextURL     = "/api/v1/next"
	prevURL     = "/api/v1/prev"
)

type Server struct {
	Config *config.Config
	//HandlerPlayerGRPC
	HandlerPlaylistHTTP *httpHandler.PlaylistHandler
	HandlerPlayerHTTP   *httpHandler.PlayerHandler
	HandlerPlaylistGRPC *grpcHandler.HandlerGRPC
}

func NewServer(p *services.PlayerService, pl *services.PlaylistService, cfg *config.Config) *Server {
	return &Server{
		Config:              cfg,
		HandlerPlayerHTTP:   &httpHandler.PlayerHandler{Service: p},
		HandlerPlaylistHTTP: &httpHandler.PlaylistHandler{Service: pl},
		HandlerPlaylistGRPC: &grpcHandler.HandlerGRPC{Service: pl},
	}
}

func (s *Server) RunServer() {
	//grpcServer := grpc.NewServer()
	//api.RegisterPlaylistServiceServer(grpcServer, s.HandlerPlaylistGRPC)
	//
	//listener, err := net.Listen(s.Config.ListenGRPC.Network, s.Config.ListenGRPC.Port)
	//if err != nil {
	//	log.Fatalf("Failed to listen: %v", err)
	//}
	//err = grpcServer.Serve(listener)
	//if err != nil {
	//	fmt.Print(err)
	//}

	httpRouter := httprouter.New()

	httpRouter.POST(songsPOST, s.HandlerPlaylistHTTP.AddSong)
	httpRouter.GET(songsURL, s.HandlerPlaylistHTTP.GetSong)
	httpRouter.PUT(songsURL, s.HandlerPlaylistHTTP.UpdateSong)
	httpRouter.DELETE(songsURL, s.HandlerPlaylistHTTP.DeleteSong)
	httpRouter.GET(playlistURL, s.HandlerPlaylistHTTP.GetPlaylist)
	httpRouter.POST(nextURL, s.HandlerPlayerHTTP.NextSong)
	httpRouter.POST(prevURL, s.HandlerPlayerHTTP.PrevSong)
	httpRouter.GET(playURL, s.HandlerPlayerHTTP.PlaySong)
	httpRouter.GET(pauseURL, s.HandlerPlayerHTTP.PauseSong)

	log.Fatal(http.ListenAndServe(s.Config.ListenHTTP.Port, httpRouter))

}
