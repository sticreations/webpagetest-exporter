package collector

import (
	"net/http"
	"reflect"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	webpagetest "github.com/sticreations/go-webpagetest"
	"github.com/sticreations/webpagetest-exporter/exporter"
)

type ResultListener func(*webpagetest.ResultData) error

type PagespeedCollector struct {
	listenerAddress string
	target          string
	checkInterval   time.Duration
	resultChannel   chan *webpagetest.ResultData
	resultListeners []ResultListener
	apiKey          string
}

func NewCollector(listenerAddress, wptApiKey string, target string, checkInterval time.Duration) *PagespeedCollector {
	return &PagespeedCollector{
		listenerAddress: listenerAddress,
		target:          target,
		checkInterval:   checkInterval,
		resultChannel:   make(chan *webpagetest.ResultData, 1),
		resultListeners: []ResultListener{},
		apiKey:          wptApiKey,
	}
}

func (e *PagespeedCollector) Start() error {
	startupMessage := "starting pagespeed exporter on listener %s for %v target(s) with re-check interval of %s"
	log.Infof(startupMessage, e.listenerAddress, e.target, e.checkInterval)

	s := &http.Server{
		Addr: e.listenerAddress,
	}

	// Register prometheus handler
	http.Handle("/metrics", promhttp.Handler())

	// Register prometheus metrics resultListeners
	e.registerListener(exporter.PrometheusMetricsListener)

	go e.watch()
	go e.collect()

	return s.ListenAndServe()
}

func (e *PagespeedCollector) registerListener(listener ResultListener) {
	e.resultListeners = append(e.resultListeners, listener)
}

func (e *PagespeedCollector) watch() {
	wpt, err := webpagetest.NewClient("https://webpagetest.org")
	if err != nil {
		log.Error("Could not create WPT Client")
	}
	for true {
		params := &webpagetest.TestSettings{
			URL:    e.target,
			Runs:   1,
			APIKey: e.apiKey,
		}
		rs, err := wpt.RunTestAndWait(*params, func(testId, status string, duration int) {
			log.Printf("The testId is %v, STATUS: %v, Time wasted: %v Seconds", testId, status, duration)
		})

		if err != nil {
			log.Errorf("Could not get Testdata: %v", err)
		}

		e.resultChannel <- rs
		time.Sleep(e.checkInterval)
	}
}

func (e *PagespeedCollector) collect() {
	for {
		select {
		case res := <-e.resultChannel:
			e.handleResult(res)
		}
	}
}

func (e *PagespeedCollector) handleResult(result *webpagetest.ResultData) {
	for _, l := range e.resultListeners {
		err := l(result)
		if err != nil {
			log.Error("listener " + reflect.TypeOf(l).String() + " thew an error while processing a result for target ")
		}
	}
}
