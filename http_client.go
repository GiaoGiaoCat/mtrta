package mtrta

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Dial      time.Duration
	KeepAlive time.Duration
	MaxConns  int
	MaxIdle   int
	Version   int
	Mock      bool
}

var (
	cfgVersion         = -1
	httpClient         *http.Client
	clientInstanceLock sync.Mutex
)

func GetHTTPClient(c *Config) *http.Client {
	if c == nil {
		return http.DefaultClient
	}

	if httpClient != nil && c.Version == cfgVersion {
		return httpClient
	}

	clientInstanceLock.Lock()
	defer clientInstanceLock.Unlock()

	if httpClient != nil && c.Version == cfgVersion {
		return httpClient
	}

	cfgVersion = c.Version

	dialer := &net.Dialer{
		Timeout:   c.Dial,
		KeepAlive: c.KeepAlive,
	}
	transport := &http.Transport{
		DialContext:         dialer.DialContext,
		MaxConnsPerHost:     c.MaxConns,
		MaxIdleConnsPerHost: c.MaxIdle,
		IdleConnTimeout:     c.KeepAlive,
	}

	httpClient = &http.Client{
		Transport: transport,
	}

	info, _ := yaml.Marshal(c)
	fmt.Printf("new client with \n%v\n", string(info))

	return httpClient
}
