package main

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	env := make(Environment)
	for _, entry := range dirEntries {
		fileName := entry.Name()

		if entry.IsDir() || strings.Contains(fileName, "=") {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			return nil, err
		}

		if info.Size() == 0 {
			env[fileName] = EnvValue{
				NeedRemove: true,
			}
		} else {
			value, err := readValueFromFile(filepath.Join(dir, fileName))
			if err != nil {
				return nil, err
			}
			env[fileName] = EnvValue{
				Value: value,
			}
		}
	}
	return env, nil
}

func readValueFromFile(filePath string) (value string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	line, _, err := reader.ReadLine()
	if err != nil {
		return
	}

	// пробелы и табуляция в конце T удаляются, терминальные нули (0x00) заменяются на перевод строки (\n)
	line = bytes.ReplaceAll(line, []byte{0x00}, []byte("\n"))
	value = strings.TrimRight(string(line), "\t ")

	return
}
