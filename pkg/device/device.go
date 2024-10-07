package device

import (
	"context"
	"errors"

	"github.com/vier21/pc-01-network-be/config"
)

var (
	JuniperDev = "juniper"
	CiscoDev   = "cisco"
)

type Device interface {
	InitRouter() IRouter
	InitSwitch() ISwitch
}

type IRouter interface {
	Ping(ctx context.Context)
	AddStaticRoute(ctx context.Context)
	AddACL(ctx context.Context)
	ShowInterface(ctx context.Context)
	EnableInterface(ctx context.Context)
	DisableInterface(ctx context.Context)
	ShowVLAN(ctx context.Context)
}

type ISwitch interface {
	CreateVlan(ctx context.Context)
	ShowVlan(ctx context.Context)
}

func GetDevice(brand string) (Device, error) {
	if brand == CiscoDev {
		return &ciscoDevice{
			PlaybookDIRRouter: config.GetConfig().PlaybookDirCiscoRouter,
			PlaybookDIRSwitch: config.GetConfig().PlaybookDirCiscoSwitch,
		}, nil
	}

	return nil, errors.New("error")
}
