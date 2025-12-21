package mocks

import (
	wifilib "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/mock"
)

type WiFiHandle struct {
	mock.Mock
}

func (_m *WiFiHandle) Interfaces() ([]*wifilib.Interface, error) {
	args := _m.Called()

	var interfaces []*wifilib.Interface
	if val := args.Get(0); val != nil {
		interfaces = val.([]*wifilib.Interface)
	}

	return interfaces, args.Error(1)
}
