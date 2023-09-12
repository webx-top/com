package com

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseRetryAfter(t *testing.T) {
	now := time.Now().UTC()
	tim := now.Format(`15:04:05`)
	year := now.AddDate(1, 0, 0).Year()
	retryAfter := fmt.Sprintf(`Wed, 21 Oct %d %s GMT`, year, tim)
	d := ParseRetryAfter(retryAfter)
	assert.Equal(t, fmt.Sprintf(`%d-10-21T%sZ`, year, now.Add(-time.Second).Format(`15:04:05`)), now.Add(d).UTC().Format(time.RFC3339))
	retryAfter = `300`
	d = ParseRetryAfter(retryAfter)
	assert.Equal(t, float64(5), d.Minutes())
}

func TestIsNetworkOrHostDown(t *testing.T) {
	r := IsNetworkOrHostDown(context.Canceled, false)
	assert.False(t, r)
	r = IsNetworkOrHostDown(context.DeadlineExceeded, false)
	assert.True(t, r)
	r = IsNetworkOrHostDown(context.DeadlineExceeded, true)
	assert.False(t, r)

	r = IsNetworkOrHostDown(&url.Error{Err: &net.DNSError{}}, true)
	assert.True(t, r)
	r = IsNetworkOrHostDown(&url.Error{Err: &net.OpError{}}, true)
	assert.True(t, r)
	r = IsNetworkOrHostDown(&url.Error{Err: net.UnknownNetworkError(`err`)}, true)
	assert.True(t, r)

	var netErr net.Error = &net.DNSError{IsTimeout: true, IsTemporary: true}
	r = IsNetworkOrHostDown(netErr, true)
	assert.True(t, r)

	r = IsNetworkOrHostDown(errors.New(`connection refused`), true)
	assert.True(t, r)
}
