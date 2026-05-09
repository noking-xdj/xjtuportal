package http

import (
	"testing"
	"xjtuportal/component/basic"
)

var (
	conf = map[string][]string{
		"1080:1083":   {"shadowsocks", "shadowsocksR", "privoxy"},
		"2800:2803":   {"Netch"},
		"7890:7893":   {"clash", "Clash for Windows"},
		"8080:8088":   {},
		"10808:10810": {"v2rayN"},
	}
)

func TestProxyCheck(t *testing.T) {
	configHelper := &basic.ConfigHelper{
		ProgramSettings: &basic.ProgramSettings{},
	}
	configHelper.ProgramSettings.ProgramConnectivitySettings.Proxy.Ports = conf
	configHelper.ProgramSettings.ProgramConnectivitySettings.Proxy.TestUrl = testUrl
	configHelper.ProgramSettings.ProgramConnectivitySettings.Proxy.Timeout = 1

	p := InitProxyHelper(basic.LoggerTemp, configHelper)

	if len(p.proxyPorts) != 24 {
		t.Fatalf("expected 24 parsed ports, got %d", len(p.proxyPorts))
	}
	if _, ok := p.proxyPorts[1080]; !ok {
		t.Error("expected range start port 1080 to be parsed")
	}
	if _, ok := p.proxyPorts[1083]; !ok {
		t.Error("expected range end port 1083 to be parsed")
	}
	if _, ok := p.proxyPorts[8080]; !ok {
		t.Error("expected single port 8080 to be parsed")
	}
}
