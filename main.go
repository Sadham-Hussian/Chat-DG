package main

import (
	"log"
	"net"
	"os"

	"github.com/Sadham-Hussian/Chat-DG/proto"
	"google.golang.org/grpc"
	glog "google.golang.org/grpc/grpclog"
)

var grpcLog glog.LoggerV2

func init() {
	grpcLog = glog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
}

// Connection handles connection
type Connection struct {
	stream proto.Broadcast_CreateStreamServer
	id     string
	active bool
	error  chan error
}

// Server which holds the list of connections
type Server struct {
	Connection []*Connection
}

// CreateStream connects the newly joined client
func (s *Server) CreateStream(pconn *proto.Connect, stream proto.Broadcast_CreateStreamServer) error {
	conn := &Connection{
		stream: stream,
		id:     pconn.User.Id,
		active: true,
		error:  make(chan error),
	}

	s.Connection = append(s.Connection, conn)

	return <-conn.error
}

func main() {
	var conn []*Connection

	server := &Server{conn}
	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("error creating the server %v", err)
	}

	grpcLog.Info("Starting the server at port :8000")

	proto.RegisterBroadcastServer(grpcServer, server)
	grpcServer.Serve(listener)
}
