package datastoreserver

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/twitchtv/twirp"

	"github.com/thingful/datastore/pkg/rpc/datastore"
)

// Server implements the EncryptedStore interface defined in our proto file.
type Server struct {
	store map[string][][]byte
}

var _ datastore.EncryptedStore = &Server{}

// WriteData is our implementation of the corresponding method on our interface
// that persists data to some store.
func (s *Server) WriteData(ctx context.Context, req *datastore.WriteRequest) (*datastore.WriteResponse, error) {
	if req.PublicKey == "" {
		return nil, twirp.InvalidArgumentError("public_key", "is required to identify the consumer for which data should be stored")
	}

	if s.store[req.PublicKey] == nil {
		s.store[req.PublicKey] = [][]byte{}
	}

	s.store[req.PublicKey] = append(s.store[req.PublicKey], req.Data)

	return &datastore.WriteResponse{}, nil
}

// ReadData is our implementation of the corresponding method on our interface
// that allows data to be read from the store.
func (s *Server) ReadData(ctx context.Context, req *datastore.ReadRequest) (*datastore.ReadResponse, error) {
	if req.PublicKey == "" {
		return nil, twirp.InvalidArgumentError("public_key", "is required to identify the consumer for which data should be returned")
	}

	// is a slice of byte slices or nil
	encryptedEvents := s.store[req.PublicKey]

	if encryptedEvents == nil {
		return nil, twirp.NotFoundError("unable to find data for the requested consumer")
	}

	events := []*datastore.EncryptedEvent{}

	// we aren't really storing data here, so just set time to now, and wrap as a protobuf Timestamp instance
	t, err := ptypes.TimestampProto(time.Now())
	if err != nil {
		return nil, err
	}

	for _, encryptedEvent := range encryptedEvents {
		event := &datastore.EncryptedEvent{
			Timestamp: t,
			Data:      encryptedEvent,
		}
		events = append(events, event)
	}

	return &datastore.ReadResponse{
		PublicKey: req.PublicKey,
		Events:    events,
	}, nil
}
