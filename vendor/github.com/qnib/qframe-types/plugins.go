package qtypes

import (
	"log"
	"fmt"
	"strings"
	"github.com/zpatrick/go-config"
)

const (
	FILTER = "filter"
	COLLECTOR = "collector"
	HANDLER = "handler"
)

type Plugin struct {
	QChan 	QChan
	Cfg 	config.Config
	Typ		string
	Version string
	Name 	string
}

func NewPlugin(qChan QChan, cfg config.Config) Plugin {
	return Plugin{
		QChan: qChan,
		Cfg: cfg,
	}
}

func logStrToInt(level string) int {
	def := 6
	switch level {
	case "error":
		return 3
	case "warn":
		return 4
	case "notice":
		return 5
	case "info":
		return 6
	case "debug":
		return 7
	default:
		return def
	}
}


func (p *Plugin) CfgString(path string) (string, error) {
	res, err := p.Cfg.String(fmt.Sprintf("%s.%s.%s", p.Typ, p.Name, path))
	return res, err
}

func (p *Plugin) CfgStringOr(path, alt string) string {
	res, err := p.CfgString(path)
	if err != nil {
		return alt
	}
	return res
}

func (p *Plugin) CfgInt(path string) (int, error) {
	res, err := p.Cfg.Int(fmt.Sprintf("%s.%s.%s", p.Typ, p.Name, path))
	return res, err
}

func (p *Plugin) CfgIntOr(path string, alt int) int {
	res, err := p.CfgInt(path)
	if err != nil {
		return alt
	}
	return res
}

func (p *Plugin) CfgBool(path string) (bool, error) {
	res, err := p.Cfg.Bool(fmt.Sprintf("%s.%s.%s", p.Typ, p.Name, path))
	return res, err
}

func (p *Plugin) CfgBoolOr(path string, alt bool) bool {
	res, err := p.CfgBool(path)
	if err != nil {
		return alt
	}
	return res
}

func (p *Plugin) GetInputs() []string {
	inStr, err := p.CfgString("inputs")
	if err != nil {
		inStr = ""
	}
	return strings.Split(inStr, ",")
}


func (p *Plugin) Log(logLevel, msg string) {
	// TODO: Setup in each Log() invocation seems rude
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	dL, _ := p.Cfg.StringOr("log.level", "info")
	dI := logStrToInt(dL)
	lI := logStrToInt(logLevel)
	if dI >= lI {
		log.Printf("[%+6s] %s >> %s", strings.ToUpper(logLevel), p.Name, msg)
	}
}

func NewNamedPlugin(qChan QChan, cfg config.Config, typ, name, version string) Plugin {
	p := Plugin{
		QChan: qChan,
		Cfg:   cfg,
	}
	p.Typ = typ
	p.Version = version
	p.Name = name
	return p
}

func (p *Plugin) StartTicker(name string, durMs int) Ticker {
	p.Log("debug", fmt.Sprintf("Start ticker '%s' with duration of %dms", name, durMs))
	ticker := NewTicker(name, durMs)
	go ticker.DispatchTicker(p.QChan)
	return ticker
}