package main

import (
	"errors"
	"io"
	"os"

	progressbar "github.com/schollz/progressbar/v3"
)

const chunkSize = 1024

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	inFile, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}

	stat, err := inFile.Stat()
	if err != nil {
		return ErrUnsupportedFile
	}

	fileSize := stat.Size()
	if offset > fileSize {
		return ErrOffsetExceedsFileSize
	}

	outFile, err := os.Create(toPath)
	if err != nil {
		return err
	}

	if limit == 0 {
		limit = fileSize
	}
	if (fileSize - offset) < limit {
		limit = fileSize - offset
	}

	bar := progressbar.Default(limit)
	buff := make([]byte, chunkSize)
	for limit > 0 {
		readBytes, err := inFile.ReadAt(buff, offset)
		if err != nil && !errors.Is(err, io.EOF) {
			return err
		}

		writeBytes := readBytes
		if int64(writeBytes) > limit {
			writeBytes = int(limit)
		}
		_, err = outFile.Write(buff[:writeBytes])
		if err != nil {
			return err
		}

		limit -= int64(readBytes)
		offset += int64(chunkSize)

		err = bar.Add(writeBytes)
		if err != nil {
			return err
		}
	}
	err = outFile.Chmod(stat.Mode().Perm())
	if err != nil {
		return err
	}

	return nil
}
