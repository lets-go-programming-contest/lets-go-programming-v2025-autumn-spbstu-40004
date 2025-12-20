package wifi

import (
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:generate mockery --name=WiFiHandle --testonly --quiet --outpkg=wifi_test --output=.

func createMockInterface(name string, mac string) *wifi.Interface {
	hwAddr, _ := net.ParseMAC(mac)
	return &wifi.Interface{
		Index:        1,
		Name:         name,
		HardwareAddr: hwAddr,
		PHY:          0,
		Device:       0,
		Type:         wifi.InterfaceTypeStation,
		Frequency:    2412,
	}
}

type testWiFiHandle struct {
	interfaces []*wifi.Interface
	err        error
}

func (t *testWiFiHandle) Interfaces() ([]*wifi.Interface, error) {
	return t.interfaces, t.err
}

func TestNew(t *testing.T) {
	handle := &testWiFiHandle{}
	service := New(handle)

	assert.NotNil(t, service)
	assert.Equal(t, handle, service.WiFi)
}

func TestWiFiService_GetAddresses(t *testing.T) {
	handle := &testWiFiHandle{
		interfaces: []*wifi.Interface{
			createMockInterface("wlan0", "00:11:22:33:44:55"),
		},
	}
	service := New(handle)

	addresses, err := service.GetAddresses()

	assert.NoError(t, err)
	require.Len(t, addresses, 1)

	expectedAddr, _ := net.ParseMAC("00:11:22:33:44:55")
	assert.Equal(t, expectedAddr, addresses[0])
}

func TestWiFiService_GetNames(t *testing.T) {
	handle := &testWiFiHandle{
		interfaces: []*wifi.Interface{
			createMockInterface("wlan0", "00:11:22:33:44:55"),
		},
	}
	service := New(handle)

	names, err := service.GetNames()

	assert.NoError(t, err)
	require.Len(t, names, 1)
	assert.Equal(t, "wlan0", names[0])
}
