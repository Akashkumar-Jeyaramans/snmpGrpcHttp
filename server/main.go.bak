package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	pb "github.com/Akashkumar-Jeyaramans/snmpGrpcHttp/v1/commands"
	"github.com/gosnmp/gosnmp"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/metadata"
)

// type snmpServer struct {
// 	pb.UnimplementedCommandServer
// }

type Logger interface{}

// CommandServer is the GRPC Service for snmp.Command
type CommandServer struct {
	pb.UnimplementedCommandServer
	ListenAddr       string
	DefaultVersion   gosnmp.SnmpVersion
	DefaultCommunity string
	Logger           Logger
	SNMPLogger       gosnmp.Logger
}

func (s *CommandServer) Get(ctx context.Context, oids *pb.OidList) (*pb.SnmpPacket, error) {

	md, _ := metadata.FromIncomingContext(ctx)
	fmt.Printf("md : %v", md)
	snmp, err := s.snmpConnectionFromMetadata(md)
	if err != nil {
		log.Printf("[ERR] %v", err)
		return nil, err
	}

	err = snmp.Connect()
	if err != nil {
		return nil, fmt.Errorf("ERR_SNMP_CONN: %v", err)
	}
	defer snmp.Conn.Close()

	packet, err := snmp.Get(oids.Oids)
	if err != nil {
		return nil, fmt.Errorf("ERR_SNMP_GET: %v", err)
	}

	vars := packet.Variables
	pbVars := make([]*pb.SnmpPDU, len(vars))
	for i, v := range vars {
		pbVars[i] = ToPbSnmpPDU(v)
	}

	return &pb.SnmpPacket{
		Error:    uint32(packet.ErrorIndex),
		Variable: pbVars,
	}, nil
}

func main() {
	mux := runtime.NewServeMux()

	err := pb.RegisterCommandHandlerServer(context.Background(), mux, &CommandServer{})
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

func (s *CommandServer) snmpConnectionFromMetadata(md metadata.MD) (snmp *gosnmp.GoSNMP, err error) {
	mdTarget := md.Get("10.6.50.8")
	if mdTarget == nil {
		return nil, fmt.Errorf("snmp-target not given %v", mdTarget)
	}
	target := mdTarget[0]

	mdPort := md.Get("161")
	port := gosnmp.Default.Port
	if mdPort != nil {
		iport, err := strconv.Atoi(mdPort[0])
		if err != nil {
			return nil, fmt.Errorf("invalid snmp-port: %s", mdPort[0])
		}
		port = uint16(iport)
	}

	community := s.DefaultCommunity
	mdCommunity := md.Get("public")
	if mdCommunity != nil {
		community = mdCommunity[0]
	}

	version := s.DefaultVersion
	mdVersion := md.Get("1")
	if mdVersion != nil {
		switch mdVersion[0] {
		case "1":
			version = gosnmp.Version1
		case "2", "2c":
			version = gosnmp.Version2c
		case "3":
			return nil, fmt.Errorf("snmp-version 3 not supported yet")
		default:
			return nil, fmt.Errorf("invalid snmp-version; %s", mdVersion[0])
		}
	}

	return &gosnmp.GoSNMP{
		Target:    target,
		Port:      port,
		Community: community,
		Version:   version,
		Timeout:   gosnmp.Default.Timeout,
		Retries:   gosnmp.Default.Retries,
		MaxOids:   gosnmp.Default.MaxOids,
		Logger:    s.SNMPLogger,
	}, nil
}

// ToPbSnmpPDU helper comverts SnmpPDU to protobuf SnmpPDU
func ToPbSnmpPDU(pdu gosnmp.SnmpPDU) (pbSnmpPdu *pb.SnmpPDU) {
	pbSnmpPdu = &pb.SnmpPDU{
		Name: pdu.Name,
		Type: pb.Asn1BER(pdu.Type),
	}

	switch pdu.Type {
	case gosnmp.OctetString:
		pbSnmpPdu.Value = &pb.SnmpPDU_Str{
			Str: string(pdu.Value.([]byte)),
		}
	case gosnmp.Gauge32:
		fallthrough
	case gosnmp.Counter32:
		fallthrough
	case gosnmp.TimeTicks:
		var val uint32
		var ok bool
		if val, ok = pdu.Value.(uint32); !ok {
			val = uint32(pdu.Value.(uint))
		}
		pbSnmpPdu.Value = &pb.SnmpPDU_UI32{
			UI32: val,
		}
	case gosnmp.Counter64:
		var val uint64
		var ok bool
		if val, ok = pdu.Value.(uint64); !ok {
			val = uint64(pdu.Value.(uint))
		}
		pbSnmpPdu.Value = &pb.SnmpPDU_UI64{
			UI64: val,
		}
	case gosnmp.Integer:
		pbSnmpPdu.Value = &pb.SnmpPDU_I32{
			I32: int32(pdu.Value.(int)),
		}
	case gosnmp.ObjectIdentifier:
		pbSnmpPdu.Value = &pb.SnmpPDU_Str{
			Str: pdu.Value.(string),
		}
	}

	return pbSnmpPdu
}

// ToGoSnmpPDU helper converts protobuf snmppdu to go snmpPDU
func ToGoSnmpPDU(pbSnmpPdu *pb.SnmpPDU) (pdu gosnmp.SnmpPDU) {
	pdu = gosnmp.SnmpPDU{
		Name: pbSnmpPdu.Name,
		Type: gosnmp.Asn1BER(pbSnmpPdu.Type),
	}

	switch pbSnmpPdu.Type {
	case pb.Asn1BER_OctetString:
		pdu.Value = pbSnmpPdu.GetStr()
	case pb.Asn1BER_TimeTicks:
		pdu.Value = pbSnmpPdu.GetUI64()
	case pb.Asn1BER_Integer:
		pdu.Value = pbSnmpPdu.GetI32()
	}

	return pdu
}
