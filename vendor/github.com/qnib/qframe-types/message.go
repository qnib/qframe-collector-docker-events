package qtypes

import (
	"strings"
	"github.com/docker/docker/api/types"
<<<<<<< HEAD
=======
	"github.com/qnib/qframe-utils"
	"fmt"
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
)

const (
	MsgCEE = "cee"
	MsgTCP = "tcp"
<<<<<<< HEAD
=======
	MsgDLOG = "docker-log"
	MsgMetric = "metric" //needs to have name,time and value field ; optional tags (key1=val1,key2=val2)
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
)


type Message struct {
	Base
	Container   types.ContainerJSON
	Name       	string            	`json:"name"`
<<<<<<< HEAD
	LogLevel       string				`json:"loglevel"`
=======
	LogLevel    string				`json:"loglevel"`
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	MessageType	string            	`json:"type"`
	Message     string            	`json:"value"`
	KV			map[string]string 	`json:"data"`
}

func NewMessage(base Base, name, mType, msg string) Message {
<<<<<<< HEAD
	return Message{
=======
	m := Message{
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
		Base: base,
		Name: name,
		Container: types.ContainerJSON{},
		LogLevel: "INFO",
		MessageType: mType,
		Message: msg,
		KV: map[string]string{},
	}
<<<<<<< HEAD
=======
	m.SourceID = int(qutils.GetGID())
	return m
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
}

func NewContainerMessage(base Base, cnt types.ContainerJSON, name, mType, msg string) Message {
	m := NewMessage(base, name, mType, msg)
	m.Container = cnt
<<<<<<< HEAD
	return m
}

=======
	m.ID = m.GenContainerMsgID()
	return m
}

// GenContainerMsgID uses "<container_id>-<time.UnixNano()>-<MSG>" and does a sha1 hash.
func (m *Message) GenContainerMsgID() string {
	s := fmt.Sprintf("%s-%d-%s", m.Container.ID, m.Time.UnixNano(), m.Message)
	return Sha1HashString(s)
}

>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
func (m *Message) GetContainerName() string {
	if m.Container.Name != "" {
		return strings.Trim(m.Container.Name, "/")
	} else {
		return "<none>"
	}
<<<<<<< HEAD
}
=======
}
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
