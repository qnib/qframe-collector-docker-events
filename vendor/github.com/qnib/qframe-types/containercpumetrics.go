package qtypes

import (
	"strconv"
<<<<<<< HEAD
	"strings"
	"github.com/elastic/beats/libbeat/common"
	"github.com/docker/docker/api/types"
	dc "github.com/fsouza/go-dockerclient"
	"fmt"
=======
	"github.com/elastic/beats/libbeat/common"
	"github.com/docker/docker/api/types"
	dc "github.com/fsouza/go-dockerclient"
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
)

// Inspired by https://github.com/elastic/beats/blob/master/metricbeat/module/docker/cpu/helper.go
type CPUStats struct {
	Base
	Container                   *types.Container
	PerCpuUsage                 common.MapStr
	TotalUsage                  float64
	UsageInKernelmode           uint64
	UsageInKernelmodePercentage float64
	UsageInUsermode             uint64
	UsageInUsermodePercentage   float64
	SystemUsage                 uint64
	SystemUsagePercentage       float64
}

func NewCPUStats(src Base, stats *dc.Stats) *CPUStats {
	return &CPUStats{
		Base: src,
		PerCpuUsage: perCpuUsage(stats),
		TotalUsage: totalUsage(stats),
		UsageInKernelmode: stats.CPUStats.CPUUsage.UsageInKernelmode,
		UsageInKernelmodePercentage: usageInKernelmode(stats),
		UsageInUsermode: stats.CPUStats.CPUUsage.UsageInUsermode,
		UsageInUsermodePercentage: usageInUsermode(stats),
		SystemUsage: stats.CPUStats.SystemCPUUsage,
		SystemUsagePercentage: systemUsage(stats),
	}
}

<<<<<<< HEAD
func (cs *CPUStats) ToMetrics(src string) []Metric {
	dim := map[string]string{
		"container_id": cs.Container.ID,
		"container_name": strings.Trim(cs.Container.Names[0], "/"),
		"image_name": cs.Container.Image,
		"command": strings.Replace(cs.Container.Command, " ", "#", -1),
		"created": fmt.Sprintf("%d", cs.Container.Created),
	}
	for k, v := range cs.Container.Labels {
		dv := strings.Replace(v, " ", "#", -1)
		dv = strings.Replace(v, ".", "_", -1)
		dim[k] = dv
	}
=======

func (cs *CPUStats) ToMetrics(src string) []Metric {
	dim := AssembleDefaultDimensions(cs.Container)
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	return []Metric{
		cs.NewExtMetric(src, "cpu.usage.kernel.percent", Gauge, cs.UsageInKernelmodePercentage, dim, cs.Time, true),
		cs.NewExtMetric(src, "cpu.usage.user.percent", Gauge, cs.UsageInUsermodePercentage, dim, cs.Time, true),
		cs.NewExtMetric(src, "cpu.system.usage.percent", Gauge, cs.SystemUsagePercentage, dim, cs.Time, true),
	}
}


func perCpuUsage(stats *dc.Stats) common.MapStr {
	var output common.MapStr
	if len(stats.CPUStats.CPUUsage.PercpuUsage) == len(stats.PreCPUStats.CPUUsage.PercpuUsage) {
		output = common.MapStr{}
		for index := range stats.CPUStats.CPUUsage.PercpuUsage {
			cpu := common.MapStr{}
			cpu["pct"] = calculateLoad(stats.CPUStats.CPUUsage.PercpuUsage[index] - stats.PreCPUStats.CPUUsage.PercpuUsage[index])
			cpu["ticks"] = stats.CPUStats.CPUUsage.PercpuUsage[index]
			output[strconv.Itoa(index)] = cpu
		}
	}
	return output
}

func totalUsage(stats *dc.Stats) float64 {
	return calculateLoad(stats.CPUStats.CPUUsage.TotalUsage - stats.PreCPUStats.CPUUsage.TotalUsage)
}

<<<<<<< HEAD
func usageInKernelmode(stats *dc.Stats) float64 {
	return calculateLoad(stats.CPUStats.CPUUsage.UsageInKernelmode - stats.PreCPUStats.CPUUsage.UsageInKernelmode)
}

func usageInUsermode(stats *dc.Stats) float64 {
	return calculateLoad(stats.CPUStats.CPUUsage.UsageInUsermode - stats.PreCPUStats.CPUUsage.UsageInUsermode)
}

func systemUsage(stats *dc.Stats) float64 {
=======
func usageInKernelmode(stats *dc.Stats) (val float64) {
	if stats.PreCPUStats.CPUUsage.UsageInKernelmode == 0 {
		return
	}
	return calculateLoad(stats.CPUStats.CPUUsage.UsageInKernelmode - stats.PreCPUStats.CPUUsage.UsageInKernelmode)
}

func usageInUsermode(stats *dc.Stats) (val float64) {
	if stats.PreCPUStats.CPUUsage.UsageInUsermode == 0 {
		return
	}
	return calculateLoad(stats.CPUStats.CPUUsage.UsageInUsermode - stats.PreCPUStats.CPUUsage.UsageInUsermode)
}

func systemUsage(stats *dc.Stats) (val float64) {
	if stats.PreCPUStats.SystemCPUUsage == 0 {
		return
	}
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	return calculateLoad(stats.CPUStats.SystemCPUUsage - stats.PreCPUStats.SystemCPUUsage)
}

func calculateLoad(value uint64) float64 {
	return float64(value) / float64(1000000000)
}

// \beats

