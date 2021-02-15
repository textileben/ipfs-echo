package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type ipfsEchoRequest struct {
	Msg        string
	Status     string
	StartedAt  time.Time
	FinishedAt time.Time
	Duration   float64
}

func NewIpfsEcho(msg string) *ipfsEchoRequest {
	return &ipfsEchoRequest{
		Msg: msg,
	}
}

func (iR *ipfsEchoRequest) Started() {
	iR.StartedAt = time.Now()
}

func (iR *ipfsEchoRequest) Finished() {
	iR.FinishedAt = time.Now()
	iR.Duration = time.Since(iR.StartedAt).Seconds()
}

//Service implements UseCase interface
type Service struct {
	ipfsEchoRequestHistogram *prometheus.HistogramVec
}

var (
	IpfsEchoAttempts = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ipfsecho_attempts",
			Help: "IPFS Echo attemps"},
		[]string{"type", "status"})
	IpfsEchoHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "ipfsecho_response_histogram",
			Help: "IPFS Echo response histogram"},
		[]string{"type", "status"})
)

func init() {
	prometheus.MustRegister(IpfsEchoAttempts)
	prometheus.MustRegister(IpfsEchoHistogram)

}
