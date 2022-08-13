package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/taaraora/hh-stub/pkg/hh"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/protobuf/proto"
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

	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("cannot generate keypair, %s", err.Error())
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

	var lsn uint64 = 1

	for {
		time.Sleep(time.Second * 5)

		req := &hh.ReportSessionsStatsRequest{
			HotspotPubKey: &hh.PublicKey{
				X: priv.PublicKey.X.Bytes(),
				Y: priv.PublicKey.Y.Bytes(),
			},
			AgwSerialNumber: "JHGHGJHKL-8897666-OLOLO",
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
					LocalSequenceNumber:   lsn,
					RecordOpeningTime:     88,
					RecordDuration:        99,
					CauseForRecordClosing: 55,
					RecordType:            12,
					ChargingId:            13,
					RawUserLocationInfo:   nil,
					UeTimezone:            "15",
				},
			},

			Signature: nil,
		}

		r, s, err := signature(req, priv)
		if err != nil {
			panic(err)
		}

		req.Signature = &hh.Signature{
			R: r.Bytes(),
			S: s.Bytes(),
		}

		_, err = client.ReportSessionsStats(ctx, req)
		if err != nil {
			log.Printf("%v", err.Error())
		} else {
			log.Printf("message is sent, lsn: %v", lsn)
		}

		lsn = lsn + 1
	}
}

func signature(tx proto.Message, priv *ecdsa.PrivateKey) (*big.Int, *big.Int, error) {
	txBytes, err := proto.Marshal(tx)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot proto marshal tx to bytes, %s", err.Error())
	}

	hash := sha256.Sum256(txBytes)
	r, s, err := ecdsa.Sign(rand.Reader, priv, hash[:])
	if err != nil {
		return nil, nil, fmt.Errorf("cannot create a signature of tx bytes, %s", err.Error())
	}

	return r, s, nil
}
