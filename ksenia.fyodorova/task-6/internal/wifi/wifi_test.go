package wifi_test

import (
	"errors"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	mywifi "github.com/lolnyok/task-6/internal/wifi"
	extwifi "github.com/mdlayher/wifi"
)

type MockWiFiHandle struct {
	mock.Mock
}

func (m *MockWiFiHandle) Interfaces() ([]*extwifi.Interface, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*extwifi.Interface), args.Error(1)
}

func TestWiFiService_GetAddresses_Success(t *testing.T) {
	mockWiFi := new(MockWiFiHandle)
	service := mywifi.New(mockWiFi)

	hwAddr1, _ := net.ParseMAC("00:11:22:33:44:55")
	hwAddr2, _ := net.ParseMAC("AA:BB:CC:DD:EE:FF")

	interfaces := []*extwifi.Interface{
		{Index: 1, Name: "wlan0", HardwareAddr: hwAddr1},
		{Index: 2, Name: "wlan1", HardwareAddr: hwAddr2},
	}

	expectedAddresses := []net.HardwareAddr{hwAddr1, hwAddr2}

	mockWiFi.On("Interfaces").Return(interfaces, nil)

	addresses, err := service.GetAddresses()

	require.NoError(t, err)
	assert.Equal(t, expectedAddresses, addresses)
	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetAddresses_Error(t *testing.T) {
	mockWiFi := new(MockWiFiHandle)
	service := mywifi.New(mockWiFi)

	testErr := errors.New("interface error")
	mockWiFi.On("Interfaces").Return([]*extwifi.Interface{}, testErr)

	addresses, err := service.GetAddresses()

	require.Error(t, err)
	require.Nil(t, addresses)
	require.ErrorContains(t, err, "getting interfaces")
	require.ErrorIs(t, err, testErr)
	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetAddresses_Empty(t *testing.T) {
	mockWiFi := new(MockWiFiHandle)
	service := mywifi.New(mockWiFi)

	mockWiFi.On("Interfaces").Return([]*extwifi.Interface{}, nil)

	addresses, err := service.GetAddresses()

	require.NoError(t, err)
	require.Empty(t, addresses)
	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetAddresses_NilInterface(t *testing.T) {
	mockWiFi := new(MockWiFiHandle)
	service := mywifi.New(mockWiFi)

	interfaces := []*extwifi.Interface{nil}
	mockWiFi.On("Interfaces").Return(interfaces, nil)

	addresses, err := service.GetAddresses()

	require.NoError(t, err)
	require.Len(t, addresses, 1)
	assert.Nil(t, addresses[0])
	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetNames_Success(t *testing.T) {
	mockWiFi := new(MockWiFiHandle)
	service := mywifi.New(mockWiFi)

	hwAddr1, _ := net.ParseMAC("00:11:22:33:44:55")
	hwAddr2, _ := net.ParseMAC("AA:BB:CC:DD:EE:FF")

	interfaces := []*extwifi.Interface{
		{Index: 1, Name: "wlan0", HardwareAddr: hwAddr1},
		{Index: 2, Name: "wlan1", HardwareAddr: hwAddr2},
	}

	expectedNames := []string{"wlan0", "wlan1"}

	mockWiFi.On("Interfaces").Return(interfaces, nil)

	names, err := service.GetNames()

	require.NoError(t, err)
	assert.Equal(t, expectedNames, names)
	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetNames_Error(t *testing.T) {
	mockWiFi := new(MockWiFiHandle)
	service := mywifi.New(mockWiFi)

	testErr := errors.New("interface error")
	mockWiFi.On("Interfaces").Return([]*extwifi.Interface{}, testErr)

	names, err := service.GetNames()

	require.Error(t, err)
	require.Nil(t, names)
	require.ErrorContains(t, err, "getting interfaces")
	require.ErrorIs(t, err, testErr)
	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetNames_Empty(t *testing.T) {
	mockWiFi := new(MockWiFiHandle)
	service := mywifi.New(mockWiFi)

	mockWiFi.On("Interfaces").Return([]*extwifi.Interface{}, nil)

	names, err := service.GetNames()

	require.NoError(t, err)
	require.Empty(t, names)
	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetNames_NilInterface(t *testing.T) {
	mockWiFi := new(MockWiFiHandle)
	service := mywifi.New(mockWiFi)

	interfaces := []*extwifi.Interface{nil}
	mockWiFi.On("Interfaces").Return(interfaces, nil)

	names, err := service.GetNames()

	require.NoError(t, err)
	require.Len(t, names, 1)
	assert.Equal(t, "", names[0])
	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetNames_WithNilName(t *testing.T) {
	mockWiFi := new(MockWiFiHandle)
	service := mywifi.New(mockWiFi)

	hwAddr, _ := net.ParseMAC("00:11:22:33:44:55")
	interfaces := []*extwifi.Interface{
		{Index: 1, Name: "", HardwareAddr: hwAddr},
	}

	expectedNames := []string{""}

	mockWiFi.On("Interfaces").Return(interfaces, nil)

	names, err := service.GetNames()

	require.NoError(t, err)
	assert.Equal(t, expectedNames, names)
	mockWiFi.AssertExpectations(t)
}
