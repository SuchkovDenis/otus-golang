package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("require exit code 1 if empty cmd", func(t *testing.T) {
		require.Equal(t, 1, RunCmd(nil, nil))
		require.Equal(t, 1, RunCmd([]string{}, nil))
	})

	t.Run("require standard return", func(t *testing.T) {
		require.Equal(t, 0, RunCmd([]string{""}, nil))
	})

	t.Run("require unset envVar", func(t *testing.T) {
		varKey, varValue := "TEST_ENV_UNSET", "TEST_VALUE"
		err := os.Setenv(varKey, varValue)
		require.NoError(t, err)
		require.Equal(t, varValue, os.Getenv(varKey))

		require.Equal(t, 0, RunCmd([]string{""}, Environment{varKey: EnvValue{NeedRemove: true}}))
		_, ok := os.LookupEnv(varKey)
		require.False(t, ok)
	})

	t.Run("require set envVar", func(t *testing.T) {
		varKey, varValue := "TEST_ENV_SET", "TEST_VALUE"
		_, ok := os.LookupEnv(varKey)
		require.False(t, ok)

		require.Equal(t, 0, RunCmd([]string{""}, Environment{varKey: EnvValue{Value: varValue}}))
		value, ok := os.LookupEnv(varKey)
		require.True(t, ok)
		require.Equal(t, varValue, value)
	})

	t.Run("require exec command", func(t *testing.T) {
		out := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		RunCmd([]string{"echo", "Hello world"}, nil)

		require.NoError(t, w.Close())
		os.Stdout = out

		var buf bytes.Buffer
		_, err := io.Copy(&buf, r)
		require.NoError(t, err)

		require.Equal(t, "Hello world\n", buf.String())
	})
}
