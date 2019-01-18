/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

//go:generate protoc -I ../pipeSql --go_out=plugins=grpc:../pipeSql ../pipeSql/pipeSql.proto

package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	// pb "google.golang.org/grpc/examples/helloworld/helloworld"
	pb "Rpc/pipeSql/pipeSql"

	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement pipeSql.GreeterServer.
type server struct{}

// DoSQL implements pipeSql.GreeterServer
func (s *server) DoSQL(ctx context.Context, in *pb.SqlRequest) (*pb.SqlReply, error) {
	log.Printf("Received: %v", in.XXname)
	return &pb.SqlReply{Message: "Handled Request: " + in.XXname}, nil
}
func (s *server) DoSQLAgain(ctx context.Context, in *pb.SqlRequest) (*pb.SqlReply, error) {
	log.Printf("(Again): Received: %v", in.XXname)
	return &pb.SqlReply{Message: "Handled Request Again: " + in.XXname}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
