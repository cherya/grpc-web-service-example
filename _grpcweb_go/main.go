package _grpcweb_go

import (
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "hello-world/protos"
	"net/http"
)

// Server represents the gRPC server
type Server struct {
}

// SayHello generates response to a HelloRequest request
func (s *Server) SayHello(context context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + in.Name}, nil
}

func corsFunc(origin string) bool {
	return origin == "http://localhost:9000"
}

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterHelloWorldServer(grpcServer, &Server{})

	wrappedGrpc := grpcweb.WrapServer(grpcServer, grpcweb.WithOriginFunc(corsFunc))

	if err := http.ListenAndServe(":9090", wrappedGrpc); err != nil {
		grpclog.Fatalf("failed starting http2 server: %v", err)
	}
}
