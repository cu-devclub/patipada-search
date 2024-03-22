package monitoring

import (
	"github.com/prometheus/client_golang/prometheus"
)

//* IF the usecase is bigger, this one can be an object and
//* passed to the usecase but right now we only count the search requests
//* so we make it simple and just use a global variable

var searchCounter prometheus.Counter

func NewMonitoring() {
	searchCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "search_counter",
		Help: "The total number of searche requests",
	})

	prometheus.MustRegister(searchCounter)
}

func GetSearchCounter() prometheus.Counter {
	return searchCounter
}
