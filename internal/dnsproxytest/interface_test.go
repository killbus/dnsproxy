package dnsproxytest_test

import (
	"github.com/killbus/dnsproxy/internal/dnsproxytest"
	"github.com/killbus/dnsproxy/upstream"
)

// type check
var _ upstream.Upstream = (*dnsproxytest.FakeUpstream)(nil)
