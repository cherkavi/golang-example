package unittesting

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// flag.Parse() - to use external flags
	os.Exit(m.Run())
}
