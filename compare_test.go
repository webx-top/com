package com

import (
	"testing"
)

func TestVersionCompare(t *testing.T) {
	if VersionCompare(`5.7.19`, `5.8`) != VersionCompareLt {
		t.Error(`Error: 5.7.19 >= 5.8`)
	}
	if VersionCompare(`5.0.19`, `5.0.2`) != VersionCompareGt {
		t.Error(`Error: 5.0.19 <= 5.0.2`)
	}
	if VersionCompare(`5.0.19`, `5.0.20`) != VersionCompareLt {
		t.Error(`Error: 5.0.19 >= 5.0.20`)
	}
	if VersionCompare(`5.10.19`, `5.5.20`) != VersionCompareGt {
		t.Error(`Error: 5.10.19 <= 5.5.20`)
	}
	if VersionCompare(`5.10.19`, `11.5.20`) != VersionCompareLt {
		t.Error(`Error: 5.10.19 >= 11.5.20`)
	}
	if VersionCompare(`6.1.1`, `6.1.1`) != VersionCompareEq {
		t.Error(`Error: 6.1.1 != 6.1.1`)
	}
	if VersionCompare(`6.1`, `6.1.1`) != VersionCompareLt {
		t.Error(`Error: 6.1 >= 6.1.1`)
	}
	if VersionCompare(`6.1.0`, `6.1.0-alpha`) != VersionCompareGt {
		t.Error(`Error: 6.1.0 <= 6.1.0-alpha`)
	}
	if VersionCompare(`6.1.0-beta`, `6.1.0-alpha`) != VersionCompareGt {
		t.Error(`Error: 6.1.0-beta <= 6.1.0-alpha`)
	}
	if VersionCompare(`6.1.0-beta`, `6.1.0-beta`) != VersionCompareEq {
		t.Error(`Error: 6.1.0-beta != 6.1.0-beta`)
	}
	if VersionCompare(`6.1.0-beta2`, `6.1.0-beta1`) != VersionCompareGt {
		t.Error(`Error: 6.1.0-beta2 <= 6.1.0-beta1`)
	}
	if VersionCompare(`6.1.0-stable`, `6.1.0-beta1`) != VersionCompareGt {
		t.Error(`Error: 6.1.0-stable <= 6.1.0-beta1`)
	}
	if VersionCompare(`6.1.0.1`, `6.1.0`) != VersionCompareGt {
		t.Error(`Error: 6.1.0.1 <= 6.1.0`)
	}
	if VersionCompare(`6.1.0`, `6.1.0.1`) != VersionCompareLt {
		t.Error(`Error: 6.1.0 >= 6.1.0.1`)
	}
	if VersionCompare(`6.1.0.1`, `6.1.0.2`) != VersionCompareLt {
		t.Error(`Error: 6.1.0.1 >= 6.1.0.2`)
	}
}
