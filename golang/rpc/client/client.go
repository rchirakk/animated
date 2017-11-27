package main

import "fmt"
import "log"
import "golang.org/x/net/context"
import "google.golang.org/grpc"
import pb "github.com/rchirakk/animated/golang/rpc/addrBook"

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	aid := 0
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewAddrTxnClient(conn)
	for { 
		aid++
		fmt.Printf("client...agent:000%v\n", aid)
		res, err := client.AddrReq(context.Background(), &pb.HelloReq{
			Hreq: "agent:00" + fmt.Sprintf("%d", aid)})
		if err != nil {
			log.Fatalf("failed to run rpc: %v", err)
		}
		fmt.Printf("agent:00%d got %+v \n", aid, res)
	}
}
