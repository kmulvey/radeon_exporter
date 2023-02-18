package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/kmulvey/path"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var deviceNameRegex = regexp.MustCompile("name$")

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

	fmt.Println("started, go to grafana to monitor")

	ticker := time.NewTicker(time.Duration(updateInterval * int(time.Millisecond)))
	for range ticker.C {

		err = collectStats(cards)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func findRadeonDevices() ([]path.Entry, error) {

	var files, err = path.List("/sys/class/hwmon", 2, path.NewRegexEntitiesFilter(deviceNameRegex))
	if err != nil {
		return nil, err
	}

	var results []path.Entry
	for _, file := range files {

		b, err := os.ReadFile(file.AbsolutePath)
		if err != nil {
			return nil, err
		}

		if strings.TrimSpace(string(b)) == "amdgpu" {
			var cardDir, err = path.NewEntry(filepath.Dir(file.AbsolutePath), 0)
			if err != nil {
				return nil, err
			}
			results = append(results, cardDir)
		}
	}

	return results, nil
}

func collectStats(cards []path.Entry) error {

	for _, card := range cards {

		var cardId = filepath.Base(card.AbsolutePath)

		for statName, promStat := range statMap {

			var value, err = parseFileAsFloat(filepath.Join(card.AbsolutePath, statName))
			if err != nil {
				return err
			}

			promStat.WithLabelValues(cardId).Set(value)
		}
	}

	return nil
}

func parseFileAsFloat(file string) (float64, error) {

	b, err := os.ReadFile(file)
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(strings.TrimSpace(string(b)), 64)
}

/*
func parseFileAsString(file string) (string, error) {

	b, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(b)), nil
}
*/
