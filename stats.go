package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// nolint:gochecknoglobals
var (
	promNamespace = "radeon_exporter"

	fanInput = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "fan1_input",
		Help:      "fan speed in RPM",
	}, []string{"card_id"})

	fanMax = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "fan1_max",
		Help:      "a maximum value Unit: revolution/max (RPM)",
	}, []string{"card_id"})

	computeFreq = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "freq1_input",
		Help:      "the gfx/compute clock in hertz",
	}, []string{"card_id"})

	memFreq = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "freq2_input",
		Help:      "the memory clock in hertz",
	}, []string{"card_id"})

	dieTemp = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "temp1_input",
		Help:      "the on die GPU temperature in millidegrees Celsius",
	}, []string{"card_id"})

	vddgfx = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "in0_input",
		Help:      "the vddgfx voltage in millivolts",
	}, []string{"card_id"})

	vddnb = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "in1_input",
		Help:      "the vddnb voltage in millivolts",
	}, []string{"card_id"})

	ppt = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "power1_input",
		Help:      "the ppt voltage in microvolts",
	}, []string{"card_id"})

	gpuBusy = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "gpu_busy_percent",
		Help:      "how busy the GPU is as a percentage",
	}, []string{"card_id"})

	memBusy = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "mem_busy_percent",
		Help:      "how busy the VRAM is as a percentage",
	}, []string{"card_id"})

	gttSize = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "mem_info_gtt_size",
		Help:      "total size of the GTT block, in bytes",
	}, []string{"card_id"})

	gttUsed = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "mem_info_gtt_used",
		Help:      "current used size of the GTT block, in bytes",
	}, []string{"card_id"})

	visibleVRAMSize = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "mem_info_vis_vram_size",
		Help:      "total amount of visible VRAM in bytes",
	}, []string{"card_id"})

	visibleVRAMUsed = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "mem_info_vis_vram_used",
		Help:      "currently used visible VRAM in bytes",
	}, []string{"card_id"})

	totalVRAMSize = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "mem_info_vram_size",
		Help:      "total amount of VRAM in bytes",
	}, []string{"card_id"})

	usedVRAM = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "mem_info_vram_used",
		Help:      "currently used VRAM in bytes",
	}, []string{"card_id"})

	// statMap maps filename to prom stat.
	statMap = map[string]*prometheus.GaugeVec{
		"fan1_input":              fanInput,
		"fan1_max":                fanMax,
		"freq1_input":             computeFreq,
		"freq2_input":             memFreq,
		"temp1_input":             dieTemp,
		"in0_input":               vddgfx,
		"in1_input":               vddnb,
		"power1_input":            ppt,
		"device/gpu_busy_percent": gpuBusy,
		"device/mem_busy_percent": memBusy,
		// "device/current_link_speed": memBusy,  HANDLE string vals
		"device/mem_info_gtt_size":      gttSize,
		"device/mem_info_gtt_used":      gttUsed,
		"device/mem_info_vis_vram_size": visibleVRAMSize,
		"device/mem_info_vis_vram_used": visibleVRAMUsed,
		"device/mem_info_vram_size":     totalVRAMSize,
		"device/mem_info_vram_used":     usedVRAM,
	}
)
