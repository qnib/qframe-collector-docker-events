package main

import (
	"log"
	"fmt"
	"time"

	"github.com/zpatrick/go-config"
	"github.com/qnib/qframe-types"
	"github.com/qnib/qframe-collector-docker-events/lib"
)

func Run(qChan qtypes.QChan, cfg config.Config, name string) {
	p, _ := qframe_collector_docker_events.New(qChan, cfg, name)
	p.Run()
}

func main() {
	qChan := qtypes.NewQChan()
	qChan.Broadcast()
	cfgMap := map[string]string{}

	cfg := config.NewConfig(
		[]config.Provider{
			config.NewStatic(cfgMap),
		},
	)
	p, err := qframe_collector_docker_events.New(qChan, *cfg, "test")
	if err != nil {
		log.Printf("[EE] Failed to create collector: %v", err)
		return
	}
	go p.Run()
	time.Sleep(2*time.Second)
	bg := qChan.Data.Join()
	for {
		qm := bg.Recv().(qtypes.QMsg)
		fmt.Printf("#### Received: %s\n", qm.Msg)
		break
	}
}
