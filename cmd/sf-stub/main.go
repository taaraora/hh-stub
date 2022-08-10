package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/taaraora/hh-stub/pkg/hh"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func main() {
	keepaliveCfg := keepalive.ClientParameters{
		//Time:                cfg.InactiveTimeout,
		//Timeout:             cfg.KeepAliveTimeout,
		//PermitWithoutStream: cfg.PermitWithoutStream,
	}

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(keepaliveCfg),
		//grpc.WithDefaultCallOptions(
		//	grpc.MaxCallRecvMsgSize(cfg.MaxMessageSize),
		//	grpc.MaxCallSendMsgSize(cfg.MaxMessageSize),
		//),
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serverURI := fmt.Sprintf("localhost:%d", 8081)

	conn, err := grpc.DialContext(ctx, serverURI, opts...)
	if err != nil {
		log.Fatalf("grpc dial err: %v", err.Error())
	}

	defer conn.Close()

	log.Printf("connection successful to host %s", serverURI)

	client := hh.NewHeliumHandlerClient(conn)

	req := &hh.ReportSessionsStatsRequest{
		HotspotPubKeyBin: []byte{5, 7, 8, 8},
		AgwSerialNumber:  "JHGHGJHKL-8897666-OLOLO",
		SessionStats: []*hh.SessionStats{
			&hh.SessionStats{
				Sid: "fd",
				Usage: []*hh.SessionStatsUsage{
					&hh.SessionStatsUsage{
						RuleId:    "22",
						BytesTx:   33,
						BytesRx:   5555,
						DroppedTx: 6767,
						DroppedRx: 2018,
					},
				},
				Imsi:                  "we",
				Imei:                  "g",
				Msisdn:                "weg",
				PlmnId:                "er",
				PgwPlmnId:             "re",
				Apnni:                 "hh",
				PdpType:               "lkl",
				UeIp:                  "3",
				SgwIp:                 "3",
				PgwIp:                 "4",
				UeIpv6:                "5",
				SgwIpv6:               "6",
				PgwIpv6:               "7",
				RatType:               100,
				SessionStartTime:      200,
				FinalRecord:           false,
				LocalSequenceNumber:   77,
				RecordOpeningTime:     88,
				RecordDuration:        99,
				CauseForRecordClosing: 55,
				RecordType:            12,
				ChargingId:            13,
				RawUserLocationInfo:   []byte{5, 7, 8, 8},
				UeTimezone:            "15",
			},
		},

		Signature: []byte{5, 7, 8, 8},
	}

	for {
		time.Sleep(time.Second * 5)

		_, err := client.ReportSessionsStats(ctx, req)
		if err != nil {
			log.Printf("%v", err.Error())
		}

	}
}
