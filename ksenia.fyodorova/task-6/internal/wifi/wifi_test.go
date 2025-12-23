package wifi_test

import (
	"errors"
	"net"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/lolnyok/task-6/internal/wifi"
	"github.com/mdlayher/wifi"
)

type MockWiFiHandle struct {
	mock.Mock
}

func (m *MockWiFiHandle) Interfaces() ([]*wifi.Interface, error) {
	args := m.Called()
	return args.Get(0).([]*wifi.Interface), args.Error(1)
}

func TestWiFiService_GetAddresses_Success(t *testing.T) {
	t.Parallel()

	mockWiFi := new(MockWiFiHandle)
	service := wifi.New(mockWiFi)

	hwAddr1, _ := net.ParseMAC("00:11:22:33:44:55")
	hwAddr2, _ := net.ParseMAC("AA:BB:CC:DD:EE:FF")

	interfaces := []*wifi.Interface{
		{
			Index:        1,
			Name:         "wlan0",
			HardwareAddr: hwAddr1,
		},
		{
			Index:        2,
			Name:         "wlan1",
			HardwareAddr: hwAddr2,
		},
	}

	expectedAddresses := []net.HardwareAddr{hwAddr1, hwAddr2}

	mockWiFi.On("Interfaces").Return(interfaces, nil)

	addresses, err := service.GetAddresses()

	require.NoError(t, err)
	require.Equal(t, expectedAddresses, addresses)
	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetAddresses_Error(t *testing.T) {
	t.Parallel()

	mockWiFi := new(MockWiFiHandle)
	service := wifi.New(mockWiFi)

	testErr := errors.New("interface error")
	mockWiFi.On("Interfaces").Return([]*wifi.Interface{}, testErr)

	addresses, err := service.GetAddresses()

	require.Error(t, err)
	require.Nil(t, addresses)
	require.ErrorContains(t, err, "getting interfaces")
	require.ErrorIs(t, err, testErr)
	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetAddresses_Empty(t *testing.T) {
	t.Parallel()

	mockWiFi := new(MockWiFiHandle)
	service := wifi.New(mockWiFi)

	mockWiFi.On("Interfaces").Return([]*wifi.Interface{}, nil)

	addresses, err := service.GetAddresses()

	require.NoError(t, err)
	require.Empty(t, addresses)
	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetNames_Success(t *testing.T) {
	t.Parallel()

	mockWiFi := new(MockWiFiHandle)
	service := wifi.New(mockWiFi)

	hwAddr1, _ := net.ParseMAC("00:11:22:33:44:55")
	hwAddr2, _ := net.ParseMAC("AA:BB:CC:DD:EE:FF")

	interfaces := []*wifi.Interface{
		{
			Index:        1,
			Name:         "wlan0",
			HardwareAddr: hwAddr1,
		},
		{
			Index:        2,
			Name:         "wlan1",
			HardwareAddr: hwAddr2,
		},
	}

	expectedNames := []string{"wlan0", "wlan1"}

	mockWiFi.On("Interfaces").Return(interfaces, nil)

	names, err := service.GetNames()

	require.NoError(t, err)
	require.Equal(t, expectedNames, names)
	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetNames_Error(t *testing.T) {
	t.Parallel()

	mockWiFi := new(MockWiFiHandle)
	service := wifi.New(mockWiFi)

	testErr := errors.New("interface error")
	mockWiFi.On("Interfaces").Return([]*wifi.Interface{}, testErr)

	names, err := service.GetNames()

	require.Error(t, err)
	require.Nil(t, names)
	require.ErrorContains(t, err, "getting interfaces")
	require.ErrorIs(t, err, testErr)
	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetNames_Empty(t *testing.T) {
	t.Parallel()

	mockWiFi := new(MockWiFiHandle)
	service := wifi.New(mockWiFi)

	mockWiFi.On("Interfaces").Return([]*wifi.Interface{}, nil)

	names, err := service.GetNames()

	require.NoError(t, err)
	require.Empty(t, names)
	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetNames_WithNilHardwareAddr(t *testing.T) {
	t.Parallel()

	mockWiFi := new(MockWiFiHandle)
	service := wifi.New(mockWiFi)

	interfaces := []*wifi.Interface{
		{
			Index:        1,
			Name:         "wlan0",
			HardwareAddr: nil,
		},
		{
			Index:        2,
			Name:         "wlan1",
			HardwareAddr: []byte{0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF},
		},
	}

	expectedNames := []string{"wlan0", "wlan1"}

	mockWiFi.On("Interfaces").Return(interfaces, nil)

	names, err := service.GetNames()

	require.NoError(t, err)
	require.Equal(t, expectedNames, names)
	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetAddresses_WithNilHardwareAddr(t *testing.T) {
	t.Parallel()

	mockWiFi := new(MockWiFiHandle)
	service := wifi.New(mockWiFi)

	hwAddr2, _ := net.ParseMAC("AA:BB:CC:DD:EE:FF")

	interfaces := []*wifi.Interface{
		{
			Index:        1,
			Name:         "wlan0",
			HardwareAddr: nil,
		},
		{
			Index:        2,
			Name:         "wlan1",
			HardwareAddr: hwAddr2,
		},
	}

	expectedAddresses := []net.HardwareAddr{nil, hwAddr2}

	mockWiFi.On("Interfaces").Return(interfaces, nil)

	// Act
	addresses, err := service.GetAddresses()

	// Assert
	require.NoError(t, err)
	require.Equal(t, expectedAddresses, addresses)
	mockWiFi.AssertExpectations(t)
}
