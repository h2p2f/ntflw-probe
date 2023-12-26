package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"

	"go.uber.org/zap"

	"github.com/h2p2f/ntflw-probe/internal/probe/config"
	"github.com/h2p2f/ntflw-probe/internal/probe/flow"
	"github.com/h2p2f/ntflw-probe/internal/probe/logger"
	"github.com/h2p2f/ntflw-probe/internal/probe/model"
)

func Run(ctx context.Context, sigint <-chan os.Signal, connectionsClosed chan<- struct{}) {

	ctx, cancel := context.WithCancel(ctx)

	conf := config.NewProbeConfig()

	log, err := logger.NewLogger(conf.LogLevel)
	if err != nil {
		cancel()
		return
	}
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			return
		}
	}(log)

	log.Info("Starting probe",
		zap.String("ip", conf.ListenAddr),
		zap.Int("port", conf.ListenPort),
		zap.Int("workers", conf.Workers))

	bufRawFlowChan := make(chan []byte, 10000)
	bufParsingResultChan := make(chan model.Packet, 10000)

	templateMap := flow.NewTemplateMap()

	for i := 0; i < conf.Workers; i++ {
		go flow.ParsePacket(ctx, templateMap, bufRawFlowChan, bufParsingResultChan, log)
	}

	ip := net.ParseIP(conf.ListenAddr)
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   ip,
		Port: conf.ListenPort,
	})
	if err != nil {
		log.Fatal("Error listening UDP", zap.Error(err))
	}

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				buf := make([]byte, 4096)
				_, _, err := conn.ReadFromUDP(buf)
				if err != nil {
					log.Info("Error reading from UDP", zap.Error(err))
				}
				bufRawFlowChan <- buf
			}

		}
	}(ctx)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case packet := <-bufParsingResultChan:
				data, err := json.Marshal(packet)
				if err != nil {
					log.Info("Error marshalling packet", zap.Error(err))
				}
				fmt.Println(string(data))
			}
		}
	}(ctx)

	log.Info("Started")

	<-sigint
	log.Info("Shutting down")

	cancel()
	time.Sleep(3 * time.Second)

	err = conn.Close()
	if err != nil {
		log.Fatal("Error closing UDP connection", zap.Error(err))
	}

	close(connectionsClosed)
	log.Info("probe shutdown gracefully")
}
