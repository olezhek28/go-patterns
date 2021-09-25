package db

import (
	"fmt"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetConnect(t *testing.T) {
	var (
		db1 = "First database"
		db2 = "Second database"

		table1 = "First table"
		table2 = "Second table"
	)

	t.Run("sync case", func(t *testing.T) {
		con := GetConnect(db1)
		if con == nil {
			t.Fatalf("First connection to db is nil")
		}

		con.SetTableName(table1)
		require.Equal(t, table1, con.GetTableName())
		require.Equal(t, db1, con.GetDBName())

		con2 := GetConnect(db2)
		require.Equal(t, con, con2)

		con2.SetTableName(table2)
		require.Equal(t, table2, con2.GetTableName())
		require.Equal(t, table2, con.GetTableName())
		require.Equal(t, db1, con2.GetDBName())
	})

	// check 'go test -race'
	t.Run("async case", func(t *testing.T) {
		con1 := GetConnect(db1)
		con2 := GetConnect(db2)

		var w sync.WaitGroup

		for i := 0; i < 3000; i++ {
			j := i

			w.Add(1)
			go func() {
				con1.SetTableName(table1 + " - " + strconv.Itoa(j))
				w.Done()
			}()

			w.Add(1)
			go func() {
				con2.SetTableName(table2 + " - " + strconv.Itoa(j))
				w.Done()
			}()
		}

		fmt.Println(con1.GetTableName())
		fmt.Println(con2.GetTableName())
	})
}
