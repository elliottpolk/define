package main

import (
	"os"
	"strings"

	"github.com/pkg/errors"
)

func mkdir(path string) error {
	if _, err := os.Stat(path); err != nil {
		if !os.IsNotExist(err) {
			return errors.Wrap(err, "unable to stat path")
		}

		//  attempt to generate if not exists
		if err := os.MkdirAll(path, 0700); err != nil {
			return errors.Wrap(err, "unable to make all directories")
		}
	}

	return nil
}

func simpl(name string) string {
	s := strings.Split(name, ",")
	return strings.TrimSpace(s[len(s)-1])
}
