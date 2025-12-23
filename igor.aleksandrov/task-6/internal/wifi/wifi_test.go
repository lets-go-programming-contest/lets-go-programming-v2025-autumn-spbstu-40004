package wifi

import (
	"errors"
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UniversalMockWiFiHandle struct {
	mock.Mock
}

func (m *UniversalMockWiFiHandle) Interfaces() ([]*wifi.Interface, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*wifi.Interface), args.Error(1)
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

	mockWiFi := new(UniversalMockWiFiHandle)
	service := New(mockWiFi)

	mockWiFi.On("Interfaces").Return([]*wifi.Interface{
		{HardwareAddr: mustMAC("00:11:22:33:44:55")},
		{HardwareAddr: mustMAC("aa:bb:cc:dd:ee:ff")},
	}, nil)

	addrs, err := service.GetAddresses()
	assert.NoError(t, err)
	assert.Len(t, addrs, 2)
	assert.Equal(t, mustMAC("00:11:22:33:44:55"), addrs[0])
	assert.Equal(t, mustMAC("aa:bb:cc:dd:ee:ff"), addrs[1])

	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetAddresses_Empty(t *testing.T) {
	t.Parallel()

	mockWiFi := new(UniversalMockWiFiHandle)
	service := New(mockWiFi)

	mockWiFi.On("Interfaces").Return([]*wifi.Interface{}, nil)

	addrs, err := service.GetAddresses()
	assert.NoError(t, err)
	assert.Empty(t, addrs)

	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetAddresses_Error(t *testing.T) {
	t.Parallel()

	mockWiFi := new(UniversalMockWiFiHandle)
	service := New(mockWiFi)

	mockWiFi.On("Interfaces").Return(nil, errors.New("fail"))

	addrs, err := service.GetAddresses()
	assert.Error(t, err)
	assert.Nil(t, addrs)
	assert.Contains(t, err.Error(), "getting interfaces")

	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetNames(t *testing.T) {
	t.Parallel()

	mockWiFi := new(UniversalMockWiFiHandle)
	service := New(mockWiFi)

	mockWiFi.On("Interfaces").Return([]*wifi.Interface{
		{Name: "wlan0"},
		{Name: "wlan1"},
		{Name: "eth0"},
	}, nil)

	names, err := service.GetNames()
	assert.NoError(t, err)
	assert.Len(t, names, 3)
	assert.Equal(t, []string{"wlan0", "wlan1", "eth0"}, names)

	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetNames_Empty(t *testing.T) {
	t.Parallel()

	mockWiFi := new(UniversalMockWiFiHandle)
	service := New(mockWiFi)

	mockWiFi.On("Interfaces").Return([]*wifi.Interface{}, nil)

	names, err := service.GetNames()
	assert.NoError(t, err)
	assert.Empty(t, names)

	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetNames_Error(t *testing.T) {
	t.Parallel()

	mockWiFi := new(UniversalMockWiFiHandle)
	service := New(mockWiFi)

	mockWiFi.On("Interfaces").Return(nil, errors.New("fail"))

	names, err := service.GetNames()
	assert.Error(t, err)
	assert.Nil(t, names)
	assert.Contains(t, err.Error(), "getting interfaces")

	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_NilHandle(t *testing.T) {
	t.Parallel()

	service := New(nil)
	assert.NotNil(t, service)
	assert.Nil(t, service.WiFi)
}
