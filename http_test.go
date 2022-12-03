package com

import (
	"fmt"
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
