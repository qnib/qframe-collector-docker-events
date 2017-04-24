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
	dc := qChan.Data.Join()
	for {
		select {
		case msg := <- dc.Read:
			switch msg.(type) {
			case qtypes.QMsg:
				qm := msg.(qtypes.QMsg)
				switch qm.Data.(type) {
				case qtypes.ContainerEvent:
					ce := qm.Data.(qtypes.ContainerEvent)
					if ce.Event.Type == "container" && ce.Event.Action == "start" {
						fmt.Printf("#### Received container.start event for: %s\n", ce.Container.Name)
						break
					}
				}

			}
		}

	}
}
