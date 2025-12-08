package mocks

import (
    "github.com/mdlayher/wifi"
    "github.com/stretchr/testify/mock"
)

type WiFiHandle struct {
    mock.Mock
}

func (_m *WiFiHandle) Interfaces() ([]*wifi.Interface, error) {
    args := _m.Called()
    
    var interfaces []*wifi.Interface
    if val := args.Get(0); val != nil {
        interfaces = val.([]*wifi.Interface)
    }
    
    return interfaces, args.Error(1)
}
