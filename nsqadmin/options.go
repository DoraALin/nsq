package nsqadmin

import (
	"github.com/absolute8511/nsq/internal/levellogger"
	"time"
)

type Options struct {
	HTTPAddress string `flag:"http-address"`

	GraphiteURL   string `flag:"graphite-url"`
	ProxyGraphite bool   `flag:"proxy-graphite"`

	UseStatsdPrefixes   bool   `flag:"use-statsd-prefixes"`
	StatsdPrefix        string `flag:"statsd-prefix"`
	StatsdCounterFormat string `flag:"statsd-counter-format"`
	StatsdGaugeFormat   string `flag:"statsd-gauge-format"`

	StatsdInterval time.Duration `flag:"statsd-interval"`

	NSQLookupdHTTPAddresses []string `flag:"lookupd-http-address" cfg:"nsqlookupd_http_addresses"`
	NSQDHTTPAddresses       []string `flag:"nsqd-http-address" cfg:"nsqd_http_addresses"`

	HTTPClientTLSInsecureSkipVerify bool   `flag:"http-client-tls-insecure-skip-verify"`
	HTTPClientTLSRootCAFile         string `flag:"http-client-tls-root-ca-file"`
	HTTPClientTLSCert               string `flag:"http-client-tls-cert"`
	HTTPClientTLSKey                string `flag:"http-client-tls-key"`

	NotificationHTTPEndpoint string `flag:"notification-http-endpoint"`

	LogDir string `flag:"log-dir" cfg:"log_dir"`
	Logger levellogger.Logger
}

func NewOptions() *Options {
	return &Options{
		HTTPAddress:         "0.0.0.0:4171",
		UseStatsdPrefixes:   true,
		StatsdPrefix:        "nsq.%s",
		StatsdCounterFormat: "stats.counters.%s.count",
		StatsdGaugeFormat:   "stats.gauges.%s",
		StatsdInterval:      60 * time.Second,
		Logger:              &levellogger.GLogger{},
	}
}
