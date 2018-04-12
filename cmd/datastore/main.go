package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/twitchtv/twirp"

	"github.com/thingful/datastore/pkg/impl/datastoreserver"
	"github.com/thingful/datastore/pkg/rpc/datastore"
)

type contextKey string

func (c contextKey) String() string {
	return fmt.Sprintf("DECODE.%s", string(c))
}

var (
	reqStartTimestampKey = contextKey("request.start")
)

func main() {
	server := datastoreserver.NewServer()

	hooks := &twirp.ServerHooks{}

	hooks.RequestReceived = func(ctx context.Context) (context.Context, error) {
		return context.WithValue(ctx, reqStartTimestampKey, time.Now()), nil
	}

	hooks.ResponseSent = func(ctx context.Context) {
		start, ok := ctx.Value(reqStartTimestampKey).(time.Time)
		if ok {
			dur := time.Now().Sub(start)
			fmt.Printf("Request completed, duration: %v\n", dur)
		}
	}

	twirpHandler := datastore.NewEncryptedStoreServer(server, hooks)

	http.ListenAndServe(":8080", twirpHandler)
}
