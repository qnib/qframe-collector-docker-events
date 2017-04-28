package qframe_collector_docker_events

import (
	"fmt"
	"log"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"github.com/docker/docker/api/types"
	"github.com/zpatrick/go-config"

	"github.com/qnib/qframe-types"
)

const (
	version = "0.1.2"
	pluginTyp = "collector"
	dockerAPI = "v1.29"
)


type Plugin struct {
	qtypes.Plugin
	engCli *client.Client
	Inventory qtypes.ContainerInventory

}

func New(qChan qtypes.QChan, cfg config.Config, name string) (Plugin, error) {
	var err error
	p := Plugin{
		Plugin: qtypes.NewNamedPlugin(qChan, cfg, pluginTyp, name, version),
	}
	return p, err
}

func (p *Plugin) Run() {
	dockerHost := p.CfgStringOr("docker-host", "unix:///var/run/docker.sock")
	// Filter start/stop event of a container
	engineCli, err := client.NewClient(dockerHost, dockerAPI, nil, nil)
	p.Inventory = qtypes.NewContainerInventory()
	if err != nil {
		p.Log("error", fmt.Sprintf("Could not connect docker/docker/client to '%s': %v", dockerHost, err))
		return
	}
	info, err := engineCli.Info(context.Background())
	if err != nil {
		p.Log("error", fmt.Sprintf("Error during Info(): %v >err> %s", info, err))
		return
	} else {
		p.Log("info", fmt.Sprintf("Connected to '%s' / v'%s'", info.Name, info.ServerVersion))
	}

	msgs, errs := engineCli.Events(context.Background(), types.EventsOptions{})
	for {
		select {
		case dMsg := <-msgs:
			qm := qtypes.NewQMsg("docker-event", "docker-events")
			qm.Msg = fmt.Sprintf("%s: %s.%s", dMsg.Actor.Attributes["name"], dMsg.Type, dMsg.Action)
			if dMsg.Type == "container" {

				cnt, err := p.Inventory.GetCntByID(dMsg.Actor.ID)
				if err != nil {
					cnt, err := engineCli.ContainerInspect(context.Background(), dMsg.Actor.ID)
					if err != nil {
						p.Log("error", fmt.Sprintf("Could not find container '%s': %v", dMsg.Actor.ID, err.Error()))
						continue
					}
					ce := qtypes.ContainerEvent{
						Event:     dMsg,
						Container: cnt,
					}
					p.Inventory.SetCntByEvent(ce)
				}
				if dMsg.Action == "die" || dMsg.Action == "destroy" {
					cnt = types.ContainerJSON{}
				}
				qm.Data = qtypes.ContainerEvent{
					Event:     dMsg,
					Container: cnt,
				}
				p.QChan.Data.Send(qm)
			}
		case dErr := <-errs:
			if dErr != nil {
				log.Printf("[EE] %v", dErr)
			}
		}
	}
}
