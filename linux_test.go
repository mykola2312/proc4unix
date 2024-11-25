//go:build linux

package proc

import "testing"

func TestQuery(t *testing.T) {
	if _, err := Query(1); err != nil {
		t.Failed()
	}
}
