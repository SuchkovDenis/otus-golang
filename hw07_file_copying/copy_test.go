package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("File not exist", func(t *testing.T) {
		inputFile := "test_data_input0.txt"
		outFile := "test_data_out0.txt"

		err := Copy(inputFile, outFile, 0, 0)
		require.Equal(t, ErrUnsupportedFile, err)
	})

	t.Run("Write all file without offset", func(t *testing.T) {
		inputFile := "test_data_input1.txt"
		outFile := "test_data_out1.txt"
		testData := []byte("Test data")

		prepareTestFiles(t, inputFile, testData)

		err := Copy(inputFile, outFile, 0, 0)
		require.NoError(t, err)

		assertData(t, outFile, testData)

		removeTestFiles(t, inputFile, outFile)
	})

	t.Run("Write all file without offset limit 1000", func(t *testing.T) {
		inputFile := "test_data_input2.txt"
		outFile := "test_data_out2.txt"
		testData := []byte("Test data")

		prepareTestFiles(t, inputFile, testData)

		err := Copy(inputFile, outFile, 0, 1000)
		require.NoError(t, err)

		assertData(t, outFile, testData)

		removeTestFiles(t, inputFile, outFile)
	})

	t.Run("Write all file offset 1 limit 4", func(t *testing.T) {
		inputFile := "test_data_input3.txt"
		outFile := "test_data_out3.txt"
		testData := []byte("Test data")

		prepareTestFiles(t, inputFile, testData)

		err := Copy(inputFile, outFile, 1, 4)
		require.NoError(t, err)

		assertData(t, outFile, testData[1:5])

		removeTestFiles(t, inputFile, outFile)
	})

	t.Run("Write all file offset exceeded", func(t *testing.T) {
		inputFile := "test_data_input4.txt"
		outFile := "test_data_out4.txt"
		testData := []byte("Test data")

		prepareTestFiles(t, inputFile, testData)

		err := Copy(inputFile, outFile, 100, 0)
		require.Equal(t, ErrOffsetExceedsFileSize, err)

		removeTestFiles(t, inputFile)
	})
}

func prepareTestFiles(t *testing.T, inputFile string, testData []byte) {
	t.Helper()
	err := os.WriteFile(inputFile, testData, os.ModePerm)
	require.NoError(t, err)
}

func removeTestFiles(t *testing.T, files ...string) {
	t.Helper()
	for _, file := range files {
		err := os.Remove(file)
		require.NoError(t, err)
	}
}

func assertData(t *testing.T, outFile string, expectedData []byte) {
	t.Helper()
	outData, err := os.ReadFile(outFile)
	require.NoError(t, err)
	require.Equal(t, expectedData, outData)
}
