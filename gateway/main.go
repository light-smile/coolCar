package main

import (
	"context"
	authpb "coolCar/service/auth/api/gen/v1"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
			runtime.MIMEWildcard,&runtime.JSONPb{
				EnumsAsInts: true,
				OrigName: true,
			},
		))
	err := authpb.RegisterAuthServiceHandlerFromEndpoint(
			c, mux, "localhost:8081",
			[]grpc.DialOption{grpc.WithInsecure()},
		)
	if err != nil {
		log.Fatalf("cannot register auth service: %v", err)
	}


	log.Fatal(http.ListenAndServe(":8080",mux))
}
