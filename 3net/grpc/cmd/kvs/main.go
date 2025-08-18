package main

import (
	"flag"
	"log"
	"net"

	"dat515/3net/grpc/kvstore"
	pb "dat515/3net/grpc/proto"

	"google.golang.org/grpc"
)

func main() {
	var (
		server   = flag.Bool("server", false, "Start gRPC server if true; otherwise start the client")
		n        = flag.Int("n", 25, "Number of key-value pairs to insert and lookup")
		endpoint = flag.String("endpoint", "localhost:12110", "Endpoint on which server runs or to which the client connects")
	)
	flag.Parse()

	if *server {
		listener, err := net.Listen("tcp", *endpoint)
		if err != nil {
			log.Fatalf("Failed to listen on %v: %v", *endpoint, err)
		}
		log.Printf("Listener started on %v\n", *endpoint)

		srv := kvstore.NewKeyValueServicesServer()
		grpcServer := grpc.NewServer()
		pb.RegisterKeyValueServiceServer(grpcServer, srv)
		log.Println("Preparing to serve incoming requests.")
		log.Fatal(grpcServer.Serve(listener))
	} else {
		client(*n, *endpoint)
	}
}
