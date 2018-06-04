package main

import (
	"flag"
	"log"
	"os"
	"time"

	"git.evb2.net/medusa/webpagetest-exporter"
)

var (
	wptApiKey       string
	listenerAddress string
	target          string
	checkInterval   time.Duration
)

func main() {
	parseFlags()

	exp := collector.NewCollector(listenerAddress, wptApiKey, target, checkInterval)
	log.Fatal(exp.Start())
}

func parseFlags() {
	flag.StringVar(&wptApiKey, "api-key", getenv("PAGESPEED_API_KEY", ""), "sets the google API key used for pagespeed")
	flag.StringVar(&listenerAddress, "listener", getenv("PAGESPEED_LISTENER", ":9271"), "sets the listener address for the exporters")
	flag.StringVar(&target, "target", getenv("PAGESPEED_TARGETS", ""), "comma separated list of targets to measure")
	intervalFlag := flag.String("interval", getenv("PAGESPEED_INTERVAL", "1h"), "check interval (e.g. 3s 4h 5d ...)")

	flag.Parse()

	if wptApiKey == "" {
		log.Fatal("api key parameter must be specified")
	}

	if duration, err := time.ParseDuration(*intervalFlag); err != nil {
		log.Fatal("could not parse the interval flag '", intervalFlag, "'")
	} else {
		checkInterval = duration
	}
}

func getenv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
