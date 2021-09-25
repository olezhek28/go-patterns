package singleton

import (
	"fmt"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetInstance(t *testing.T) {
	var (
		firstName  = "First name"
		secondName = "Second name"
	)

	t.Run("sync case", func(t *testing.T) {
		s := GetInstance()
		if s == nil {
			t.Fatalf("First sigletone is nil")
		}

		s.SetName(firstName)
		require.Equal(t, firstName, s.GetName())

		s2 := GetInstance()
		require.Equal(t, s, s2)

		s2.SetName(secondName)
		require.Equal(t, secondName, s.GetName())
	})

	t.Run("async case", func(t *testing.T) {
		s1 := GetInstance()
		s2 := GetInstance()

		var w sync.WaitGroup

		for i := 0; i < 3000; i++ {
			j := i

			w.Add(1)
			go func() {
				s1.SetName(firstName + " - " + strconv.Itoa(j))
				w.Done()
			}()

			w.Add(1)
			go func() {
				s2.SetName(secondName + " - " + strconv.Itoa(j))
				w.Done()
			}()
		}

		fmt.Println(s1.GetName())
		fmt.Println(s2.GetName())
	})
}
