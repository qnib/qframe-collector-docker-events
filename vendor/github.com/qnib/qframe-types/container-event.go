package qtypes

import (
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types"

<<<<<<< HEAD
)
type ContainerEvent struct {
	Base
	Message   string
	Container types.ContainerJSON
	Event events.Message
=======
	"strings"
)
type ContainerEvent struct {
	Base
	Message   	string
	Container 	types.ContainerJSON
	Event 		events.Message
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
}

func NewContainerEvent(base Base, cnt types.ContainerJSON, event events.Message) ContainerEvent {
	return ContainerEvent{
		Base: base,
		Container: cnt,
		Event: event,
	}
}
<<<<<<< HEAD
=======


func (ce *ContainerEvent) GetContainerName() string {
	if ce.Container.Name != "" {
		return strings.Trim(ce.Container.Name, "/")
	} else {
		return "<none>"
	}
}
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
