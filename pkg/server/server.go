package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/thingful/decodestore/pkg/internal/datastoreserver"
	"github.com/thingful/decodestore/pkg/rpc/datastore"
)

type Server struct {
	srv    *http.Server
	logger log.Logger
}

func NewServer(addr string) *Server {
	rpc := datastoreserver.NewServer()
	twirpHandler := datastore.NewEncryptedStoreServer(rpc, nil)

	srv := &http.Server{Addr: addr, Handler: twirpHandler}

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "time", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	return &Server{
		srv:    srv,
		logger: logger,
	}
}

func (s *Server) Start() {
	s.logger.Log("msg", "starting server", "addr", s.srv.Addr)

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	go func() {
		if err := s.srv.ListenAndServe(); err != nil {
			s.logger.Log("msg", "listen error", "error", err)
		}
	}()

	<-stopChan
	s.logger.Log("msg", "stopping server")

	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	s.srv.Shutdown(ctx)
}
