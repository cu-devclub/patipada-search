package monitoring

import (
	"github.com/prometheus/client_golang/prometheus"
)

//* IF the usecase is bigger, this one can be an object and
//* passed to the usecase but right now we only count the search requests
//* so we make it simple and just use a global variable

var searchCounter struct {
	draftCounter   prometheus.Counter
	confirmCounter prometheus.Counter
}

func NewMonitoring() {

	searchDraftCounter := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "search_draft_counter",
		Help: "The total number of search requests with status draft",
	})

	searchConfirmCounter := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "search_confirm_counter",
		Help: "The total number of search requests with status confirm",
	})

	prometheus.MustRegister(searchDraftCounter)

	prometheus.MustRegister(searchConfirmCounter)

	searchCounter.draftCounter = searchDraftCounter
	searchCounter.confirmCounter = searchConfirmCounter
}

func MonitoringSearch(searchStatus string) {
	switch searchStatus {
	case "draft":
		searchCounter.draftCounter.Inc()
	case "confirm":
		searchCounter.confirmCounter.Inc()
	}
}
