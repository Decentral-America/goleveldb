package leveldb

import (
	"testing"

	"github.com/Decentral-America/goleveldb/leveldb/testutil"
)

func TestLevelDB(t *testing.T) {
	testutil.RunSuite(t, "LevelDB Suite")
}
