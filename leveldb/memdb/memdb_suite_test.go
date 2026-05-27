package memdb

import (
	"testing"

	"github.com/Decentral-America/goleveldb/leveldb/testutil"
)

func TestMemDB(t *testing.T) {
	testutil.RunSuite(t, "MemDB Suite")
}
