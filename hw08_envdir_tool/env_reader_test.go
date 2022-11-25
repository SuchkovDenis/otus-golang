package main

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	environment, err := ReadDir(filepath.Join("testdata", "env"))
	expected := Environment{
		"BAR":   EnvValue{Value: "bar"},
		"EMPTY": EnvValue{Value: ""},
		"FOO":   EnvValue{Value: "   foo\nwith new line"},
		"HELLO": EnvValue{Value: "\"hello\""},
		"UNSET": EnvValue{NeedRemove: true},
	}

	require.NoError(t, err)
	require.Equal(t, expected, environment)
}
