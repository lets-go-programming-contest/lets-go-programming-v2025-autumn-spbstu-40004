package wifi_test

import (
	"errors"
	"net"
	"testing"

	wifiInternal "github.com/belyaevEDU/task-6/internal/wifi"
	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var errGettingInterfaces = errors.New("getting interfaces")

func TestGetAddresses(t *testing.T) {
	t.Parallel()

	mockWifi := NewWiFiHandle(t)
	service := wifiInternal.New(mockWifi)

	interfaces := []*wifi.Interface{
		{HardwareAddr: net.HardwareAddr{0x52, 0x69, 0x48}},
		{HardwareAddr: net.HardwareAddr{0xAA, 0xBB, 0xCC}},
	}

	mockWifi.On("Interfaces").Return(interfaces, nil)

	addresses, err := service.GetAddresses()

	require.NoError(t, err)
	assert.Len(t, addresses, len(interfaces))

	for index, elem := range interfaces {
		assert.Equal(t, elem.HardwareAddr, addresses[index])
	}

	mockWifi.AssertExpectations(t)
}

func TestGetAddressesInterfacesError(t *testing.T) {
	t.Parallel()

	mockWifi := NewWiFiHandle(t)
	service := wifiInternal.New(mockWifi)

	mockWifi.On("Interfaces").Return(nil, errGettingInterfaces)

	addresses, err := service.GetAddresses()

	require.Error(t, err)
	assert.Contains(t, err.Error(), errGettingInterfaces.Error())
	assert.Nil(t, addresses)

	mockWifi.AssertExpectations(t)
}

func TestGetNames(t *testing.T) {
	t.Parallel()

	mockWifi := NewWiFiHandle(t)
	service := wifiInternal.New(mockWifi)

	interfaces := []*wifi.Interface{
		{Name: "en0"},
		{Name: "bridge0"},
	}

	mockWifi.On("Interfaces").Return(interfaces, nil)

	names, err := service.GetNames()

	require.NoError(t, err)
	assert.Len(t, names, len(interfaces))

	for index, elem := range interfaces {
		assert.Equal(t, elem.Name, names[index])
	}

	mockWifi.AssertExpectations(t)
}

func TestGetNamesInterfacesError(t *testing.T) {
	t.Parallel()

	mockWifi := NewWiFiHandle(t)
	service := wifiInternal.New(mockWifi)

	mockWifi.On("Interfaces").Return(nil, errGettingInterfaces)

	names, err := service.GetNames()

	require.Error(t, err)
	assert.Contains(t, err.Error(), errGettingInterfaces.Error())
	assert.Nil(t, names)

	mockWifi.AssertExpectations(t)
}
