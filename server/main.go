package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "github.com/Akashkumar-Jeyaramans/snmpGrpcHttp/v1/commands"
	g "github.com/gosnmp/gosnmp"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type getServer struct {
	pb.UnimplementedGetCommandServer
}

func (s getServer) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	// Default is a pointer to a GoSNMP struct that contains sensible defaults
	// eg port 161, community public, etc
	fmt.Printf("ipaddress %v", in.Target)
	fmt.Printf("oid %v", in.Oid)

	g.Default.Target = in.Target
	err := g.Default.Connect()
	if err != nil {
		log.Fatalf("Connect() err: %v", err)
	}
	defer g.Default.Conn.Close()

	oids := []string{in.Oid}
	result, err2 := g.Default.Get(oids) // Get() accepts up to g.MAX_OIDS
	if err2 != nil {
		log.Fatalf("Get() err: %v", err2)
	}
	var z string
	var k string
	for i, variable := range result.Variables {
		fmt.Printf("%d: oid: %s ", i, variable.Name)
		z = variable.Name
		switch variable.Type {
		case g.OctetString:
			fmt.Printf("string: %s\n", string(variable.Value.([]byte)))
			k = string(variable.Value.([]byte))
		default:
			fmt.Printf("number: %d\n", g.ToBigInt(variable.Value))
		}
	}
	return &pb.GetResponse{
		Name:  z,
		Value: k,
	}, nil
}

func main() {
	mux := runtime.NewServeMux()

	err := pb.RegisterGetCommandHandlerServer(context.Background(), mux, &getServer{})
	if err != nil {
		log.Fatal(err)
	}

	server := http.Server{
		Handler: mux,
	}
	// creating a listener for server
	l, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("server listening : %v", l)
	}
	// start server
	err = server.Serve(l)
	if err != nil {
		log.Fatal(err)
	}

}
