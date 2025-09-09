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
	// este campo es requerido por la interfaz generada de gRPC
	pb.UnimplementedKeyValueServiceServer
}

// NewKeyValueServicesServer crea un servidor con el mapa inicializado.
func NewKeyValueServicesServer() *keyValueServicesServer {
	return &keyValueServicesServer{
		kv: make(map[string]string),
	}
}

// Insert intenta insertar una nueva (key,value).
// Si la key ya existe, NO la pisa y devuelve success=false.
// Si no existe, la crea y success=true.
func (s *keyValueServicesServer) Insert(ctx context.Context, req *pb.InsertRequest) (*pb.InsertResponse, error) {
	key := req.GetKey()
	val := req.GetValue()

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.kv[key]; exists {
		resp := pb.InsertResponse_builder{
			Success: proto.Bool(false),
		}
		return resp.Build(), nil
	}

	s.kv[key] = val
	resp := pb.InsertResponse_builder{
		Success: proto.Bool(true),
	}
	return resp.Build(), nil
}

// Lookup devuelve el valor asociado a la key; si no existe, devuelve "".
func (s *keyValueServicesServer) Lookup(ctx context.Context, req *pb.LookupRequest) (*pb.LookupResponse, error) {
	key := req.GetKey()

	s.mu.RLock()
	defer s.mu.RUnlock()

	val := s.kv[key] // devuelve "" si no existe
	resp := pb.LookupResponse_builder{
		Value: proto.String(val),
	}
	return resp.Build(), nil
}

// Keys devuelve todas las keys ordenadas alfabéticamente.
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
