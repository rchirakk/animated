package main

import "fmt"
import "log"
import "net"
import "time"
import "golang.org/x/net/context"
import "google.golang.org/grpc"
import pb "github.com/rchirakk/animated/golang/rpc/addrBook"

const (
	port = ":50051"
)

type server struct{}

func (s *server) AddrReq(ctx context.Context, req *pb.HelloReq) (*pb.Person, error) {
	fmt.Printf("addrReq ho ho ho \n")
	time.Sleep(200 * time.Millisecond)
	return &pb.Person{Name:"tumpadi tum tum tum"}, nil
}

func main() {
	fmt.Printf("server...\n")
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAddrTxnServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
