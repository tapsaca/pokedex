package main

import "os"

func callbackExit(config *config) error {
	os.Exit(0)
	return nil
}