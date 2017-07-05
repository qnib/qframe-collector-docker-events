package qtypes


import (
<<<<<<< HEAD
	"strings"
	"github.com/docker/docker/api/types"
	dc "github.com/fsouza/go-dockerclient"
	"fmt"
=======
	"github.com/docker/docker/api/types"
	dc "github.com/fsouza/go-dockerclient"
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	"math"
)

// Inspired by https://github.com/elastic/beats/blob/master/metricbeat/module/docker/cpu/helper.go
type MemoryStats struct {
	Base
	Container   *types.Container
	Failcnt   	float64
	Limit     	float64
	MaxUsage  	float64
	TotalRss  	float64
	TotalRssP 	float64
	Usage     	float64
	UsageP 		float64
}

func NewMemoryStats(src Base, stats *dc.Stats) MemoryStats {
	return MemoryStats{
		Base:      src,
		Failcnt:   float64(stats.MemoryStats.Failcnt),
		Limit:     float64(stats.MemoryStats.Limit),
		MaxUsage:  float64(stats.MemoryStats.MaxUsage),
		TotalRss:  float64(stats.MemoryStats.Stats.TotalRss),
		TotalRssP: calcUsage(float64(stats.MemoryStats.Stats.TotalRss), float64(stats.MemoryStats.Limit)),
		Usage:     float64(stats.MemoryStats.Usage),
		UsageP:    calcUsage(float64(stats.MemoryStats.Usage), float64(stats.MemoryStats.Limit)),
	}
}

func (ms *MemoryStats) ToMetrics(src string) []Metric {
<<<<<<< HEAD
	dim := map[string]string{
		"container_id": ms.Container.ID,
		"container_name": strings.Trim(ms.Container.Names[0], "/"),
		"image_name": ms.Container.Image,
		"command": strings.Replace(ms.Container.Command, " ", "#", -1),
		"created": fmt.Sprintf("%d", ms.Container.Created),
	}
	for k, v := range ms.Container.Labels {
		dv := strings.Replace(v, " ", "#", -1)
		dv = strings.Replace(v, ".", "_", -1)
		dim[k] = dv
	}
=======
	dim := AssembleDefaultDimensions(ms.Container)
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	return []Metric{
		ms.NewExtMetric(src, "memory.usage.percent", Gauge, ms.UsageP, dim, ms.Time, true),
		ms.NewExtMetric(src, "memory.total_rss.percent", Gauge, ms.TotalRssP, dim, ms.Time, true),
		ms.NewExtMetric(src, "memory.total_rss.bytes", Gauge, ms.TotalRss, dim, ms.Time, true),
		ms.NewExtMetric(src, "memory.usage.bytes", Gauge, ms.Usage, dim, ms.Time, true),
		ms.NewExtMetric(src, "memory.failcnt", Gauge, ms.Failcnt, dim, ms.Time, true),
		ms.NewExtMetric(src, "memory.limit.bytes", Gauge, ms.Limit, dim, ms.Time, true),
	}
}

func calcUsage(frac, all float64) float64 {
	v := float64(frac / all)
	if math.IsNaN(v) {
		v = 0.0
	}
	return v
}
