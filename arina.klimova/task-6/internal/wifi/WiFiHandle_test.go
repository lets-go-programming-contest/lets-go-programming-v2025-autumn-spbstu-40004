package wifi

import (
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"
)

func TestWiFiHandleImplementation(t *testing.T) {
	handle := &testWiFiHandle{}

	var _ WiFiHandle = handle
	assert.True(t, true)
}

func TestNewWiFiService(t *testing.T) {
	handle := &testWiFiHandle{}
	service := New(handle)

	assert.NotNil(t, service)
	assert.Equal(t, handle, service.WiFi)
}

type anotherWiFiHandle struct{}

func (a *anotherWiFiHandle) Interfaces() ([]*wifi.Interface, error) {
	return []*wifi.Interface{}, nil
}

func TestAnotherWiFiHandleImplementation(t *testing.T) {
	handle := &anotherWiFiHandle{}
	var _ WiFiHandle = handle
	assert.True(t, true)
}
