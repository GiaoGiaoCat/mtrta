package mtrta

import (
	"testing"
	"time"

	yaml "gopkg.in/yaml.v2"
)

func TestConfigMarshal(t *testing.T) {
	c := &Config{
		Dial:      10 * time.Second,
		KeepAlive: 1 * time.Second,
		MaxConns:  10,
		MaxIdle:   10,
		Version:   0,
	}
	bb, _ := yaml.Marshal(c)
	t.Log(string(bb))
}

func TestConfigUnMarshal(t *testing.T) {
	cfgText := `
dial: 10s
keepalive: 1s
maxconns: 10
maxidle: 10
version: 0
`

	var c Config
	err := yaml.Unmarshal([]byte(cfgText), &c)
	t.Log(err)
	t.Log(c)
}
