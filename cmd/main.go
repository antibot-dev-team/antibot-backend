package main

import (
	"fmt"
	"github.com/antibot-dev-team/antibot-backend/antibot"
	"github.com/antibot-dev-team/antibot-backend/antibot/http"
	"github.com/ardanlabs/conf"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const (
	namespace = "Antibot"
	logPrefix = "main"
)

func main() {
	logger := logrus.New()

	if err := run(logger); err != nil {
		logger.Errorf("%s: error: %s", logPrefix, err)
		os.Exit(1)
	}
}

func run(logger *logrus.Logger) error {
	// =========================================================================
	// Configuration

	var cfg antibot.Config
	cfg.Version.SVN = "1.0.0-dev"

	if err := conf.Parse(os.Args[1:], namespace, &cfg); err != nil {
		switch err {
		case conf.ErrHelpWanted:
			usage, err := conf.Usage(namespace, &cfg)
			if err != nil {
				return errors.Wrap(err, "generating config usage")
			}
			fmt.Println(usage)
			return nil
		case conf.ErrVersionWanted:
			version, err := conf.VersionString(namespace, &cfg)
			if err != nil {
				return errors.Wrap(err, "generating config version")
			}
			fmt.Println(version)
			return nil
		}
		return errors.Wrap(err, "parsing config")
	}

	// =========================================================================
	// Logging

	logger.SetOutput(os.Stdout)
	ll, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		ll = logrus.InfoLevel
		logger.Warnf("Unknown loglevel %s, used INFO instead", cfg.LogLevel)
	}
	logger.SetLevel(ll)

	// =========================================================================
	// App Starting

	logger.Infof("%s: Started", logPrefix)
	defer logger.Infof("%s: Completed", logPrefix)

	out, err := conf.String(&cfg)
	if err != nil {
		return errors.Wrap(err, "generating config for output")
	}
	logger.Infof("%s: Config:\n%v\n", logPrefix, out)

	logger.Infof("%s: Antibot started", logPrefix)

	// =========================================================================
	// Start Antibot HTTP server

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	handler := http.HandleAPI(cfg.Version.SVN, logger, shutdown, &cfg)

	s := fasthttp.Server{
		Handler:      handler,
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
	}

	serverErrors := make(chan error, 1)

	go func() {
		logger.Infof("%s: %s is listening on %s", logPrefix, namespace, net.JoinHostPort(cfg.Web.Host, cfg.Web.Port))
		serverErrors <- s.ListenAndServe(net.JoinHostPort(cfg.Web.Host, cfg.Web.Port))
	}()

	select {
	case err := <-serverErrors:
		return errors.Wrap(err, "server error")

	case sig := <-shutdown:
		logger.Infof("%s: %v: Start shutdown", logPrefix, sig)
		if err := s.Shutdown(); err != nil {
			return errors.Wrap(err, "could not stop server gracefully")
		}
		logger.Infof("%s: %v: Completed shutdown", logPrefix, sig)
	}

	return nil
}
