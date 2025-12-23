package wifi_test

import (
	"errors"
	"net"
	"testing"

	mywifi "github.com/MrMels625/task-6/internal/wifi"
	wifilib "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var errFail = errors.New("fail")

type MockWiFiHandle struct {
	mock.Mock
}

func (m *MockWiFiHandle) Interfaces() ([]*wifilib.Interface, error) {
	args := m.Called()

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	val, ok := args.Get(0).([]*wifilib.Interface)
	if !ok {
		return nil, errors.New("type assertion failed")
	}

	return val, args.Error(1)
}

func mustMAC(s string) net.HardwareAddr {
	mac, err := net.ParseMAC(s)
	if err != nil {
		panic(err)
	}

	return mac
}

func TestWiFiService_GetAddresses(t *testing.T) {
	t.Parallel()

	mockWiFi := &MockWiFiHandle{}
	service := mywifi.New(mockWiFi)

	mockWiFi.On("Interfaces").Return([]*wifilib.Interface{
		{HardwareAddr: mustMAC("00:11:22:33:44:55")},
		{HardwareAddr: mustMAC("aa:bb:cc:dd:ee:ff")},
	}, nil)

	addrs, err := service.GetAddresses()
	require.NoError(t, err)
	assert.Len(t, addrs, 2)
	assert.Equal(t, mustMAC("00:11:22:33:44:55"), addrs[0])
	assert.Equal(t, mustMAC("aa:bb:cc:dd:ee:ff"), addrs[1])

	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetAddresses_Empty(t *testing.T) {
	t.Parallel()

	mockWiFi := &MockWiFiHandle{}
	service := mywifi.New(mockWiFi)

	mockWiFi.On("Interfaces").Return([]*wifilib.Interface{}, nil)

	addrs, err := service.GetAddresses()
	require.NoError(t, err)
	assert.Empty(t, addrs)

	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetAddresses_InterfaceWithNilMAC(t *testing.T) {
	t.Parallel()

	mockWiFi := &MockWiFiHandle{}
	service := mywifi.New(mockWiFi)

	mockWiFi.On("Interfaces").Return([]*wifilib.Interface{
		{HardwareAddr: nil},
	}, nil)

	addrs, err := service.GetAddresses()
	require.NoError(t, err)
	assert.Len(t, addrs, 1)
	assert.Nil(t, addrs[0])

	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetAddresses_Error(t *testing.T) {
	t.Parallel()

	mockWiFi := &MockWiFiHandle{}
	service := mywifi.New(mockWiFi)

	mockWiFi.On("Interfaces").Return(nil, errFail)

	addrs, err := service.GetAddresses()
	require.Error(t, err)
	assert.Nil(t, addrs)
	assert.Contains(t, err.Error(), "getting interfaces")

	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetNames(t *testing.T) {
	t.Parallel()

	mockWiFi := &MockWiFiHandle{}
	service := mywifi.New(mockWiFi)

	mockWiFi.On("Interfaces").Return([]*wifilib.Interface{
		{Name: "wlan0"},
		{Name: "wlan1"},
		{Name: "eth0"},
	}, nil)

	names, err := service.GetNames()
	require.NoError(t, err)
	assert.Len(t, names, 3)
	assert.Equal(t, []string{"wlan0", "wlan1", "eth0"}, names)

	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetNames_Empty(t *testing.T) {
	t.Parallel()

	mockWiFi := &MockWiFiHandle{}
	service := mywifi.New(mockWiFi)

	mockWiFi.On("Interfaces").Return([]*wifilib.Interface{}, nil)

	names, err := service.GetNames()
	require.NoError(t, err)
	assert.Empty(t, names)

	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetNames_EmptyName(t *testing.T) {
	t.Parallel()

	mockWiFi := &MockWiFiHandle{}
	service := mywifi.New(mockWiFi)

	mockWiFi.On("Interfaces").Return([]*wifilib.Interface{
		{Name: ""},
	}, nil)

	names, err := service.GetNames()
	require.NoError(t, err)
	assert.Equal(t, []string{""}, names)

	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetNames_DuplicateNames(t *testing.T) {
	t.Parallel()

	mockWiFi := &MockWiFiHandle{}
	service := mywifi.New(mockWiFi)

	mockWiFi.On("Interfaces").Return([]*wifilib.Interface{
		{Name: "wlan0"},
		{Name: "wlan0"},
	}, nil)

	names, err := service.GetNames()
	require.NoError(t, err)
	assert.Equal(t, []string{"wlan0", "wlan0"}, names)

	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetNames_Error(t *testing.T) {
	t.Parallel()

	mockWiFi := &MockWiFiHandle{}
	service := mywifi.New(mockWiFi)

	mockWiFi.On("Interfaces").Return(nil, errFail)

	names, err := service.GetNames()
	require.Error(t, err)
	assert.Nil(t, names)
	assert.Contains(t, err.Error(), "getting interfaces")

	mockWiFi.AssertExpectations(t)
}

func TestNew(t *testing.T) {
	t.Parallel()

	mockWiFi := &MockWiFiHandle{}
	service := mywifi.New(mockWiFi)

	assert.NotNil(t, service)
	assert.Equal(t, mockWiFi, service.WiFi)
}

func TestNew_NilHandle(t *testing.T) {
	t.Parallel()

	service := mywifi.New(nil)
	assert.NotNil(t, service)
	assert.Nil(t, service.WiFi)
}
