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

//go:generate protoc -I ../logmonitorning --go_out=plugins=grpc:../logmonitorning ../logmonitorning/logmonitorning.proto

// Package main implements a server for Logger service.
package main

import (
	"context"
	"log"
	"net"
        "fmt"
        "io/ioutil"
	"google.golang.org/grpc"

)

const (
	port = ":50051"
)

// server is used to implement logmonitorning.LoggerServer.
type server struct {
	pb.UnimplementedLoggerServer
}

// DashBoardLogManagement implements logmonitorning.LoggerServer
func (s *server) DashBoardLogManagement(ctx context.Context, in *pb.LogRequest) (*pb.LogReply, error) {
	log.Printf("Received: %v", in.GetQuery())
        data, err := ioutil.ReadFile("access.log")
        if err != nil {
           fmt.Println("File reading error", err)
        }
        fmt.Println("Contents of file:", string(data))

	return &pb.LogReply{Message: "Hello " +string(data)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLoggerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
