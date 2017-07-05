package main

import (
	"fmt"
	"golang.org/x/net/context"
	"log"
	"time"

	dt "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/zpatrick/go-config"

	"github.com/qnib/qframe-collector-docker-events/lib"
	"github.com/qnib/qframe-types"
)

const (
	dockerHost = "unix:///var/run/docker.sock"
	dockerAPI  = "v1.29"
)

func Run(qChan qtypes.QChan, cfg *config.Config, name string) {
	p, _ := qframe_collector_docker_events.New(qChan, cfg, name)
	p.Run()
}

func initConfig() (config *container.Config) {
	return &container.Config{Image: "alpine", Volumes: nil, Cmd: []string{"/bin/sleep", "5"}, AttachStdout: false}
}

func hConfig() (config *container.HostConfig) {
	return &container.HostConfig{AutoRemove: true}
}
func startCnt(cli *client.Client, name string) {
	time.Sleep(2 * time.Second)
	// Start container
	create, err := cli.ContainerCreate(context.Background(), initConfig(), hConfig(), nil, name)
	if err != nil {
		fmt.Println(err)
	}
	err = cli.ContainerStart(context.Background(), create.ID, dt.ContainerStartOptions{})
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	qChan := qtypes.NewQChan()
	qChan.Broadcast()
	cfgMap := map[string]string{
		"log.level": "debug",
	}

	cfg := config.NewConfig(
		[]config.Provider{
			config.NewStatic(cfgMap),
		},
	)
	// EngineCli
	engineCli, err := client.NewClient(dockerHost, dockerAPI, nil, nil)
	if err != nil {
		log.Println("Could not connect to /var/run/docker.sock")
	}
	p, err := qframe_collector_docker_events.New(qChan, cfg, "docker-events")
	if err != nil {
		log.Printf("[EE] Failed to create collector: %v", err)
		return
	}
	go p.Run()
	// Start Containers
	cntName := fmt.Sprintf("TestCnt%d", time.Now().Unix())
	go startCnt(engineCli, cntName)

	dc := qChan.Data.Join()
	doStop := false
	for {
		select {
		case msg := <-dc.Read:
			switch msg.(type) {
			case qtypes.QMsg:
				qm := msg.(qtypes.QMsg)
				switch qm.Data.(type) {
				case qtypes.ContainerEvent:
					ce := qm.Data.(qtypes.ContainerEvent)
					if ce.Event.Type == "container" && ce.Event.Action == "start" {
						fmt.Printf("#### Received container.start event for: %s\n", ce.Container.Name)
						doStop = true
					}
				}

			}
		}
		if doStop {
			break
		}
	}
}
