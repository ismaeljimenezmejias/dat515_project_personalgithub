package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "dat515/3net/grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func client(n int, endpoint string) {
	// 1. Conectar con el servidor gRPC
	conn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("No se pudo conectar a %s: %v", endpoint, err)
	}
	defer conn.Close()

	c := pb.NewKeyValueServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 2. Insertar n pares key-value
	fmt.Printf("== Insertando %d pares ==\n", n)
	for i := 1; i <= n; i++ {
		key := fmt.Sprintf("%d", i)
		value := fmt.Sprintf("val-%d", i)

		req := pb.InsertRequest_builder{
			Key:   &key,
			Value: &value,
		}.Build()

		resp, err := c.Insert(ctx, req)
		if err != nil {
			log.Fatalf("Insert fallo para key=%s: %v", key, err)
		}
		fmt.Printf("Insert(%s,%s) => success=%t\n", key, value, resp.GetSuccess())
	}

	// 3. Lookup de n keys
	fmt.Printf("\n== Lookup de %d keys ==\n", n)
	for i := 1; i <= n; i++ {
		key := fmt.Sprintf("%d", i)

		req := pb.LookupRequest_builder{
			Key: &key,
		}.Build()

		resp, err := c.Lookup(ctx, req)
		if err != nil {
			log.Fatalf("Lookup fallo para key=%s: %v", key, err)
		}
		fmt.Printf("Lookup(%s) => %q\n", key, resp.GetValue())
	}

	// (opcional) 4. Pedir Keys y mostrarlas
	fmt.Printf("\n== Keys() ==\n")
	keysResp, err := c.Keys(ctx, pb.KeysRequest_builder{}.Build())
	if err != nil {
		log.Fatalf("Keys fallo: %v", err)
	}
	fmt.Println(keysResp.GetKeys())
}
