package state_test

import (
	"os"
	"testing"

	"github.com/romakingwolf/patriot/core/types"
)

func TestMain(m *testing.M) {
	types.RegisterMockEvidencesGlobal()
	os.Exit(m.Run())
}
