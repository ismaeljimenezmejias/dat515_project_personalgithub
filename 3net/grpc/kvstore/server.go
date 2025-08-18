package kvstore

import (
	"context"

	pb "dat515/3net/grpc/proto"

	"google.golang.org/protobuf/proto"
)

type keyValueServicesServer struct {
	// TODO(student): Add fields if needed
	kv map[string]string
	// this must be included in implementers of the pb.KeyValueServicesServer interface
	pb.UnimplementedKeyValueServiceServer
}

// NewKeyValueServicesServer returns an initialized KeyValueServicesServer
func NewKeyValueServicesServer() *keyValueServicesServer {
	return &keyValueServicesServer{
		kv: make(map[string]string),
	}
}

// Insert inserts a key-value pair from the request into the server's map, and
// returns a response to the client indicating whether or not the insert was successful.
func (s *keyValueServicesServer) Insert(ctx context.Context, req *pb.InsertRequest) (*pb.InsertResponse, error) {
	s.kv[req.GetKey()] = req.GetValue()
	resp := pb.InsertResponse_builder{
		Success: proto.Bool(true),
	}
	return resp.Build(), nil
}

// Lookup returns a response to containing the value corresponding to the request's key.
// If the key is not found, the response's value is empty.
func (s *keyValueServicesServer) Lookup(ctx context.Context, req *pb.LookupRequest) (*pb.LookupResponse, error) {
	// TODO(student): Implement function Lookup

	resp := pb.LookupResponse_builder{
		Value: proto.String("Initial value"),
	}
	return resp.Build(), nil
}

// Keys returns a response to containing a slice of all the keys in the server's map.
// The returned slice is sorted.
func (s *keyValueServicesServer) Keys(ctx context.Context, req *pb.KeysRequest) (*pb.KeysResponse, error) {
	// TODO(student): Implement function Keys

	resp := pb.KeysResponse_builder{
		Keys: []string{"Initial", "value"},
	}
	return resp.Build(), nil
}
