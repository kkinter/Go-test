package main

import (
	"os"
	"testing"
	"wepapp/pkg/repository/dbrepo"
)

var app application

func TestMain(m *testing.M) {
	pathToTemplates = "./../../templates/"
	app.Session = getSession()

	app.DB = &dbrepo.TestDBRepo{}
	os.Exit(m.Run())
}
