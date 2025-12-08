package wifi_test

import (
	"errors"
	"net"
	"testing"

	"github.com/15446-rus75/task-6/internal/wifi"
	wifimock "github.com/15446-rus75/task-6/internal/wifi/_mocks"

	wifilib "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var errFail = errors.New("fail")

func TestGetAddresses(t *testing.T) {
	t.Parallel()

	mockWiFi := new(wifimock.WiFiHandle)
	service := wifi.New(mockWiFi)

	interfaces := []*wifilib.Interface{
		{HardwareAddr: net.HardwareAddr{0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF}},
		{HardwareAddr: net.HardwareAddr{0x11, 0x22, 0x33, 0x44, 0x55, 0x66}},
	}

	mockWiFi.On("Interfaces").Return(interfaces, nil)

	addrs, err := service.GetAddresses()

	require.NoError(t, err)
	assert.Len(t, addrs, 2)
	assert.Equal(t, interfaces[0].HardwareAddr, addrs[0])
	assert.Equal(t, interfaces[1].HardwareAddr, addrs[1])

	mockWiFi.AssertExpectations(t)
}

func TestGetAddresses_Empty(t *testing.T) {
	t.Parallel()

	mockWiFi := new(wifimock.WiFiHandle)
	service := wifi.New(mockWiFi)

	mockWiFi.On("Interfaces").Return([]*wifilib.Interface{}, nil)

	addrs, err := service.GetAddresses()

	require.NoError(t, err)
	assert.Empty(t, addrs)

	mockWiFi.AssertExpectations(t)
}

func TestGetAddresses_Error(t *testing.T) {
	t.Parallel()

	mockWiFi := new(wifimock.WiFiHandle)
	service := wifi.New(mockWiFi)

	mockWiFi.On("Interfaces").Return(nil, errFail)

	addrs, err := service.GetAddresses()

	require.Error(t, err)
	assert.Nil(t, addrs)
	assert.Contains(t, err.Error(), "getting interfaces")

	mockWiFi.AssertExpectations(t)
}

func TestGetNames(t *testing.T) {
	t.Parallel()

	mockWiFi := new(wifimock.WiFiHandle)
	service := wifi.New(mockWiFi)

	interfaces := []*wifilib.Interface{
		{Name: "wlan0"},
		{Name: "wlan1"},
		{Name: "eth0"},
	}

	mockWiFi.On("Interfaces").Return(interfaces, nil)

	names, err := service.GetNames()

	require.NoError(t, err)
	assert.Len(t, names, 3)
	assert.Equal(t, []string{"wlan0", "wlan1", "eth0"}, names)

	mockWiFi.AssertExpectations(t)
}

func TestGetNames_Empty(t *testing.T) {
	t.Parallel()

	mockWiFi := new(wifimock.WiFiHandle)
	service := wifi.New(mockWiFi)

	mockWiFi.On("Interfaces").Return([]*wifilib.Interface{}, nil)

	names, err := service.GetNames()

	require.NoError(t, err)
	assert.Empty(t, names)

	mockWiFi.AssertExpectations(t)
}

func TestGetNames_Error(t *testing.T) {
	t.Parallel()

	mockWiFi := new(wifimock.WiFiHandle)
	service := wifi.New(mockWiFi)

	mockWiFi.On("Interfaces").Return(nil, errFail)

	names, err := service.GetNames()

	require.Error(t, err)
	assert.Nil(t, names)
	assert.Contains(t, err.Error(), "getting interfaces")

	mockWiFi.AssertExpectations(t)
}
