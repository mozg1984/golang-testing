package test

import (
	"os"
	"testing"

	"github.com/federicoleon/golang-restclient/rest"
	"github.com/mozg1984/golang-testing/src/api/app"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	go app.StartApp()
	os.Exit(m.Run())
}
