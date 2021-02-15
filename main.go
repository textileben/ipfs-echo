package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/textileben/ipfs-echo/metrics"
	"github.com/textileben/ipfs-echo/version"
	"go.uber.org/zap"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	listenAddress = kingpin.Flag(
		"web.listen-address",
		"Address on which to expose metrics and web interface.",
	).Default(":9100").String()
	remoteAPI = kingpin.Flag(
		"remote.API",
		"Network address of remote IPFS API endpoint",
	).Required().String()
	localAPI = kingpin.Flag(
		"local.API",
		"Network address of local IPFS API endpoint",
	).Default("localhost:5001").String()
	versionFlag = kingpin.Flag(
		"version",
		"Display binary version.",
	).Default("False").Bool()
)

func main() {
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	if *versionFlag {
		fmt.Printf("(version=%s, gitcommit=%s)\n", version.Version, version.GitCommit)
		fmt.Printf("(go=%s, user=%s, date=%s)\n", version.GoVersion, version.BuildUser, version.BuildDate)
		os.Exit(0)
	}
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("config", zap.String("listenAddress", *listenAddress), zap.String("remoteAPI", *remoteAPI))
	logger.Info("starting ipfs-echo", zap.String("version", version.Info()))

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
		<head><title>ipfs-echo exporter</title></head>
		<body>
		<h1>ipfs-echo exporter</h1>
		<p><a href="/metrics">Metrics</a></p>
		</body>
		</html>`))
	})
	go runEcho()
	if err := http.ListenAndServe(*listenAddress, nil); err != nil {
		logger.Fatal("Failed to listen", zap.Error(err))
	}

}

func runEcho() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	for {
		msg := RandomString(24)
		logger.Info(
			"Generated a message",
			zap.String("msg", msg),
		)

		rr := metrics.NewIpfsEcho(msg)
		rr.Started()
		remoteShell := shell.NewShellWithClient(*remoteAPI, &http.Client{Timeout: 5 * time.Second})

		cid, err := remoteShell.AddWithOpts(strings.NewReader(msg), false, false) // Don't pin, don't use raw leaves
		if err != nil {
			logger.Error(
				"Failed to add the msg to the remote",
				zap.Error(err),
			)
			metrics.IpfsEchoAttempts.WithLabelValues(
				"remote", "fail",
			).Add(1.0)
			rr.Status = "failed"
		} else {
			logger.Info(
				"Successfully added the msg to the remote",
				zap.String("cid", cid),
			)
			metrics.IpfsEchoAttempts.WithLabelValues(
				"remote", "success",
			).Add(1.0)
			rr.Status = "success"
		}
		rr.Finished()
		metrics.IpfsEchoHistogram.WithLabelValues(
			"remote", rr.Status,
		).Observe(rr.Duration)

		if rr.Status == "success" {
			lr := metrics.NewIpfsEcho(msg)
			lr.Started()
			localShell := shell.NewShellWithClient(*localAPI, &http.Client{Timeout: 5 * time.Second})
			obj, err := localShell.ObjectGet(cid)
			if err != nil {
				logger.Error(
					"Failed to get the object from CID",
					zap.String("cid", cid),
					zap.Error(err),
				)
				metrics.IpfsEchoAttempts.WithLabelValues(
					"local", "fail",
				).Add(1.0)
				rr.Status = "failed"
			} else {
				logger.Info(
					"Successfully get the object",
					zap.String("obj", obj.Data),
				)
				metrics.IpfsEchoAttempts.WithLabelValues(
					"local", "success",
				).Add(1.0)
				rr.Status = "success"
			}
			lr.Finished()
			metrics.IpfsEchoHistogram.WithLabelValues(
				"local", lr.Status,
			).Observe(lr.Duration)
		}
		time.Sleep(30 * time.Second)
	}

}
