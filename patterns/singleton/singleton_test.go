package singleton

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetInstance(t *testing.T) {
	var (
		firstName  = "First name"
		secondName = "Second name"
	)

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
}
