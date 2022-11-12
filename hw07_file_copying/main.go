package main

import (
	"flag"
	"fmt"
)

var (
	from, to      string
	limit, offset int64
)

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to write to")
	flag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
}

func main() {
	flag.Parse()

	if validation, message := validateInput(from, to, limit, offset); !validation {
		fmt.Println(message)
		return
	}

	if err := Copy(from, to, offset, limit); err != nil {
		fmt.Println(err)
	}
}

func validateInput(from, to string, limit, offset int64) (bool, string) {
	if from == "" {
		return false, "--from arg is required"
	}
	if to == "" {
		return false, "--to arg is required"
	}
	if limit < 0 {
		return false, "--limit illegal arg: must be not negative"
	}
	if offset < 0 {
		return false, "--offset illegal arg: must be not negative"
	}
	return true, ""
}
