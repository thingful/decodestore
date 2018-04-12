package datastoreserver

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/twitchtv/twirp"

	"github.com/thingful/decodestore/pkg/rpc/datastore"
)

// event is an internal type for recording an encrypted message for a consumer.
type event struct {
	timestamp time.Time
	data      []byte
}

// Server implements the EncryptedStore interface defined in our proto file. It
// just wraps a simple string/event map which we use to simulate the simplest
// possible data store.
type Server struct {
	store map[string]*event
}

// verify the adherence to the interface at compile time
var _ datastore.EncryptedStore = &Server{}

// NewServer instantiates and returns a new Server instance that implements our
// EncryptedStore interface.
func NewServer() datastore.EncryptedStore {
	s := &Server{
		store: make(map[string]*event),
	}

	return s
}

// WriteData is our implementation of the corresponding method on our interface
// that persists data to some store.
func (s *Server) WriteData(ctx context.Context, req *datastore.WriteRequest) (*datastore.WriteResponse, error) {
	if req.PublicKey == "" {
		return nil, twirp.InvalidArgumentError("public_key", "is required to identify the consumer for which data should be stored")
	}

	ev := &event{
		timestamp: time.Now(),
		data:      req.Data,
	}

	s.store[req.PublicKey] = ev

	return &datastore.WriteResponse{}, nil
}

// ReadData is our implementation of the corresponding method on our interface
// that allows data to be read from the store.
func (s *Server) ReadData(ctx context.Context, req *datastore.ReadRequest) (*datastore.ReadResponse, error) {
	if req.PublicKey == "" {
		return nil, twirp.InvalidArgumentError("public_key", "is required to identify the consumer for which data should be returned")
	}

	ev := s.store[req.PublicKey]
	if ev == nil {
		return nil, twirp.NotFoundError("unable to find data for the requested consumer")
	}

	t, err := ptypes.TimestampProto(ev.timestamp)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	events := []*datastore.EncryptedEvent{
		&datastore.EncryptedEvent{
			EventTime: t,
			Data:      ev.data,
		},
	}

	return &datastore.ReadResponse{
		PublicKey: req.PublicKey,
		Events:    events,
	}, nil
}
