package main

import (
	"authentication/data"
	"os"
	"testing"
)

var testApp *Config

func TestMain(m *testing.M) {
	repo := data.NewPostgresTestRepository(nil)
	testApp = &Config{Repo: repo}
	os.Exit(m.Run())
}
