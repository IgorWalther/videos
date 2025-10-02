package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGreater(t *testing.T) {
	require.Equal(t, "Hello, World!", greater("World"))
}
