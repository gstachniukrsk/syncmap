package syncmap_test

import (
	"runtime"
	"syncmap"
	"testing"
)

func TestSwapWorks(t *testing.T) {
	t.Run("swap works", func(t *testing.T) {
		// check if go version is below 1.20
		goVersion := runtime.Version()

		if goVersion < "go1.20" {
			t.Skip("go version is below 1.20")
		}

		m := syncmap.Map{}
		m.Store("key", "value")
		m.CompareAndSwap("key", "value", "new value")
	})
}
