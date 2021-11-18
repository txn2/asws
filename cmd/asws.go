package main

import (
	"flag"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

var (
	ipEnv                   = getEnv("IP", "127.0.0.1")
	portEnv                 = getEnv("PORT", "8080")
	staticDirEnv            = getEnv("STATIC_DIR", "./www")
	staticPathEnv           = getEnv("STATIC_PATH", "/")
	notFoundRedirectEnv     = getEnv("NOT_FOUND_REDIRECT", "false")
	notFoundRedirectPathEnv = getEnv("NOT_FOUND_REDIRECT_PATH", "/")
	notFoundFileEnv         = getEnv("NOT_FOUND_FILE", "./www/404.html")
	fsEnabledEnv            = getEnv("FS_ENABLED", "false")
	fsDirEnv                = getEnv("FS_DIR", "./files")
	fsPathEnv               = getEnv("FS_PATH", "/files")
	debugEnv                = getEnv("DEBUG", "false")
	metricsEnv              = getEnv("METRICS", "true")
	metricsPortEnv          = getEnv("METRICS_PORT", "2112")
)

var Version = "0.0.0"
var Service = "asws"

func main() {

	var (
		ip                   = flag.String("ip", ipEnv, "bind ip")
		port                 = flag.String("port", portEnv, "port")
		staticDir            = flag.String("staticDir", staticDirEnv, "static dir")
		staticPath           = flag.String("staticPath", staticPathEnv, "static path")
		notFoundRedirect     = flag.String("notFoundRedirect", notFoundRedirectEnv, "redirect on not found?")
		notFoundRedirectPath = flag.String("notFoundRedirectPath", notFoundRedirectPathEnv, "not found redirect path")
		notFoundFile         = flag.String("notFoundFile", notFoundFileEnv, "not found file to serve")
		fsEnabled            = flag.String("fsEnabled", fsEnabledEnv, "filesystem enabled")
		fsDir                = flag.String("fsDir", fsDirEnv, "filesystem directory")
		fsPath               = flag.String("fsPath", fsPathEnv, "filesystem path")
		debug                = flag.String("debug", debugEnv, "debug")
		metrics              = flag.String("metrics", metricsEnv, "metrics")
		metricsPort          = flag.String("metricsPort", metricsPortEnv, "metrics port")
	)
	flag.Parse()

	// add some useful info to metrics
	promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "www",
		Name:      "info",
		ConstLabels: prometheus.Labels{
			"go_version": runtime.Version(),
			"version":    Version,
			"service":    Service,
		},
	}).Inc()

	gin.SetMode(gin.ReleaseMode)

	if *debug == "true" {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()

	zapCfg := zap.NewProductionConfig()
	zapCfg.DisableCaller = true
	zapCfg.DisableStacktrace = true

	logger, _ := zapCfg.Build()

	if *debug == "true" {
		logger, _ = zap.NewDevelopment()
	}

	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	if *fsEnabled == "true" {
		r.StaticFS(*fsPath, http.Dir(*fsDir))
	}

	r.Static(*staticPath, *staticDir)

	r.NoRoute(func(c *gin.Context) {
		if *notFoundRedirect == "true" {
			c.Redirect(http.StatusTemporaryRedirect, *notFoundRedirectPath)
			c.Abort()
			return
		}

		content, err := os.ReadFile(*notFoundFile)
		if err != nil {
			logger.Error("404 content not found", zap.String("file", *notFoundFile))
			c.String(http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}

		c.Writer.WriteHeader(http.StatusNotFound)
		_, err = c.Writer.Write(content)
		if err != nil {
			logger.Error("404 write error", zap.Error(err))
			c.String(http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}
	})

	// metrics server (run in go routine)
	if *metrics == "true" {
		go func() {
			http.Handle("/metrics", promhttp.Handler())

			logger.Info("Starting ASWS Metrics Server",
				zap.String("type", "start_asws_metrics"),
				zap.String("version", Version),
				zap.String("port", *metricsPort),
				zap.String("ip", *ip),
			)

			err := http.ListenAndServe(*ip+":"+*metricsPort, nil)
			if err != nil {
				logger.Fatal("Error Starting ASWS Metrics Server", zap.Error(err))
				os.Exit(1)
			}
		}()
	}

	logger.Info("Starting ASWS Server",
		zap.String("type", "start_asws"),
		zap.String("version", Version),
		zap.String("port", *port),
		zap.String("ip", *ip),
	)

	// Gin web server
	err := r.Run(*ip + ":" + *port)
	if err != nil {
		logger.Fatal(err.Error())
	}
}

// getEnv gets an environment variable or sets a default if
// one does not exist.
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}

	return value
}
