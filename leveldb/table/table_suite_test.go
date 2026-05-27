package table

import (
	"testing"

	"github.com/Decentral-America/goleveldb/leveldb/testutil"
)

func TestTable(t *testing.T) {
	testutil.RunSuite(t, "Table Suite")
}
