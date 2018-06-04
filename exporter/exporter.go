package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	webpagetest "github.com/sticreations/test-go-webpacktest"
)

var (
	metricsLabels = []string{"target", "view"}

	ttfbScore = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "webspeed_ttfb",
		Help: "WebpageTest Time to First Byte for First view in ms",
	}, metricsLabels)

	loadTime = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "webspeed_loadTime",
		Help: "WebpageTest Time for Load in ms",
	}, metricsLabels)

	bytesOut = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "webspeed_loadTime",
		Help: "WebpageTest Bytes Out",
	}, metricsLabels)

	bytesIn = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "webspeed_loadTime",
		Help: "WebpageTest Bytes In",
	}, metricsLabels)

	chromeTimingRedirectStart = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "webspeed_loadTime",
		Help: "WebpageTest Bytes In",
	}, metricsLabels)

	chromeTimingRedirectEnd = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "webspeed_loadTime",
		Help: "WebpageTest Bytes In",
	}, metricsLabels)

	chromeTimingFetchStart = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "webspeed_loadTime",
		Help: "WebpageTest Bytes In",
	}, metricsLabels)

	chromeTimingResponseEnd = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "webspeed_loadTime",
		Help: "WebpageTest Bytes In",
	}, metricsLabels)

	chromeTimingDomLoading = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "webspeed_loadTime",
		Help: "WebpageTest Bytes In",
	}, metricsLabels)

	chromeTimingDomComplete = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "webspeed_loadTime",
		Help: "WebpageTest Bytes In",
	}, metricsLabels)

	chromeTimingFirstPaint = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "webspeed_loadTime",
		Help: "WebpageTest Bytes In",
	}, metricsLabels)

	chromeTimingDomInteractive = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "webspeed_loadTime",
		Help: "WebpageTest Bytes In",
	}, metricsLabels)

	fullyLoaded = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "webspeed_loadTime",
		Help: "WebpageTest Bytes In",
	}, metricsLabels)

	domElements = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "webspeed_loadTime",
		Help: "WebpageTest Bytes In",
	}, metricsLabels)

	requestCount = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "webspeed_loadTime",
		Help: "WebpageTest Bytes In",
	}, metricsLabels)
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(ttfbScore)
	prometheus.MustRegister(loadTime)
	prometheus.MustRegister(bytesOut)
	prometheus.MustRegister(bytesIn)
	prometheus.MustRegister(chromeTimingRedirectStart)
	prometheus.MustRegister(chromeTimingRedirectEnd)
	prometheus.MustRegister(chromeTimingFetchStart)
	prometheus.MustRegister(chromeTimingResponseEnd)
	prometheus.MustRegister(chromeTimingDomLoading)
	prometheus.MustRegister(chromeTimingDomComplete)
	prometheus.MustRegister(chromeTimingFirstPaint)
	prometheus.MustRegister(chromeTimingDomInteractive)
	prometheus.MustRegister(fullyLoaded)
	prometheus.MustRegister(requestCount)

}

func PrometheusMetricsListener(result *webpagetest.ResultData) error {
	for _, run := range result.Runs {
		writePrometheusData(run.FirstView, []string{result.URL, "first"})
		writePrometheusData(run.RepeatView, []string{result.URL, "first"})
	}

	if len(result.Runs) > 0 {
		r := result.Runs["1"].FirstView
		ttfbScore.WithLabelValues(result.URL).Set(float64(r.Steps[0].TTFB))
	}

	//	pageSpeedStaticResourceCount.WithLabelValues(result.Target, string(result.Strategy)).Set(float64(result.PageStats.NumberStaticResources))

	return nil
}

func writePrometheusData(testView webpagetest.TestView, labelValues []string) {
	if len(testView.Steps) > 0 {
		res := testView.Steps[0]
		ttfbScore.WithLabelValues(labelValues...).Set(float64(res.TTFB))
		loadTime.WithLabelValues(labelValues...).Set(float64(res.LoadTime))
		bytesIn.WithLabelValues(labelValues...).Set(float64(res.BytesIn))
		bytesOut.WithLabelValues(labelValues...).Set(float64(res.BytesOut))
		chromeTimingRedirectStart.WithLabelValues(labelValues...).Set(float64(res.ChromeUserTimingRedirectStart))
		chromeTimingRedirectEnd.WithLabelValues(labelValues...).Set(float64(res.ChromeUserTimingRedirectEnd))
		chromeTimingFetchStart.WithLabelValues(labelValues...).Set(float64(res.ChromeUserTimingFetchStart))
		chromeTimingResponseEnd.WithLabelValues(labelValues...).Set(float64(res.ChromeUserTimingResponseEnd))
		chromeTimingDomLoading.WithLabelValues(labelValues...).Set(float64(res.ChromeUserTimingDomLoading))
		chromeTimingDomComplete.WithLabelValues(labelValues...).Set(float64(res.ChromeUserTimingDomComplete))
		chromeTimingFirstPaint.WithLabelValues(labelValues...).Set(float64(res.ChromeUserTimingFirstPaint))
		chromeTimingDomInteractive.WithLabelValues(labelValues...).Set(float64(res.ChromeUserTimingDomInteractive))
		fullyLoaded.WithLabelValues(labelValues...).Set(float64(res.FullyLoaded))
		domElements.WithLabelValues(labelValues...).Set(float64(res.DomElements))
		requestCount.WithLabelValues(labelValues...).Set(float64(res.RequestsFull))

	}
}
