// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rapid_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/flyingmutant/rapid"
)

// ParseDate parses dates in the YYYY-MM-DD format.
func ParseDate(s string) (int, int, int, error) {
	if len(s) != 10 {
		return 0, 0, 0, fmt.Errorf("%q has wrong length: %v instead of 10", s, len(s))
	}

	if s[4] != '-' || s[7] != '-' {
		return 0, 0, 0, fmt.Errorf("'-' separators expected in %q", s)
	}

	y, err := strconv.Atoi(s[0:4])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("failed to parse year: %v", err)
	}

	m, err := strconv.Atoi(s[5:7])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("failed to parse month: %v", err)
	}

	d, err := strconv.Atoi(s[8:10])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("failed to parse day: %v", err)
	}

	return y, m, d, nil
}

// Rename to TestParseDate to make an actual (failing) test.
func TestParseDate(t *testing.T) {
	rapid.Check(t, func(t *rapid.T, y int, m int, d int) {
		s := fmt.Sprintf("%04d-%02d-%02d", y, m, d)

		y_, m_, d_, err := ParseDate(s)
		if err != nil {
			t.Fatalf("failed to parse date %q: %v", s, err)
		}

		if y_ != y || m_ != m || d_ != d {
			t.Fatalf("got back wrong date: (%d, %d, %d)", y_, m_, d_)
		}
	}, rapid.IntsRange(0, 9999), rapid.IntsRange(1, 12), rapid.IntsRange(1, 31))
}
