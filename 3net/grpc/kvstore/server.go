package kvstore

import (
	"context"
	"sort"
	"sync"

	pb "dat515/3net/grpc/proto"
	"google.golang.org/protobuf/proto"
)

type keyValueServicesServer struct {
	mu sync.RWMutex
	kv map[string]string
	pb.UnimplementedKeyValueServiceServer
}

func NewKeyValueServicesServer() *keyValueServicesServer {
	return &keyValueServicesServer{kv: make(map[string]string)}
}

func (s *keyValueServicesServer) Insert(ctx context.Context, req *pb.InsertRequest) (*pb.InsertResponse, error) {
	key := req.GetKey()
	val := req.GetValue()

	s.mu.Lock()
	s.kv[key] = val // OVERWRITE SIEMPRE (lo que piden los tests)
	s.mu.Unlock()

	resp := pb.InsertResponse_builder{
		Success: proto.Bool(true),
	}
	return resp.Build(), nil
}

func (s *keyValueServicesServer) Lookup(ctx context.Context, req *pb.LookupRequest) (*pb.LookupResponse, error) {
	key := req.GetKey()

	s.mu.RLock()
	val := s.kv[key] // "" si no existe
	s.mu.RUnlock()

	resp := pb.LookupResponse_builder{
		Value: proto.String(val),
	}
	return resp.Build(), nil
}

func (s *keyValueServicesServer) Keys(ctx context.Context, req *pb.KeysRequest) (*pb.KeysResponse, error) {
	s.mu.RLock()
	keys := make([]string, 0, len(s.kv))
	for k := range s.kv {
		keys = append(keys, k)
	}
	s.mu.RUnlock()

	sort.Strings(keys)

	resp := pb.KeysResponse_builder{
		Keys: keys,
	}
	return resp.Build(), nil
}
