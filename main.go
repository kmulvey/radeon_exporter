package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/kmulvey/path"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.szostok.io/version"
	"go.szostok.io/version/printer"
)

// https://docs.kernel.org/gpu/amdgpu/thermal.html
// https://www.kernel.org/doc/html/v5.6/hwmon/lochnagar.html
// https://kernel.org/doc/html/latest/gpu/amdgpu/driver-misc.html#mem-info-vram-total

func main() {

	var addr string
	var v, h bool
	var updateInterval int
	flag.StringVar(&addr, "port", ":9200", "address to bind to for metrics server")
	flag.IntVar(&updateInterval, "update-interval", 1000, "how often to collect metrics (in milliseconds)")
	flag.BoolVar(&v, "version", false, "print version")
	flag.BoolVar(&v, "v", false, "print version")
	flag.BoolVar(&h, "help", false, "print options")
	flag.Parse()

	if h {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if v {
		var verPrinter = printer.New()
		var info = version.Get()
		if err := verPrinter.PrintInfo(os.Stdout, info); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	go func() {
		http.Handle("/metrics", promhttp.Handler())

		var server = &http.Server{
			Addr:         addr,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}

		if err := server.ListenAndServe(); err != nil {
			log.Fatal("http server error: ", err)
		}
	}()

	var cards, err = findRadeonDevices()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("started, go to Grafana to monitor")

	ticker := time.NewTicker(time.Duration(updateInterval * int(time.Millisecond)))
	for range ticker.C {
		err = collectStats(statMap, cards) // statMap is defined in stats.go
		if err != nil {
			log.Fatal(err)
		}
	}
}

// findRadeonDevices searches hwmon to find AMD GPUs.
func findRadeonDevices() ([]path.Entry, error) {

	var dirs, err = path.List("/sys/class/hwmon", 1, false)
	if err != nil {
		return nil, fmt.Errorf("failed to list hwmon directories: %w", err)
	}

	var files []path.Entry
	for _, dir := range dirs {
		if entry, err := path.NewEntry(filepath.Join(dir.String(), "name"), 0); err == nil {
			files = append(files, entry)
		}
	}

	var results []path.Entry
	for _, file := range files {
		b, err := os.ReadFile(file.AbsolutePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read file %s: %w", file.AbsolutePath, err)
		}

		if strings.TrimSpace(string(b)) == "amdgpu" {
			var cardDir, err = path.NewEntry(filepath.Dir(file.AbsolutePath), 0)
			if err != nil {
				return nil, fmt.Errorf("failed to create path entry for %s: %w", file.AbsolutePath, err)
			}
			results = append(results, cardDir)
		}
	}

	return results, nil
}

// collectStats is given a map of hwmon files => Prometheus stat as well as an array of GPUs.
// It then calls parseFileAsFloat to get the value and publishes the stat.
func collectStats(statMap map[string]*prometheus.GaugeVec, cards []path.Entry) error {

	for _, card := range cards {
		var cardID = filepath.Base(card.AbsolutePath)

		for statName, promStat := range statMap {
			var value, err = parseFileAsFloat(filepath.Join(card.AbsolutePath, statName))
			if err != nil {
				return err
			}
			promStat.WithLabelValues(cardID).Set(value)
		}
	}

	return nil
}

// parseFileAsFloat reads a given hwmon file and returns its value as a float64.
func parseFileAsFloat(file string) (float64, error) {

	var b, err = os.ReadFile(file)
	if err != nil {
		//nolint:nilerr
		return 0, fmt.Errorf("failed to read file %s: %w", file, err)
	}

	value, err := strconv.ParseFloat(strings.TrimSpace(string(b)), 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse float from file %s: %w", file, err)
	}

	return value, nil
}
