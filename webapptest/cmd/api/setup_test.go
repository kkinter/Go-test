package main

import (
	"os"
	"testing"
	"wepapp/pkg/repository/dbrepo"
)

var app application

var expiredToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiYXVkIjoiZXhhbXBsZS5jb20iLCJleHAiOjE2ODQ3ODQxOTQsImlzcyI6ImV4YW1wbGUuY29tIiwibmFtZSI6IkpvaG4gRG9lIiwic3ViIjoiMSJ9.7c8NtPD_RfTgV2Pz12eirvt-E5mT7Ggdu_wxVzs-xOk"

func TestMain(m *testing.M) {
	app.DB = &dbrepo.TestDBRepo{}
	app.Domain = "example.com"
	app.JWTSecret = "jwtSecret"

	os.Exit(m.Run())
}
