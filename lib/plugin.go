package qframe_collector_docker_events

import (
	"fmt"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"github.com/zpatrick/go-config"
	"github.com/docker/docker/api/types"

	"github.com/qnib/qframe-types"
	"github.com/qnib/qframe-inventory/lib"
)

const (
	version = "0.2.1"
	pluginTyp = qtypes.COLLECTOR
	dockerAPI = "v1.29"
)


type Plugin struct {
	qtypes.Plugin
	engCli *client.Client

}

func New(qChan qtypes.QChan, cfg config.Config, name string) (Plugin, error) {
	var err error
	p := Plugin{
		Plugin: qtypes.NewNamedPlugin(qChan, cfg, pluginTyp, name, version),
	}
	return p, err
}

func (p *Plugin) Run() {
	p.Log("info", fmt.Sprintf("Start docker-events collector v%s", p.Version))
	ctx := context.Background()
	dockerHost := p.CfgStringOr("docker-host", "unix:///var/run/docker.sock")
	// Filter start/stop event of a container
	engineCli, err := client.NewClient(dockerHost, dockerAPI, nil, nil)
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
	// Inventory Init
	inv := qframe_inventory.NewInventory()
	// Fire events for already started containers
	cnts, _ := engineCli.ContainerList(ctx, types.ContainerListOptions{})
	for _, cnt := range cnts {
		cJson, err := engineCli.ContainerInspect(ctx, cnt.ID)
		if err != nil {
			continue
		}
		p.Log("debug", fmt.Sprintf("Already running container %s: SetItem(%s)", cJson.Name, cJson.ID))
		inv.SetItem(cnt.ID, cJson)
	}
	msgs, errs := engineCli.Events(context.Background(), types.EventsOptions{})
	for {
		select {
		case dMsg := <-msgs:
			qm := qtypes.NewQMsg("docker-event", "docker-events")
			qm.Msg = fmt.Sprintf("%s: %s.%s", dMsg.Actor.Attributes["name"], dMsg.Type, dMsg.Action)
			qm.Data = dMsg
			if dMsg.Type == "container" {
				cJson, err := inv.GetItem(dMsg.Actor.ID)
				if err != nil {
					if dMsg.Action == "die" || dMsg.Action == "destroy" {
						p.Log("error", fmt.Sprintf("Container %s just '%s' without having an entry in the Inventory", dMsg.Actor.ID, dMsg.Action))
						continue
					}
					if dMsg.Action == "start" {
						cJson, err := engineCli.ContainerInspect(ctx, dMsg.Actor.ID)
						if err != nil {
							p.Log("error", fmt.Sprintf("Could not inspect '%s'", dMsg.Actor.ID))
							continue
						}
						qm.Data = qtypes.ContainerEvent{
							Event:     dMsg,
							Container: cJson,
						}
						p.Log("debug", fmt.Sprintf("Just started container %s: SetItem(%s)", cJson.Name, cJson.ID))
						inv.SetItem(dMsg.Actor.ID, cJson)
						p.QChan.Data.Send(qm)
						continue
					}
					continue
				}
				p.Log("debug", "Container was found in the inventory...")
				qm.Data = qtypes.ContainerEvent{
					Event:     dMsg,
					Container: cJson,
				}
				p.QChan.Data.Send(qm)
				continue
			}
		case dErr := <-errs:
			if dErr != nil {
				p.Log("error", dErr.Error())
			}
		}
	}
}
