package dbuser

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/vier21/pc-01-network-be/config"
)

func TestConnection(t *testing.T) {
	config.InitConfig(".")
	logger := logrus.New()
	client, cancel, ctx, err := NewMongoConnection(logger)
	defer CloseConnection(ctx, logger, client, cancel)
	
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if client == nil {
		t.Fail()
	}

}
