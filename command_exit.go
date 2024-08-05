package main

import "os"

func callbackExit(config *config, args ...string) error {
	os.Exit(0)
	return nil
}