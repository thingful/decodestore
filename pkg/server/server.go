package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/thingful/datastore/pkg/internal/datastoreserver"
	"github.com/thingful/datastore/pkg/rpc/datastore"
)

type Server struct {
	srv *http.Server
}

func NewServer(addr string) *Server {
	rpc := datastoreserver.NewServer()
	twirpHandler := datastore.NewEncryptedStoreServer(rpc, nil)

	srv := &http.Server{Addr: addr, Handler: twirpHandler}

	return &Server{
		srv: srv,
	}
}

func (s *Server) Start() {
	log.Println("Starting server...")

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	go func() {
		if err := s.srv.ListenAndServe(); err != nil {
			log.Printf("listen error: %s\n", err)
		}
	}()

	<-stopChan
	log.Println("Shutting down server...")

	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	s.srv.Shutdown(ctx)
}
